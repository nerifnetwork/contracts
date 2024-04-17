// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface ISlashingVoting {
    struct BaseProposalInfo {
        address validator;
        string reason;
        uint256 epochId;
        uint256 votingStartTime;
        uint256 votingEndTime;
    }

    struct DetailedProposalInfo {
        BaseProposalInfo baseProposalInfo;
        bool isExecuted;
        address[] votedValidators;
    }

    struct SlashingProposalData {
        address validator;
        string reason;
        uint256 epochId;
        uint256 votingStartTime;
        uint256 votingEndTime;
        bool isExecuted;
        EnumerableSet.AddressSet votedValidatorsSet;
    }

    event ProposalCreated(uint256 indexed proposalId, address validator);
    event ProposalVoted(uint256 indexed proposalId, address indexed voter);
    event ProposalExecuted(uint256 proposalId, address validator);

    function setVotingThresholdPercentage(uint256 _votingThresholdPercentage) external;

    function createProposal(address _validatorAddr, string calldata _reason) external returns (uint256);

    function vote(uint256 _proposalId) external;

    function getDetailedProposalInfo(uint256 _proposalId) external view returns (DetailedProposalInfo memory);

    function getBaseProposalInfo(uint256 _proposalId) external view returns (BaseProposalInfo memory);

    function isProposalExecuted(uint256 _proposalId) external view returns (bool);

    function hasPendingSlashingProposal(address _userAddr) external view returns (bool);

    function getValidatorsPercentage(uint256 _currentVotes) external view returns (uint256);
}
