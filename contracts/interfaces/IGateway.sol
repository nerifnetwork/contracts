// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// IGateway represents the customer gateway contract behaviour.
interface IGateway {
    struct ConfigInfo {
        uint256[] knownWorkflows;
        address[] knownWorkflowOwners;
        address[] knownCustomerContracts;
    }

    struct UpdateConfigData {
        UpdateKnownWorkflowsEntry[] updateKnownWorkflowsEntries;
        UpdateKnownAddressesEntry[] updateKnownWorkflowOwnersEntries;
        UpdateKnownAddressesEntry[] updateKnownCustomerContractsEntries;
    }

    struct UpdateKnownWorkflowsEntry {
        uint256 knownWorkflowToUpdate;
        bool isAdding;
    }

    struct UpdateKnownAddressesEntry {
        address addressToUpdate;
        bool isAdding;
    }

    function initalize(address _registryAddr, address _gatewayOwnerAddr) external;

    // perform is the entrypoint function of all customer contracts.
    // This function accepts the workflow ID and the end customer contract
    // execution payload.
    // The function checks that the given transaction is permitted and can be forwarded
    // next to the end customer contract.
    function perform(
        uint256 _id,
        address _target,
        bytes calldata _payload
    ) external;

    function setRegistry(address _registry) external;

    function updateConfigData(UpdateConfigData memory _updateConfigData) external;

    function updateKnownWorkflows(UpdateKnownWorkflowsEntry[] memory _updateKnownWorkflowsEntries) external;

    function updateKnownWorkflowOwners(UpdateKnownAddressesEntry[] memory _updateKnownWorkflowOwnersEntries) external;

    function updateKnownCustomerContracts(UpdateKnownAddressesEntry[] memory _updateKnownCustomerContractsEntries)
        external;

    function getConfigInfo() external view returns (ConfigInfo memory);
}
