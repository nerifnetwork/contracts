// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// INetworkAddress represents the behaviuor of the contract that holds the network collective address
// generated during DKG process.
interface INetworkAddress {
    // getAddress returns the current collective network address
    function getAddress() external view returns (address);
}
