// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// Registry is the internal smart contract needed to secure the network and support important features of Nerif.
contract Registry {
    // MIN_REQUIRED_NODES is the minimum number of nodes needed to bootstrap the network.
    uint8 public constant MIN_REQUIRED_NODES = 4;

    // WorkflowStatus represents the workflow status.
    enum WorkflowStatus {
        ACTIVE,
        PAUSED,
        CANCELLED
    }

    // NodeRegistered is emitted when a node has been registered.
    event NodeRegistered(address node);

    // NodeApproved is emitted when a node has been approved by network participant.
    event NodeApproved(address node, address approver);

    // NodeActivated is emitted when a node has been approved by majority of the network.
    event NodeActivated(address node, address approver);

    // NodeUnregistered is emitted when a node has been approved by network participant.
    event NodeUnregistered(address node);

    // WorkflowRegistered is emitted when a workflow has been registered.
    event WorkflowRegistered(address owner, uint256 id, bytes hash, bytes signature);

    // ChangeWorkflowStatus is emitted when a workflow status has been changed.
    event ChangeWorkflowStatus(uint256 id, WorkflowStatus status);

    // Workflow represents the workflow metadata model.
    struct Workflow {
        uint64 id;
        address owner;
        bytes hash;
        bytes signature;
        WorkflowStatus status;
    }

    // _pendingNodes is the mapping between pending Nerif Network Nodes and approvers
    mapping(address => address[]) internal _pendingNodes;

    // _activeNodes is the list of active Nerif Network Nodes addresses
    address[] internal _activeNodes;

    // _workflows is the list of registered workflows
    mapping(uint64 => Workflow) internal _workflows;

    // _balances is the list of workflow owner balances;
    mapping(address => uint64) internal _balances;

    constructor(address[] initialNodes) public {
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

        // TODO: Make sure the given node has staked tokens within the staking contract

        // Add node to the pending list with 0 approvals
        _pendingNodes[node] = new address[](0);

        emit NodeRegistered(msg.sender);
    }

    // approveRegistration approves PENDING registration by the given public key.
    // Is the registration got >= 2/3 network approvals, the status gets changed to ACTIVE.
    // The transaction sender public key is used as an approver public key.
    //  - "node" is the node address (public key) to be approved.
    function approveRegistration(address node) publicKey {
        // TODO: Make sure the given node exists in the pending list
        // TODO: Make sure the current tx sender has not approved it yet
        // TODO: Add current tx sender to the approvers list
        // TODO: If the number of approvers >= 2/3 of all network participants, activate node and remove from pending

        // TODO: Node operator must cover approval costs

        // This event must be emitted only after the majority of the network has approved the registration.
        emit NodeApproved(node, msg.sender);
    }

    // unregisterNode unregisters an existing Nerif Network Node from the list of network participants.
    // The transaction sender address is used as a public key of the node.
    // Emits an event so Nerif Network will get informed about it and will exclude the node from the network.
    //  - "node" is the node address (public key), must be the same as the invoker address.
    function unregisterNode(address node) public {
        require(node == msg.sender, "Node address must be equal to tx sender address");

        // TODO: Delete node from the _activeNodes list

        emit NodeUnregistered(node);
    }

    // fundBalance funds the balance of the sender's public key with the given amount.
    function fundBalance(address workflowOwner) public payable {
        // TODO: Add workflowOwner to _balances list with the given amount
    }

    // fundBalance funds the balance of the sender's public key with the given amount.
    function getBalance(address workflowOwner) public returns (uint64 balance) {
        return _balances[workflowOwner];
    }

    // withdrawBalance withdraws the remaining balance of the sender's public key.
    function withdrawBalance(address workflowOwner) public {
        // TODO: Withdraw remaining balance of the given owner to its address
    }

    // registerWorkflow registers a new workflow metadata.
    // This request is allowed for the workflow owner only.
    //  - "id" is the workflow identifier.
    //  - "hash" is the workflow hash.
    //  - "signature" is the workflow hash signature made by workflow owner.
    // The given signature must correspond to the given hash and created by
    // the transaction sender.
    function registerWorkflow(uint256 id, address owner, bytes hash, bytes signature) public {
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

        // TODO: Delete workflow from the map

        emit ChangeWorkflowStatus(id, WorkflowStatus.CANCELLED);
    }

    // perform performs the contract execution defined in the registered workflow.
    // The function checks that the given performance transaction was signed by the majority
    // of the network so the workflow owner could be charged and the transaction
    // with the given payload could be passed to the customer's contract.
    //  - "payload" is the encoded transaction payload:
    //    TODO: Define payload structure
    function perform(uint256 id, bytes payload, bytes signature) public {
        // TODO: Make sure the workflow exists
        // TODO: Make consensus check
        // TODO: Make sure workflow owner has enough funds
        // TODO: Make sure the given transaction was not performed yet
        // TODO: Send transaction to the specified contract
        // TODO: Charge workflow owner's balance
        // TODO: Emit performance event
    }

    // consensusCheck is the public function of _consensusCheck.
    // It could be used to verify that an action during the workflow execution is non-malicious.
    function consensusCheck(bytes data, bytes signature) public returns (bool verified) {
        return _consensusCheck(data, signature);
    }

    // _consensusCheck checks that the given data was signed by majority of the network.
    function _consensusCheck(bytes data, bytes signature) returns (bool verified) {
        // TODO: Implement
        return true;
    }

    // _signatureCheck checks that the given data corresponds to the given signature
    // and was signed by the given address.
    function _signatureCheck(address signer, bytes data, bytes signature) returns (bool verified) {
        // TODO: Implement
        return true;
    }
}