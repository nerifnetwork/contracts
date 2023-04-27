// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../interfaces/INetworkAddress.sol";

contract MockNetworkAddress is INetworkAddress {
    address public networkAddress;

    constructor(address addr) {
        networkAddress = addr;
    }

    function getAddress() external view returns (address) {
        return networkAddress;
    }
}
