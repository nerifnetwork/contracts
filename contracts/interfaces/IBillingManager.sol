// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title IBillingManager
 * @notice Interface for the Billing Manager contract, responsible for managing user funds,
 * workflow executions funds locking, and network rewards
 */
interface IBillingManager {
    /**
     * @notice Enum representing the status of a workflow execution
     * @param NONE The default status indicating no execution status
     * @param PENDING The status indicating the execution is pending
     * @param COMPLETED The status indicating the execution is completed
     */
    enum WorkflowExecutionStatus {
        NONE,
        PENDING,
        COMPLETED
    }

    /**
     * @notice Struct containing data representing a user's funds
     * @param userFundBalance The total balance of the user's funds
     * @param userLockedBalance The amount of the user's funds locked in pending workflow executions
     * @param pendingWorkflowExecutionIds The set of IDs of pending workflow executions
     */
    struct UserFundsData {
        uint256 userFundBalance;
        uint256 userLockedBalance;
        EnumerableSet.UintSet pendingWorkflowExecutionIds;
    }

    /**
     * @notice Struct containing information about a user's funds
     * @param userFundBalance The total balance of the user's funds
     * @param userLockedBalance The amount of the user's funds locked in pending workflow executions
     * @param pendingWorkflowExecutionIds The array of IDs of pending workflow executions
     */
    struct UserFundsInfo {
        uint256 userFundBalance;
        uint256 userLockedBalance;
        uint256[] pendingWorkflowExecutionIds;
    }

    /**
     * @notice Struct containing information about a workflow execution
     * @param workflowId The ID of the associated workflow
     * @param executionLockedAmount The amount of funds locked for the execution
     * @param executionAmount The amount of funds transferred for the execution
     * @param workflowOwner The address of the owner of the workflow
     * @param status The status of the workflow execution
     */
    struct WorkflowExecutionInfo {
        uint256 workflowId;
        uint256 executionLockedAmount;
        uint256 executionAmount;
        address workflowOwner;
        WorkflowExecutionStatus status;
    }

    /**
     * @notice Event emitted when a user's balance is funded
     * @param userAddr The address of the user
     * @param userCurrentBalance The current balance of the user
     * @param fundedAmount The amount funded to the user's balance
     */
    event BalanceFunded(address indexed userAddr, uint256 userCurrentBalance, uint256 fundedAmount);

    /**
     * @notice Event emitted when funds are withdrawn from a user's balance
     * @param userAddr The address of the user
     * @param withdrawnAmount The amount withdrawn from the user's balance
     */
    event UserFundsWithdrawn(address indexed userAddr, uint256 withdrawnAmount);

    /**
     * @notice Event emitted when network rewards are withdrawn
     * @param recipientAddr The address of the recipient
     * @param rewardsAmount The amount of rewards withdrawn
     */
    event RewardsWithdrawn(address indexed recipientAddr, uint256 rewardsAmount);

    /**
     * @notice Event emitted when funds are locked for a workflow execution
     * @param workflowId The ID of the associated workflow
     * @param userAddr The address of the user initiating the execution
     * @param workflowExecutionId The ID of the workflow execution
     * @param executionLockedAmount The amount of funds locked for the execution
     */
    event ExecutionFundsLocked(
        uint256 indexed workflowId,
        address indexed userAddr,
        uint256 workflowExecutionId,
        uint256 executionLockedAmount
    );

    /**
     * @notice Event emitted when a workflow execution is completed
     * @param workflowExecutionId The ID of the workflow execution
     * @param executionAmount The amount transferred for the execution
     */
    event ExecutionCompleted(uint256 workflowExecutionId, uint256 executionAmount);

    /**
     * @notice Event emitted when an unexpected execution amount is found
     * @param workflowExecutionId The ID of the workflow execution
     * @param executionLockedAmount The amount of funds locked for the execution
     * @param executionAmount The amount transferred for the execution
     */
    event UnexpectedExecutionAmountFound(
        uint256 workflowExecutionId,
        uint256 executionLockedAmount,
        uint256 executionAmount
    );

    /**
     * @notice Locks funds for a workflow execution
     * @param _workflowId The ID of the associated workflow
     * @param _executionLockedAmount The amount of funds to be locked for the execution
     */
    function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) external;

    /**
     * @notice Completes a workflow execution
     * @param _workflowExecutionId The ID of the workflow execution to be completed
     * @param _executionAmount The amount transferred for the execution
     */
    function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) external;

    /**
     * @notice Funds the msg.sender balance
     */
    function fundBalance() external payable;

    /**
     * @notice Funds the contract balance and assigns the funds to a specified recipient
     * @param _recipientAddr The address of the recipient
     */
    function fundBalance(address _recipientAddr) external payable;

    /**
     * @notice Withdraws funds from the user's balance
     * @param _amountToWithdraw The amount to be withdrawn
     */
    function withdrawFunds(uint256 _amountToWithdraw) external;

    /**
     * @notice Withdraws network rewards
     */
    function withdrawNetworkRewards() external;

    /**
     * @notice Retrieves the status of a workflow execution
     * @param _workflowExecutionId The ID of the workflow execution
     * @return The status of the workflow execution
     */
    function getWorkflowExecutionStatus(uint256 _workflowExecutionId) external view returns (WorkflowExecutionStatus);

    /**
     * @notice Retrieves the owner of a workflow execution
     * @param _workflowExecutionId The ID of the workflow execution
     * @return The address of the owner of the workflow execution
     */
    function getWorkflowExecutionOwner(uint256 _workflowExecutionId) external view returns (address);

    /**
     * @notice Retrieves the workflow ID associated with the given execution ID
     * @param _workflowExecutionId The execution ID for which to retrieve the associated workflow ID
     * @return The workflow ID associated with the given execution ID
     */
    function getExecutionWorkflowId(uint256 _workflowExecutionId) external view returns (uint256);

    /**
     * @notice Retrieves the funds information of a user
     * @param _userAddr The address of the user
     * @return Information about the user's funds
     */
    function getUserFundsInfo(address _userAddr) external view returns (UserFundsInfo memory);

    /**
     * @notice Retrieves information about a workflow execution
     * @param _workflowExecutionId The ID of the workflow execution
     * @return Information about the workflow execution
     */
    function getWorkflowExecutionInfo(
        uint256 _workflowExecutionId
    ) external view returns (WorkflowExecutionInfo memory);

    /**
     * @notice Retrieves the available funds of a user
     * @param _userAddr The address of the user
     * @return The available funds of the user
     */
    function getUserAvailableFunds(address _userAddr) external view returns (uint256);
}
