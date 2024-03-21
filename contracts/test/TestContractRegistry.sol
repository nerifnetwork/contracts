// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "../core/ContractsRegistry.sol";

contract TestContractsRegistry is ContractsRegistry {
    function version() external pure returns (string memory) {
        return "v2.0.0";
    }
}
