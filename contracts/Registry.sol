// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

// Registry is the internal smart contract needed to secure the network and support important features of Nerif.
contract Registry {
    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;

    // MIN_APPROVAL_THRESHOLD is the node approval threshold, i.e. 2/3 of the network or 66%
    uint32 public constant MIN_APPROVAL_THRESHOLD = 66;

    // MIN_REQUIRED_NODES is the minimum number of nodes needed to bootstrap the network.
    uint8 public constant MIN_REQUIRED_NODES = 4;

    // WorkflowStatus represents the workflow status.
    enum WorkflowStatus {
        ACTIVE,
        PAUSED,
        CANCELLED
    }

    // BalanceFunded is emitted when the balance of a workflow owner has been funded
    event BalanceFunded(address workflowOwner, uint256 amount);

    // BalanceWithdrawn is emitted when the balance of a workflow owner has been withdrawn
    event BalanceWithdrawn(address workflowOwner, uint256 amount);

    // NodeRegistered is emitted when a node has been registered.
    event NodeRegistered(address node);

    // NodeApproved is emitted when a node has been approved by network participant.
    event NodeApproved(address node, address approver);

    // NodeActivated is emitted when a node has been approved by majority of the network.
    event NodeActivated(address node);

    // NodeUnregistered is emitted when a node has been approved by network participant.
    event NodeUnregistered(address node);

    // WorkflowRegistered is emitted when a workflow has been registered.
    event WorkflowRegistered(address owner, uint256 id, bytes hash, bytes signature);

    // ChangeWorkflowStatus is emitted when a workflow status has been changed.
    event ChangeWorkflowStatus(uint256 id, WorkflowStatus status);

    // Performance is emitted when a client contract was executed
    event Performance(uint256 id, uint256 gasUsed);

    // Workflow represents the workflow metadata model.
    struct Workflow {
        uint256 id;
        address owner;
        bytes hash;
        bytes signature;
        WorkflowStatus status;
    }

    // PerformPayload represents the structure of perform payload
    struct PerformPayload {
        uint256 id;
        uint256 gasAmount;
        address target;
        bytes data;
    }

    // _pendingNodes is the mapping between pending Nerif Network Nodes and approvers
    mapping(address => address[]) internal _pendingNodes;

    // _activeNodes is the list of active Nerif Network Nodes addresses
    address[] internal _activeNodes;

    // _workflows is the list of registered workflows
    mapping(uint256 => Workflow) internal _workflows;

    // _balances is the list of workflow owner balances;
    mapping(address => uint256) internal _balances;

    constructor(address[] memory initialNodes) {
        require(initialNodes.length == MIN_REQUIRED_NODES, "Not enough nodes provided");

        // TODO: Check that the given tokens have staked tokens within the staking contract

        _activeNodes = initialNodes;
    }

    // registerNode registers the newjoiner's public key with the PENDING registration status.
    // The transaction sender address is used as a public key of the node.
    // This request can be approved by existing network participants using the function below.
    //  - "node" is the node address (public key), must be the same as the invoker address.
    function registerNode(address node) public {
        require(node == msg.sender, "Node address must be equal to tx sender address");

        // Make sure this node does not exist in pending nodes list
        require(_pendingNodes[node].length == 0, "Node with the given address was already registered");

        // Make sure this node does not exist in active nodes list
        bool isActivatedNode = false;
        for (uint i = 0; i < _activeNodes.length; i++) {
            if (_activeNodes[i] == node) {
                isActivatedNode = true;
                break;
            }
        }
        require(!isActivatedNode, "Node with the given address is already active");

        // TODO: Make sure the given node has staked tokens within the staking contract

        // Add node to the pending list with 0 approvals
        _pendingNodes[node] = new address[](0);

        emit NodeRegistered(msg.sender);
    }

    // approveRegistration approves PENDING registration by the given public key.
    // Is the registration got >= 2/3 network approvals, the status gets changed to ACTIVE.
    // The transaction sender public key is used as an approver public key.
    //  - "node" is the node address (public key) to be approved.
    function approveRegistration(address node) public {
        // Make sure the given sender is an active node
        bool isActivatedNode = false;
        for (uint i = 0; i < _activeNodes.length; i++) {
            if (_activeNodes[i] == msg.sender) {
                isActivatedNode = true;
                break;
            }
        }
        require(isActivatedNode, "Operation is not permitted");

        // Make sure the given node exists in the pending list
        require(_pendingNodes[node].length > 0, "Node with the given address does not exist");

        // Make sure the current tx sender has not approved it yet
        bool alreadyApproved = false;
        for (uint i = 0; i < _pendingNodes[node].length; i++) {
            if (_pendingNodes[node][i] == msg.sender) {
                alreadyApproved = true;
                break;
            }
        }
        require(!alreadyApproved, "Node has been already approved by tx sender");

        // Add current tx sender to the approvers list
        _pendingNodes[node].push(msg.sender);
        emit NodeApproved(node, msg.sender);

        // If the number of approvers >= 2/3 of all network participants, activate node and remove from pending
        uint32 currentPercentage = uint32(100) * uint32(_pendingNodes[node].length) / uint32(_activeNodes.length);
        if (currentPercentage >= MIN_APPROVAL_THRESHOLD) {
            // Delete from pending nodes list
            delete _pendingNodes[node];

            // Add to active nodes list
            _activeNodes.push(node);

            // This event must be emitted only after the majority of the network has approved the registration.
            emit NodeActivated(node);
        }

        // TODO: Node operator must cover approval costs
    }

    // unregisterNode unregisters an existing Nerif Network Node from the list of network participants.
    // The transaction sender address is used as a public key of the node.
    // Emits an event so Nerif Network will get informed about it and will exclude the node from the network.
    //  - "node" is the node address (public key), must be the same as the invoker address.
    function unregisterNode(address node) public {
        require(node == msg.sender, "Node address must be equal to tx sender address");

        // Unregister node
        bool isUnregistered = _unregisterNode(node);
        require(isUnregistered, "Node with the given address does not exist");

        emit NodeUnregistered(node);
    }

    // fundBalance funds the balance of the sender's public key with the given amount.
    function fundBalance(address workflowOwner) public payable {
        require(msg.sender == workflowOwner, "Operation is not permitted");

        // Update the balance value
        _balances[msg.sender] += msg.value;

        emit BalanceFunded(workflowOwner, msg.value);
    }

    // fundBalance funds the balance of the sender's public key with the given amount.
    function getBalance(address workflowOwner) view public returns (uint256 balance) {
        return _balances[workflowOwner];
    }

    // withdrawBalance withdraws the remaining balance of the sender's public key.
    // TODO: Handle cases when the withdrawal happens during the workflow execution.
    function withdrawBalance(address workflowOwner) public {
        require(msg.sender == workflowOwner, "Operation is not permitted");

        address payable sender = payable(msg.sender);
        uint256 balance = _balances[sender];

        // Ensure the sender has a positive balance
        require(balance > 0, "No balance to withdraw");

        // Update the sender's balance
        _balances[sender] = 0;

        // Transfer the balance to the sender
        sender.transfer(balance);

        // Emit an event to log the withdrawal transaction
        emit BalanceWithdrawn(sender, balance);
    }

    // registerWorkflow registers a new workflow metadata.
    // This request is allowed for the workflow owner only.
    //  - "id" is the workflow identifier.
    //  - "hash" is the workflow hash.
    //  - "signature" is the workflow hash signature made by workflow owner.
    // The given signature must correspond to the given hash and created by
    // the transaction sender.
    function registerWorkflow(uint256 id, address owner, bytes calldata hash, bytes calldata signature) public {
        require(owner == msg.sender, "Workflow owner must be equal to tx sender address");

        // Check the given signature
        require(_signatureCheck(owner, hash, signature), "Signature has not been verified");

        // Store workflow with ACTIVE status
        _workflows[id] = Workflow(id, owner, hash, signature, WorkflowStatus.ACTIVE);

        emit WorkflowRegistered(msg.sender, id, hash, signature);
    }

    // pauseWorkflow pauses an existing active workflow.
    // This operation is allowed for workflow owner ONLY.
    //  - "id" is the workflow identifier.
    function pauseWorkflow(uint256 id) public {
        // Find the workflow in the list
        Workflow storage workflow = _workflows[id];
        require(workflow.id != 0, "Workflow with the given ID does not exist");

        // Check workflow ownership
        require(workflow.owner == msg.sender, "Operation is not permitted");

        // Check current workflow status
        require(workflow.status == WorkflowStatus.ACTIVE, "Only active workflows could be paused");

        // Update status
        _workflows[id].status = WorkflowStatus.PAUSED;

        emit ChangeWorkflowStatus(id, WorkflowStatus.PAUSED);
    }

    // resumeWorkflow resumes an existing paused workflow.
    // This operation is allowed for workflow owner ONLY.
    //  - "id" is the workflow identifier.
    function resumeWorkflow(uint256 id) public {
        // Find the workflow in the list
        Workflow storage workflow = _workflows[id];
        require(workflow.id != 0, "Workflow with the given ID does not exist");

        // Check workflow ownership
        require(workflow.owner == msg.sender, "Operation is not permitted");

        // Check current workflow status
        require(workflow.status == WorkflowStatus.PAUSED, "Only paused workflows could be resumed");

        // Update status
        _workflows[id].status = WorkflowStatus.ACTIVE;

        emit ChangeWorkflowStatus(id, WorkflowStatus.ACTIVE);
    }

    // cancelWorkflow cancels an existing workflow.
    // This operation is allowed for workflow owner ONLY.
    //  - "id" is the workflow identifier.
    function cancelWorkflow(uint256 id) public {
        // Find the workflow in the list
        Workflow storage workflow = _workflows[id];
        require(workflow.id != 0, "Workflow with the given ID does not exist");

        // Check workflow ownership
        require(workflow.owner == msg.sender, "Operation is not permitted");

        // Delete workflow from the map
        delete _workflows[id];

        emit ChangeWorkflowStatus(id, WorkflowStatus.CANCELLED);
    }

    // perform performs the contract execution defined in the registered workflow.
    // The function checks that the given performance transaction was signed by the majority
    // of the network so the workflow owner could be charged and the transaction
    // with the given payload could be passed to the customer's contract.
    //  - "workflowId" is the workflow ID
    //  - "gasAmount" is the maximum number of gas used to execute the transaction
    //  - "data" is the contract call data
    //  - "target" is the client contract address
    //  - "signature" is the payload signature
    function perform(uint256 workflowId, uint256 gasAmount, bytes memory data, address target, bytes memory signature) public {
        // Make sure the tx sender is in the active nodes list
        bool isActivatedNode = false;
        for (uint i = 0; i < _activeNodes.length; i++) {
            if (_activeNodes[i] == msg.sender) {
                isActivatedNode = true;
                break;
            }
        }
        require(isActivatedNode, "Operation is not permitted");

        // Make sure the given payload was signed by the network
        bytes memory payload = abi.encode(workflowId, gasAmount, data, target);
        require(_consensusCheck(payload, signature), "Consensus check failed");

        // Get a workflow by ID
        Workflow storage workflow = _workflows[workflowId];

        // Make sure workflow exists
        require(workflow.id > 0, "Workflow does not exist");

        // Make sure workflow owner has enough funds
        require(_balances[workflow.owner] > 0, "Not enough funds on balance");

        // Execute client's contract
        uint256 gasUsed = gasleft();
        bool success = _callWithExactGas(gasAmount, target, data);
        gasUsed -= gasleft();

        // Charge workflow owner _balances
        _balances[workflow.owner] -= gasUsed;

        // TODO: Make sure the given transaction was not performed yet

        // Emit performance event
        emit Performance(workflowId, gasUsed);
    }

    // consensusCheck is the public function of _consensusCheck.
    // It could be used to verify that an action during the workflow execution is non-malicious.
    function consensusCheck(bytes memory data, bytes memory signature) public returns (bool verified) {
        return _consensusCheck(data, signature);
    }

    // _consensusCheck checks that the given data was signed by majority of the network.
    function _consensusCheck(bytes memory data, bytes memory signature) internal returns (bool verified) {
        // TODO: Implement
        return true;
    }

    // _signatureCheck checks that the given data corresponds to the given signature
    // and was signed by the given address.
    function _signatureCheck(address signer, bytes memory data, bytes memory signature) internal returns (bool verified) {
        // TODO: Implement
        return true;
    }

    // _unregisterNode deletes the element from _activeNodes by the given value
    function _unregisterNode(address node) internal returns (bool) {
        // Delete the given node from the active nodes list
        for (uint i = 0; i < _activeNodes.length; i++) {
            if (_activeNodes[i] == node) {
                _activeNodes[i] = _activeNodes[_activeNodes.length - 1];
                _activeNodes.pop();
                return true;
            }
        }

        // Delete the given node from the pending nodes list
        if (_pendingNodes[node].length > 0) {
            delete _pendingNodes[node];
            return true;
        }

        return false;
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