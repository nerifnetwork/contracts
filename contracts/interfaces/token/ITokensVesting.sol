// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@solarity/solidity-lib/finance/vesting/Vesting.sol";

/**
 * @notice Interface for managing TokensVesting contract
 */
interface ITokensVesting {
    /**
     * @notice Creates a new vesting schedule
     * @dev Only contract owner can call this function
     * @param _baseSchedule The base schedule information for the vesting schedule
     * @return The ID of the created vesting schedule
     */
    function createSchedule(Vesting.BaseSchedule memory _baseSchedule) external returns (uint256);

    /**
     * @notice Creates a new token vesting
     * @dev Only contract owner can call this function
     * @param _vesting The data for configuring the token vesting
     * @return The ID of the created token vesting
     */
    function createVesting(Vesting.VestingData memory _vesting) external returns (uint256);
}
