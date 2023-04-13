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
        PENDING,
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

    // NodeRegistrationCancelled is emitted when a node registration has been cancelled.
    event NodeRegistrationCancelled(address node);

    // WorkflowRegistered is emitted when a workflow has been registered.
    event WorkflowRegistered(address owner, uint256 id, bytes hash, bytes signature);

    // WorkflowActivated is emitted when a workflow status has been changed from PENDING to ACTIVE.
    // This event gets emitted when the workflow has been successfully registered on all supported networks.
    event WorkflowActivated(uint256 id);

    // WorkflowPaused is emitted when a workflow status has been changed to PAUSED.
    event WorkflowPaused(uint256 id);

    // WorkflowResumed is emitted when a workflow status has been changed from PAUSED to ACTIVE.
    event WorkflowResumed(uint256 id);

    // WorkflowCancelled is emitted when a workflow status has been cancelled.
    event WorkflowCancelled(uint256 id);

    // Performance is emitted when a client contract was executed
    event Performance(uint256 id, uint256 gasUsed, bool success);

    // Config contains the configuration options
    struct Config {
        // performanceOverhead is the cost of the performance transaction excluding the client contract call.
        // Numbers are different depending on the chain.
        // Metrics:
        //  - Ethereum Goerli execution: 49,022
        //  - BSC Testnet execution:     46,922
        //  - Polygon Mumbai execution:  31,476
        uint256 performanceOverhead;

        // performancePremiumThreshold is the network premium threshold in percents.
        // Numbers are different depending on the chain.
        // Metrics:
        //  - Ethereum Goerli execution: 10%
        //  - BSC Testnet execution:     10%
        //  - Polygon Mumbai execution:  10%
        uint8 performancePremiumThreshold;

        // registrationOverhead is the cost of the workflow registration.
        // Numbers are different depending on the chain.
        // Metrics:
        //  - BSC Testnet cost:     115,520 (registration, sidechain)
        //  - Ethereum Goerli cost: 122,420 (registration, sidechain)
        //  - Polygon Mumbai cost:  67,282 (approval, mainchain)
        uint256 registrationOverhead;

        // cancellationOverhead is the cost of the workflow cancellation.
        // Numbers are different depending on the chain.
        // Metrics:
        //  - BSC Testnet cost:     48,568
        //  - Ethereum Goerli cost: 50,168
        //  - Polygon Mumbai cost:
        uint256 cancellationOverhead;
    }

    // Workflow represents the workflow metadata model.
    struct Workflow {
        uint256 id;
        address owner;
        bytes hash;
        bytes signature;
        WorkflowStatus status;
        bool isInternal;
        uint256 totalSpent;
    }

    // PerformPayload represents the structure of perform payload
    struct PerformPayload {
        uint256 id;
        uint256 gasAmount;
        address target;
        bytes data;
    }

    // config contains system configuration
    Config public config;

    // _pendingNodes is the mapping between pending Nerif Network Nodes and approvers
    // This is used on the mainchain only.
    mapping(address => address[]) internal _pendingNodes;

    // _activeNodes is the list of active Nerif Network Nodes addresses
    address[] internal _activeNodes;

    // _workflows is the list of registered workflows
    mapping(uint256 => Workflow) internal _workflows;

    // _balances is the list of workflow owner balances;
    mapping(address => uint256) internal _balances;

    // _isMainChain is the indicator whether the contract is deployed on the main chain
    bool internal _isMainChain;

    constructor(
        address[] memory initialNodes,
        bool isMainChain
    ) {
        require(initialNodes.length == MIN_REQUIRED_NODES, "Not enough nodes provided");

        // TODO: Check that the given tokens have staked tokens within the staking contract

        _activeNodes = initialNodes;
        _isMainChain = isMainChain;

        // Define internal workflows
        // TODO: Implement more elegant way

        // Register an activated node on the sidechain
        Workflow memory nodeActivationWorkflow = Workflow(97772660256052158796229912779610582517, address(0), "", "", WorkflowStatus.ACTIVE, true, 0);
        _workflows[nodeActivationWorkflow.id] = nodeActivationWorkflow;

        // Unregister an active node on the sidechain
        Workflow memory nodeDeactivationWorkflow = Workflow(117764555324547669208370722903305523582, address(0), "", "", WorkflowStatus.ACTIVE, true, 0);
        _workflows[nodeDeactivationWorkflow.id] = nodeDeactivationWorkflow;

        // Activate workflow on the mainchain side
        Workflow memory workflowCreationWorkflow = Workflow(40505927788353901442144037336646356013, address(0), "", "", WorkflowStatus.ACTIVE, true, 0);
        _workflows[workflowCreationWorkflow.id] = workflowCreationWorkflow;

        // Cancel workflow on the sidechain side
        Workflow memory workflowCancellationWorkflow = Workflow(219775546284901721155783592958414245131, address(0), "", "", WorkflowStatus.ACTIVE, true, 0);
        _workflows[workflowCancellationWorkflow.id] = workflowCancellationWorkflow;
    }

    // setConfig sets the given configuration
    // TODO: Only the network could update the config. Must pass consensus.
    function setConfig(
        Config calldata _config
    ) public onlyNodeOperator {
        config = _config;
    }

    function isMainChain() public view returns (bool) {
        return _isMainChain;
    }

    function getActiveNodes() public view returns (address[] memory) {
        return _activeNodes;
    }

    // registerNode registers the newjoiner with the PENDING registration status.
    // Meaning this function is allowed for active node operators only.
    // This request can be approved by existing network participants using the function below.
    //  - "node" is the newjoiner node address (public key).
    // Permissions:
    //  - Anyone can request to join the network.
    //  - Only network can execute this function on SIDECHAIN.
    function registerNode(
        address node
    ) public onlyMsgSenderOrNetwork(node) {
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

        // Is it is sidechain, it is assumed the Nerif Network got an approval from the mainchain
        // and sent a sync transaction to the sidechain so the node has to be immediately activated
        if (_isMainChain) {
            // Add node to the pending list with 0 approvals
            // FIXME: The first element is the zero address, should be fixed
            _pendingNodes[node] = new address[](1);
            _pendingNodes[node][0] = msg.sender;

            emit NodeRegistered(node);
        } else {
            // Activate the node
            _activateNode(node);
        }
    }

    // approveRegistration approves PENDING registration by the given public key.
    // Is the registration got >= 2/3 network approvals, the status gets changed to ACTIVE.
    // Arguments.
    //  - "node" is the node address (public key) to be approved.
    // Permissions:
    //  - Only active node operators can execute this function
    //  - Not allowed on sidechain
    function approveRegistration(
        address node
    ) public onlyNodeOperator onlyMainchain {
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

            // Activate the node
            _activateNode(node);
        }

        // TODO: Node operator must cover approval costs
    }

    // _activateNode activates the given node and emitts the event
    function _activateNode(address node) internal {
        // Make sure the node is not activated yet
        // Make sure the current tx sender has not approved it yet
        bool alreadyActive = false;
        for (uint i = 0; i < _activeNodes.length; i++) {
            if (_activeNodes[i] == node) {
                alreadyActive = true;
                break;
            }
        }
        require(!alreadyActive, "Node is already active");

        // Add node to the active nodes list
        _activeNodes.push(node);

        emit NodeActivated(node);
    }

    // unregisterNode unregisters an existing Nerif Network Node from the active nodes list.
    // Emits an event so Nerif Network will get informed about it and will exclude the node from the network.
    // Arguments:
    //  - "node" is the node address (public key) to unregister.
    // Permissions:
    //  - Only node can unregister itself
    //  - Only network can unregister the node on SIDECHAIN
    function unregisterNode(
        address node
    ) public onlyMsgSenderOrNetwork(node) {
        bool found = false;
        for (uint i = 0; i < _activeNodes.length; i++) {
            if (_activeNodes[i] == node) {
                _activeNodes[i] = _activeNodes[_activeNodes.length - 1];
                _activeNodes.pop();
                found = true;
                break;
            }
        }

        require(found, "Node not found");

        emit NodeUnregistered(node);
    }

    // cancelNodeRegistration cancels node registration - deletes the given node from the pending nodes list.
    // Arguments:
    //  - "node" is the node address (public key).
    // Permissions:
    //  - Only node can unregister itself
    //  - Only network can unregister the node on SIDECHAIN
    function cancelNodeRegistration(
        address node
    ) public onlyMainchain onlyMsgSender(node) {
        require(_pendingNodes[node].length > 0, "Node does not exist in the pending nodes list");

        delete _pendingNodes[node];

        emit NodeRegistrationCancelled(node);
    }

    // fundBalance funds the balance of the sender's public key with the given amount.
    // Permissions:
    //  - Anyone can fund balance.
    function fundBalance(
        address addr
    ) public payable onlyMsgSender(addr) {
        // Update the balance value
        _balances[msg.sender] += msg.value;

        emit BalanceFunded(addr, msg.value);
    }

    // getBalance returns the balance for the given address.
    // Permissions:
    //  - Anyone can get a balance.
    function getBalance(
        address addr
    ) view public returns (uint256 balance) {
        return _balances[addr];
    }

    // withdrawBalance withdraws the remaining balance of the sender's public key.
    // Permissions:
    //  - Only balance owner can withdraw its balance.
    // TODO: Handle cases when the withdrawal happens during the workflow execution.
    function withdrawBalance(
        address addr
    ) public onlyMsgSender(addr) {
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
    // Arguments:
    //  - "id" is the workflow identifier.
    //  - "owner" is the workflow owner address.
    //  - "hash" is the workflow hash.
    //  - "signature" is the workflow hash signature made by workflow owner.
    // The given signature must correspond to the given hash and created by the transaction sender.
    // Permissions:
    //  - Only workflow owner can register a workflow on MAINCHAIN.
    //  - Only network can register a workflow on SIDECHAIN.
    function registerWorkflow(
        uint256 id,
        address owner,
        bytes calldata hash,
        bytes calldata signature
    ) public onlyMsgSenderOrNetwork(owner) {
        // Check the given signature
        require(_signatureCheck(owner, hash, signature), "Signature has not been verified");

        // Select the proper state
        WorkflowStatus workflowStatus = WorkflowStatus.ACTIVE;
        if (_isMainChain) {
            // It will be active after the workflow got registered on all supported networks.
            workflowStatus = WorkflowStatus.PENDING;
        }

        // Store workflow with ACTIVE status
        _workflows[id] = Workflow(id, owner, hash, signature, workflowStatus, false, 0);

        emit WorkflowRegistered(msg.sender, id, hash, signature);
    }

    // activateWorkflow updates the workflow state from PENDING to ACTIVE.
    // Arguments:
    //  - "id" is the workflow identifier.
    //  - "status" is the workflow status.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only network can execute it.
    function activateWorkflow(
        uint256 id
    ) public onlyMainchain onlyNetwork {
        // Find the workflow in the list
        Workflow storage workflow = _workflows[id];

        // Must be PENDING
        require(workflow.status == WorkflowStatus.PENDING, "Workflow must be pending");

        // Update status
        _workflows[id].status = WorkflowStatus.ACTIVE;

        emit WorkflowActivated(id);
    }

    // pauseWorkflow pauses an existing active workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can pause an existing active workflow.
    function pauseWorkflow(
        uint256 id
    ) public onlyMainchain onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow storage workflow = _workflows[id];

        // Check current workflow status
        require(workflow.status == WorkflowStatus.ACTIVE, "Only active workflows could be paused");

        // Update status
        _workflows[id].status = WorkflowStatus.PAUSED;

        emit WorkflowPaused(id);
    }

    // resumeWorkflow resumes an existing paused workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Permitted on MAINCHAIN only.
    //  - Only workflow owner can resume an existing active workflow.
    function resumeWorkflow(
        uint256 id
    ) public onlyMainchain onlyWorkflowOwner(id) {
        // Find the workflow in the list
        Workflow storage workflow = _workflows[id];

        // Check current workflow status
        require(workflow.status == WorkflowStatus.PAUSED, "Only paused workflows could be resumed");

        // Update status
        _workflows[id].status = WorkflowStatus.ACTIVE;

        emit WorkflowResumed(id);
    }

    // cancelWorkflow cancels an existing workflow.
    // Arguments:
    //  - "id" is the workflow identifier.
    // Permissions:
    //  - Only workflow owner can cancel an existing active workflow on MAINCHAIN.
    //  - Only network can cancel a workflow on SIDECHAIN.
    function cancelWorkflow(
        uint256 id
    ) public onlyWorkflowOwnerOrNetwork(id) {
        // Find the workflow in the list
        Workflow storage workflow = _workflows[id];

        // Check current workflow status
        require(workflow.status != WorkflowStatus.CANCELLED, "Workflow is already cancelled");

        // Update status
        _workflows[id].status = WorkflowStatus.CANCELLED;

        emit WorkflowCancelled(id);
    }

    // getWorkflow returns the workflow by the given ID.
    // Permissions:
    //  - Anyone can read a workflow
    function getWorkflow(
        uint256 id
    ) view public returns (Workflow memory workflow) {
        return _workflows[id];
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
    //  - "signature" is the payload signature
    // Permissions:
    //  - Only network can execute this function
    function perform(
        uint256 workflowId,
        uint256 gasAmount,
        bytes memory data,
        address target,
        bytes memory signature
    ) public onlyNodeOperator {
        // Make sure the given payload was signed by the network
        bytes memory payload = abi.encode(workflowId, gasAmount, data, target);
        require(_consensusCheck(payload, signature), "Consensus check failed");

        // Get a workflow by ID
        Workflow storage workflow = _workflows[workflowId];

        // Make sure workflow exists
        require(workflow.id > 0, "Workflow does not exist");

        // Make sure the workflow is not paused
        require(workflow.status == WorkflowStatus.ACTIVE, "Workflow must be active");

        if (!workflow.isInternal) {
            // Make sure workflow owner has enough funds
            require(_balances[workflow.owner] > 0, "Not enough funds on balance");

            // Cannot self-execute if not internal
            require(address(this) != target, "Operation is not permitted");
        }

        // TODO: Make sure the given transaction was not performed yet

        // Execute client's contract
        uint256 gasUsed = gasleft();
        bool success = _callWithExactGas(gasAmount, target, data);
        gasUsed -= gasleft();

        // Calculate amount to charge
        uint256 amountToCharge = gasUsed;
        if (config.performanceOverhead > 0) {
            amountToCharge += config.performanceOverhead;
        }
        if (config.performancePremiumThreshold > 0) {
            amountToCharge += amountToCharge / uint256(config.performancePremiumThreshold);
        }

        amountToCharge = amountToCharge * tx.gasprice;

        if (!workflow.isInternal) {
            // Charge workflow owner balance
            _balances[workflow.owner] -= amountToCharge;

            // Add balance to the executer's address
            _balances[msg.sender] += amountToCharge;

            // Update total spent amount of the current workflow
            _workflows[workflowId].totalSpent += amountToCharge;
        }

        // Emit performance event
        emit Performance(workflowId, amountToCharge, success);
    }

    // consensusCheck is the public function of _consensusCheck.
    // It could be used to verify that an action during the workflow execution is non-malicious.
    // Permissions:
    // - Anyone can do a consensus check
    function consensusCheck(
        bytes memory data,
        bytes memory signature
    ) public view returns (bool verified) {
        return _consensusCheck(data, signature);
    }

    // _consensusCheck checks that the given data was signed by majority of the network.
    function _consensusCheck(
        bytes memory data,
        bytes memory signature
    ) internal view returns (bool verified) {
        // TODO: Implement
        return true;
    }

    // _signatureCheck checks that the given data corresponds to the given signature
    // and was signed by the given address.
    function _signatureCheck(
        address signer,
        bytes memory data,
        bytes memory signature
    ) internal view returns (bool verified) {
        // TODO: Implement
        return true;
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

    modifier onlyWorkflowOwnerOrNetwork(uint256 id) {
        if (_isMainChain) {
            Workflow storage workflow = _workflows[id];
            require(workflow.owner == msg.sender, "Operation is not permitted");
            _;
        } else {
            // Only network can execute the function on the sidechain.
            // The transaction must come from the network after reaching consensus.
            // Basically, the transaction must come from the registry contract itself,
            // namely from the perform function after passing all checks.
            require(address(this) == msg.sender, "Operation is not permitted");
            _;
        }
    }

    modifier onlyMsgSenderOrNetwork(address addr) {
        if (_isMainChain) {
            require(addr == msg.sender, "Operation is not permitted");
            _;
        } else {
            // Only network can execute the function on the sidechain.
            // The transaction must come from the network after reaching consensus.
            // Basically, the transaction must come from the registry contract itself,
            // namely from the perform function after passing all checks.
            require(address(this) == msg.sender, "Operation is not permitted");
            _;
        }
    }

    modifier onlyNetwork {
        // The transaction must come from the network after reaching consensus.
        // Basically, the transaction must come from the registry contract itself,
        // namely from the perform function after passing all checks.
        require(address(this) == msg.sender, "Operation is not permitted");
        _;
    }

    // onlyNodeOperator allows only active nodes to execute the transaction
    modifier onlyNodeOperator {
        bool found = false;
        for (uint i = 0; i < _activeNodes.length; i++) {
            if (_activeNodes[i] == msg.sender) {
                found = true;
                break;
            }
        }
        require(found, "Operation is not permitted");
        _;
    }

    // onlyWorkflowOwner checks that the given tx sender is an actual workflow owner
    modifier onlyWorkflowOwner(uint256 id) {
        Workflow storage workflow = _workflows[id];
        require(workflow.owner == msg.sender, "Operation is not permitted");
        _;
    }

    // onlyMsgSender checks that the given address is the message sender one
    modifier onlyMsgSender(address addr) {
        require(addr == msg.sender, "Operation is not permitted");
        _;
    }

    modifier onlyMainchain {
        require(_isMainChain, "Operation is not permitted");
        _;
    }
}