// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IStaking {
    enum ValidatorStatus {
        INACTIVE,
        ACTIVE,
        SLASHED
    }

    struct ValidatorData {
        uint256 stake;
        ValidatorStatus status;
    }

    struct UserStakeInfo {
        address userAddr;
        uint256 userStakedAmount;
    }

    struct WithdrawalAnnouncement {
        uint256 tokensAmount;
        uint256 withdrawalTime;
    }

    event MinimalStakeUpdated(uint256 minimalStake);
    event TokensStaked(address indexed tokensSender, address indexed tokensRecipient, uint256 stakeAmount);
    event WithdrawalAnnounced(address indexed userAddr, uint256 tokensAmount, uint256 withdrawalTime);
    event TokensWithdrawn(address indexed userAddr, uint256 tokensAmount);

    function setMinimalStake(uint256 _minimalStake) external;

    function updateWhitelistedUsers(address[] calldata _usersToUpdate, bool _isAdding) external;

    function slashValidator(address _validator) external;

    function addRewardsToStake(address _validator, uint256 _amount) external;

    function announceWithdrawal(uint256 _amount) external;

    function withdraw() external;

    function stake(uint256 _stakeAmount) external;

    function stakeWithPermit(
        uint256 _stakeAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external;

    function minimalStake() external view returns (uint256);

    function totalStake() external view returns (uint256);

    function getUsersStakeInfo(
        uint256 _offset,
        uint256 _limit
    ) external view returns (UserStakeInfo[] memory _usersStakeInfoArr);

    function getStake(address _validator) external view returns (uint256);

    function getWithdrawalAnnouncement(address _userAddr) external view returns (WithdrawalAnnouncement memory);

    function isUserWhitelisted(address _userAddr) external view returns (bool);

    function hasWithdrawalAnnouncement(address _userAddr) external view returns (bool);
}
