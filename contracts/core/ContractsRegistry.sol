// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

import "@solarity/solidity-lib/contracts-registry/presets/OwnableContractsRegistry.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/ISignerAddress.sol";

contract ContractsRegistry is IContractsRegistry, OwnableContractsRegistry, UUPSUpgradeable {
    string public constant DKG_NAME = "DKG";
    string public constant STAKING_NAME = "STAKING";
    string public constant SLASHING_VOTING_NAME = "SLASHING_VOTING";
    string public constant REWARDS_DISTRIBUTION_POOL_NAME = "REWARDS_DISTRIBUTION_POOL";

    string public constant BILLING_MANAGER_NAME = "BILLING_MANAGER";
    string public constant REGISTRY_NAME = "REGISTRY";
    string public constant GATEWAY_FACTORY_NAME = "GATEWAY_FACTORY";

    string public constant NERIF_TOKEN_NAME = "NERIF_TOKEN";
    string public constant TOKENS_VESTING_NAME = "TOKENS_VESTING";

    string public constant SIGNER_GETTER_NAME = "SIGNER_GETTER";

    bool public override isMainChain;

    function setIsMainChain(bool _newValue) external onlyOwner {
        isMainChain = _newValue;
    }

    function getSigner() external view override returns (address _signer) {
        if (hasContract(SIGNER_GETTER_NAME)) {
            _signer = ISignerAddress(getContract(SIGNER_GETTER_NAME)).getSignerAddress();
        }
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

    function getSignerGetterContract() external view override returns (address) {
        return getContract(SIGNER_GETTER_NAME);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
