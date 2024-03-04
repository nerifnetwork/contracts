// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import "@openzeppelin/contracts/interfaces/IERC20.sol";

import "@solarity/solidity-lib/libs/arrays/Paginator.sol";
import "@solarity/solidity-lib/libs/data-structures/StringSet.sol";

import "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import "../interfaces/IBillingManager.sol";
import "../interfaces/SignerOwnable.sol";

import "./Registry.sol";

contract BillingManager is IBillingManager, Initializable, SignerOwnable, UUPSUpgradeable {
    using EnumerableSet for *;
    using StringSet for StringSet.Set;
    using SafeERC20 for IERC20;

    Registry public registry;

    string public nativeDepositAssetKey;
    uint256 public nextWorkflowExecutionId;

    EnumerableSet.AddressSet internal _existingUsers;
    StringSet.Set internal _supportedDepositAssetKeys;

    mapping(string => DepositAssetData) internal _depositAssetsData;
    mapping(address => UserData) internal _usersData;

    mapping(uint256 => WorkflowExecutionInfo) internal _workflowsExecutionInfo;

    modifier onlyExistingDepositAsset(string memory _depositAssetKey) {
        _onlyExistingDepositAsset(_depositAssetKey);
        _;
    }

    modifier onlyEnabledDepositAsset(string memory _depositAssetKey) {
        _onlyEnabledDepositAsset(_depositAssetKey);
        _;
    }

    receive() external payable onlyEnabledDepositAsset(nativeDepositAssetKey) {
        _depositAsset(nativeDepositAssetKey, msg.sender, msg.sender, msg.value, true);
    }

    function initialize(
        address _registryAddr,
        address _signerGetterAddress,
        DepositAssetInfo calldata _nativeDepositAssetInfo
    ) external initializer {
        registry = Registry(_registryAddr);

        nativeDepositAssetKey = _nativeDepositAssetInfo.depositAssetKey;
        _addDepositAsset(_nativeDepositAssetInfo);

        _setSignerGetter(_signerGetterAddress);
    }

    function addDepositAssets(DepositAssetInfo[] calldata _depositAssetInfoArr) external onlySigner {
        for (uint256 i = 0; i < _depositAssetInfoArr.length; i++) {
            _addDepositAsset(_depositAssetInfoArr[i]);
        }
    }

    function updateWorkflowExecutionDiscount(
        string calldata _depositAssetKey,
        uint256 _newWorkflowExecutionDiscount
    ) external override onlyExistingDepositAsset(_depositAssetKey) onlySigner {
        uint256 currentWorkflowExecutionDiscount = _depositAssetsData[_depositAssetKey].workflowExecutionDiscount;

        _depositAssetsData[_depositAssetKey].workflowExecutionDiscount = _newWorkflowExecutionDiscount;

        emit WorkflowExecutionDiscountUpdated(
            _depositAssetKey,
            _newWorkflowExecutionDiscount,
            currentWorkflowExecutionDiscount
        );
    }

    function updateDepositAssetEnabledStatus(
        string calldata _depositAssetKey,
        bool _newEnabledStatus
    ) external override onlyExistingDepositAsset(_depositAssetKey) onlySigner {
        require(
            _depositAssetsData[_depositAssetKey].isEnabled != _newEnabledStatus,
            "BillingManager: Invalid new enabled status"
        );

        _depositAssetsData[_depositAssetKey].isEnabled = _newEnabledStatus;

        emit DepositAssetEnabledStatusUpdated(_depositAssetKey, _newEnabledStatus);
    }

    function lockExecutionFunds(
        string calldata _depositAssetKey,
        uint256 _workflowId,
        uint256 _executionLockedAmount
    ) external override onlySigner onlyExistingDepositAsset(_depositAssetKey) {
        require(registry.isWorkflowExist(_workflowId), "BillingManager: Workflow does not exist");

        address workflowOwner = registry.getWorkflowOwner(_workflowId);

        require(
            getUserAvailableFunds(workflowOwner, _depositAssetKey) >= _executionLockedAmount,
            "BillingManager: Not enough available funds to lock"
        );

        uint256 currentWorkflowExecutionId = nextWorkflowExecutionId++;

        UserDepositData storage userDepositData = _usersData[workflowOwner].userDepositsData[_depositAssetKey];

        userDepositData.userLockedAmount += _executionLockedAmount;
        userDepositData.pendingWorkflowExecutionIds.add(currentWorkflowExecutionId);

        _workflowsExecutionInfo[currentWorkflowExecutionId] = WorkflowExecutionInfo(
            _depositAssetKey,
            _workflowId,
            _executionLockedAmount,
            0,
            workflowOwner,
            WorkflowExecutionStatus.PENDING
        );

        emit ExecutionFundsLocked(
            _depositAssetKey,
            _workflowId,
            workflowOwner,
            currentWorkflowExecutionId,
            _executionLockedAmount
        );
    }

    function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) external override onlySigner {
        WorkflowExecutionInfo storage workflowExecutionInfo = _workflowsExecutionInfo[_workflowExecutionId];

        require(
            workflowExecutionInfo.status == WorkflowExecutionStatus.PENDING,
            "BillingManager: Not a pending workflow execution"
        );

        string memory depositAssetKey = workflowExecutionInfo.depositAssetKey;
        UserDepositData storage userDepositData = _usersData[workflowExecutionInfo.workflowOwner].userDepositsData[
            depositAssetKey
        ];

        uint256 executionLockedAmount = workflowExecutionInfo.executionLockedAmount;

        if (executionLockedAmount < _executionAmount) {
            _executionAmount = Math.min(
                _executionAmount,
                userDepositData.userDepositedAmount - userDepositData.userLockedAmount + executionLockedAmount
            );

            emit UnexpectedExecutionAmountFound(_workflowExecutionId, executionLockedAmount, _executionAmount);
        }

        workflowExecutionInfo.status = WorkflowExecutionStatus.COMPLETED;
        workflowExecutionInfo.executionAmount = _executionAmount;

        userDepositData.userDepositedAmount -= _executionAmount;
        userDepositData.userLockedAmount -= executionLockedAmount;
        userDepositData.pendingWorkflowExecutionIds.remove(_workflowExecutionId);

        _depositAssetsData[depositAssetKey].networkRewards += _executionAmount;

        registry.updateWorkflowTotalSpent(depositAssetKey, workflowExecutionInfo.workflowId, _executionAmount);

        emit ExecutionCompleted(_workflowExecutionId, _executionAmount);
    }

    function deposit(
        string memory _depositAssetKey,
        address _recipientAddr,
        uint256 _depositAmount
    ) external payable override onlyExistingDepositAsset(_depositAssetKey) onlyEnabledDepositAsset(_depositAssetKey) {
        bool isNative = isNativeDepositAsset(_depositAssetKey);

        if (isNative) {
            require(_depositAmount == msg.value, "BillingManager: Invalid msg.value");
        } else {
            require(msg.value == 0, "BillingManager: Not a native deposit asset");
        }

        _depositAsset(_depositAssetKey, msg.sender, _recipientAddr, _depositAmount, isNative);
    }

    function depositWithPermit(
        string memory _depositAssetKey,
        address _recipientAddr,
        uint256 _depositAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external override onlyExistingDepositAsset(_depositAssetKey) onlyEnabledDepositAsset(_depositAssetKey) {
        require(
            !isNativeDepositAsset(_depositAssetKey),
            "BillingManager: Unable to deposit native currency with permit"
        );

        require(isDepositAssetPermitable(_depositAssetKey), "BillingManager: Deposit asset is not permitable");

        IERC20Permit(_depositAssetsData[_depositAssetKey].tokenAddr).permit(
            msg.sender,
            address(this),
            _depositAmount,
            _sigExpirationTime,
            _v,
            _r,
            _s
        );

        _depositAsset(_depositAssetKey, msg.sender, _recipientAddr, _depositAmount, false);
    }

    function withdrawFunds(
        string memory _depositAssetKey,
        uint256 _amountToWithdraw
    ) external override onlyExistingDepositAsset(_depositAssetKey) {
        require(_amountToWithdraw > 0, "BillingManager: Zero amount to withdraw");
        require(
            getUserAvailableFunds(msg.sender, _depositAssetKey) >= _amountToWithdraw,
            "BillingManager: Not enough available funds to withdraw"
        );

        UserDepositData storage userDepositData = _usersData[msg.sender].userDepositsData[_depositAssetKey];

        uint256 newUserFundsAmount = userDepositData.userDepositedAmount - _amountToWithdraw;

        userDepositData.userDepositedAmount = newUserFundsAmount;

        if (newUserFundsAmount == 0) {
            _existingUsers.remove(msg.sender);
        }

        _sendFunds(_depositAssetKey, msg.sender, _amountToWithdraw);

        emit UserFundsWithdrawn(msg.sender, _amountToWithdraw);
    }

    function withdrawNetworkRewards(
        string memory _depositAssetKey
    ) external override onlyExistingDepositAsset(_depositAssetKey) {
        DepositAssetData storage depositAssetData = _depositAssetsData[_depositAssetKey];

        uint256 amountToWithdraw = depositAssetData.networkRewards;

        require(amountToWithdraw > 0, "BillingManager: No network rewards to withdraw");

        delete depositAssetData.networkRewards;

        address signerAddr = signerGetter.getSignerAddress();

        require(signerAddr != address(0), "BillingManager: Zero signer address");

        _sendFunds(_depositAssetKey, signerAddr, amountToWithdraw);

        emit RewardsWithdrawn(signerAddr, amountToWithdraw);
    }

    function getTotalUsersCount() external view override returns (uint256) {
        return _existingUsers.length();
    }

    function getSupportedDepositAssetKeys() external view returns (string[] memory) {
        return _supportedDepositAssetKeys.values();
    }

    function getDepositAssetsInfo(
        string[] calldata _depositAssetKeysArr
    ) external view returns (DepositAssetInfo[] memory _depositAssetsInfoArr) {
        _depositAssetsInfoArr = new DepositAssetInfo[](_depositAssetKeysArr.length);

        for (uint256 i = 0; i < _depositAssetKeysArr.length; i++) {
            _depositAssetsInfoArr[i] = DepositAssetInfo(
                _depositAssetKeysArr[i],
                _depositAssetsData[_depositAssetKeysArr[i]]
            );
        }
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

    function getUsersDepositInfo(
        string memory _depositAssetKey,
        uint256 _offset,
        uint256 _limit
    ) external view override returns (UserDepositInfo[] memory _usersInfoArr) {
        uint256 to = Paginator.getTo(_existingUsers.length(), _offset, _limit);

        _usersInfoArr = new UserDepositInfo[](to - _offset);

        for (uint256 i = _offset; i < to; i++) {
            _usersInfoArr[i - _offset] = getUserDepositInfo(_existingUsers.at(i), _depositAssetKey);
        }
    }

    function getWorkflowExecutionInfo(
        uint256 _workflowExecutionId
    ) external view override returns (WorkflowExecutionInfo memory) {
        return _workflowsExecutionInfo[_workflowExecutionId];
    }

    function getUserDepositAssetKeys(address _userAddr) external view returns (string[] memory) {
        return _usersData[_userAddr].depositAssetKeys.values();
    }

    function getNetworkRewards(string memory _depositAssetKey) external view returns (uint256) {
        return _depositAssetsData[_depositAssetKey].networkRewards;
    }

    function getUserDepositInfo(
        address _userAddr,
        string memory _depositAssetKey
    ) public view override returns (UserDepositInfo memory) {
        UserDepositData storage userDepositData = _usersData[_userAddr].userDepositsData[_depositAssetKey];

        return
            UserDepositInfo(
                _userAddr,
                userDepositData.userDepositedAmount,
                userDepositData.userLockedAmount,
                userDepositData.pendingWorkflowExecutionIds.values()
            );
    }

    function getUserAvailableFunds(
        address _userAddr,
        string memory _depositAssetKey
    ) public view override returns (uint256) {
        UserDepositData storage userDepositData = _usersData[_userAddr].userDepositsData[_depositAssetKey];

        return userDepositData.userDepositedAmount - userDepositData.userLockedAmount;
    }

    function isDepositAssetSupported(string memory _depositAssetKey) public view override returns (bool) {
        return _supportedDepositAssetKeys.contains(_depositAssetKey);
    }

    function isNativeDepositAsset(string memory _depositAssetKey) public view override returns (bool) {
        return Strings.equal(_depositAssetKey, nativeDepositAssetKey);
    }

    function isDepositAssetEnabled(string memory _depositAssetKey) public view override returns (bool) {
        return _depositAssetsData[_depositAssetKey].isEnabled;
    }

    function isDepositAssetPermitable(string memory _depositAssetKey) public view override returns (bool) {
        return _depositAssetsData[_depositAssetKey].isPermitable;
    }

    function _authorizeUpgrade(address) internal virtual override onlySigner {}

    function _addDepositAsset(DepositAssetInfo memory _depositAssetInfo) internal {
        require(
            !isDepositAssetSupported(_depositAssetInfo.depositAssetKey),
            "BillingManager: Deposit asset already exists"
        );
        require(
            _depositAssetInfo.depositAssetData.tokenAddr != address(0) ||
                Strings.equal(_depositAssetInfo.depositAssetKey, nativeDepositAssetKey),
            "BillingManager: Invalid deposit asset token address"
        );
        require(bytes(_depositAssetInfo.depositAssetKey).length > 0, "BillingManager: Invalid deposit asset key");
        require(
            _depositAssetInfo.depositAssetData.networkRewards == 0,
            "BillingManager: Invalid init network rewards value"
        );

        _depositAssetsData[_depositAssetInfo.depositAssetKey] = _depositAssetInfo.depositAssetData;
        _supportedDepositAssetKeys.add(_depositAssetInfo.depositAssetKey);

        emit DepositAssetAdded(_depositAssetInfo.depositAssetKey, _depositAssetInfo.depositAssetData);
    }

    function _depositAsset(
        string memory _depositAssetKey,
        address _tokensSenderAddr,
        address _recipientAddr,
        uint256 _depositAmount,
        bool _isNative
    ) internal {
        require(_depositAmount > 0, "BillingManager: Zero deposit amount");

        if (!_isNative) {
            IERC20(_depositAssetsData[_depositAssetKey].tokenAddr).safeTransferFrom(
                _tokensSenderAddr,
                address(this),
                _depositAmount
            );
        }

        UserData storage userData = _usersData[_recipientAddr];

        userData.userDepositsData[_depositAssetKey].userDepositedAmount += _depositAmount;
        userData.depositAssetKeys.add(_depositAssetKey);

        _existingUsers.add(_recipientAddr);

        emit AssetDeposited(_depositAssetKey, _tokensSenderAddr, _recipientAddr, _depositAmount);
    }

    function _sendFunds(string memory _depositAssetKey, address _recipientAddr, uint256 _amountToSend) internal {
        if (isNativeDepositAsset(_depositAssetKey)) {
            Address.sendValue(payable(_recipientAddr), _amountToSend);
        } else {
            IERC20(_depositAssetsData[_depositAssetKey].tokenAddr).safeTransfer(_recipientAddr, _amountToSend);
        }
    }

    function _onlyExistingDepositAsset(string memory _depositAssetKey) internal view {
        require(isDepositAssetSupported(_depositAssetKey), "BillingManager: Deposit asset does not exist");
    }

    function _onlyEnabledDepositAsset(string memory _depositAssetKey) internal view {
        require(isDepositAssetEnabled(_depositAssetKey), "BillingManager: Deposit asset is disabled");
    }
}
