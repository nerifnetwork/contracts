// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// ContractKeys keeps all contract keys/names that can be used to store or get the contract address
// from the contract registry.
abstract contract ContractKeys {
    string public constant STAKING_KEY = "staking";
    string public constant DKG_KEY = "dkg";
    string public constant SUPPORTED_TOKENS_KEY = "supported-tokens";
    string public constant SLASHING_VOTING_KEY = "slashing-voting";
    string public constant REWARD_DISTRIBUTION_POOL_KEY = "reward-distribution-pool";
}
