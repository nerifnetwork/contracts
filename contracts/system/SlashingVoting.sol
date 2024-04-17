// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/utils/Globals.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/system/IDKG.sol";
import "../interfaces/system/IStaking.sol";
import "../interfaces/system/ISlashingVoting.sol";

contract SlashingVoting is ISlashingVoting, Initializable, AbstractDependant {
    using EnumerableSet for EnumerableSet.AddressSet;

    IContractsRegistry internal _contractsRegistry;
    IStaking internal _staking;
    IDKG internal _dkg;

    uint256 public lastProposalId;

    uint256 public votingThresholdPercentage;

    mapping(uint256 => SlashingProposalData) internal _proposalsData;
    mapping(address => uint256) internal _pendingSlashingProposals;

    modifier onlySigner() {
        _onlySigner();
        _;
    }

    modifier onlyValidator() {
        _onlyValidator();
        _;
    }

    function initialize(uint256 _votingThresholdPercentage) external initializer {
        votingThresholdPercentage = _votingThresholdPercentage;
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _contractsRegistry = contractsRegistry;
        _staking = IStaking(contractsRegistry.getStakingContract());
        _dkg = IDKG(contractsRegistry.getDKGContract());
    }

    // solhint-disable-next-line ordering
    function setVotingThresholdPercentage(uint256 _votingThresholdPercentage) external onlySigner {
        votingThresholdPercentage = _votingThresholdPercentage;
    }

    function createProposal(address _validatorAddr, string calldata _reason) external onlyValidator returns (uint256) {
        require(_dkg.isActiveValidator(_validatorAddr), "SlashingVoting: Target is not active validator");
        require(!hasPendingSlashingProposal(_validatorAddr), "SlashingVoting: Validator already has pending proposal");

        uint256 newProposalId = ++lastProposalId;

        SlashingProposalData storage proposalData = _proposalsData[newProposalId];

        proposalData.validator = _validatorAddr;
        proposalData.reason = _reason;

        IDKG.MainEpochInfo memory epochInfo = _dkg.createProposal();

        proposalData.epochId = epochInfo.epochId;
        proposalData.votingStartTime = epochInfo.epochStartTime;
        proposalData.votingEndTime = epochInfo.dkgGenPeriodEndTime;

        _pendingSlashingProposals[_validatorAddr] = newProposalId;

        if (epochInfo.epochStartTime <= block.timestamp) {
            _vote(newProposalId);
        }

        return newProposalId;
    }

    function vote(uint256 _proposalId) external onlyValidator {
        _vote(_proposalId);
    }

    function hasPendingSlashingProposal(address _userAddr) public view returns (bool) {
        return _pendingSlashingProposals[_userAddr] != 0;
    }

    function _vote(uint256 _proposalId) internal {
        SlashingProposalData storage proposalData = _proposalsData[_proposalId];

        require(proposalData.votingStartTime <= block.timestamp, "SlashingVoting: Voting hasn't started yet");
        require(proposalData.votedValidatorsSet.add(msg.sender), "SlashingVoting: Validator has already voted");

        uint256 currentVotesCount = ++proposalData.slashingProposalVotesCount;

        if (_getValidatorsPercentage(currentVotesCount) >= votingThresholdPercentage) {
            address validatorAddr = proposalData.validator;

            _dkg.slashValidator(validatorAddr);
            _staking.slashValidator(validatorAddr);

            delete _pendingSlashingProposals[validatorAddr];

            emit ProposalExecuted(_proposalId, validatorAddr);
        }

        emit ProposalVoted(_proposalId, msg.sender);
    }

    function _getValidatorsPercentage(uint256 _currentVotes) internal view returns (uint256) {
        return (_currentVotes * PERCENTAGE_100) / _dkg.getActiveValidatorsCount();
    }

    function _onlySigner() internal view {
        require(_contractsRegistry.getSigner() == msg.sender, "SlashingVoting: Not a signer");
    }

    function _onlyValidator() internal view {
        require(_dkg.isActiveValidator(msg.sender), "SlashingVoting: Not an active system validator");
    }
}
