// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

/**
 * @notice Interface for managing NerifToken
 */
interface INerifToken {
    /**
     * @notice Mints tokens and assigns them to a specified recipient
     * @dev Only contract owner can call this function
     * @param _recipient The address of the recipient to whom the tokens will be minted
     * @param _tokensAmount The amount of tokens to be minted
     */
    function ownerMint(address _recipient, uint256 _tokensAmount) external;

    /**
     * @notice Mints tokens and assigns them to a specified recipient
     * @dev Only vesting contract can call this function
     * @param _recipient The address of the recipient to whom the tokens will be minted
     * @param _tokensAmount The amount of tokens to be minted
     */
    function vestingMint(address _recipient, uint256 _tokensAmount) external;
}
