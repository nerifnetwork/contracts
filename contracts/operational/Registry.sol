// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../interfaces/SignerOwnable.sol";
import "../interfaces/IGateway.sol";
import "./RegistryGateway.sol";
import "./RegistryWorkflow.sol";
import "./RegistryBalance.sol";

// TODO: 1. Charge user for the workflow registration and cancellation in order to perform the internal workflow.
//          Charged amount to the reward distribution pool.
// TODO: 2. Reward distribution on mainchain and sidechain logic.
//          Mainchain: send to the reward distribution pool.

// Registry is the internal smart contract needed to secure the network and support important features of Nerif.
contract Registry is Initializable, SignerOwnable, RegistryGateway, RegistryWorkflow, RegistryBalance {
    // Config contains the configuration options
    struct Config {
        // performanceOverhead is the cost of the performance transaction excluding the client contract call.
        uint256 performanceOverhead;
        // performancePremiumThreshold is the network premium threshold in percents.
        uint8 performancePremiumThreshold;
        // registrationOverhead is the cost of the workflow registration.
        uint256 registrationOverhead;
        // cancellationOverhead is the cost of the workflow cancellation.
        uint256 cancellationOverhead;
        // maxWorkflowsPerAccount is the maximum number of workflows per user.
        // 0 value means there is no limit.
        uint16 maxWorkflowsPerAccount;
    }

    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;
    string internal constant GATEWAY_PERFORM_FUNC_SIGNATURE = "perform(uint256,address,bytes)";

    bool public isMainChain;
    Config public config;
    uint256 public networkRewards;

    event BalanceFunded(address addr, uint256 amount);
    event BalanceWithdrawn(address addr, uint256 amount);
    event RewardsWithdrawn(address addr, uint256 amount);
    event GatewaySet(address owner, address gateway);
    event WorkflowRegistered(address owner, uint256 id, bytes hash);
    event WorkflowStatusChanged(uint256 id, WorkflowStatus status);
    event Performance(uint256 id, uint256 gasUsed, bool success);

    // onlyMainchain permits transactions on the mainchain only
    modifier onlyMainchain() {
        require(isMainChain, "Registry: operation is not permitted");
        _;
    }

    // onlyRegistry permits transaction coming from the contract itself.
    // This is needed for those transaction that must pass the performance process.
    modifier onlyRegistry() {
        require(address(this) == msg.sender, "Registry: operation is not permitted");
        _;
    }

    // onlyMsgSender checks that the given address is the transaction sender one.
    modifier onlyMsgSender(address addr) {
        require(addr == msg.sender, "Registry: operation is not permitted");
        _;
    }

    // onlyMsgSender modifier for mainchain OR onlyRegistry for sidechain
    modifier onlyMsgSenderOrRegistry(address addr) {
        if (isMainChain) {
            require(addr == msg.sender, "Registry: operation is not permitted");
            _;
        } else {
            require(address(this) == msg.sender, "Registry: operation is not permitted");
            _;
        }
    }

    // onlyWorkflowOwnerOrRegistry permits operation for the workflow owner if it is mainnet
    // or for the network on sidechains.
    modifier onlyWorkflowOwnerOrRegistry(uint256 id) {
        if (isMainChain) {
            Workflow memory workflow = getWorkflow(id);
            require(workflow.owner == msg.sender, "Registry: operation is not permitted");
            _;
        } else {
            // Only network can execute the function on the sidechain.
            // The transaction must come from the network after reaching consensus.
            // Basically, the transaction must come from the registry contract itself,
            // namely from the perform function after passing all checks.
            require(address(this) == msg.sender, "Registry: operation is not permitted");
            _;
        }
    }

    function initialize(
        bool _isMainChain,
        address _signerGetterAddress,
        Workflow[] calldata _internalWorkflows,
        Config calldata _config
    ) external initializer {
        isMainChain = _isMainChain;
        _setSignerGetter(_signerGetterAddress);

        for (uint256 i = 0; i < _internalWorkflows.length; i++) {
            _addWorkflow(_internalWorkflows[i]);
        }

        config = _config;
    }

    // setConfig sets the given configuration
    function setConfig(Config calldata _config) external onlySigner {
        config = _config;
    }

    // fundBalance funds the balance of the sender's address with the given amount.
    function fundBalance() external payable {
        uint256 currentBalance = getBalance(msg.sender);
        _setBalance(msg.sender, currentBalance + msg.value);
        emit BalanceFunded(msg.sender, msg.value);
    }

    // withdrawBalance withdraws the remaining balance of the sender's public key.
    // Permissions:
    //  - Only balance owner can withdraw its balance.
    // TODO: Handle cases when the withdrawal happens during the workflow execution. Introduce withdrawal request.
    function withdrawBalance() external {
        address payable sender = payable(msg.sender);
        uint256 balance = getBalance(sender);

        // Ensure the sender has a positive balance
        require(balance > 0, "Registry: no balance to withdraw");

        // Update the sender's balance
        _setBalance(sender, 0);

        // Transfer the balance to the sender
        sender.transfer(balance);

        // Emit an event to log the withdrawal transaction
        emit BalanceWithdrawn(sender, balance);
    }

    // withdrawRewards sends network rewards to the rewards withdrawal address
    function withdrawRewards() external {
        require(networkRewards > 0, "Registry: nothing to withdraw");
        require(address(signerGetter) != address(0x0), "Registry: signer storage address is not specified");

        address payable addr = payable(signerGetter.getSignerAddress());
        require(addr != address(0x0), "Registry: withdrawal address is not specified");

        // Transfer rewards
        addr.transfer(networkRewards);
        networkRewards = 0;

        // Emit an event to log the withdrawal transaction
        emit RewardsWithdrawn(addr, networkRewards);
    }

    // setGateway sets gateway for the given owner address.
    function setGateway(address gateway) external {
        _setGateway(msg.sender, IGateway(gateway));
        emit GatewaySet(msg.sender, gateway);
    }

    // pauseWorkflow pauses an existing active workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can pause an existing active workflow.
    function pauseWorkflow(uint256 id) external onlyMainchain onlyExistingWorkflow(id) onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Check current workflow status
        require(workflow.status == WorkflowStatus.ACTIVE, "Registry: only active workflows could be paused");

        // Update status
        workflow.status = WorkflowStatus.PAUSED;
        require(_updateWorkflow(workflow), "Registry: failed to update workflow");

        emit WorkflowStatusChanged(id, WorkflowStatus.PAUSED);
    }

    // resumeWorkflow resumes an existing paused workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can resume an existing active workflow.
    function resumeWorkflow(uint256 id) external onlyMainchain onlyExistingWorkflow(id) onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Check current workflow status
        require(workflow.status == WorkflowStatus.PAUSED, "Registry: only paused workflows could be resumed");

        // Update status
        workflow.status = WorkflowStatus.ACTIVE;
        require(_updateWorkflow(workflow), "Registry: failed to update workflow");

        emit WorkflowStatusChanged(id, WorkflowStatus.ACTIVE);
    }

    // perform performs the contract execution defined in the registered workflow.
    // The function checks that the given performance transaction was signed by the majority
    // of the network so the workflow owner could be charged and the transaction
    // with the given payload could be passed to the customer's contract.
    // Arguments:
    //  - "workflowId" is the workflow ID
    //  - "gasAmount" is the maximum number of gas used to execute the transaction
    //  - "data" is the contract call data
    //  - "target" is the client contract address
    // Permissions:
    //  - Only network can execute this function.
    function perform(
        uint256 workflowId,
        uint256 gasAmount,
        bytes calldata data,
        address target
    ) external onlySigner onlyExistingWorkflow(workflowId) {
        // Get a workflow by ID
        Workflow memory workflow = getWorkflow(workflowId);
        require(workflow.id > 0, "Registry: workflow not found");

        // Make sure the workflow is not paused
        require(workflow.status == WorkflowStatus.ACTIVE, "Registry: workflow must be active");

        // Get current balance of workflow owner
        uint256 currentBalance = getBalance(workflow.owner);

        if (!workflow.isInternal) {
            // Make sure workflow owner has enough funds
            require(currentBalance > 0, "Registry: not enough funds on balance");

            // Cannot self-execute if not internal
            require(address(this) != target, "Registry: operation is not permitted");
        }

        // TODO: Make sure the given transaction was not performed yet

        // Execute client's contract through gateway
        uint256 gasUsed = gasleft();
        bool success = false;
        if (address(this) == target) {
            success = _callWithExactGas(gasAmount, target, data);
        } else {
            // Get workflow owner's gateway
            IGateway existingGateway = getGateway(workflow.owner); // TODO: Make sure it is not zero address

            // Execute customer contract through its gateway
            success = _callWithExactGas(
                gasAmount,
                address(existingGateway),
                abi.encodeWithSignature(GATEWAY_PERFORM_FUNC_SIGNATURE, workflowId, target, data)
            );
        }
        gasUsed -= gasleft();

        // Adding performance overhead if exists
        if (config.performanceOverhead > 0) {
            gasUsed += config.performanceOverhead;
        }

        // Adding performance premium
        if (config.performancePremiumThreshold > 0) {
            gasUsed += gasUsed / uint256(config.performancePremiumThreshold);
        }

        // Calculate amount to charge
        uint256 amountToCharge = gasUsed * tx.gasprice;

        if (!workflow.isInternal) {
            // Make sure owner has enough funds
            require(currentBalance >= amountToCharge, "Registry: not enough funds on balance");

            // Charge workflow owner balance
            _setBalance(workflow.owner, currentBalance - amountToCharge);

            // Move amount to the network rewards balance
            networkRewards += amountToCharge;
        }

        // Update total spent amount of the current workflow
        workflow.totalSpent += amountToCharge;
        require(_updateWorkflow(workflow), "Registry: failed to update workflow");

        // Emit performance event
        emit Performance(workflowId, gasUsed, success);
    }

    // registerWorkflow registers a new workflow metadata.
    // Arguments:
    //  - "id" is the workflow identifier.
    //  - "owner" is the workflow owner address.
    //  - "hash" is the workflow hash.
    // The given signature must correspond to the given hash and created by the transaction sender.
    // Permissions:
    //  - Only workflow owner can register a workflow on MAINCHAIN.
    //  - Only network can register a workflow on SIDECHAIN throught the regular performance process.
    function registerWorkflow(
        uint256 id,
        address owner,
        bytes calldata hash
    ) public onlyMsgSenderOrRegistry(owner) {
        // Check if the given workflow owner has a gateway registered.
        IGateway existingGateway = getGateway(owner);
        require(address(existingGateway) != address(0x0), "Registry: gateway not found");

        // Check if the given sender has capacity to create one more workflow
        if (isMainChain && config.maxWorkflowsPerAccount > 0) {
            require(
                _workflowsPerAddress(msg.sender) < config.maxWorkflowsPerAccount,
                "Registry: reached max workflows capacity"
            );
        }

        // Use ACTIVE workflow status by default for sidechains
        WorkflowStatus workflowStatus = WorkflowStatus.ACTIVE;

        // Or set the PENDING one for the mainchain
        if (isMainChain) {
            workflowStatus = WorkflowStatus.PENDING;
        }

        // Store a new workflow
        require(_addWorkflow(Workflow(id, owner, hash, workflowStatus, false, 0)), "Registry: failed to add workflow");

        // Emmit the event
        emit WorkflowRegistered(msg.sender, id, hash);
    }

    // activateWorkflow updates the workflow state from PENDING to ACTIVE.
    // Arguments:
    //  - "id" is the workflow identifier.
    //  - "status" is the workflow status.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only network can execute it through the regular performance process.
    function activateWorkflow(uint256 id) public onlyMainchain onlyRegistry onlyExistingWorkflow(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Must be PENDING
        require(workflow.status == WorkflowStatus.PENDING, "Registry: workflow must be pending");

        // Update status
        workflow.status = WorkflowStatus.ACTIVE;
        require(_updateWorkflow(workflow), "Registry: failed to update workflow");

        emit WorkflowStatusChanged(id, WorkflowStatus.ACTIVE);
    }

    // cancelWorkflow cancels an existing workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Only workflow owner can cancel an existing active workflow on MAINCHAIN.
    //  - Only network can cancel a workflow on SIDECHAIN through the regular performance process.
    function cancelWorkflow(uint256 id) public onlyExistingWorkflow(id) onlyWorkflowOwnerOrRegistry(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Check current workflow status
        require(workflow.status != WorkflowStatus.CANCELLED, "Registry: workflow is already cancelled");

        // Update status
        workflow.status = WorkflowStatus.CANCELLED;
        require(_updateWorkflow(workflow), "Registry: failed to update workflow");

        emit WorkflowStatusChanged(id, WorkflowStatus.CANCELLED);
    }

    // getWorkflowOwnerBalance returns the current balance of the given workflow ID.
    function getWorkflowOwnerBalance(uint256 id) public view returns (uint256) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Return just 1 for internal workflows so the workflow engine can identify that the workflow owner has funds.
        // TODO: Implement internal workflows logic in more accurate way
        if (workflow.isInternal) {
            return 1;
        }

        require(workflow.owner != address(0x0), "Registry: workflow does not exist");

        // Return owner's balance
        return getBalance(workflow.owner);
    }

    // _callWithExactGas calls target address with exactly gasAmount gas and data as calldata
    // or reverts if at least gasAmount gas is not available
    function _callWithExactGas(
        uint256 gasAmount,
        address target,
        bytes memory data
    ) private returns (bool success) {
        assembly {
            let g := gas()

            // Compute g -= PERFORM_GAS_CUSHION and check for underflow
            if lt(g, PERFORM_GAS_CUSHION) {
                revert(0, 0)
            }

            g := sub(g, PERFORM_GAS_CUSHION)

            // if g - g//64 <= gasAmount, revert
            // (we subtract g//64 because of EIP-150)
            if iszero(gt(sub(g, div(g, 64)), gasAmount)) {
                revert(0, 0)
            }

            // solidity calls check that a contract actually exists at the destination, so we do the same
            if iszero(extcodesize(target)) {
                revert(0, 0)
            }

            // call and return whether we succeeded. ignore return data
            success := call(gasAmount, target, 0, add(data, 0x20), mload(data), 0, 0)
        }
        return success;
    }
}
