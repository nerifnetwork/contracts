// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/IGateway.sol";

abstract contract RegistryGateway {
    struct Gateway {
        IGateway gateway;
        address owner;
    }

    mapping(address => uint256) private indexMap;
    Gateway[] private gateways;

    function getGateways() external view returns (Gateway[] memory) {
        return gateways;
    }

    function getGateway(address owner) public view returns (IGateway) {
        if (!_hasGateway(owner)) {
            return IGateway(address(0x0));
        }

        return gateways[indexMap[owner] - 1].gateway;
    }

    function _setGateway(address owner, IGateway gateway) internal {
        if (_hasGateway(owner)) {
            gateways[indexMap[owner] - 1] = Gateway(gateway, owner);
        } else {
            gateways.push(Gateway(gateway, owner));
            indexMap[owner] = gateways.length;
        }

        _checkGatewayEntry(owner);
    }

    function _checkGatewayEntry(address owner) private view {
        uint256 index = indexMap[owner];
        assert(index <= gateways.length);

        if (_hasGateway(owner)) {
            assert(index > 0);
        } else {
            assert(index == 0);
        }
    }

    function _hasGateway(address owner) private view returns (bool) {
        return indexMap[owner] > 0;
    }
}
