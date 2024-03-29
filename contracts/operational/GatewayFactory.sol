// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/proxy/beacon/ProxyBeacon.sol";
import "@solarity/solidity-lib/proxy/beacon/PublicBeaconProxy.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/operational/IGateway.sol";
import "../interfaces/operational/IGatewayFactory.sol";

contract GatewayFactory is IGatewayFactory, OwnableUpgradeable, AbstractDependant {
    ProxyBeacon public gatewayBeacon;

    address public registryAddr;

    modifier onlyRegistry() {
        require(registryAddr == msg.sender, "GatewayFactory: operation is not permitted");
        _;
    }

    function initialize(address _gatewayImpl) external initializer {
        __Ownable_init();

        gatewayBeacon = new ProxyBeacon();

        _setNewImplementation(_gatewayImpl);
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        registryAddr = contractsRegistry.getRegistryContract();
    }

    // solhint-disable-next-line ordering
    function setNewImplementation(address _newGatewayImpl) external override onlyOwner {
        _setNewImplementation(_newGatewayImpl);
    }

    function deployGateway(address _gatewayOwner) external override onlyRegistry returns (address) {
        address newGatewayProxyAddr = address(new PublicBeaconProxy(address(gatewayBeacon), ""));

        IGateway(newGatewayProxyAddr).initialize(msg.sender, _gatewayOwner);

        emit GatewayDeployed(newGatewayProxyAddr, _gatewayOwner);

        return newGatewayProxyAddr;
    }

    function getGatewayImplementation() external view override returns (address) {
        return gatewayBeacon.implementation();
    }

    function _setNewImplementation(address _newImplementation) internal {
        if (gatewayBeacon.implementation() != _newImplementation) {
            gatewayBeacon.upgradeTo(_newImplementation);
        }
    }
}
