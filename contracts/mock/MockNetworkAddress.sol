// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../interfaces/ISignerAddress.sol";

contract MockNetworkAddress is ISignerAddress {
    address public networkAddress;

    constructor(address addr) {
        networkAddress = addr;
    }

    function getSignerAddress() external view returns (address) {
        return networkAddress;
    }
}
