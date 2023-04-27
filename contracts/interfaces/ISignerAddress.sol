// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// ISignerAddress represents the behavior of the contract that holds the network collective address
// generated during DKG process.
interface ISignerAddress {
    // getSignerAddress returns the current signer address
    function getSignerAddress() external view returns (address);
}
