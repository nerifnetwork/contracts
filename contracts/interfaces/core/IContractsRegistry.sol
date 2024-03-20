// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IContractsRegistry {
    function injectDependencies(string calldata _name) external;

    function injectDependenciesWithData(string calldata _name, bytes calldata _data) external;

    function upgradeContract(string calldata _name, address _newImplementation) external;

    function upgradeContractAndCall(string calldata _name, address _newImplementation, bytes calldata _data) external;

    function addContract(string calldata _name, address _contractAddress) external;

    function addProxyContract(string calldata _name, address _contractAddress) external;

    function addProxyContractAndCall(string calldata _name, address _contractAddress, bytes calldata _data) external;

    function justAddProxyContract(string calldata _name, address _contractAddress) external;

    function removeContract(string calldata _name) external;

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
}
