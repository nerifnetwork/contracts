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
import "../interfaces/system/IDKG.sol";
import "../interfaces/system/IStaking.sol";
import "../interfaces/system/ISlashingVoting.sol";
import "../interfaces/operational/IBillingManager.sol";

contract Staking is IStaking, Initializable, AbstractDependant {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;
    using SetHelper for EnumerableSet.AddressSet;

    IContractsRegistry internal _contractsRegistry;
    IDKG internal _dkg;
    ISlashingVoting internal _slashingVoting;
    IBillingManager internal _billingManager;

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

    modifier onlySlashingVoting() {
        _onlySlashingVoting();
        _;
    }

    modifier onlyWhitelistedUser() {
        _onlyWhitelistedUser();
        _;
    }

    modifier onlyNotSlashed() {
        _onlyNotSlashed();
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
        _slashingVoting = ISlashingVoting(contractsRegistry.getSlashingVotingContract());
        _billingManager = IBillingManager(contractsRegistry.getBillingManagerContract());
    }

    // solhint-disable-next-line ordering
    function setMinimalStake(uint256 _minimalStake) external override onlySigner {
        _setMinimalStake(_minimalStake);
    }

    function updateWhitelistedUsers(address[] calldata _usersToUpdate, bool _isAdding) external override onlySigner {
        if (_isAdding) {
            _usersWhitelist.add(_usersToUpdate);
        } else {
            _usersWhitelist.remove(_usersToUpdate);
        }
    }

    function slashValidator(address _validator) external override onlySlashingVoting {
        string memory nerifTokenDepositAssetKey = _billingManager.nerifTokenDepositAssetKey();

        require(
            _billingManager.getDepositAssetTokenAddr(nerifTokenDepositAssetKey) == address(stakeToken),
            "Staking: Stake token not a NERIF token"
        );

        _billingManager.addSlashedTokens(nerifTokenDepositAssetKey, _usersStake[_validator]);
    }

    function stake(uint256 _stakeAmount) external override onlyWhitelistedUser onlyNotSlashed {
        _stake(msg.sender, msg.sender, _stakeAmount);
    }

    function stakeWithPermit(
        uint256 _stakeAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external override onlyWhitelistedUser onlyNotSlashed {
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

    function announceWithdrawal() external override onlyNotSlashed {
        uint256 userStakedAmount = _usersStake[msg.sender];

        require(userStakedAmount > 0, "Staking: Nothing to withdraw");
        require(!hasWithdrawalAnnouncement(msg.sender), "Staking: User already has withdrawal announcement");

        uint256 withdrawalEpochId;

        if (userStakedAmount >= minimalStake) {
            withdrawalEpochId = _dkg.announceValidatorExit(msg.sender);
        } else {
            withdrawalEpochId = _dkg.getActiveEpochId();
        }

        _withdrawalAnnouncements[msg.sender] = WithdrawalAnnouncement(userStakedAmount, withdrawalEpochId);

        emit WithdrawalAnnounced(msg.sender, userStakedAmount, withdrawalEpochId);
    }

    function withdraw() external override onlyNotSlashed {
        require(hasWithdrawalAnnouncement(msg.sender), "Staking: User does not have withdrawal announcement");
        require(
            _withdrawalAnnouncements[msg.sender].withdrawalEpochId <= _dkg.getActiveEpochId(),
            "Staking: The time for withdrawal has not come"
        );

        uint256 withdrawalAmount = _withdrawalAnnouncements[msg.sender].tokensAmount;

        totalStake -= withdrawalAmount;

        delete _usersStake[msg.sender];
        delete _withdrawalAnnouncements[msg.sender];

        if (_dkg.isValidator(msg.sender)) {
            _dkg.removeValidator(msg.sender);
        }

        stakeToken.safeTransfer(msg.sender, withdrawalAmount);

        emit TokensWithdrawn(msg.sender, withdrawalAmount);
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

    function getWhitelistedUsers() external view returns (address[] memory) {
        return _usersWhitelist.values();
    }

    function isUserWhitelisted(address _userAddr) public view returns (bool) {
        return _usersWhitelist.contains(_userAddr);
    }

    function hasWithdrawalAnnouncement(address _userAddr) public view override returns (bool) {
        return _withdrawalAnnouncements[_userAddr].withdrawalEpochId > 0;
    }

    function _setMinimalStake(uint256 _minimalStake) internal {
        minimalStake = _minimalStake;

        emit MinimalStakeUpdated(_minimalStake);
    }

    function _stake(address _tokenSenderAddr, address _stakeRecipientAddr, uint256 _stakeAmount) internal {
        require(_stakeAmount > 0, "Staking: Zero stake amount");
        require(stakeToken.balanceOf(_tokenSenderAddr) >= _stakeAmount, "Staking: Not enough tokens to stake");
        require(!_dkg.isValidator(_stakeRecipientAddr), "Staking: User is already a validator");
        require(!hasWithdrawalAnnouncement(_stakeRecipientAddr), "Staking: User has a withdrawal announcement");

        uint256 newStakeAmount = _usersStake[_stakeRecipientAddr] + _stakeAmount;

        if (newStakeAmount >= minimalStake) {
            _dkg.addValidator(_stakeRecipientAddr);
        }

        _usersStake[_stakeRecipientAddr] = newStakeAmount;
        totalStake += _stakeAmount;

        stakeToken.safeTransferFrom(_tokenSenderAddr, address(this), _stakeAmount);

        emit TokensStaked(_tokenSenderAddr, _stakeRecipientAddr, _stakeAmount);
    }

    function _onlySigner() internal view {
        require(msg.sender == _contractsRegistry.getSigner(), "Staking: Not a signer");
    }

    function _onlySlashingVoting() internal view {
        require(msg.sender == address(_slashingVoting), "Staking: Not a slashing voting address");
    }

    function _onlyWhitelistedUser() internal view {
        require(isUserWhitelisted(msg.sender), "Staking: Not a whitelisted user");
    }

    function _onlyNotSlashed() internal view {
        require(!_dkg.isValidatorSlashed(msg.sender), "Staking: Validator was slashed");
    }
}
