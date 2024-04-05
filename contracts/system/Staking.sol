// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/system/IStaking.sol";
import "../interfaces/system/IDKG.sol";

import "./SlashingVoting.sol";
import "./RewardDistributionPool.sol";

contract Staking is IStaking, Initializable, AbstractDependant {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

    IContractsRegistry internal _contractsRegistry;
    IDKG internal _dkg;
    SlashingVoting internal _slashingVoting;
    RewardDistributionPool internal _rewardsDistributionPool;

    IERC20 public stakeToken;

    uint256 public override minimalStake;
    uint256 public override withdrawalPeriod;
    uint256 public override totalStake;

    EnumerableSet.AddressSet internal _validators;

    mapping(address => ValidatorData) internal _validatorsData;
    mapping(address => WithdrawalAnnouncement) internal _withdrawalAnnouncements;

    modifier onlySigner() {
        _onlySigner();
        _;
    }

    modifier onlyNotSlashed() {
        require(!isValidatorSlashed(msg.sender), "Staking: validator is slashed");
        _;
    }

    modifier onlySlashingVoting() {
        require(msg.sender == address(_slashingVoting), "Staking: not a slashing voting");
        _;
    }

    modifier onlyRewardDistributionPool() {
        require(msg.sender == address(_rewardsDistributionPool), "Staking: only RewardDistributionPool contract");
        _;
    }

    function initialize(address _stakeToken, uint256 _minimalStake, uint256 _withdrawalPeriod) external initializer {
        stakeToken = IERC20(_stakeToken);

        _setMinimalStake(_minimalStake);
        _setWithdrawalPeriod(_withdrawalPeriod);
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _contractsRegistry = contractsRegistry;
        _dkg = IDKG(contractsRegistry.getDKGContract());
        _slashingVoting = SlashingVoting(contractsRegistry.getSlashingVotingContract());
        _rewardsDistributionPool = RewardDistributionPool(
            payable(contractsRegistry.getRewardsDistributionPoolContract())
        );
    }

    // solhint-disable-next-line ordering
    function setMinimalStake(uint256 _minimalStake) external override onlySigner {
        _setMinimalStake(_minimalStake);
    }

    function setWithdrawalPeriod(uint256 _withdrawalPeriod) external override onlySigner {
        _setWithdrawalPeriod(_withdrawalPeriod);
    }

    function addRewardsToStake(address _validator, uint256 _amount) external override onlyRewardDistributionPool {
        _validatorsData[_validator].stake += _amount;
        totalStake += _amount;
    }

    function slash(address _validator) external override onlySlashingVoting {
        _updateValidatorStatus(_validator, ValidatorStatus.SLASHED);
    }

    function announceWithdrawal(uint256 _amount) external override onlyNotSlashed {
        require(_amount > 0, "Staking: amount must be greater than zero");
        require(_amount <= _validatorsData[msg.sender].stake, "Staking: amount must be <= to stake");

        _withdrawalAnnouncements[msg.sender] = WithdrawalAnnouncement(_amount, block.timestamp);

        if (_validatorsData[msg.sender].stake - _amount < minimalStake && _validators.contains(msg.sender)) {
            _updateValidatorStatus(msg.sender, ValidatorStatus.INACTIVE);
        }
    }

    function revokeWithdrawal() external override onlyNotSlashed {
        require(hasWithdrawalAnnouncement(msg.sender), "Staking: user does not have withdrawal announcement");

        uint256 amount = _withdrawalAnnouncements[msg.sender].amount;

        delete _withdrawalAnnouncements[msg.sender];

        if (_validatorsData[msg.sender].stake + amount >= minimalStake && !_validators.contains(msg.sender)) {
            _updateValidatorStatus(msg.sender, ValidatorStatus.ACTIVE);
        }
    }

    function withdraw() external override onlyNotSlashed {
        require(hasWithdrawalAnnouncement(msg.sender), "Staking: user does not have withdrawal announcement");
        require(
            _withdrawalAnnouncements[msg.sender].time + withdrawalPeriod <= block.timestamp,
            "Staking: withdrawal period not passed"
        );

        uint256 withdrawalAmount = _withdrawalAnnouncements[msg.sender].amount;

        _validatorsData[msg.sender].stake -= withdrawalAmount;
        totalStake -= withdrawalAmount;

        delete _withdrawalAnnouncements[msg.sender];

        stakeToken.safeTransfer(msg.sender, withdrawalAmount);

        emit TokensWithdrawn(msg.sender, withdrawalAmount);
    }

    function stake(uint256 _stakeAmount) external override onlyNotSlashed {
        _stake(msg.sender, msg.sender, _stakeAmount);
    }

    function stakeWithPermit(
        uint256 _stakeAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external override onlyNotSlashed {
        IERC20Permit(address(stakeToken)).permit(
            msg.sender,
            address(this),
            _stakeAmount,
            _sigExpirationTime,
            _v,
            _r,
            _s
        );

        _stake(msg.sender, msg.sender, _stakeAmount);
    }

    function getValidators() external view override returns (address[] memory) {
        return _validators.values();
    }

    function getValidatorsCount() external view override returns (uint256) {
        return _validators.length();
    }

    function getValidatorsInfo(
        uint256 _offset,
        uint256 _limit
    ) external view override returns (ValidatorInfo[] memory _validatorsInfoArr) {
        uint256 to = Paginator.getTo(_validators.length(), _offset, _limit);

        _validatorsInfoArr = new ValidatorInfo[](to - _offset);

        for (uint256 i = _offset; i < to; i++) {
            address currentValidatorAddr = _validators.at(i);

            _validatorsInfoArr[i - _offset] = ValidatorInfo(
                currentValidatorAddr,
                _validatorsData[currentValidatorAddr]
            );
        }
    }

    function getStake(address _validator) external view override returns (uint256) {
        return _validatorsData[_validator].stake;
    }

    function getWithdrawalAnnouncement(
        address _userAddr
    ) external view override returns (WithdrawalAnnouncement memory) {
        return _withdrawalAnnouncements[_userAddr];
    }

    function isValidatorActive(address _validator) public view override returns (bool) {
        return getValidatorStatus(_validator) == ValidatorStatus.ACTIVE;
    }

    function isValidatorSlashed(address _validator) public view override returns (bool) {
        return getValidatorStatus(_validator) == ValidatorStatus.SLASHED;
    }

    function getValidatorStatus(address _validator) public view override returns (ValidatorStatus) {
        return _validatorsData[_validator].status;
    }

    function hasWithdrawalAnnouncement(address _userAddr) public view override returns (bool) {
        return _withdrawalAnnouncements[_userAddr].time > 0;
    }

    function _setMinimalStake(uint256 _minimalStake) internal {
        minimalStake = _minimalStake;

        emit MinimalStakeUpdated(_minimalStake);
    }

    function _setWithdrawalPeriod(uint256 _withdrawalPeriod) internal {
        withdrawalPeriod = _withdrawalPeriod;

        emit WithdrawalPeriodUpdated(_withdrawalPeriod);
    }

    function _updateValidatorStatus(address _validatorToUpdate, ValidatorStatus _newStatus) internal {
        bool success;

        if (_newStatus == ValidatorStatus.INACTIVE || _newStatus == ValidatorStatus.SLASHED) {
            success = _validators.remove(_validatorToUpdate);
        } else {
            success = _validators.add(_validatorToUpdate);
        }

        require(success, "Staking: Failed to update validator status");

        _validatorsData[_validatorToUpdate].status = _newStatus;

        _dkg.updateActiveValidators();
    }

    function _updateValidators(address _validatorToUpdate, bool _isAdding) internal {
        bool success = _isAdding ? _validators.add(_validatorToUpdate) : _validators.remove(_validatorToUpdate);

        require(success, "Staking: Invalid validator address to update");

        _dkg.updateActiveValidators();
    }

    function _stake(address _tokenSenderAddr, address _stakeRecipientAddr, uint256 _stakeAmount) internal {
        require(_stakeAmount > 0, "Staking: Zero stake amount");
        require(stakeToken.balanceOf(_tokenSenderAddr) >= _stakeAmount, "Staking: Not enough tokens to stake");

        ValidatorData storage validatorData = _validatorsData[_stakeRecipientAddr];

        if (validatorData.status == ValidatorStatus.INACTIVE && validatorData.stake + _stakeAmount >= minimalStake) {
            validatorData.status = ValidatorStatus.ACTIVE;

            _updateValidators(msg.sender, true);
        }

        stakeToken.safeTransferFrom(_tokenSenderAddr, address(this), _stakeAmount);

        validatorData.stake += _stakeAmount;
        totalStake += _stakeAmount;

        emit TokensStaked(_tokenSenderAddr, _stakeRecipientAddr, _stakeAmount);
    }

    function _onlySigner() internal view {
        require(_contractsRegistry.getSigner() == msg.sender, "Staking: Not a signer");
    }
}
