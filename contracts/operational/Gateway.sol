// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "../interfaces/IGateway.sol";
import "../interfaces/IRegistry.sol";

import "./Registry.sol";

contract Gateway is IGateway, OwnableUpgradeable {
    using EnumerableSet for *;

    Registry public registry;

    EnumerableSet.UintSet internal _knownWorkflows;
    EnumerableSet.AddressSet internal _knownWorkflowOwners;
    EnumerableSet.AddressSet internal _knownCustomerContracts;

    // onlyRegistry permits transactions coming from the registry contract.
    modifier onlyRegistry() {
        require(address(registry) == msg.sender, "Gateway: operation is not permitted");
        _;
    }

    // onlyAllowedWorkflow allows transactions passed through checks based on config.
    modifier onlyAllowedWorkflow(uint256 _id, address _target) {
        _onlyAllowedWorkflow(_id, _target);
        _;
    }

    function initialize(address _registryAddr, address _gatewayOwnerAddr) external initializer {
        __Ownable_init();

        registry = Registry(_registryAddr);
        transferOwnership(_gatewayOwnerAddr);
    }

    function perform(
        uint256 _id,
        address _target,
        bytes calldata _payload
    ) external override onlyRegistry onlyAllowedWorkflow(_id, _target) {
        (bool success, ) = _target.call(_payload);
        require(success, "Gateway: failed to execute customer contract");
    }

    function setRegistry(address _registry) external override onlyOwner {
        registry = Registry(_registry);
    }

    function updateConfigData(UpdateConfigData memory _updateConfigData) external override onlyOwner {
        _updateKnownWorkflows(_updateConfigData.updateKnownWorkflowsEntries);
        _updateKnownAddresses(_knownWorkflowOwners, _updateConfigData.updateKnownWorkflowOwnersEntries);
        _updateKnownAddresses(_knownCustomerContracts, _updateConfigData.updateKnownCustomerContractsEntries);
    }

    function updateKnownWorkflows(UpdateKnownWorkflowsEntry[] memory _updateKnownWorkflowsEntries)
        external
        override
        onlyOwner
    {
        _updateKnownWorkflows(_updateKnownWorkflowsEntries);
    }

    function updateKnownWorkflowOwners(UpdateKnownAddressesEntry[] memory _updateKnownWorkflowOwnersEntries)
        external
        override
        onlyOwner
    {
        _updateKnownAddresses(_knownWorkflowOwners, _updateKnownWorkflowOwnersEntries);
    }

    function updateKnownCustomerContracts(UpdateKnownAddressesEntry[] memory _updateKnownCustomerContractsEntries)
        external
        override
        onlyOwner
    {
        _updateKnownAddresses(_knownCustomerContracts, _updateKnownCustomerContractsEntries);
    }

    function getConfigInfo() external view override returns (ConfigInfo memory) {
        return ConfigInfo(_knownWorkflows.values(), _knownWorkflowOwners.values(), _knownCustomerContracts.values());
    }

    function _updateKnownWorkflows(UpdateKnownWorkflowsEntry[] memory _updateKnownWorkflowsEntries) internal {
        for (uint256 i = 0; i < _updateKnownWorkflowsEntries.length; i++) {
            if (_updateKnownWorkflowsEntries[i].isAdding) {
                _knownWorkflows.add(_updateKnownWorkflowsEntries[i].knownWorkflowToUpdate);
            } else {
                _knownWorkflows.remove(_updateKnownWorkflowsEntries[i].knownWorkflowToUpdate);
            }
        }
    }

    function _updateKnownAddresses(
        EnumerableSet.AddressSet storage _addressSet,
        UpdateKnownAddressesEntry[] memory _knownAddressesEntries
    ) internal {
        for (uint256 i = 0; i < _knownAddressesEntries.length; i++) {
            if (_knownAddressesEntries[i].isAdding) {
                _addressSet.add(_knownAddressesEntries[i].addressToUpdate);
            } else {
                _addressSet.remove(_knownAddressesEntries[i].addressToUpdate);
            }
        }
    }

    function _onlyAllowedWorkflow(uint256 _id, address _target) internal view {
        address workflowOwner = registry.getWorkflowOwner(_id);

        // Gateway owner's workflows can be executed anytime
        bool allowed = owner() == workflowOwner;

        if (
            !allowed &&
            (_knownWorkflowOwners.contains(workflowOwner) ||
                _knownWorkflows.contains(_id) ||
                _knownCustomerContracts.contains(_target))
        ) {
            allowed = true;
        }

        require(allowed, "Gateway: is not allowed workflow");
    }
}
