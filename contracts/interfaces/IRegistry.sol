// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../operational/RegistryWorkflow.sol";

/**
 * @title IRegistry
 * @notice Interface for the Registry contract, responsible for managing workflows, gateways, and configurations
 */
interface IRegistry {
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

    function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) external;

    /**
     * @notice Sets the address of the GatewayFactory contract
     * @param _newGatewayFactory The address of the new GatewayFactory contract
     */
    function setGatewayFactory(address _newGatewayFactory) external;

    /**
     * @notice Sets the address of the gateway contract for the msg.sender address
     * @param gateway The address of the gateway contract to be set
     */
    function setGateway(address gateway) external;

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
     * @notice Pauses a workflow with the given ID
     * @param id The ID of the workflow to be paused
     */
    function pauseWorkflow(uint256 id) external;

    /**
     * @notice Resumes a paused workflow with the given ID
     * @param id The ID of the workflow to be resumed
     */
    function resumeWorkflow(uint256 id) external;

    /**
     * @notice Performs an action on a workflow execution
     * @param workflowId The ID of the workflow
     * @param workflowExecutionId The ID of the specific execution within the workflow
     * @param gasAmount The gas amount for the transaction
     * @param data The transaction data
     * @param target The address of the target contract
     */
    function perform(
        uint256 workflowId,
        uint256 workflowExecutionId,
        uint256 gasAmount,
        bytes calldata data,
        address target
    ) external;

    /**
     * @notice Registers a new workflow
     * @param id The ID of the workflow
     * @param owner The address of the owner registering the workflow
     * @param hash The hash of the workflow
     * @param requireGateway The boolean indicating whether a gateway is required for the workflow
     */
    function registerWorkflow(
        uint256 id,
        address owner,
        bytes calldata hash,
        bool requireGateway
    ) external;

    /**
     * @notice Activates a registered workflow
     * @param id The ID of the workflow to be activated
     */
    function activateWorkflow(uint256 id) external;

    /**
     * @notice Cancels a registered workflow
     * @param id The ID of the workflow to be canceled
     */
    function cancelWorkflow(uint256 id) external;
}
