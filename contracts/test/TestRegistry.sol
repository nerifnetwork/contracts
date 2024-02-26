// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../interfaces/IGateway.sol";

import "../operational/Registry.sol";

contract TestRegistry is Registry {
    function callPerform(address _gatewayAddr, uint256 _id, address _target, bytes calldata _payload) external {
        IGateway(_gatewayAddr).perform(_id, _target, _payload);
    }
}
