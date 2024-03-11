// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

/**
 * @title IGatewayFactory
 * @notice Interface for the GatewayFactory contract, responsible for deploying new instances of Gateway contracts
 */
interface IGatewayFactory {
    /**
     * @notice Event emitted when a new Gateway contract is deployed
     * @param gatewayAddr The address of the deployed Gateway contract
     * @param gatewayOwnerAddr The address of the owner of the deployed Gateway contract
     */
    event GatewayDeployed(address gatewayAddr, address gatewayOwnerAddr);

    /**
     * @notice Sets the new implementation address for the Gateway contract
     * @param _newGatewayImpl The address of the new implementation contract
     */
    function setNewImplementation(address _newGatewayImpl) external;

    /**
     * @notice Deploys a new Gateway contract with the provided owner address
     * @param _gatewayOwner The address of the owner for the new Gateway contract
     * @return Address of the newly deployed Gateway contract
     */
    function deployGateway(address _gatewayOwner) external returns (address);

    /**
     * @notice Retrieves the address of the current Gateway implementation contract
     * @return Address of the current Gateway implementation contract
     */
    function getGatewayImplementation() external view returns (address);
}
