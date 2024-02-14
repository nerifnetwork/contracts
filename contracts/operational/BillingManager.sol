// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "../interfaces/IBillingManager.sol";

import "./Registry.sol";

contract BillingManager is IBillingManager, OwnableUpgradeable {
    using EnumerableSet for EnumerableSet.UintSet;

    Registry public registry;

    uint256 public networkRewards;
    uint256 public nextWorkflowExecutionId;

    mapping(address => UserFundsData) internal _usersFundsData;
    mapping(uint256 => WorkflowExecutionInfo) internal _workflowsExecutionInfo;

    function initialize(address _registryAddr) external initializer {
        __Ownable_init();

        registry = Registry(_registryAddr);
    }

    function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) external override onlyOwner {
        address workflowOwner = registry.getWorkflowOwner(_workflowId);

        UserFundsData storage userFundsData = _usersFundsData[workflowOwner];

        require(
            getUserAvailableFunds(workflowOwner) >= _executionLockedAmount,
            "BillingManager: Not enough funds to lock"
        );

        uint256 currentWorkflowExecutionId = nextWorkflowExecutionId++;

        userFundsData.userLockedBalance += _executionLockedAmount;
        userFundsData.pendingWorkflowExecutionIds.add(currentWorkflowExecutionId);

        _workflowsExecutionInfo[currentWorkflowExecutionId] = WorkflowExecutionInfo(
            _workflowId,
            _executionLockedAmount,
            0,
            workflowOwner,
            WorkflowExecutionStatus.PENDING
        );

        emit ExecutionFundsLocked(_workflowId, workflowOwner, currentWorkflowExecutionId, _executionLockedAmount);
    }

    function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) external override onlyOwner {
        WorkflowExecutionInfo storage workflowExecutionInfo = _workflowsExecutionInfo[_workflowExecutionId];

        require(
            workflowExecutionInfo.status == WorkflowExecutionStatus.PENDING,
            "BillingManager: Not a pending workflow execution"
        );
        require(
            workflowExecutionInfo.executionLockedAmount >= _executionAmount,
            "BillingManager: Execution amount > locked amount"
        );

        // TODO: What should be done if _executionAmount > _executionLockedAmount?

        workflowExecutionInfo.status = WorkflowExecutionStatus.COMPLETED;
        workflowExecutionInfo.executionAmount = _executionAmount;

        UserFundsData storage userFundsData = _usersFundsData[workflowExecutionInfo.workflowOwner];

        userFundsData.userFundBalance -= _executionAmount;
        userFundsData.userLockedBalance -= workflowExecutionInfo.executionLockedAmount;
        userFundsData.pendingWorkflowExecutionIds.remove(_workflowExecutionId);

        networkRewards += _executionAmount;

        emit ExecutionCompleted(_workflowExecutionId, _executionAmount);
    }

    function fundBalance() external payable override {
        _fundBalance(msg.sender);
    }

    function fundBalance(address _recipientAddr) external payable override {
        _fundBalance(_recipientAddr);
    }

    function withdrawFunds(uint256 _amountToWithdraw) external override {
        require(
            getUserAvailableFunds(msg.sender) >= _amountToWithdraw,
            "BillingManager: Not enough available funds to withdraw"
        );

        _usersFundsData[msg.sender].userFundBalance -= _amountToWithdraw;

        Address.sendValue(payable(msg.sender), _amountToWithdraw);

        emit UserFundsWithdrawn(msg.sender, _amountToWithdraw);
    }

    function withdrawNetworkRewards() external override onlyOwner {
        uint256 amountToWithdraw = networkRewards;

        require(amountToWithdraw > 0, "BillingManager: No network rewards to withdraw");

        delete networkRewards;

        Address.sendValue(payable(msg.sender), amountToWithdraw);

        emit RewardsWithdrawn(msg.sender, amountToWithdraw);
    }

    function getUserAvailableFunds(address _userAddr) public view override returns (uint256) {
        return _usersFundsData[_userAddr].userFundBalance - _usersFundsData[_userAddr].userLockedBalance;
    }

    function _fundBalance(address _recipientAddr) internal {
        require(msg.value > 0, "BillingManager: Zero funds to add");

        uint256 newUserBalance = _usersFundsData[_recipientAddr].userFundBalance + msg.value;

        _usersFundsData[_recipientAddr].userFundBalance = newUserBalance;

        emit BalanceFunded(_recipientAddr, newUserBalance, msg.value);
    }
}
