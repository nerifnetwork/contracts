// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/libs/arrays/Paginator.sol";
import "@solarity/solidity-lib/libs/arrays/SetHelper.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/system/IStaking.sol";
import "../interfaces/system/IDKG.sol";

import "./RewardDistributionPool.sol";

contract Staking is IStaking, Initializable, AbstractDependant {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;
    using SetHelper for EnumerableSet.AddressSet;

    IContractsRegistry internal _contractsRegistry;
    IDKG internal _dkg;
    RewardDistributionPool internal _rewardsDistributionPool;

    IERC20 public stakeToken;

    uint256 public override minimalStake;
    uint256 public override totalStake;

    EnumerableSet.AddressSet internal _usersWhitelist;

    mapping(address => uint256) internal _usersStake;
    mapping(address => WithdrawalAnnouncement) internal _withdrawalAnnouncements;

    modifier onlySigner() {
        _onlySigner();
        _;
    }

    modifier onlyWhitelistedUser() {
        _onlyWhitelistedUser();
        _;
    }

    modifier onlyRewardDistributionPool() {
        require(msg.sender == address(_rewardsDistributionPool), "Staking: only RewardDistributionPool contract");
        _;
    }

    function initialize(
        address _stakeToken,
        uint256 _minimalStake,
        address[] calldata _whitelistedUsers
    ) external initializer {
        stakeToken = IERC20(_stakeToken);

        _setMinimalStake(_minimalStake);
        _usersWhitelist.add(_whitelistedUsers);
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _contractsRegistry = contractsRegistry;
        _dkg = IDKG(contractsRegistry.getDKGContract());
        _rewardsDistributionPool = RewardDistributionPool(
            payable(contractsRegistry.getRewardsDistributionPoolContract())
        );
    }

    // solhint-disable-next-line ordering
    function setMinimalStake(uint256 _minimalStake) external override onlySigner {
        _setMinimalStake(_minimalStake);
    }

    function updateWhitelistedUsers(address[] calldata _usersToUpdate, bool _isAdding) external onlySigner {
        if (_isAdding) {
            _usersWhitelist.add(_usersToUpdate);
        } else {
            _usersWhitelist.remove(_usersToUpdate);
        }
    }

    function addRewardsToStake(address _validator, uint256 _amount) external override onlyRewardDistributionPool {
        _usersStake[_validator] += _amount;
        totalStake += _amount;
    }

    function announceWithdrawal(uint256 _amountToAnnounce) external override {
        require(_amountToAnnounce > 0, "Staking: Amount must be greater than zero");
        require(!hasWithdrawalAnnouncement(msg.sender), "Staking: user already has withdrawal announcement");

        uint256 userStakedAmount = _usersStake[msg.sender];

        require(_amountToAnnounce <= userStakedAmount, "Staking: Not enough stake");

        uint256 withdrawalTime = block.timestamp;

        if (userStakedAmount - minimalStake < _amountToAnnounce) {
            withdrawalTime = _dkg.announceValidatorExit(msg.sender);
        }

        _withdrawalAnnouncements[msg.sender] = WithdrawalAnnouncement(_amountToAnnounce, withdrawalTime);

        emit WithdrawalAnnounced(msg.sender, _amountToAnnounce, withdrawalTime);
    }

    function withdraw() external override {
        require(hasWithdrawalAnnouncement(msg.sender), "Staking: User does not have withdrawal announcement");
        require(
            _withdrawalAnnouncements[msg.sender].withdrawalTime <= block.timestamp,
            "Staking: The time for withdrawal has not come"
        );

        uint256 withdrawalAmount = _withdrawalAnnouncements[msg.sender].tokensAmount;

        _usersStake[msg.sender] -= withdrawalAmount;
        totalStake -= withdrawalAmount;

        delete _withdrawalAnnouncements[msg.sender];

        _dkg.removeValidator(msg.sender);

        stakeToken.safeTransfer(msg.sender, withdrawalAmount);

        emit TokensWithdrawn(msg.sender, withdrawalAmount);
    }

    function stake(uint256 _stakeAmount) external override onlyWhitelistedUser {
        _stake(msg.sender, msg.sender, _stakeAmount);
    }

    function stakeWithPermit(
        uint256 _stakeAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external override onlyWhitelistedUser {
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

    function getUsersStakeInfo(
        uint256 _offset,
        uint256 _limit
    ) external view override returns (UserStakeInfo[] memory _usersStakeInfoArr) {
        uint256 to = Paginator.getTo(_usersWhitelist.length(), _offset, _limit);

        _usersStakeInfoArr = new UserStakeInfo[](to - _offset);

        for (uint256 i = _offset; i < to; i++) {
            address currentUserAddr = _usersWhitelist.at(i);

            _usersStakeInfoArr[i - _offset] = UserStakeInfo(currentUserAddr, _usersStake[currentUserAddr]);
        }
    }

    function getStake(address _userAddr) external view override returns (uint256) {
        return _usersStake[_userAddr];
    }

    function getWithdrawalAnnouncement(
        address _userAddr
    ) external view override returns (WithdrawalAnnouncement memory) {
        return _withdrawalAnnouncements[_userAddr];
    }

    function isUserWhitelisted(address _userAddr) public view returns (bool) {
        return _usersWhitelist.contains(_userAddr);
    }

    function hasWithdrawalAnnouncement(address _userAddr) public view override returns (bool) {
        return _withdrawalAnnouncements[_userAddr].withdrawalTime > 0;
    }

    function _setMinimalStake(uint256 _minimalStake) internal {
        minimalStake = _minimalStake;

        emit MinimalStakeUpdated(_minimalStake);
    }

    function _stake(address _tokenSenderAddr, address _stakeRecipientAddr, uint256 _stakeAmount) internal {
        require(_stakeAmount > 0, "Staking: Zero stake amount");
        require(stakeToken.balanceOf(_tokenSenderAddr) >= _stakeAmount, "Staking: Not enough tokens to stake");

        uint256 newStakeAmount = _usersStake[_stakeRecipientAddr] + _stakeAmount;

        if (!_dkg.isValidator(_stakeRecipientAddr) && newStakeAmount >= minimalStake) {
            _dkg.addValidator(_stakeRecipientAddr);
        }

        stakeToken.safeTransferFrom(_tokenSenderAddr, address(this), _stakeAmount);

        _usersStake[_stakeRecipientAddr] = newStakeAmount;
        totalStake += _stakeAmount;

        emit TokensStaked(_tokenSenderAddr, _stakeRecipientAddr, _stakeAmount);
    }

    function _onlySigner() internal view {
        require(_contractsRegistry.getSigner() == msg.sender, "Staking: Not a signer");
    }

    function _onlyWhitelistedUser() internal view {
        require(isUserWhitelisted(msg.sender), "Staking: Not a whitelisted user");
    }
}
