// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IContractsRegistry {
    function setIsMainChain(bool _newValue) external;

    function isMainChain() external view returns (bool);

    function getSigner() external view returns (address);

    function getDKGContract() external view returns (address);

    function getStakingContract() external view returns (address);

    function getSlashingVotingContract() external view returns (address);

    function getRewardsDistributionPoolContract() external view returns (address);

    function getBillingManagerContract() external view returns (address);

    function getRegistryContract() external view returns (address);

    function getGatewayFactoryContract() external view returns (address);

    function getNerifTokenContract() external view returns (address);

    function getTokensVestingContract() external view returns (address);

    function getSignerGetterContract() external view returns (address);
}
