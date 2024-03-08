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
}
