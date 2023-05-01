// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../interfaces/ISignerAddress.sol";
import "../interfaces/IRegistry.sol";
import "../interfaces/IGateway.sol";

// TODO: 1. Charge user for the workflow registration and cancellation in order to perform the internal workflow.
//          Charged amount to the reward distribution pool.

// Registry is the internal smart contract needed to secure the network and support important features of Nerif.
contract Registry is Initializable, IRegistry {
    // Config contains the configuration options
    struct Config {
        // performanceOverhead is the cost of the performance transaction excluding the client contract call.
        // Numbers are different depending on the sidechain.
        // Metrics:
        //  - Ethereum Goerli execution: 49,022
        //  - BSC Testnet execution:     46,922
        //  - Polygon Mumbai execution:  31,476
        uint256 performanceOverhead;
        // performancePremiumThreshold is the network premium threshold in percents.
        // Numbers are different depending on the sidechain.
        // Metrics:
        //  - Ethereum Goerli execution: 10%
        //  - BSC Testnet execution:     10%
        //  - Polygon Mumbai execution:  10%
        uint8 performancePremiumThreshold;
        // registrationOverhead is the cost of the workflow registration.
        // Numbers are different depending on the sidechain.
        // Metrics:
        //  - BSC Testnet cost:     115,520 (registration, sidechain)
        //  - Ethereum Goerli cost: 122,420 (registration, sidechain)
        //  - Polygon Mumbai cost:  67,282 (approval, mainchain)
        uint256 registrationOverhead;
        // cancellationOverhead is the cost of the workflow cancellation.
        // Numbers are different depending on the sidechain.
        // Metrics:
        //  - BSC Testnet cost:     48,568
        //  - Ethereum Goerli cost: 50,168
        //  - Polygon Mumbai cost:
        uint256 cancellationOverhead;
        // maxWorkflowsPerAccount is the maximum number of workflows per user.
        // 0 value means there is no limit.
        uint16 maxWorkflowsPerAccount;
    }

    struct PerformPayload {
        uint256 id;
        uint256 gasAmount;
        address target;
        bytes data;
    }

    struct Gateway {
        IGateway gateway;
        address owner;
    }

    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;
    string internal constant GATEWAY_PERFORM_FUNC_SIGNATURE = "perform(uint256,address,bytes)";

    Config public config;
    Gateway[] internal gateways;
    mapping(uint256 => Workflow) internal workflows;
    mapping(address => uint16) internal workflowsPerAddress;
    mapping(address => uint256) internal balances;
    bool public isMainChain;
    ISignerAddress public networkAddress;
    uint256 public networkRewards;

    event BalanceFunded(address workflowOwner, uint256 amount);
    event BalanceWithdrawn(address workflowOwner, uint256 amount);
    event GatewaySet(address workflowOwner, address gateway);
    event WorkflowRegistered(address owner, uint256 id, bytes hash);
    event WorkflowActivated(uint256 id);
    event WorkflowPaused(uint256 id);
    event WorkflowResumed(uint256 id);
    event WorkflowCancelled(uint256 id);
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

    // onlyNetwork permits transactions coming from the collective network address.
    modifier onlyNetwork() {
        require(networkAddress.getSignerAddress() == msg.sender, "Registry: operation is not permitted");
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

    // onlyExistingWorkflow checks that the given workflow exists.
    modifier onlyExistingWorkflow(uint256 id) {
        Workflow storage workflow = workflows[id];
        require(workflow.id > 0, "Registry: workflow does not exist");
        _;
    }

    // onlyWorkflowOwner checks that the given tx sender is an actual workflow owner.
    modifier onlyWorkflowOwner(uint256 id) {
        Workflow storage workflow = workflows[id];
        require(workflow.id > 0, "Registry: workflow does not exist");
        _;
    }

    // onlyWorkflowOwnerOrRegistry permits operation for the workflow owner if it is mainnet
    // or for the network on sidechains.
    modifier onlyWorkflowOwnerOrRegistry(uint256 id) {
        if (isMainChain) {
            Workflow storage workflow = workflows[id];
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

    function initialize(address _networkAddress, bool _isMainChain) external initializer {
        networkAddress = ISignerAddress(_networkAddress);
        isMainChain = _isMainChain;

        // Define internal workflows
        // TODO: Implement more elegant way

        // Activate workflow on the mainchain side
        Workflow memory workflowCreationWorkflow = Workflow(
            40505927788353901442144037336646356013,
            address(0),
            "",
            WorkflowStatus.ACTIVE,
            true,
            0
        );
        workflows[workflowCreationWorkflow.id] = workflowCreationWorkflow;

        // Cancel workflow on the sidechain side
        Workflow memory workflowCancellationWorkflow = Workflow(
            219775546284901721155783592958414245131,
            address(0),
            "",
            WorkflowStatus.ACTIVE,
            true,
            0
        );
        workflows[workflowCancellationWorkflow.id] = workflowCancellationWorkflow;
    }

    // setConfig sets the given configuration
    function setConfig(Config calldata _config) external onlyNetwork {
        config = _config;
    }

    // fundBalance funds the balance of the sender's public key with the given amount.
    function fundBalance(address addr) external payable onlyMsgSender(addr) {
        // Update the balance value
        balances[msg.sender] += msg.value;

        emit BalanceFunded(addr, msg.value);
    }

    // withdrawBalance withdraws the remaining balance of the sender's public key.
    // Permissions:
    //  - Only balance owner can withdraw its balance.
    // TODO: Handle cases when the withdrawal happens during the workflow execution. Introduce withdrawal request.
    function withdrawBalance(address addr) external onlyMsgSender(addr) {
        address payable sender = payable(msg.sender);
        uint256 balance = balances[sender];

        // Ensure the sender has a positive balance
        require(balance > 0, "Registry: no balance to withdraw");

        // Update the sender's balance
        balances[sender] = 0;

        // Transfer the balance to the sender
        sender.transfer(balance);

        // Emit an event to log the withdrawal transaction
        emit BalanceWithdrawn(sender, balance);
    }

    // setGateway sets gateway for the given owner address.
    function setGateway(address owner, address gateway) external onlyMsgSender(owner) {
        (IGateway existingGateway, uint256 index) = getGateway(owner);
        if (address(existingGateway) != address(0x0)) {
            gateways[index] = Gateway(IGateway(gateway), owner);
        } else {
            gateways.push(Gateway(IGateway(gateway), owner));
        }

        emit GatewaySet(owner, gateway);
    }

    // pauseWorkflow pauses an existing active workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can pause an existing active workflow.
    function pauseWorkflow(uint256 id) external onlyMainchain onlyExistingWorkflow(id) onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow memory workflow = workflows[id];

        // Check current workflow status
        require(workflow.status == WorkflowStatus.ACTIVE, "Registry: only active workflows could be paused");

        // Update status
        workflows[id].status = WorkflowStatus.PAUSED;

        emit WorkflowPaused(id);
    }

    // resumeWorkflow resumes an existing paused workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can resume an existing active workflow.
    function resumeWorkflow(uint256 id) external onlyMainchain onlyExistingWorkflow(id) onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow memory workflow = workflows[id];

        // Check current workflow status
        require(workflow.status == WorkflowStatus.PAUSED, "Registry: only paused workflows could be resumed");

        // Update status
        workflows[id].status = WorkflowStatus.ACTIVE;

        emit WorkflowResumed(id);
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
    ) external onlyNetwork onlyExistingWorkflow(workflowId) {
        // Get a workflow by ID
        Workflow memory workflow = workflows[workflowId];

        // Make sure the workflow is not paused
        require(workflow.status == WorkflowStatus.ACTIVE, "Registry: workflow must be active");

        if (!workflow.isInternal) {
            // Make sure workflow owner has enough funds
            require(balances[workflow.owner] > 0, "Registry: not enough funds on balance");

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
            // Check the target contract is not a gateway one
            require(!_gatewayExists(target), "Registry: gateway cannot be executed");

            // Get workflow owner's gateway
            (IGateway existingGateway, ) = getGateway(workflow.owner);

            // Execute customer contract through its gateway
            success = _callWithExactGas(
                gasAmount,
                address(existingGateway),
                abi.encodeWithSignature(GATEWAY_PERFORM_FUNC_SIGNATURE, workflowId, target, data)
            );
        }
        gasUsed -= gasleft();

        // Calculate amount to charge
        uint256 amountToCharge = gasUsed;

        // Adding performance overhead if exists
        if (config.performanceOverhead > 0) {
            amountToCharge += config.performanceOverhead;
        }

        // Adding performance premium
        if (config.performancePremiumThreshold > 0) {
            amountToCharge += amountToCharge / uint256(config.performancePremiumThreshold);
        }

        // Calculate amount to charge
        amountToCharge = amountToCharge * tx.gasprice;

        if (!workflow.isInternal) {
            // Make sure owner has enough funds
            require(balances[workflow.owner] >= amountToCharge, "Registry: not enough funds on balance");

            // Charge workflow owner balance
            balances[workflow.owner] -= amountToCharge;

            // Move amount to the network rewards balance
            networkRewards += amountToCharge;
        }

        // Update total spent amount of the current workflow
        workflows[workflowId].totalSpent += amountToCharge;

        // Emit performance event
        emit Performance(workflowId, amountToCharge, success);
    }

    // getWorkflow returns the workflow by the given ID.
    // Permissions:
    //  - Anyone can read a workflow
    function getWorkflow(uint256 id) external view returns (Workflow memory workflow) {
        return workflows[id];
    }

    // getBalance returns the balance for the given address.
    function getBalance(address addr) external view returns (uint256 balance) {
        return balances[addr];
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
        (IGateway existingGateway, ) = getGateway(owner);
        require(address(existingGateway) != address(0x0), "Registry: gateway not found");

        // Check if the given sender has capacity to create one more workflow
        if (isMainChain && config.maxWorkflowsPerAccount > 0) {
            require(
                workflowsPerAddress[msg.sender] < config.maxWorkflowsPerAccount,
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
        workflows[id] = Workflow(id, owner, hash, workflowStatus, false, 0);
        workflowsPerAddress[msg.sender]++;

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
        Workflow memory workflow = workflows[id];

        // Must be PENDING
        require(workflow.status == WorkflowStatus.PENDING, "Registry: workflow must be pending");

        // Update status
        workflows[id].status = WorkflowStatus.ACTIVE;

        emit WorkflowActivated(id);
    }

    // cancelWorkflow cancels an existing workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Only workflow owner can cancel an existing active workflow on MAINCHAIN.
    //  - Only network can cancel a workflow on SIDECHAIN through the regular performance process.
    function cancelWorkflow(uint256 id) public onlyExistingWorkflow(id) onlyWorkflowOwnerOrRegistry(id) {
        // Find the workflow in the list
        Workflow memory workflow = workflows[id];

        // Check current workflow status
        require(workflow.status != WorkflowStatus.CANCELLED, "Registry: workflow is already cancelled");

        // Update status
        workflows[id].status = WorkflowStatus.CANCELLED;
        workflowsPerAddress[msg.sender]--;

        emit WorkflowCancelled(id);
    }

    // getGateway returns gateway address of the given owner.
    function getGateway(address owner) public view returns (IGateway gateway, uint256 index) {
        for (uint256 i = 0; i < gateways.length; i++) {
            if (gateways[i].owner == owner) {
                gateway = gateways[i].gateway;
                index = i;
                break;
            }
        }
        return (gateway, index);
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

    // _gatewayExists returns true if a gateway with the given address exists
    function _gatewayExists(address gateway) private view returns (bool exists) {
        for (uint256 i = 0; i < gateways.length; i++) {
            if (address(gateways[i].gateway) == gateway) {
                exists = true;
                break;
            }
        }
        return exists;
    }
}
