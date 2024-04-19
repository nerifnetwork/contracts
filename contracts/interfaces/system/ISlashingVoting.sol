// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @notice Interface for managing slashing proposals and voting
 */
interface ISlashingVoting {
    /**
     * @notice Represents the base information of a proposal
     * @param validator The address of the validator whose slashing will be voted for
     * @param reason The reason for the proposal
     * @param epochId The epoch ID associated with the proposal
     * @param votingStartTime The start time of the voting period for the proposal
     * @param votingEndTime The end time of the voting period for the proposal
     */
    struct BaseProposalInfo {
        address validator;
        string reason;
        uint256 epochId;
        uint256 votingStartTime;
        uint256 votingEndTime;
    }

    /**
     * @notice Represents detailed information of a proposal
     * @param baseProposalInfo The base information of the proposal
     * @param isExecuted The boolean indicating whether the proposal has been executed
     * @param votedValidators The array containing the addresses of validators who voted for the proposal
     */
    struct DetailedProposalInfo {
        BaseProposalInfo baseProposalInfo;
        bool isExecuted;
        address[] votedValidators;
    }

    /**
     * @notice Represents data associated with a slashing proposal
     * @param validator The address of the validator whose slashing will be voted for
     * @param reason The reason for the slashing proposal
     * @param epochId The epoch ID associated with the slashing proposal
     * @param votingStartTime The start time of the voting period for the slashing proposal
     * @param votingEndTime The end time of the voting period for the slashing proposal
     * @param isExecuted The boolean indicating whether the slashing proposal has been executed
     * @param votedValidatorsSet The set containing the addresses of validators who voted for the slashing proposal
     */
    struct SlashingProposalData {
        address validator;
        string reason;
        uint256 epochId;
        uint256 votingStartTime;
        uint256 votingEndTime;
        bool isExecuted;
        EnumerableSet.AddressSet votedValidatorsSet;
    }

    /**
     * @notice Emitted when a new proposal is created
     * @param proposalId The ID of the newly created proposal
     * @param validator The address of the validator whose slashing will be voted for
     */
    event ProposalCreated(uint256 indexed proposalId, address validator);

    /**
     * @notice Emitted when a vote is cast for a proposal
     * @param proposalId The ID of the proposal being voted on
     * @param voter The address of the validator who cast the vote
     */
    event ProposalVoted(uint256 indexed proposalId, address indexed voter);

    /**
     * @notice Emitted when a proposal is executed
     * @param proposalId The ID of the executed proposal
     * @param validator The address of the slashed validator
     */
    event ProposalExecuted(uint256 proposalId, address validator);

    /**
     * @notice Sets the threshold percentage required for a proposal to be executed
     * @param _votingThresholdPercentage The threshold percentage for proposal execution
     */
    function setVotingThresholdPercentage(uint256 _votingThresholdPercentage) external;

    /**
     * @notice Creates a new proposal
     * @param _validatorAddr The address of the validator whose slashing will be voted for
     * @param _reason The reason for the proposal
     * @return The ID of the newly created proposal
     */
    function createProposal(address _validatorAddr, string calldata _reason) external returns (uint256);

    /**
     * @notice Casts a vote for a proposal
     * @param _proposalId The ID of the proposal to vote on
     */
    function vote(uint256 _proposalId) external;

    /**
     * @notice Retrieves detailed information about a proposal
     * @param _proposalId The ID of the proposal
     * @return Detailed information about the proposal
     */
    function getDetailedProposalInfo(uint256 _proposalId) external view returns (DetailedProposalInfo memory);

    /**
     * @notice Retrieves the base information of a proposal
     * @param _proposalId The ID of the proposal
     * @return The base information of the proposal
     */
    function getBaseProposalInfo(uint256 _proposalId) external view returns (BaseProposalInfo memory);

    /**
     * @notice Checks if a proposal has been executed
     * @param _proposalId The ID of the proposal
     * @return True if the proposal has been executed, false otherwise
     */
    function isProposalExecuted(uint256 _proposalId) external view returns (bool);

    /**
     * @notice Checks if a user has pending slashing proposals
     * @param _userAddr The address of the user
     * @return True if the user has pending slashing proposals, false otherwise
     */
    function hasPendingSlashingProposal(address _userAddr) external view returns (bool);

    /**
     * @notice Calculates the percentage of validators based on the current votes count
     * @param _currentVotes The current number of votes
     * @return The percentage of validators
     */
    function getValidatorsPercentage(uint256 _currentVotes) external view returns (uint256);
}
