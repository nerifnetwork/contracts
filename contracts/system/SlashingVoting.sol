// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/SignerOwnable.sol";
import "../interfaces/system/IDKG.sol";

import "./Staking.sol";

enum SlashingReason {
    REASON_NO_RECENT_BLOCKS,
    REASON_DKG_INACTIVITY,
    REASON_DKG_VIOLATION,
    REASON_SIGNING_INACTIVITY,
    REASON_SIGNING_VIOLATION
}

enum SlashingReasonGroup {
    NONE,
    REASON_GROUP_BLOCKS,
    REASON_GROUP_DKG,
    REASON_GROUP_SIGNING
}

// SlashingVoting represents the validator slashing mechanism.
// It allows validators to vote for slashing of a specific validator.
contract SlashingVoting is Initializable, AbstractDependant {
    struct SlashingProposal {
        address validator;
        string reason;
        mapping(address => bool) slashingProposalVotes;
        uint256 slashingProposalVoteCounts;
    }

    IContractsRegistry internal _contractsRegistry;
    Staking internal _staking;
    IDKG internal _dkg;

    SlashingProposal[] public proposals;

    uint256 public epochPeriod;
    uint256 public slashingThresold;
    uint256 public slashingEpochs;

    // Votes
    mapping(bytes32 => mapping(address => bool)) public votes;
    mapping(bytes32 => uint256) public voteCounts;

    // Bans
    mapping(bytes32 => bool) public bans;
    mapping(uint256 => mapping(address => mapping(SlashingReason => bool))) public bansByReason;
    mapping(uint256 => mapping(address => uint256)) public bansByEpoch;
    mapping(uint256 => mapping(SlashingReason => address[])) public bannedValidators;

    event VotedWithReason(address voter, address validator, SlashingReason reason);
    event BannedWithReason(address validator, SlashingReason reason);
    event SlashedWithReason(address validator);

    event ProposalCreated(uint256 proposalId, address validator);
    event ProposalVoted(uint256 proposalId, address validator, address voter);
    event ProposalExecuted(uint256 proposalId, address validator);

    modifier onlySigner() {
        _onlySigner();
        _;
    }

    modifier onlyValidator() {
        _onlyValidator();
        _;
    }

    function initialize(uint256 _epochPeriod, uint256 _slashingThresold, uint256 _slashingEpochs) external initializer {
        epochPeriod = _epochPeriod;
        slashingThresold = _slashingThresold;
        slashingEpochs = _slashingEpochs;
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _contractsRegistry = contractsRegistry;
        _staking = Staking(contractsRegistry.getStakingContract());
        _dkg = IDKG(contractsRegistry.getDKGContract());
    }

    // solhint-disable-next-line ordering
    function voteWithReason(address _validator, SlashingReason _reason, bytes calldata _nonce) external onlyValidator {
        bytes32 voteHash = votingHashWithReason(_validator, _reason, _nonce);

        require(_staking.isValidatorActive(_validator) == true, "SlashingVoting: target is not active validator");
        require(bans[voteHash] == false, "SlashingVoting: validator is already banned");
        require(votes[voteHash][msg.sender] == false, "SlashingVoting: voter is already voted against given validator");

        votes[voteHash][msg.sender] = true;
        voteCounts[voteHash]++;
        emit VotedWithReason(msg.sender, _validator, _reason);

        uint256 epoch = currentEpoch();
        address[] memory validators = _staking.getValidators();
        if (voteCounts[voteHash] >= (validators.length / 2 + 1)) {
            bans[voteHash] = true;
            bansByReason[epoch][_validator][_reason] = true;
            bansByEpoch[epoch][_validator]++;
            bannedValidators[epoch][_reason].push(_validator);
            emit BannedWithReason(_validator, _reason);

            if (_reason == SlashingReason.REASON_DKG_INACTIVITY || _reason == SlashingReason.REASON_DKG_VIOLATION) {
                _dkg.updateActiveValidators();
            }
        }

        if (shouldShash(epoch, _validator)) {
            _staking.slash(_validator);
            emit SlashedWithReason(_validator);
        }
    }

    function createProposal(address _validator, string memory _reason) external onlyValidator {
        SlashingProposal storage newProposal = proposals.push();

        newProposal.validator = _validator;
        newProposal.reason = _reason;

        uint256 proposalId = proposals.length - 1;
        emit ProposalCreated(proposalId, _validator);

        voteProposal(proposalId);
    }

    function voteProposal(uint256 _proposalId) public onlyValidator {
        require(_proposalId < proposals.length, "SlashingVoting: proposal doesn't exist!");

        SlashingProposal storage proposal = proposals[_proposalId];

        require(
            _staking.isValidatorActive(proposal.validator) == true,
            "SlashingVoting: target is not active validator"
        );
        require(
            proposals[_proposalId].slashingProposalVotes[msg.sender] == false,
            "SlashingVoting: you already voted in this proposal"
        );

        proposals[_proposalId].slashingProposalVotes[msg.sender] = true;
        proposals[_proposalId].slashingProposalVoteCounts++;

        address[] memory validators = _staking.getValidators();
        if (proposals[_proposalId].slashingProposalVoteCounts >= (validators.length / 2 + 1)) {
            _staking.slash(proposal.validator);
            emit ProposalExecuted(_proposalId, proposal.validator);
        }
        emit ProposalVoted(_proposalId, proposal.validator, msg.sender);
    }

    function setEpochPeriod(uint256 _epochPeriod) public onlySigner {
        epochPeriod = _epochPeriod;
    }

    function setSlashingThresold(uint256 _slashingThresold) public onlySigner {
        slashingThresold = _slashingThresold;
    }

    function setSlashingEpochs(uint256 _slashingEpochs) public onlySigner {
        slashingEpochs = _slashingEpochs;
    }

    function isBannedByReason(address _validator, SlashingReason _reason) public view returns (bool) {
        return bansByReason[currentEpoch()][_validator][_reason];
    }

    function shouldShash(uint256 _epoch, address _validator) public view returns (bool) {
        if (_epoch < slashingEpochs) {
            return false;
        }

        uint256 totalBans;
        for (uint256 i = _epoch; i > _epoch - slashingEpochs; i--) {
            uint256 bansInEpoch = bansByEpoch[i][_validator];
            if (bansInEpoch == 0) {
                return false;
            }

            totalBans += bansInEpoch;
        }

        return totalBans >= slashingThresold;
    }

    function getBansByEpoch(uint256 _epoch, address _validator) public view returns (uint256) {
        return bansByEpoch[_epoch][_validator];
    }

    function getBannedValidatorsByReason(SlashingReason _reason) public view returns (address[] memory) {
        return bannedValidators[currentEpoch()][_reason];
    }

    function currentEpoch() public view returns (uint256) {
        return epochByBlock(block.number);
    }

    function epochByBlock(uint256 _blockNumber) public view returns (uint256) {
        return _blockNumber / epochPeriod;
    }

    function votingHashWithReason(
        address _validator,
        SlashingReason _reason,
        bytes calldata _nonce
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_validator, _reason, _nonce));
    }

    function _onlySigner() internal view {
        require(_contractsRegistry.getSigner() == msg.sender, "SlashingVoting: Not a signer");
    }

    function _onlyValidator() internal view {
        require(_staking.isValidatorActive(msg.sender), "SlashingVoting: Not a system validator");
    }
}
