// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/libs/data-structures/memory/Vector.sol";

import "../interfaces/core/ISignerAddress.sol";
import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/system/IDKG.sol";
import "../interfaces/system/IStaking.sol";
import "../interfaces/system/ISlashingVoting.sol";

contract DKG is IDKG, Initializable, AbstractDependant {
    using EnumerableSet for EnumerableSet.AddressSet;
    using Vector for Vector.AddressVector;
    using ECDSA for *;

    uint256 public constant FIRST_EPOCH_ID = 1;

    IStaking internal _staking;
    ISlashingVoting internal _slashingVoting;

    uint256 public updatesCollectionEpochDuration;
    uint256 public dkgGenerationEpochDuration;
    uint256 public guaranteedWorkingEpochDuration;

    uint256 internal _activeEpochId;
    address internal _activeSigner;

    EnumerableSet.AddressSet internal _validators;

    mapping(uint256 => DKGEpochData) internal _dkgEpochsData;
    mapping(address => ValidatorData) internal _validatorsData;
    mapping(address => bool) internal _slashedValidators;

    modifier onlyActiveValidator() {
        _onlyActiveValidator();
        _;
    }

    modifier onlyStaking() {
        _onlyStaking();
        _;
    }

    modifier onlySlashingVoting() {
        _onlySlashingVoting();
        _;
    }

    function initialize(
        uint256 _updatesCollectionEpochDuration,
        uint256 _dkgGenerationEpochDuration,
        uint256 _guaranteedWorkingEpochDuration
    ) external initializer {
        updatesCollectionEpochDuration = _updatesCollectionEpochDuration;
        dkgGenerationEpochDuration = _dkgGenerationEpochDuration;
        guaranteedWorkingEpochDuration = _guaranteedWorkingEpochDuration;

        _createNewEpoch();
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _staking = IStaking(contractsRegistry.getStakingContract());
        _slashingVoting = ISlashingVoting(contractsRegistry.getSlashingVotingContract());
    }

    // solhint-disable-next-line ordering
    function addValidator(address _validatorToAdd) external override onlyStaking {
        require(!isValidator(_validatorToAdd), "DKG: Validator already exists");

        uint256 startValidationEpoch = _collectUpdates();

        _validatorsData[_validatorToAdd] = ValidatorData(startValidationEpoch, type(uint256).max);
        _validators.add(_validatorToAdd);

        emit NewValidatorAdded(_validatorToAdd, startValidationEpoch);
    }

    function announceValidatorExit(
        address _validatorAddr
    ) external override onlyStaking returns (uint256 _endValidationEpoch) {
        require(isActiveValidator(_validatorAddr), "DKG: Validator is not active");
        require(
            _validatorsData[_validatorAddr].endValidationEpochId == type(uint256).max,
            "DKG: Exit of the validator has already been announced"
        );

        _endValidationEpoch = _collectUpdates();

        _validatorsData[_validatorAddr].endValidationEpochId = _endValidationEpoch;

        emit ValidatorExitAnnounced(_validatorAddr, _endValidationEpoch);
    }

    function createProposal() external override onlySlashingVoting returns (uint256) {
        _collectUpdates();

        return _dkgEpochsData[_activeEpochId].dkgGenPeriodStartTime;
    }

    function slashValidator(address _validatorAddr) external override onlySlashingVoting {
        require(!_slashedValidators[_validatorAddr], "DKG: Validator has already slashed");

        _slashedValidators[_validatorAddr] = true;

        _removeValidator(_validatorAddr);

        emit ValidatorSlashed(_validatorAddr);
    }

    function voteSigner(
        uint256 _epochId,
        address _signerAddress,
        bytes memory _signature
    ) external override onlyActiveValidator {
        require(getEpochStatus(_epochId) == DKGEpochStatuses.DKG_GENERATION, "DKG: Not a DKG generation period");

        address recoveredSigner = bytes("verify").toEthSignedMessageHash().recover(_signature);
        require(recoveredSigner == _signerAddress, "DKG: Signature is invalid");

        DKGEpochData storage epochData = _dkgEpochsData[_epochId];

        require(epochData.signerVotes[msg.sender] == address(0), "DKG: Already voted");

        uint256 newVotesCount = ++epochData.signerVotesCount[_signerAddress];
        epochData.signerVotes[msg.sender] = _signerAddress;

        emit SignerVoted(_epochId, msg.sender, _signerAddress);

        if (newVotesCount > getActiveValidatorsCount() / 2 && _activeSigner != _signerAddress) {
            epochData.epochSigner = _signerAddress;
            _activeSigner = _signerAddress;

            _createNewEpoch();
            updateAllValidators();

            emit SignerAddressUpdated(_epochId, _signerAddress);
        }
    }

    function updateAllValidators() public override {
        address[] memory allCurrentValidators = _validators.values();

        for (uint256 i = 0; i < allCurrentValidators.length; ++i) {
            if (_validatorsData[allCurrentValidators[i]].endValidationEpochId <= _activeEpochId) {
                _removeValidator(allCurrentValidators[i]);
            }
        }
    }

    function getSignerAddress() external view override returns (address) {
        return _activeSigner;
    }

    function getActiveEpochId() external view returns (uint256) {
        return _activeEpochId;
    }

    function getDKGEpochInfo(uint256 _epochId) external view override returns (DKGEpochInfo memory) {
        DKGEpochData storage epochData = _dkgEpochsData[_epochId];

        return
            DKGEpochInfo(
                _epochId,
                epochData.epochStartTime,
                epochData.dkgGenPeriodStartTime,
                epochData.epochSigner,
                getEpochStatus(_epochId)
            );
    }

    function getValidatorInfo(address _validator) external view override returns (ValidatorInfo memory) {
        ValidatorData storage validatorData = _validatorsData[_validator];

        return ValidatorInfo(_validator, validatorData.startValidationEpochId, validatorData.endValidationEpochId);
    }

    function getSignerVotesCount(uint256 _epochId, address _signerAddr) external view override returns (uint256) {
        return _dkgEpochsData[_epochId].signerVotesCount[_signerAddr];
    }

    function getValidatorVote(uint256 _epochId, address _validatorAddr) external view override returns (address) {
        return _dkgEpochsData[_epochId].signerVotes[_validatorAddr];
    }

    function getAllValidators() external view override returns (address[] memory) {
        return _validators.values();
    }

    function getActiveValidators() external view override returns (address[] memory) {
        Vector.AddressVector memory validatorsVector = Vector.newAddress();

        uint256 allValidatorsCount = _validators.length();

        for (uint256 i = 0; i < allValidatorsCount; ++i) {
            address currentValidator = _validators.at(i);

            if (isActiveValidator(currentValidator)) {
                validatorsVector.push(currentValidator);
            }
        }

        return validatorsVector.toArray();
    }

    function getAllValidatorsCount() external view override returns (uint256) {
        return _validators.length();
    }

    function getActiveValidatorsCount() public view override returns (uint256 _activeValidatorsCount) {
        uint256 allValidatorsCount = _validators.length();

        for (uint256 i = 0; i < allValidatorsCount; ++i) {
            if (isActiveValidator(_validators.at(i))) {
                _activeValidatorsCount++;
            }
        }
    }

    function getActiveEpochStatus() public view override returns (DKGEpochStatuses) {
        return getEpochStatus(_activeEpochId);
    }

    function getEpochStatus(uint256 _epochId) public view override returns (DKGEpochStatuses) {
        if (_epochId > _activeEpochId) {
            return DKGEpochStatuses.NONE;
        } else if (_epochId != _activeEpochId) {
            return DKGEpochStatuses.FINISHED;
        }

        uint256 epochStartTime = _dkgEpochsData[_epochId].epochStartTime;

        if ((epochStartTime += guaranteedWorkingEpochDuration) >= block.timestamp) {
            return DKGEpochStatuses.GUARANTEED_WORKING;
        }

        uint256 dkgGenPeriodStartTime = _dkgEpochsData[_epochId].dkgGenPeriodStartTime;

        if (dkgGenPeriodStartTime != 0) {
            if (dkgGenPeriodStartTime > block.timestamp) {
                return DKGEpochStatuses.UPDATES_COLLECTION;
            } else if (dkgGenPeriodStartTime + dkgGenerationEpochDuration >= block.timestamp) {
                return DKGEpochStatuses.DKG_GENERATION;
            }
        }

        return DKGEpochStatuses.ACTIVE;
    }

    function getGuaranteedWorkingPeriodEndTime(uint256 _epochId) public view returns (uint256) {
        return _dkgEpochsData[_epochId].epochStartTime + guaranteedWorkingEpochDuration;
    }

    function getDKGPeriodEndTime(uint256 _epochId) public view override returns (uint256) {
        require(_dkgEpochsData[_epochId].dkgGenPeriodStartTime != 0, "DKG: DKG generation period does not exist");

        return _dkgEpochsData[_epochId].dkgGenPeriodStartTime + dkgGenerationEpochDuration;
    }

    function isActiveValidator(address _validatorAddr) public view override returns (bool) {
        ValidatorData memory validatorData = _validatorsData[_validatorAddr];

        return
            !_slashedValidators[_validatorAddr] &&
            validatorData.startValidationEpochId <= _activeEpochId &&
            validatorData.endValidationEpochId > _activeEpochId;
    }

    function isValidator(address _validatorAddr) public view override returns (bool) {
        return _validators.contains(_validatorAddr);
    }

    function isValidatorSlashed(address _validatorAddr) public view returns (bool) {
        return _slashedValidators[_validatorAddr];
    }

    function _collectUpdates() internal returns (uint256) {
        DKGEpochStatuses activeEpochStatus = getActiveEpochStatus();

        require(
            activeEpochStatus != DKGEpochStatuses.DKG_GENERATION,
            "DKG: Unable to collect updates during the DKG generation period"
        );

        uint256 currentEpochId = _activeEpochId;

        if (
            _dkgEpochsData[currentEpochId].dkgGenPeriodStartTime == 0 ||
            getDKGPeriodEndTime(currentEpochId) < block.timestamp
        ) {
            uint256 updatesCollectionStartTime = activeEpochStatus != DKGEpochStatuses.GUARANTEED_WORKING
                ? block.timestamp
                : getGuaranteedWorkingPeriodEndTime(currentEpochId);

            _dkgEpochsData[currentEpochId].dkgGenPeriodStartTime =
                updatesCollectionStartTime +
                updatesCollectionEpochDuration;
        }

        return currentEpochId == FIRST_EPOCH_ID ? currentEpochId : ++currentEpochId;
    }

    function _createNewEpoch() internal {
        uint256 newEpochId = ++_activeEpochId;

        _dkgEpochsData[newEpochId].epochStartTime = block.timestamp;
        emit NewEpochCreated(newEpochId, block.timestamp);
    }

    function _removeValidator(address _validatorToRemove) internal {
        _validators.remove(_validatorToRemove);
        delete _validatorsData[_validatorToRemove];

        emit ValidatorRemoved(_validatorToRemove);
    }

    function _onlyActiveValidator() internal view {
        require(isActiveValidator(msg.sender), "DKG: Not an active validator");
    }

    function _onlyStaking() internal view {
        require(msg.sender == address(_staking), "DKG: Not a staking address");
    }

    function _onlySlashingVoting() internal view {
        require(msg.sender == address(_slashingVoting), "DKG: Not a slashing voting address");
    }
}
