// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import "../operational/BillingManager.sol";

contract TestBillingManager is BillingManager {
    function version() external pure returns (string memory) {
        return "v2.0.0";
    }
}
