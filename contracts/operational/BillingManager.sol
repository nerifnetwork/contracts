// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "../interfaces/IBillingManager.sol";
import "../interfaces/SignerOwnable.sol";

import "./Registry.sol";

contract BillingManager is IBillingManager, Initializable, SignerOwnable {
    using EnumerableSet for EnumerableSet.UintSet;

    Registry public registry;

    uint256 public networkRewards;
    uint256 public nextWorkflowExecutionId;

    mapping(address => UserFundsData) internal _usersFundsData;
    mapping(uint256 => WorkflowExecutionInfo) internal _workflowsExecutionInfo;

    function initialize(address _registryAddr, address _signerGetterAddress) external initializer {
        registry = Registry(_registryAddr);

        _setSignerGetter(_signerGetterAddress);
    }

    function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) external override onlySigner {
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

    function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) external override onlySigner {
        WorkflowExecutionInfo storage workflowExecutionInfo = _workflowsExecutionInfo[_workflowExecutionId];

        require(
            workflowExecutionInfo.status == WorkflowExecutionStatus.PENDING,
            "BillingManager: Not a pending workflow execution"
        );

        uint256 executionLockedAmount = workflowExecutionInfo.executionLockedAmount;
        require(executionLockedAmount >= _executionAmount, "BillingManager: Execution amount > locked amount");

        UserFundsData storage userFundsData = _usersFundsData[workflowExecutionInfo.workflowOwner];

        if (executionLockedAmount < _executionAmount) {
            _executionAmount = Math.min(_executionAmount, userFundsData.userFundBalance);

            emit UnexpectedExecutionAmountFound(_workflowExecutionId, executionLockedAmount, _executionAmount);
        }

        workflowExecutionInfo.status = WorkflowExecutionStatus.COMPLETED;
        workflowExecutionInfo.executionAmount = _executionAmount;

        userFundsData.userFundBalance -= _executionAmount;
        userFundsData.userLockedBalance -= executionLockedAmount;
        userFundsData.pendingWorkflowExecutionIds.remove(_workflowExecutionId);

        networkRewards += _executionAmount;

        registry.updateWorkflowTotalSpent(workflowExecutionInfo.workflowId, _executionAmount);

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

    function withdrawNetworkRewards() external override {
        uint256 amountToWithdraw = networkRewards;

        require(amountToWithdraw > 0, "BillingManager: No network rewards to withdraw");

        delete networkRewards;

        address signerAddr = signerGetter.getSignerAddress();

        require(signerAddr != address(0), "BillingManager: Zero signer address");

        Address.sendValue(payable(signerAddr), amountToWithdraw);

        emit RewardsWithdrawn(signerAddr, amountToWithdraw);
    }

    function getWorkflowExecutionStatus(uint256 _workflowExecutionId) external view returns (WorkflowExecutionStatus) {
        return _workflowsExecutionInfo[_workflowExecutionId].status;
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
