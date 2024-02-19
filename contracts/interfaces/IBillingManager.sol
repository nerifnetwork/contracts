// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface IBillingManager {
    enum WorkflowExecutionStatus {
        NONE,
        PENDING,
        COMPLETED
    }

    struct UserFundsData {
        uint256 userFundBalance;
        uint256 userLockedBalance;
        EnumerableSet.UintSet pendingWorkflowExecutionIds;
    }

    struct UserFundsInfo {
        uint256 userFundBalance;
        uint256 userLockedBalance;
        uint256[] pendingWorkflowExecutionIds;
    }

    struct WorkflowExecutionInfo {
        uint256 workflowId;
        uint256 executionLockedAmount;
        uint256 executionAmount;
        address workflowOwner;
        WorkflowExecutionStatus status;
    }

    event BalanceFunded(address indexed userAddr, uint256 userCurrentBalance, uint256 fundedAmount);
    event UserFundsWithdrawn(address indexed userAddr, uint256 withdrawnAmount);
    event RewardsWithdrawn(address indexed recipientAddr, uint256 rewardsAmount);
    event ExecutionFundsLocked(
        uint256 indexed workflowId,
        address indexed userAddr,
        uint256 workflowExecutionId,
        uint256 executionLockedAmount
    );
    event ExecutionCompleted(uint256 workflowExecutionId, uint256 executionAmount);
    event UnexpectedExecutionAmountFound(
        uint256 workflowExecutionId,
        uint256 executionLockedAmount,
        uint256 executionAmount
    );

    function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) external;

    function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) external;

    function fundBalance() external payable;

    function fundBalance(address _recipientAddr) external payable;

    function withdrawFunds(uint256 _amountToWithdraw) external;

    function withdrawNetworkRewards() external;

    function getWorkflowExecutionStatus(uint256 _workflowExecutionId) external view returns (WorkflowExecutionStatus);

    function getWorkflowExecutionOwner(uint256 _workflowExecutionId) external view returns (address);

    function getUserFundsInfo(address _userAddr) external view returns (UserFundsInfo memory);

    function getWorkflowExecutionInfo(uint256 _workflowExecutionId)
        external
        view
        returns (WorkflowExecutionInfo memory);

    function getUserAvailableFunds(address _userAddr) external view returns (uint256);
}
