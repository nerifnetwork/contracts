// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

/**
 * @title IGateway
 * @notice Interface for the Gateway contract, responsible for managing known workflows and customer contracts
 */
interface IGateway {
    /**
     * @notice Struct for storing configuration information including known workflows and addresses
     * @param knownWorkflows The array of known workflow IDs
     * @param knownWorkflowOwners The array of addresses corresponding to the owners of known workflows
     * @param knownCustomerContracts The array of addresses corresponding to known customer contracts
     */
    struct ConfigInfo {
        uint256[] knownWorkflows;
        address[] knownWorkflowOwners;
        address[] knownCustomerContracts;
    }

    /**
     * @notice Struct for updating configuration data including known workflows and addresses
     * @param updateKnownWorkflowsEntries The array of entries for updating known workflows
     * @param updateKnownWorkflowOwnersEntries The array of entries for updating known workflow owners
     * @param updateKnownCustomerContractsEntries The array of entries for updating known customer contracts
     */
    struct UpdateConfigData {
        UpdateKnownWorkflowsEntry[] updateKnownWorkflowsEntries;
        UpdateKnownAddressesEntry[] updateKnownWorkflowOwnersEntries;
        UpdateKnownAddressesEntry[] updateKnownCustomerContractsEntries;
    }

    /**
     * @notice Struct for updating known workflows
     * @param knownWorkflowToUpdate The ID of the known workflow to update
     * @param isAdding The flag indicating whether to add or remove the known workflow
     */
    struct UpdateKnownWorkflowsEntry {
        uint256 knownWorkflowToUpdate;
        bool isAdding;
    }

    /**
     * @notice Struct for updating known addresses
     * @param addressToUpdate The address to update
     * @param isAdding The flag indicating whether to add or remove the address
     */
    struct UpdateKnownAddressesEntry {
        address addressToUpdate;
        bool isAdding;
    }

    /**
     * @notice Initializes the Gateway contract with the provided registry and owner addresses
     * @param _registryAddr The address of the registry contract
     * @param _gatewayOwnerAddr The address of the owner of the Gateway contract
     */
    function initialize(address _registryAddr, address _gatewayOwnerAddr) external;

    /**
     * @notice Entry point function for all customer contracts
     * @dev This function accepts the workflow ID and the end customer contract execution payload
     * The function checks that the given transaction is permitted and can be forwarded to the end customer contract
     * @param _id The ID of the workflow
     * @param _target The address of the end customer contract
     * @param _payload The payload to execute on the end customer contract
     */
    function perform(
        uint256 _id,
        address _target,
        bytes calldata _payload
    ) external;

    /**
     * @notice Sets the registry address for the Gateway contract
     * @param _registry The new address of the registry contract
     */
    function setRegistry(address _registry) external;

    /**
     * @notice Updates configuration data for the Gateway contract
     * @param _updateConfigData The struct containing the updated configuration data
     */
    function updateConfigData(UpdateConfigData memory _updateConfigData) external;

    /**
     * @notice Updates known workflows based on the provided entries
     * @param _updateKnownWorkflowsEntries The array of entries for updating known workflows
     */
    function updateKnownWorkflows(UpdateKnownWorkflowsEntry[] memory _updateKnownWorkflowsEntries) external;

    /**
     * @notice Updates known workflow owners based on the provided entries
     * @param _updateKnownWorkflowOwnersEntries The array of entries for updating known workflow owners
     */
    function updateKnownWorkflowOwners(UpdateKnownAddressesEntry[] memory _updateKnownWorkflowOwnersEntries) external;

    /**
     * @notice Updates known customer contracts based on the provided entries
     * @param _updateKnownCustomerContractsEntries The array of entries for updating known customer contracts
     */
    function updateKnownCustomerContracts(UpdateKnownAddressesEntry[] memory _updateKnownCustomerContractsEntries)
        external;

    /**
     * @notice Retrieves configuration information for the Gateway contract
     * @return ConfigInfo Struct containing configuration information
     */
    function getConfigInfo() external view returns (ConfigInfo memory);
}
