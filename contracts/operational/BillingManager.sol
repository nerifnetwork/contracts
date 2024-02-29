// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import "../interfaces/IBillingManager.sol";
import "../interfaces/SignerOwnable.sol";

import "./Registry.sol";

contract BillingManager is IBillingManager, Initializable, SignerOwnable, UUPSUpgradeable {
    using EnumerableSet for *;

    Registry public registry;

    uint256 public networkRewards;
    uint256 public nextWorkflowExecutionId;

    EnumerableSet.AddressSet internal _existingUsers;

    mapping(address => UserFundsData) internal _usersFundsData;
    mapping(uint256 => WorkflowExecutionInfo) internal _workflowsExecutionInfo;

    receive() external payable {
        _fundBalance(msg.sender);
    }

    function initialize(address _registryAddr, address _signerGetterAddress) external initializer {
        registry = Registry(_registryAddr);

        _setSignerGetter(_signerGetterAddress);
    }

    function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) external override onlySigner {
        require(registry.isWorkflowExist(_workflowId), "BillingManager: Workflow does not exist");

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

        UserFundsData storage userFundsData = _usersFundsData[workflowExecutionInfo.workflowOwner];
        uint256 executionLockedAmount = workflowExecutionInfo.executionLockedAmount;

        if (executionLockedAmount < _executionAmount) {
            _executionAmount = Math.min(
                _executionAmount,
                userFundsData.userFundBalance - userFundsData.userLockedBalance + executionLockedAmount
            );

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

        uint256 newUserFundsAmount = _usersFundsData[msg.sender].userFundBalance - _amountToWithdraw;

        _usersFundsData[msg.sender].userFundBalance = newUserFundsAmount;

        if (newUserFundsAmount == 0) {
            _existingUsers.remove(msg.sender);
        }

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

    function getTotalUsersCount() external view override returns (uint256) {
        return _existingUsers.length();
    }

    function getWorkflowExecutionStatus(
        uint256 _workflowExecutionId
    ) external view override returns (WorkflowExecutionStatus) {
        return _workflowsExecutionInfo[_workflowExecutionId].status;
    }

    function getWorkflowExecutionOwner(uint256 _workflowExecutionId) external view override returns (address) {
        return _workflowsExecutionInfo[_workflowExecutionId].workflowOwner;
    }

    function getExecutionWorkflowId(uint256 _workflowExecutionId) external view override returns (uint256) {
        return _workflowsExecutionInfo[_workflowExecutionId].workflowId;
    }

    function getExistingUsers(uint256 _offset, uint256 _limit) external view override returns (address[] memory) {
        return Paginator.part(_existingUsers, _offset, _limit);
    }

    function getUsersFundsInfo(
        uint256 _offset,
        uint256 _limit
    ) external view override returns (UserFundsInfo[] memory _usersInfoArr) {
        uint256 to = Paginator.getTo(_existingUsers.length(), _offset, _limit);

        _usersInfoArr = new UserFundsInfo[](to - _offset);

        for (uint256 i = _offset; i < to; i++) {
            _usersInfoArr[i - _offset] = getUserFundsInfo(_existingUsers.at(i));
        }
    }

    function getWorkflowExecutionInfo(
        uint256 _workflowExecutionId
    ) external view override returns (WorkflowExecutionInfo memory) {
        return _workflowsExecutionInfo[_workflowExecutionId];
    }

    function getUserFundsInfo(address _userAddr) public view override returns (UserFundsInfo memory) {
        UserFundsData storage userFundsData = _usersFundsData[_userAddr];

        return
            UserFundsInfo(
                _userAddr,
                userFundsData.userFundBalance,
                userFundsData.userLockedBalance,
                userFundsData.pendingWorkflowExecutionIds.values()
            );
    }

    function getUserAvailableFunds(address _userAddr) public view override returns (uint256) {
        return _usersFundsData[_userAddr].userFundBalance - _usersFundsData[_userAddr].userLockedBalance;
    }

    function _authorizeUpgrade(address) internal virtual override onlySigner {}

    function _fundBalance(address _recipientAddr) internal {
        require(msg.value > 0, "BillingManager: Zero funds to add");

        uint256 newUserBalance = _usersFundsData[_recipientAddr].userFundBalance + msg.value;

        _usersFundsData[_recipientAddr].userFundBalance = newUserBalance;
        _existingUsers.add(_recipientAddr);

        emit BalanceFunded(_recipientAddr, newUserBalance, msg.value);
    }
}
