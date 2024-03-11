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

    struct ValidatorInfo {
        address validatorAddr;
        ValidatorData validatorData;
    }

    struct WithdrawalAnnouncement {
        uint256 amount;
        uint256 time;
    }

    event MinimalStakeUpdated(uint256 minimalStake);
    event WithdrawalPeriodUpdated(uint256 withdrawalPeriod);
    event TokensStaked(address indexed tokensSender, address indexed tokensRecipient, uint256 stakeAmount);
    event TokensWithdrawn(address indexed userAddr, uint256 tokensAmount);

    function setMinimalStake(uint256 _minimalStake) external;

    function setWithdrawalPeriod(uint256 _withdrawalPeriod) external;

    function addRewardsToStake(address _validator, uint256 _amount) external;

    function slash(address _validator) external;

    function announceWithdrawal(uint256 _amount) external;

    function revokeWithdrawal() external;

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

    function withdrawalPeriod() external view returns (uint256);

    function totalStake() external view returns (uint256);

    function getValidators() external view returns (address[] memory);

    function getValidatorsCount() external view returns (uint256);

    function getValidatorsInfo(uint256 _offset, uint256 _limit) external view returns (ValidatorInfo[] memory);

    function getStake(address _validator) external view returns (uint256);

    function getWithdrawalAnnouncement(address _userAddr) external view returns (WithdrawalAnnouncement memory);

    function isValidatorActive(address _validator) external view returns (bool);

    function isValidatorSlashed(address _validator) external view returns (bool);

    function getValidatorStatus(address _validator) external view returns (ValidatorStatus);

    function hasWithdrawalAnnouncement(address _userAddr) external view returns (bool);
}
