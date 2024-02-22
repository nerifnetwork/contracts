// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

/**
 * @title IRegistry
 * @notice Interface for the Registry contract, responsible for managing workflows, gateways, and configurations
 */
interface IRegistry {
    /**
     * @dev Enum representing the status of a workflow
     * @param NONE The initial status
     * @param PENDING The workflow is pending activation
     * @param ACTIVE The workflow is active and operational
     * @param PAUSED The workflow is temporarily paused
     * @param CANCELLED The workflow is cancelled and no longer operational
     */
    enum WorkflowStatus {
        NONE,
        PENDING,
        ACTIVE,
        PAUSED,
        CANCELLED
    }

    /**
     * @dev Struct representing a workflow
     * @param id The unique identifier of the workflow
     * @param owner The address of the owner of the workflow
     * @param hash The hash representing the content or configuration of the workflow
     * @param status The status of the workflow (see WorkflowStatus enum)
     * @param totalSpent The total amount spent on the workflow
     */
    struct Workflow {
        uint256 id;
        address owner;
        bytes hash;
        WorkflowStatus status;
        uint256 totalSpent;
    }

    /**
     * @dev Struct containing information required to register a new workflow
     * @param id The unique identifier of the workflow
     * @param workflowOwner The address of the owner of the workflow
     * @param hash The hash representing the content or configuration of the workflow
     * @param requireGateway The flag indicating whether the workflow requires a gateway
     */
    struct RegisterWorkflowInfo {
        uint256 id;
        address workflowOwner;
        bytes hash;
        bool requireGateway;
    }

    /**
     * @notice Event emitted when a new gateway is set
     * @param owner The address of the owner setting the gateway
     * @param gateway The address of the newly set gateway
     */
    event GatewaySet(address owner, address gateway);

    /**
     * @notice Event emitted when a workflow is registered
     * @param owner The address of the owner registering the workflow
     * @param id The ID of the registered workflow
     * @param hash The hash of the registered workflow
     */
    event WorkflowRegistered(address owner, uint256 id, bytes hash);

    /**
     * @notice Event emitted when the status of a workflow is changed
     * @param id The ID of the workflow whose status changed
     * @param status The new status of the workflow
     */
    event WorkflowStatusChanged(uint256 id, WorkflowStatus status);

    /**
     * @notice Event emitted upon completion of a workflow execution
     * @param workflowId The ID of the executed workflow
     * @param workflowExecutionId The ID of the specific execution within the workflow
     * @param success The boolean indicating whether the execution was successful
     */
    event Performance(uint256 workflowId, uint256 workflowExecutionId, bool success);

    /**
     * @notice Sets the maximum number of workflows allowed per account
     * @param _newMaxWorkflowsPerAccount The new maximum number of workflows per account to be set
     */
    function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) external;

    /**
     * @notice Sets the address of the GatewayFactory contract
     * @param _newGatewayFactory The address of the new GatewayFactory contract
     */
    function setGatewayFactory(address _newGatewayFactory) external;

    /**
     * @notice Sets the address of the gateway contract for the msg.sender address
     * @param _gateway The address of the gateway contract to be set
     */
    function setGateway(address _gateway) external;

    /**
     * @notice Deploys a new gateway contract and sets it as the current gateway for the msg.sender address
     * @return The address of the newly deployed gateway contract
     */
    function deployAndSetGateway() external returns (address);

    /**
     * @notice Updates the total amount spent on a workflow execution
     * @param _workflowId The ID of the workflow
     * @param _workflowExecutionAmount The amount spent on the workflow execution
     */
    function updateWorkflowTotalSpent(uint256 _workflowId, uint256 _workflowExecutionAmount) external;

    /**
     * @notice Pauses workflows with the given IDs
     * @param _workflowIds The array of workflow IDs to be paused
     */
    function pauseWorkflows(uint256[] calldata _workflowIds) external;

    /**
     * @notice Resumes paused workflows with the given IDs
     * @param _workflowIds The array of workflow IDs to be resumed
     */
    function resumeWorkflows(uint256[] calldata _workflowIds) external;

    /**
     * @notice Performs an action on a workflow execution
     * @param _workflowId The ID of the workflow
     * @param _workflowExecutionId The ID of the specific execution within the workflow
     * @param _gasAmount The gas amount for the transaction
     * @param _data The transaction data
     * @param _target The address of the target contract
     */
    function perform(
        uint256 _workflowId,
        uint256 _workflowExecutionId,
        uint256 _gasAmount,
        bytes calldata _data,
        address _target
    ) external;

    /**
     * @notice Registers new workflows with the provided information
     * @param _registerWorkflowInfoArr The array containing information to register new workflows
     */
    function registerWorkflows(RegisterWorkflowInfo[] calldata _registerWorkflowInfoArr) external;

    /**
     * @notice Activates registered workflows with the given IDs
     * @param _workflowIds The array of workflow IDs to be activated
     */
    function activateWorkflows(uint256[] calldata _workflowIds) external;

    /**
     * @notice Cancels registered workflows with the given IDs
     * @param _workflowIds The array of workflow IDs to be canceled
     */
    function cancelWorkflows(uint256[] calldata _workflowIds) external;
}
