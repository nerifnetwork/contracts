// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

/**
 * @notice Interface for accessing addresses of various contracts registered in the system
 */
interface IContractsRegistry {
    /**
     * @notice Sets whether the current chain is the main chain
     * @param _newValue The new value indicating whether the current chain is the main chain
     */
    function setIsMainChain(bool _newValue) external;

    /**
     * @notice Retrieves whether the current chain is the main chain
     * @return True if the current chain is the main chain, otherwise false
     */
    function isMainChain() external view returns (bool);

    /**
     * @notice Retrieves the address of the system signer contract
     * @return The address of the system signer contract
     */
    function getSigner() external view returns (address);

    /**
     * @notice Retrieves the address of the DKG (Distributed Key Generation) contract
     * @return The address of the DKG contract
     */
    function getDKGContract() external view returns (address);

    /**
     * @notice Retrieves the address of the staking contract
     * @return The address of the staking contract
     */
    function getStakingContract() external view returns (address);

    /**
     * @notice Retrieves the address of the slashing voting contract
     * @return The address of the slashing voting contract
     */
    function getSlashingVotingContract() external view returns (address);

    /**
     * @notice Retrieves the address of the billing manager contract
     * @return The address of the billing manager contract
     */
    function getBillingManagerContract() external view returns (address);

    /**
     * @notice Retrieves the address of the registry contract
     * @return The address of the registry contract
     */
    function getRegistryContract() external view returns (address);

    /**
     * @notice Retrieves the address of the gateway factory contract
     * @return The address of the gateway factory contract
     */
    function getGatewayFactoryContract() external view returns (address);

    /**
     * @notice Retrieves the address of the Nerif token contract
     * @return The address of the Nerif token contract
     */
    function getNerifTokenContract() external view returns (address);

    /**
     * @notice Retrieves the address of the tokens vesting contract
     * @return The address of the tokens vesting contract
     */
    function getTokensVestingContract() external view returns (address);

    /**
     * @notice Retrieves the address of the signer getter contract
     * @return The address of the signer getter contract
     */
    function getSignerGetterContract() external view returns (address);
}
