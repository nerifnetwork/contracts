// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../operational/Gateway.sol";

contract TestGateway is Gateway {
    function version() external pure returns (string memory) {
        return "2.0.0";
    }
}
