// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "@solarity/solidity-lib/proxy/beacon/ProxyBeacon.sol";
import "@solarity/solidity-lib/proxy/beacon/PublicBeaconProxy.sol";

import "../interfaces/IGateway.sol";
import "../interfaces/IGatewayFactory.sol";

contract GatewayFactory is IGatewayFactory, OwnableUpgradeable {
    ProxyBeacon public gatewayBeacon;

    address public registryAddr;

    modifier onlyRegistry() {
        require(registryAddr == msg.sender, "GatewayFactory: operation is not permitted");
        _;
    }

    function initialize(address _registryAddr, address _gatewayImpl) external initializer {
        __Ownable_init();

        gatewayBeacon = new ProxyBeacon();
        registryAddr = _registryAddr;

        _setNewImplementation(_gatewayImpl);
    }

    function setNewImplementation(address _newGatewayImpl) external override onlyOwner {
        _setNewImplementation(_newGatewayImpl);
    }

    function deployGateway(address _gatewayOwner) external override onlyRegistry returns (address) {
        address newGatewayProxyAddr = address(new PublicBeaconProxy(address(gatewayBeacon), ""));

        IGateway(newGatewayProxyAddr).initialize(msg.sender, _gatewayOwner);

        return newGatewayProxyAddr;
    }

    function _setNewImplementation(address _newImplementation) internal {
        if (gatewayBeacon.implementation() != _newImplementation) {
            gatewayBeacon.upgradeTo(_newImplementation);
        }
    }
}
