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

    modifier onlyActiveValidator() {
        _onlyActiveValidator();
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
    function setVotingThresholdPercentage(uint256 _votingThresholdPercentage) external override onlySigner {
        votingThresholdPercentage = _votingThresholdPercentage;
    }

    function createProposal(
        address _validatorAddr,
        string calldata _reason
    ) external override onlyActiveValidator returns (uint256) {
        require(_dkg.isActiveValidator(_validatorAddr), "SlashingVoting: Target is not an active validator");
        require(!hasPendingSlashingProposal(_validatorAddr), "SlashingVoting: Validator already has pending proposal");

        uint256 newProposalId = ++lastProposalId;

        SlashingProposalData storage proposalData = _proposalsData[newProposalId];

        proposalData.validator = _validatorAddr;
        proposalData.reason = _reason;

        proposalData.epochId = _dkg.getActiveEpochId();
        proposalData.votingStartTime = block.timestamp;
        proposalData.votingEndTime = _dkg.createProposal();

        _pendingSlashingProposals[_validatorAddr] = newProposalId;

        _vote(newProposalId);

        emit ProposalCreated(newProposalId, _validatorAddr);

        return newProposalId;
    }

    function vote(uint256 _proposalId) external override onlyActiveValidator {
        _vote(_proposalId);
    }

    function getDetailedProposalInfo(uint256 _proposalId) external view override returns (DetailedProposalInfo memory) {
        SlashingProposalData storage proposalData = _proposalsData[_proposalId];

        return
            DetailedProposalInfo(
                getBaseProposalInfo(_proposalId),
                proposalData.isExecuted,
                proposalData.votedValidatorsSet.values()
            );
    }

    function getBaseProposalInfo(uint256 _proposalId) public view override returns (BaseProposalInfo memory) {
        SlashingProposalData storage proposalData = _proposalsData[_proposalId];

        return
            BaseProposalInfo(
                proposalData.validator,
                proposalData.reason,
                proposalData.epochId,
                proposalData.votingStartTime,
                proposalData.votingEndTime
            );
    }

    function isProposalExecuted(uint256 _proposalId) public view override returns (bool) {
        return _proposalsData[_proposalId].isExecuted;
    }

    function hasPendingSlashingProposal(address _userAddr) public view override returns (bool) {
        uint256 pendingProposalId = _pendingSlashingProposals[_userAddr];

        return pendingProposalId != 0 && _proposalsData[pendingProposalId].votingEndTime > block.timestamp;
    }

    function getValidatorsPercentage(uint256 _currentVotes) public view override returns (uint256) {
        return (_currentVotes * PERCENTAGE_100) / _dkg.getActiveValidatorsCount();
    }

    function _vote(uint256 _proposalId) internal {
        SlashingProposalData storage proposalData = _proposalsData[_proposalId];

        require(!proposalData.isExecuted, "SlashingVoting: Proposal has already executed");
        require(
            proposalData.votingStartTime <= block.timestamp && block.timestamp < proposalData.votingEndTime,
            "SlashingVoting: Voting hasn't started yet or finished"
        );
        require(proposalData.votedValidatorsSet.add(msg.sender), "SlashingVoting: Validator has already voted");

        if (getValidatorsPercentage(proposalData.votedValidatorsSet.length()) >= votingThresholdPercentage) {
            address validatorAddr = proposalData.validator;

            _dkg.slashValidator(validatorAddr);
            _staking.slashValidator(validatorAddr);

            proposalData.isExecuted = true;
            delete _pendingSlashingProposals[validatorAddr];

            emit ProposalExecuted(_proposalId, validatorAddr);
        }

        emit ProposalVoted(_proposalId, msg.sender);
    }

    function _onlySigner() internal view {
        require(_contractsRegistry.getSigner() == msg.sender, "SlashingVoting: Not a signer");
    }

    function _onlyActiveValidator() internal view {
        require(_dkg.isActiveValidator(msg.sender), "SlashingVoting: Not an active system validator");
    }
}
