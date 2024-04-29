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
import "../interfaces/operational/IRegistry.sol";

contract BillingManager is IBillingManager, AbstractDependant, EIP712Upgradeable {
    using EnumerableSet for *;
    using StringSet for StringSet.Set;
    using SafeERC20 for IERC20;

    // solhint-disable-next-line var-name-mixedcase
    bytes32 private constant _FUNDS_WITHDRAW_TYPEHASH =
        keccak256(
            // solhint-disable-next-line max-line-length
            "FundsWithdraw(address userAddr,bytes32 depositAssetsHash,bytes32 amountsHash,uint256 nonce,uint256 deadline)"
        );

    // solhint-disable-next-line var-name-mixedcase
    bytes32 private constant _REWARDS_WITHDRAW_TYPEHASH =
        keccak256(
            // solhint-disable-next-line max-line-length
            "RewardsWithdraw(address userAddr,bytes32 depositAssetsHash,bytes32 amountsHash,uint256 nonce,uint256 deadline)"
        );

    IContractsRegistry internal _contractsRegistry;
    IRegistry internal _registry;

    string public override nativeDepositAssetKey;
    string public override nerifTokenDepositAssetKey;

    StringSet.Set internal _supportedDepositAssetKeys;

    mapping(string => DepositAssetData) internal _depositAssetsData;
    mapping(address => mapping(uint256 => bool)) internal _usersNonces;

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

    function initialize(
        DepositAssetInfo calldata _nativeDepositAssetInfo,
        DepositAssetInfo calldata _nerifTokenDepositAssetInfo
    ) external initializer {
        __EIP712_init("BillingManager", "1");

        nativeDepositAssetKey = _nativeDepositAssetInfo.depositAssetKey;
        nerifTokenDepositAssetKey = _nerifTokenDepositAssetInfo.depositAssetKey;

        _addDepositAsset(_nativeDepositAssetInfo);
        _addDepositAsset(_nerifTokenDepositAssetInfo);
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _contractsRegistry = contractsRegistry;
        _registry = IRegistry(contractsRegistry.getRegistryContract());
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
        SigData calldata _sigData
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
            _sigData.sigExpirationTime,
            _sigData.v,
            _sigData.r,
            _sigData.s
        );

        _depositAsset(_depositAssetKey, msg.sender, _recipientAddr, _depositAmount, false);
    }

    function withdrawFunds(WithdrawData calldata _withdrawData) external {
        _withdraw(msg.sender, _FUNDS_WITHDRAW_TYPEHASH, _withdrawData);

        emit FundsWithdrawn(
            msg.sender,
            _withdrawData.depositAssetKeys,
            _withdrawData.withdrawAmounts,
            _withdrawData.nonce
        );
    }

    function withdrawRewards(WithdrawData calldata _withdrawData) external {
        _withdraw(msg.sender, _REWARDS_WITHDRAW_TYPEHASH, _withdrawData);

        emit RewardsWithdrawn(
            msg.sender,
            _withdrawData.depositAssetKeys,
            _withdrawData.withdrawAmounts,
            _withdrawData.nonce
        );
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

    function getTotalDepositAssetAmount(string calldata _depositAssetKey) external view returns (uint256) {
        return _depositAssetsData[_depositAssetKey].totalAssetAmount;
    }

    function isUserNonceUsed(address _userAddr, uint256 _nonce) public view returns (bool) {
        return _usersNonces[_userAddr][_nonce];
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
            _depositAssetInfo.depositAssetData.totalAssetAmount == 0,
            "BillingManager: Invalid init total asset amount value"
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

        _depositAssetsData[_depositAssetKey].totalAssetAmount += _depositAmount;

        emit AssetDeposited(_depositAssetKey, _tokensSenderAddr, _recipientAddr, _depositAmount);
    }

    function _withdraw(address _userAddr, bytes32 _structTypehash, WithdrawData calldata _withdrawData) internal {
        require(block.timestamp <= _withdrawData.sigData.sigExpirationTime, "BillingManager: Expired deadline");
        require(!isUserNonceUsed(msg.sender, _withdrawData.nonce), "BillingManager: Nonce has already been used");

        bytes32 rewardsWithdrawStructHash = _getStructHash(
            _structTypehash,
            _userAddr,
            _withdrawData.depositAssetKeys,
            _withdrawData.withdrawAmounts,
            _withdrawData.sigData.sigExpirationTime,
            _withdrawData.nonce
        );

        _onlySigner(
            ECDSA.recover(
                _hashTypedDataV4(rewardsWithdrawStructHash),
                _withdrawData.sigData.v,
                _withdrawData.sigData.r,
                _withdrawData.sigData.s
            )
        );

        for (uint256 i = 0; i < _withdrawData.depositAssetKeys.length; i++) {
            _withdrawAsset(_userAddr, _withdrawData.depositAssetKeys[i], _withdrawData.withdrawAmounts[i]);
        }

        _usersNonces[msg.sender][_withdrawData.nonce] = true;
    }

    function _withdrawAsset(address _recipientAddr, string calldata _assetKey, uint256 _withdrawAmount) internal {
        _onlyExistingDepositAsset(_assetKey);
        _onlyEnoughDepositAsset(_assetKey, _withdrawAmount);
        _sendFunds(_assetKey, _recipientAddr, _withdrawAmount);

        _depositAssetsData[_assetKey].totalAssetAmount -= _withdrawAmount;
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

    function _onlyEnoughDepositAsset(string memory _depositAssetKey, uint256 _neededAmount) internal view {
        require(
            _depositAssetsData[_depositAssetKey].totalAssetAmount >= _neededAmount,
            "BillingManager: Not enough deposit asset"
        );
    }

    function _onlyEnabledDepositAsset(string memory _depositAssetKey) internal view {
        require(isDepositAssetEnabled(_depositAssetKey), "BillingManager: Deposit asset is disabled");
    }

    function _getStructHash(
        bytes32 _structTypehash,
        address _userAddr,
        string[] calldata _depositAssetKeys,
        uint256[] calldata _withdrawAmounts,
        uint256 _sigExpirationTime,
        uint256 _nonce
    ) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    _structTypehash,
                    _userAddr,
                    keccak256(abi.encode(_depositAssetKeys)),
                    keccak256(abi.encode(_withdrawAmounts)),
                    _nonce,
                    _sigExpirationTime
                )
            );
    }
}
