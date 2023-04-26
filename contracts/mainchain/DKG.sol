// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./Staking.sol";
import "./ContractKeys.sol";
import "./ContractRegistry.sol";
import "./SlashingVoting.sol";

struct GenerationInfo {
    address signer;
    address[] validators;
    uint256 deadline;
    mapping(address => bool) isValidator;
    mapping(address => address) signerVotes;
    mapping(address => uint256) signerVoteCounts;
    mapping(uint256 => RoundData) roundData;
}

struct RoundData {
    uint256 count;
    mapping(address => bytes) data;
}

// DKG represents the on-sidechain logic needed to perform distributed key generation process done by validators.
// Once the DKG process is finished, a new collective/sender address could be stored.
contract DKG is ContractKeys, Initializable {
    using ECDSA for bytes;
    using ECDSA for bytes32;

    enum GenerationStatus {
        PENDING,
        EXPIRED,
        ACTIVE
    }

    ContractRegistry public contractRegistry;

    mapping(address => uint256) public signerToGeneration;

    GenerationInfo[] public generations;
    uint256 public lastActiveGeneration;
    uint256 public deadlinePeriod;

    event RoundDataProvided(uint256 generation, uint256 round, address validator);
    event RoundDataFilled(uint256 generation, uint256 round);

    event ValidatorsUpdated(uint256 generation, address[] validators);
    event SignerVoted(uint256 generation, address validator, address collectiveSigner);
    event SignerAddressUpdated(uint256 generation, address signerAddress);

    event ThresholdSignerUpdated(address signer);

    modifier onlyDKGValidator(uint256 _generation) {
        require(
            generations.length > _generation && generations[_generation].isValidator[msg.sender],
            "DKG: not a validator"
        );
        _;
    }

    modifier roundIsFilled(uint256 _generation, uint256 _round) {
        require(
            _round == 0 ||
                generations[_generation].roundData[_round].count == generations[_generation].validators.length,
            "DKG: round was not filled"
        );
        _;
    }

    modifier roundNotProvided(uint256 _generation, uint256 _round) {
        require(
            generations[_generation].roundData[_round].data[msg.sender].length == 0,
            "DKG: round data already provided"
        );
        _;
    }

    modifier onlySigner() {
        require(msg.sender == generations[lastActiveGeneration].signer, "DKG: not a active signer");
        _;
    }

    modifier onlyPending(uint256 _generation) {
        require(getStatus(_generation) == GenerationStatus.PENDING, "DKG: not a pending generation");
        _;
    }

    function initialize(address _contractRegistry, uint256 _deadlinePeriod) external initializer {
        generations.push();
        generations[0].signer = msg.sender;
        signerToGeneration[msg.sender] = 0;
        contractRegistry = ContractRegistry(_contractRegistry);
        deadlinePeriod = _deadlinePeriod;
    }

    function updateGeneration() external {
        uint256 newGeneration = generations.length;
        GenerationInfo storage oldGenerationInfo = generations[newGeneration - 1];

        uint256 validatorsCount = 0;
        bool newValidatorsAdded = false;
        address[] memory stakingValidators = _stakingContract().getValidators();
        address[] memory newValidators = new address[](stakingValidators.length);
        for (uint256 i = 0; i < stakingValidators.length; i++) {
            address validator = stakingValidators[i];

            // Validators banned by DKG reasons do not participate
            // in key generation for some time
            if (
                _isBannedByReason(validator, SlashingReason.REASON_DKG_INACTIVITY) ||
                _isBannedByReason(validator, SlashingReason.REASON_DKG_VIOLATION)
            ) {
                continue;
            }

            if (!oldGenerationInfo.isValidator[validator]) {
                newValidatorsAdded = true;
            }

            newValidators[validatorsCount] = validator;
            validatorsCount++;
        }

        uint256 oldValidatorsCount = oldGenerationInfo.validators.length;
        if (
            // Distributed key generation algorithm requires at least 2 participants
            validatorsCount < 2 ||
            // Validator count same as previous and there is no new validators,
            // meaning both arrays the same, no need to create new DKG generation
            (validatorsCount == oldValidatorsCount && !newValidatorsAdded)
        ) {
            return;
        }

        generations.push();
        for (uint256 i = 0; i < validatorsCount; i++) {
            generations[newGeneration].validators.push(newValidators[i]);
            generations[newGeneration].isValidator[newValidators[i]] = true;
        }

        generations[newGeneration].deadline = block.number + deadlinePeriod;
        lastActiveGeneration = newGeneration;

        emit ValidatorsUpdated(newGeneration, newValidators);
        emit RoundDataFilled(newGeneration, 0);
    }

    function roundBroadcast(
        uint256 _generation,
        uint256 _round,
        bytes memory _rawData
    )
        external
        onlyDKGValidator(_generation)
        roundIsFilled(_generation, _round - 1)
        roundNotProvided(_generation, _round)
        onlyPending(_generation)
    {
        generations[_generation].roundData[_round].count++;
        generations[_generation].roundData[_round].data[msg.sender] = _rawData;
        emit RoundDataProvided(_generation, _round, msg.sender);
        if (generations[_generation].roundData[_round].count == generations[_generation].validators.length) {
            emit RoundDataFilled(_generation, _round);
        }
    }

    function voteSigner(
        uint256 _generation,
        address _signerAddress,
        bytes memory _signature
    ) external onlyDKGValidator(_generation) roundIsFilled(_generation, 3) {
        GenerationInfo storage generationInfo = generations[_generation];
        require(generationInfo.deadline >= block.number, "DKG: voting is ended");

        address recoveredSigner = bytes("verify").toEthSignedMessageHash().recover(_signature);
        require(recoveredSigner == _signerAddress, "DKG: signature is invalid");

        require(generationInfo.signerVotes[msg.sender] == address(0), "DKG: already voted");

        generationInfo.signerVotes[msg.sender] = _signerAddress;
        generationInfo.signerVoteCounts[_signerAddress]++;

        emit SignerVoted(_generation, msg.sender, _signerAddress);

        bool enoughVotes = _enoughVotes(_generation, generationInfo.signerVoteCounts[_signerAddress]);
        bool signerChanged = generationInfo.signer != _signerAddress;
        if (enoughVotes && signerChanged) {
            generationInfo.signer = _signerAddress;
            signerToGeneration[_signerAddress] = _generation;
            emit SignerAddressUpdated(_generation, _signerAddress);
        }
    }

    function isRoundFilled(uint256 _generation, uint256 _round) external view returns (bool) {
        return generations[_generation].roundData[_round].count == generations[_generation].validators.length;
    }

    function getRoundBroadcastCount(uint256 _generation, uint256 _round) external view returns (uint256) {
        return generations[_generation].roundData[_round].count;
    }

    function getRoundBroadcastData(
        uint256 _generation,
        uint256 _round,
        address _validator
    ) external view returns (bytes memory) {
        return generations[_generation].roundData[_round].data[_validator];
    }

    function getCurrentValidators() external view returns (address[] memory) {
        return generations[generations.length - 1].validators;
    }

    function getSignerAddress() external view returns (address) {
        return generations[lastActiveGeneration].signer;
    }

    function getGenerationsCount() external view returns (uint256) {
        return generations.length;
    }

    function isCurrentValidator(address _validator) external view returns (bool) {
        return this.isValidator(lastActiveGeneration, _validator);
    }

    function isValidator(uint256 _generation, address _validator) external view returns (bool) {
        if (generations.length > _generation) {
            return generations[_generation].isValidator[_validator];
        }

        return false;
    }

    function getValidatorsCount(uint256 _generation) external view returns (uint256) {
        return generations[_generation].validators.length;
    }

    function setDeadlinePeriod(uint256 _deadlinePeriod) public onlySigner {
        deadlinePeriod = _deadlinePeriod;
    }

    function getStatus(uint256 _generation) public view returns (GenerationStatus) {
        if (generations[_generation].signer != address(0)) {
            return GenerationStatus.ACTIVE;
        }

        if (generations[_generation].deadline >= block.number) {
            return GenerationStatus.PENDING;
        }

        return GenerationStatus.EXPIRED;
    }

    function getValidators(uint256 _generation) public view returns (address[] memory) {
        return generations[_generation].validators;
    }

    function _enoughVotes(uint256 _generation, uint256 votes) private view returns (bool) {
        return votes > (generations[_generation].validators.length / 2);
    }

    function _stakingContract() private view returns (Staking) {
        return Staking(payable(contractRegistry.getContract(STAKING_KEY)));
    }

    function _slashingVotingContract() private view returns (SlashingVoting) {
        return SlashingVoting(contractRegistry.getContract(SLASHING_VOTING_KEY));
    }

    function _isBannedByReason(address _validator, SlashingReason _reason) private view returns (bool) {
        return _slashingVotingContract().isBannedByReason(_validator, _reason);
    }
}
