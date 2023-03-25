// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// Registry is the internal smart contract needed to secure the network and support important features of Nerif.
contract Registry {
    // NodeStatus represents the node status.
    enum NodeStatus {
        PENDING,
        ACTIVE,
        UNREGISTERED
    }

    // WorkflowStatus represents the workflow status.
    enum WorkflowStatus {
        ACTIVE,
        PAUSED,
        CANCELLED
    }

    // NodeRegistered is emitted when a node has been registered.
    event NodeRegistered(address addr);

    // ChangeNodeStatus is emitted when a node status has been changed.
    event ChangeNodeStatus(address addr, NodeStatus status);

    // WorkflowRegistered is emitted when a workflow has been registered.
    event WorkflowRegistered(address owner, uint256 id, bytes hash, bytes signature);

    // ChangeWorkflowStatus is emitted when a workflow status has been changed.
    event ChangeWorkflowStatus(uint256 id, WorkflowStatus status);

    // registerNode registers the newjoiner's public key with the PENDING registration status.
    // The transaction sender address is used as a public key of the node.
    // This request can be approved by existing network participants using the function below.
    function registerNode() public {
        // TODO: Implement
        emit NodeRegistered(msg.sender);
    }

    // approveRegistration approves PENDING registration by the given public key.
    // Is the registration got >= 2/3 network approvals, the status gets changed to ACTIVE.
    // The transaction sender public key is used as an approver public key.
    function approveRegistration(bytes publicKey) publicKey {
        // TODO: Implement

        // This event must be emitted only after the majority of the network has approved the registration.
        emit ChangeNodeStatus(msg.sender, NodeStatus.ACTIVE);
    }

    // unregisterNode unregisters an existing Nerif Network Node from the list of network participants.
    // The transaction sender address is used as a public key of the node.
    // Emits an event so Nerif Network will get informed about it and will exclude the node from the network.
    function unregisterNode() public {
        // TODO: Implement
        emit ChangeNodeStatus(msg.sender, NodeStatus.UNREGISTERED);
    }

    // fundBalance funds the balance of the sender's public key with the given amount.
    function fundBalance() public {
        // TODO: Implement
    }

    // withdrawBalance withdraws the remaining balance of the sender's public key.
    function withdrawBalance() public{
        // TODO: Implement
    }

    // registerWorkflow registers a new workflow metadata.
    // This request is allowed for the workflow owner only.
    //  - "id" is the workflow identifier.
    //  - "hash" is the workflow hash.
    //  - "signature" is the workflow hash signature made by workflow owner.
    // The given signature must correspond to the given hash and created by
    // the transaction sender.
    function registerWorkflow(uint256 id, bytes hash, bytes signature) public {
        // TODO: Implement
        emit WorkflowRegistered(msg.sender, id, hash, signature);
    }

    // pauseWorkflow pauses an existing active workflow.
    // This operation is allowed for workflow owner ONLY.
    //  - "id" is the workflow identifier.
    function pauseWorkflow(uint256 id) public {
        // TODO: Implement
        emit ChangeWorkflowStatus(id, WorkflowStatus.PAUSED);
    }

    // resumeWorkflow resumes an existing paused workflow.
    // This operation is allowed for workflow owner ONLY.
    //  - "id" is the workflow identifier.
    function resumeWorkflow(uint256 id) public {
        // TODO: Implement
        emit ChangeWorkflowStatus(id, WorkflowStatus.ACTIVE);
    }

    // cancelWorkflow cancels an existing workflow.
    // This operation is allowed for workflow owner ONLY.
    //  - "id" is the workflow identifier.
    function cancelWorkflow(uint256 id) public {
        // TODO: Implement
        emit ChangeWorkflowStatus(id, WorkflowStatus.CANCELLED);
    }

    // perform performs the contract execution defined in the registered workflow.
    // The function checks that the given performance transaction was signed by the majority
    // of the network so the workflow owner could be charged and the transaction
    // with the given payload could be passed to the customer's contract.
    //  - "payload" is the encoded transaction payload.
    function perform(bytes payload) public {
        // TODO: Implement
    }
}