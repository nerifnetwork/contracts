// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

//
/**
 * @notice ISignerAddress represents the behavior of the contract that holds the network collective address
 * generated during DKG process.
 */
interface ISignerAddress {
    /**
     * @notice Returns the current collective signer address
     * @return The current collective signer address
     */
    function getSignerAddress() external view returns (address);
}
