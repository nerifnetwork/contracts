// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@solarity/solidity-lib/finance/vesting/Vesting.sol";

interface ITokensVesting {
    function createSchedule(Vesting.BaseSchedule memory _baseSchedule) external returns (uint256);

    function createVesting(Vesting.VestingData memory _vesting) external returns (uint256);
}
