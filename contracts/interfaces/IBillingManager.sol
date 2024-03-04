// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@solarity/solidity-lib/libs/data-structures/StringSet.sol";
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

    struct DepositAssetData {
        address tokenAddr;
        uint256 workflowExecutionDiscount;
        uint256 networkRewards;
        bool isPermitable;
        bool isEnabled;
    }

    struct DepositAssetInfo {
        string depositAssetKey;
        DepositAssetData depositAssetData;
    }

    struct UserData {
        StringSet.Set depositAssetKeys;
        mapping(string => UserDepositData) userDepositsData;
    }

    struct UserDepositData {
        uint256 userDepositedAmount;
        uint256 userLockedAmount;
        EnumerableSet.UintSet pendingWorkflowExecutionIds;
    }

    /**
     * @notice Struct containing information about a user's funds
     * @param userAddr The address of the user
     * @param pendingWorkflowExecutionIds The array of IDs of pending workflow executions
     */
    struct UserDepositInfo {
        address userAddr;
        uint256 userDepositedAmount;
        uint256 userLockedAmount;
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
        string depositAssetKey;
        uint256 workflowId;
        uint256 executionLockedAmount;
        uint256 executionAmount;
        address workflowOwner;
        WorkflowExecutionStatus status;
    }

    event DepositAssetAdded(string depositAssetKey, DepositAssetData depositAssetData);
    event WorkflowExecutionDiscountUpdated(
        string indexed depositAssetKey,
        uint256 newWorkflowExecutionDiscount,
        uint256 prevWorkflowExecutionDiscount
    );
    event DepositAssetEnabledStatusUpdated(string indexed depositAssetKey, bool newEnabledStatus);
    event AssetDeposited(
        string indexed depositAssetKey,
        address tokensSender,
        address depositRecipient,
        uint256 depositAmount
    );

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
        string indexed depositAssetKey,
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

    function addDepositAssets(DepositAssetInfo[] memory _depositAssetInfoArr) external;

    function updateWorkflowExecutionDiscount(
        string memory _depositAssetKey,
        uint256 _newWorkflowExecutionDiscount
    ) external;

    function updateDepositAssetEnabledStatus(string memory _depositAssetKey, bool _newEnabledStatus) external;

    /**
     * @notice Locks funds for a workflow execution
     * @param _workflowId The ID of the associated workflow
     * @param _executionLockedAmount The amount of funds to be locked for the execution
     */
    function lockExecutionFunds(
        string memory _depositAssetKey,
        uint256 _workflowId,
        uint256 _executionLockedAmount
    ) external;

    /**
     * @notice Completes a workflow execution
     * @param _workflowExecutionId The ID of the workflow execution to be completed
     * @param _executionAmount The amount transferred for the execution
     */
    function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) external;

    function deposit(string memory _depositAssetKey, address _recipientAddr, uint256 _depositAmount) external payable;

    function depositWithPermit(
        string memory _depositAssetKey,
        address _recipientAddr,
        uint256 _depositAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external;

    /**
     * @notice Withdraws funds from the user's balance
     * @param _amountToWithdraw The amount to be withdrawn
     */
    function withdrawFunds(string memory _depositAssetKey, uint256 _amountToWithdraw) external;

    /**
     * @notice Withdraws network rewards
     */
    function withdrawNetworkRewards(string memory _depositAssetKey) external;

    /**
     * @notice Retrieves the total count of registered users
     * @return uint256 The total count of registered users
     */
    function getTotalUsersCount() external view returns (uint256);

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
     * @notice Retrieves an array of existing user addresses within the specified range
     * @param _offset The starting index of users to retrieve
     * @param _limit The maximum number of users to retrieve
     * @return An array containing user addresses
     */
    function getExistingUsers(uint256 _offset, uint256 _limit) external view returns (address[] memory);

    /**
     * @notice Retrieves an array of user funds information within the specified range
     * @param _offset The starting index of user funds information to retrieve
     * @param _limit The maximum number of user funds information to retrieve
     * @return _usersInfoArr An array containing information about user funds
     */
    function getUsersDepositInfo(
        string memory _depositAssetKey,
        uint256 _offset,
        uint256 _limit
    ) external view returns (UserDepositInfo[] memory _usersInfoArr);

    /**
     * @notice Retrieves information about a workflow execution
     * @param _workflowExecutionId The ID of the workflow execution
     * @return Information about the workflow execution
     */
    function getWorkflowExecutionInfo(
        uint256 _workflowExecutionId
    ) external view returns (WorkflowExecutionInfo memory);

    /**
     * @notice Retrieves the funds information of a user
     * @param _userAddr The address of the user
     * @return Information about the user's funds
     */
    function getUserDepositInfo(
        address _userAddr,
        string memory _depositAssetKey
    ) external view returns (UserDepositInfo memory);

    /**
     * @notice Retrieves the available funds of a user
     * @param _userAddr The address of the user
     * @return The available funds of the user
     */
    function getUserAvailableFunds(address _userAddr, string memory _depositAssetKey) external view returns (uint256);
}
