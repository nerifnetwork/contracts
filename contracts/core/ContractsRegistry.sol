// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractContractsRegistry.sol";

import "../interfaces/SignerOwnable.sol";
import "../interfaces/core/IContractsRegistry.sol";

contract ContractsRegistry is IContractsRegistry, AbstractContractsRegistry, SignerOwnable, UUPSUpgradeable {
    string public constant DKG_NAME = "DKG";
    string public constant STAKING_NAME = "STAKING";
    string public constant SLASHING_VOTING_NAME = "SLASHING_VOTING";
    string public constant REWARDS_DISTRIBUTION_POOL_NAME = "REWARDS_DISTRIBUTION_POOL";

    string public constant BILLING_MANAGER_NAME = "BILLING_MANAGER";
    string public constant REGISTRY_NAME = "REGISTRY";
    string public constant GATEWAY_FACTORY_NAME = "GATEWAY_FACTORY";

    string public constant NERIF_TOKEN_NAME = "NERIF_TOKEN";
    string public constant TOKENS_VESTING_NAME = "TOKENS_VESTING";

    bool public override isMainChain;

    function initialize(address _signerGetterAddress, bool _isMainChain) external initializer {
        __ContractsRegistry_init();

        _setSignerGetter(_signerGetterAddress);
        isMainChain = _isMainChain;
    }

    function injectDependencies(string calldata _name) external override onlySigner {
        _injectDependencies(_name);
    }

    function injectDependenciesWithData(string calldata _name, bytes calldata _data) external override onlySigner {
        _injectDependenciesWithData(_name, _data);
    }

    function upgradeContract(string calldata _name, address _newImplementation) external override onlySigner {
        _upgradeContract(_name, _newImplementation);
    }

    function upgradeContractAndCall(
        string calldata _name,
        address _newImplementation,
        bytes calldata _data
    ) external override onlySigner {
        _upgradeContractAndCall(_name, _newImplementation, _data);
    }

    function addContract(string calldata _name, address _contractAddress) external override onlySigner {
        _addContract(_name, _contractAddress);
    }

    function addProxyContract(string calldata _name, address _contractAddress) external override onlySigner {
        _addProxyContract(_name, _contractAddress);
    }

    function addProxyContractAndCall(
        string calldata _name,
        address _contractAddress,
        bytes calldata _data
    ) external override onlySigner {
        _addProxyContractAndCall(_name, _contractAddress, _data);
    }

    function justAddProxyContract(string calldata _name, address _contractAddress) external override onlySigner {
        _justAddProxyContract(_name, _contractAddress);
    }

    function removeContract(string calldata _name) external override onlySigner {
        _removeContract(_name);
    }

    function getSigner() external view override returns (address) {
        return signerGetter.getSignerAddress();
    }

    function getDKGContract() external view override returns (address) {
        return getContract(DKG_NAME);
    }

    function getStakingContract() external view override returns (address) {
        return getContract(STAKING_NAME);
    }

    function getSlashingVotingContract() external view override returns (address) {
        return getContract(SLASHING_VOTING_NAME);
    }

    function getRewardsDistributionPoolContract() external view override returns (address) {
        return getContract(REWARDS_DISTRIBUTION_POOL_NAME);
    }

    function getBillingManagerContract() external view override returns (address) {
        return getContract(BILLING_MANAGER_NAME);
    }

    function getRegistryContract() external view override returns (address) {
        return getContract(REGISTRY_NAME);
    }

    function getGatewayFactoryContract() external view override returns (address) {
        return getContract(GATEWAY_FACTORY_NAME);
    }

    function getNerifTokenContract() external view override returns (address) {
        return getContract(NERIF_TOKEN_NAME);
    }

    function getTokensVestingContract() external view override returns (address) {
        return getContract(TOKENS_VESTING_NAME);
    }

    function _authorizeUpgrade(address) internal override onlySigner {}
}
