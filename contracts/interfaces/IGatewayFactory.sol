// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IGatewayFactory {
    event GatewayDeployed(address gatewayAddr, address gatewayOwnerAddr);

    function setNewImplementation(address _newGatewayImpl) external;

    function deployGateway(address _gatewayOwner) external returns (address);

    function getGatewayImplementation() external view returns (address);
}
