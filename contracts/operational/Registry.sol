// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../interfaces/SignerOwnable.sol";
import "../interfaces/IGateway.sol";
import "../interfaces/IGatewayFactory.sol";
import "../interfaces/IBillingManager.sol";
import "../interfaces/IRegistry.sol";

import "./RegistryGateway.sol";
import "./RegistryWorkflow.sol";

// Registry is the internal smart contract needed to secure the network and support important features of Nerif.
contract Registry is IRegistry, Initializable, SignerOwnable, RegistryGateway, RegistryWorkflow {
    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;
    string internal constant GATEWAY_PERFORM_FUNC_SIGNATURE = "perform(uint256,address,bytes)";

    IGatewayFactory public gatewayFactory;
    IBillingManager public billingManager;

    bool public isMainChain;
    Config public config;

    // onlyMainchain permits transactions on the mainchain only
    modifier onlyMainchain() {
        require(isMainChain, "Registry: operation is not permitted");
        _;
    }

    // onlyMsgSenderOrSigner modifier for message sender or collective address only
    modifier onlyMsgSenderOrSigner(address addr) {
        if (isMainChain) {
            require(addr == msg.sender, "Registry: operation is not permitted");
        } else {
            require(signerGetter.getSignerAddress() == msg.sender, "Registry: operation is not permitted");
        }
        _;
    }

    // onlyWorkflowOwnerOrSigner permits operation for the workflow owner if it is mainnet
    // or for the network collective address.
    modifier onlyWorkflowOwnerOrSigner(uint256 id) {
        if (isMainChain) {
            Workflow memory workflow = getWorkflow(id);
            require(workflow.owner == msg.sender, "Registry: operation is not permitted");
        } else {
            // Only network can execute the function on the sidechain.
            // The transaction must come from the network after reaching consensus.
            // Basically, the transaction must come from the registry contract itself,
            // namely from the perform function after passing all checks.
            require(signerGetter.getSignerAddress() == msg.sender, "Registry: operation is not permitted");
        }
        _;
    }

    modifier onlyBillingManager() {
        require(msg.sender == address(billingManager), "Registry: sender is not a billing manager");
        _;
    }

    function initialize(
        bool _isMainChain,
        address _signerGetterAddress,
        address _gatewayFactoryAddr,
        address _billingManagerAddr,
        Config calldata _config
    ) external initializer {
        isMainChain = _isMainChain;
        config = _config;

        gatewayFactory = IGatewayFactory(_gatewayFactoryAddr);
        billingManager = IBillingManager(_billingManagerAddr);

        _setSignerGetter(_signerGetterAddress);
    }

    // setConfig sets the given configuration
    function setConfig(Config calldata _config) external override onlySigner {
        config = _config;
    }

    function setGatewayFactory(address _newGatewayFactory) external override onlySigner {
        gatewayFactory = IGatewayFactory(_newGatewayFactory);
    }

    // setGateway sets gateway for the given owner address.
    function setGateway(address gateway) external override {
        _setGateway(msg.sender, IGateway(gateway));
        emit GatewaySet(msg.sender, gateway);
    }

    function deployAndSetGateway() external override returns (address) {
        address newGatewayAddr = gatewayFactory.deployGateway(msg.sender);

        _setGateway(msg.sender, IGateway(newGatewayAddr));
        emit GatewaySet(msg.sender, newGatewayAddr);

        return newGatewayAddr;
    }

    function updateWorkflowTotalSpent(uint256 _workflowId, uint256 _workflowExecutionAmount)
        external
        override
        onlyBillingManager
    {
        Workflow memory workflow = getWorkflow(_workflowId);

        workflow.totalSpent += _workflowExecutionAmount;
        _updateWorkflow(workflow);
    }

    // pauseWorkflow pauses an existing active workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can pause an existing active workflow.
    function pauseWorkflow(uint256 id) external override onlyMainchain onlyExistingWorkflow(id) onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Check current workflow status
        require(workflow.status == WorkflowStatus.ACTIVE, "Registry: only active workflows could be paused");

        // Update status
        workflow.status = WorkflowStatus.PAUSED;
        _updateWorkflow(workflow);

        emit WorkflowStatusChanged(id, WorkflowStatus.PAUSED);
    }

    // resumeWorkflow resumes an existing paused workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can resume an existing active workflow.
    function resumeWorkflow(uint256 id) external override onlyMainchain onlyExistingWorkflow(id) onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Check current workflow status
        require(workflow.status == WorkflowStatus.PAUSED, "Registry: only paused workflows could be resumed");

        // Update status
        workflow.status = WorkflowStatus.ACTIVE;
        _updateWorkflow(workflow);

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
        uint256 workflowExecutionId,
        uint256 gasAmount,
        bytes calldata data,
        address target
    ) external override onlySigner {
        // Get a workflow by ID
        Workflow memory workflow = getWorkflow(workflowId);

        // Make sure the workflow is not paused
        require(workflow.status == WorkflowStatus.ACTIVE, "Registry: workflow must be active");

        require(
            billingManager.getWorkflowExecutionOwner(workflowExecutionId) == workflow.owner &&
                billingManager.getWorkflowExecutionStatus(workflowExecutionId) ==
                IBillingManager.WorkflowExecutionStatus.PENDING,
            "Registry: invalid workflow execution ID"
        );

        // Cannot self-execute if not internal
        require(address(this) != target, "Registry: operation is not permitted");

        // Execute client's contract through gateway
        // Get workflow owner's gateway
        IGateway existingGateway = getGateway(workflow.owner);

        require(address(existingGateway) != address(0), "Registry: zero gateway address");

        // Execute customer contract through its gateway
        bool success = _callWithExactGas(
            gasAmount,
            address(existingGateway),
            abi.encodeWithSignature(GATEWAY_PERFORM_FUNC_SIGNATURE, workflowId, target, data)
        );

        // Emit performance event
        emit Performance(workflowId, workflowExecutionId, success);
    }

    // registerWorkflow registers a new workflow metadata.
    // Arguments:
    //  - "id" is the workflow identifier.
    //  - "owner" is the workflow owner address.
    //  - "hash" is the workflow hash.
    // The given signature must correspond to the given hash and created by the transaction sender.
    // Permissions:
    //  - Only workflow owner can register a workflow on MAINCHAIN.
    //  - Only network can register a workflow on SIDECHAIN through the regular performance process.
    function registerWorkflow(
        uint256 id,
        address owner,
        bytes calldata hash,
        bool requireGateway
    ) external override onlyMsgSenderOrSigner(owner) {
        // Check if the given workflow owner has a gateway registered.
        if (requireGateway) {
            IGateway existingGateway = getGateway(owner);
            require(address(existingGateway) != address(0x0), "Registry: gateway not found");
        }

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
        require(_addWorkflow(Workflow(id, owner, hash, workflowStatus, 0)), "Registry: failed to add workflow");

        // Emmit the event
        emit WorkflowRegistered(owner, id, hash);
    }

    // activateWorkflow updates the workflow state from PENDING to ACTIVE.
    // Arguments:
    //  - "id" is the workflow identifier.
    //  - "status" is the workflow status.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only network can execute it through the regular performance process.
    function activateWorkflow(uint256 id) external override onlyMainchain onlySigner onlyExistingWorkflow(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Must be PENDING
        require(workflow.status == WorkflowStatus.PENDING, "Registry: workflow must be pending");

        // Update status
        workflow.status = WorkflowStatus.ACTIVE;
        _updateWorkflow(workflow);

        emit WorkflowStatusChanged(id, WorkflowStatus.ACTIVE);
    }

    // cancelWorkflow cancels an existing workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Only workflow owner can cancel an existing active workflow on MAINCHAIN.
    //  - Only network can cancel a workflow on SIDECHAIN through the regular performance process.
    function cancelWorkflow(uint256 id) external override onlyExistingWorkflow(id) onlyWorkflowOwnerOrSigner(id) {
        // Find the workflow in the list
        Workflow memory workflow = getWorkflow(id);

        // Check current workflow status
        require(workflow.status != WorkflowStatus.CANCELLED, "Registry: workflow is already cancelled");

        // Update status
        workflow.status = WorkflowStatus.CANCELLED;
        _updateWorkflow(workflow);

        emit WorkflowStatusChanged(id, WorkflowStatus.CANCELLED);
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
