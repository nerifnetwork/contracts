// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import "@openzeppelin/contracts/interfaces/IERC20.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/libs/arrays/Paginator.sol";
import "@solarity/solidity-lib/libs/data-structures/StringSet.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/operational/IBillingManager.sol";

import "./Registry.sol";

contract BillingManager is IBillingManager, AbstractDependant, EIP712Upgradeable {
    using EnumerableSet for *;
    using StringSet for StringSet.Set;
    using SafeERC20 for IERC20;

    // solhint-disable-next-line var-name-mixedcase
    bytes32 private constant _WITHDRAW_TYPEHASH =
        keccak256(
            "Withdraw(address userAddr,string depositAssetKey,uint256 withdrawAmount,uint256 nonce,uint256 deadline)"
        );

    IContractsRegistry internal _contractsRegistry;
    Registry internal _registry;

    string public nativeDepositAssetKey;

    EnumerableSet.AddressSet internal _existingUsers;
    StringSet.Set internal _supportedDepositAssetKeys;

    mapping(string => DepositAssetData) internal _depositAssetsData;
    mapping(address => UserData) internal _usersData;

    modifier onlySigner() {
        _onlySigner(msg.sender);
        _;
    }

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

    function initialize(DepositAssetInfo calldata _nativeDepositAssetInfo) external initializer {
        __EIP712_init("BillingManager", "1");

        nativeDepositAssetKey = _nativeDepositAssetInfo.depositAssetKey;

        _addDepositAsset(_nativeDepositAssetInfo);
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _contractsRegistry = contractsRegistry;
        _registry = Registry(contractsRegistry.getRegistryContract());
    }

    // solhint-disable-next-line ordering
    function addDepositAssets(DepositAssetInfo[] calldata _depositAssetInfoArr) external onlySigner {
        for (uint256 i = 0; i < _depositAssetInfoArr.length; i++) {
            _addDepositAsset(_depositAssetInfoArr[i]);
        }
    }

    function updateWorkflowExecutionDiscount(
        string calldata _depositAssetKey,
        uint256 _newWorkflowExecutionDiscount
    ) external override onlySigner onlyExistingDepositAsset(_depositAssetKey) {
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
    ) external override onlySigner onlyExistingDepositAsset(_depositAssetKey) {
        require(
            _depositAssetsData[_depositAssetKey].isEnabled != _newEnabledStatus,
            "BillingManager: Invalid new enabled status"
        );

        _depositAssetsData[_depositAssetKey].isEnabled = _newEnabledStatus;

        emit DepositAssetEnabledStatusUpdated(_depositAssetKey, _newEnabledStatus);
    }

    function networkWithdraw(NetworkWithdrawInfo[] calldata _networkWithdrawArr) external override onlySigner {
        for (uint256 i = 0; i < _networkWithdrawArr.length; i++) {
            NetworkWithdrawInfo calldata withdrawInfo = _networkWithdrawArr[i];

            _onlyExistingDepositAsset(withdrawInfo.depositAssetKey);
            _onlyEnoughFunds(withdrawInfo.depositAssetKey, withdrawInfo.userAddr, withdrawInfo.amountToWithdraw);

            _updateUserDepositData(
                withdrawInfo.depositAssetKey,
                withdrawInfo.userAddr,
                withdrawInfo.amountToWithdraw,
                false
            );

            _depositAssetsData[withdrawInfo.depositAssetKey].networkRewards += withdrawInfo.amountToWithdraw;

            emit NetworkWithdrawCompleted(
                withdrawInfo.depositAssetKey,
                withdrawInfo.userAddr,
                withdrawInfo.amountToWithdraw
            );
        }
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
        uint256 _amountToWithdraw,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external override {
        require(block.timestamp <= _sigExpirationTime, "BillingManager: Expired deadline");

        uint256 currentNonce = _usersData[msg.sender].withdrawNonce++;

        bytes32 withdrawStructHash = keccak256(
            abi.encode(
                _WITHDRAW_TYPEHASH,
                msg.sender,
                keccak256(bytes(_depositAssetKey)),
                _amountToWithdraw,
                currentNonce,
                _sigExpirationTime
            )
        );

        _onlySigner(ECDSA.recover(_hashTypedDataV4(withdrawStructHash), _v, _r, _s));

        _onlyEnoughFunds(_depositAssetKey, msg.sender, _amountToWithdraw);
        _updateUserDepositData(_depositAssetKey, msg.sender, _amountToWithdraw, false);
        _sendFunds(_depositAssetKey, msg.sender, _amountToWithdraw);

        emit UserFundsWithdrawn(_depositAssetKey, msg.sender, _amountToWithdraw);
    }

    function withdrawNetworkRewards(
        string memory _depositAssetKey
    ) external override onlyExistingDepositAsset(_depositAssetKey) {
        DepositAssetData storage depositAssetData = _depositAssetsData[_depositAssetKey];

        uint256 amountToWithdraw = depositAssetData.networkRewards;

        require(amountToWithdraw > 0, "BillingManager: No network rewards to withdraw");

        delete depositAssetData.networkRewards;

        address signerAddr = _contractsRegistry.getSigner();

        require(signerAddr != address(0), "BillingManager: Zero signer address");

        _sendFunds(_depositAssetKey, signerAddr, amountToWithdraw);

        emit RewardsWithdrawn(_depositAssetKey, signerAddr, amountToWithdraw);
    }

    function getTotalUsersCount() external view override returns (uint256) {
        return _existingUsers.length();
    }

    function getSupportedDepositAssetKeys() external view override returns (string[] memory) {
        return _supportedDepositAssetKeys.values();
    }

    function getDepositAssetsInfo(
        string[] calldata _depositAssetKeysArr
    ) external view override returns (DepositAssetInfo[] memory _depositAssetsInfoArr) {
        _depositAssetsInfoArr = new DepositAssetInfo[](_depositAssetKeysArr.length);

        for (uint256 i = 0; i < _depositAssetKeysArr.length; i++) {
            _depositAssetsInfoArr[i] = DepositAssetInfo(
                _depositAssetKeysArr[i],
                _depositAssetsData[_depositAssetKeysArr[i]]
            );
        }
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

    function getUserWithdrawNonce(address _userAddr) external view override returns (uint256) {
        return _usersData[_userAddr].withdrawNonce;
    }

    function getUserDepositAssetKeys(address _userAddr) external view override returns (string[] memory) {
        return _usersData[_userAddr].depositAssetKeys.values();
    }

    function getNetworkRewards(string memory _depositAssetKey) external view override returns (uint256) {
        return _depositAssetsData[_depositAssetKey].networkRewards;
    }

    function getUserDepositInfo(
        address _userAddr,
        string memory _depositAssetKey
    ) public view override returns (UserDepositInfo memory) {
        return UserDepositInfo(_userAddr, _usersData[_userAddr].userDepositsAmount[_depositAssetKey]);
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

        _updateUserDepositData(_depositAssetKey, _recipientAddr, _depositAmount, true);

        emit AssetDeposited(_depositAssetKey, _tokensSenderAddr, _recipientAddr, _depositAmount);
    }

    function _updateUserDepositData(
        string memory _depositAssetKey,
        address _userAddr,
        uint256 _amountToUpdate,
        bool _isAdding
    ) internal {
        require(_amountToUpdate > 0, "BillingManager: Zero amount to update");

        UserData storage userData = _usersData[_userAddr];

        if (_isAdding) {
            userData.userDepositsAmount[_depositAssetKey] += _amountToUpdate;
            userData.depositAssetKeys.add(_depositAssetKey);

            _existingUsers.add(_userAddr);
        } else {
            uint256 newUserFundsAmount = userData.userDepositsAmount[_depositAssetKey] - _amountToUpdate;

            userData.userDepositsAmount[_depositAssetKey] = newUserFundsAmount;

            if (newUserFundsAmount == 0) {
                _existingUsers.remove(_userAddr);
            }
        }
    }

    function _sendFunds(string memory _depositAssetKey, address _recipientAddr, uint256 _amountToSend) internal {
        if (isNativeDepositAsset(_depositAssetKey)) {
            Address.sendValue(payable(_recipientAddr), _amountToSend);
        } else {
            IERC20(_depositAssetsData[_depositAssetKey].tokenAddr).safeTransfer(_recipientAddr, _amountToSend);
        }
    }

    function _onlySigner(address _userAddr) internal view {
        require(_contractsRegistry.getSigner() == _userAddr, "BillingManager: Not a signer");
    }

    function _onlyExistingDepositAsset(string memory _depositAssetKey) internal view {
        require(isDepositAssetSupported(_depositAssetKey), "BillingManager: Deposit asset does not exist");
    }

    function _onlyEnabledDepositAsset(string memory _depositAssetKey) internal view {
        require(isDepositAssetEnabled(_depositAssetKey), "BillingManager: Deposit asset is disabled");
    }

    function _onlyEnoughFunds(string memory _depositAssetKey, address _userAddr, uint256 _neededAmount) internal view {
        require(
            _usersData[_userAddr].userDepositsAmount[_depositAssetKey] >= _neededAmount,
            "BillingManager: Not enough deposited funds"
        );
    }
}
