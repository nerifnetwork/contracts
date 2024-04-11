// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/libs/data-structures/memory/Vector.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/system/IDKG.sol";
import "../interfaces/system/IStaking.sol";
import "../interfaces/ISignerAddress.sol";

import "hardhat/console.sol";

contract DKG is IDKG, Initializable, AbstractDependant {
    using EnumerableSet for EnumerableSet.AddressSet;
    using Vector for Vector.AddressVector;
    using ECDSA for *;

    IStaking internal _staking;

    uint256 public updatesCollectionEpochDuration;
    uint256 public dkgGenerationEpochDuration;
    uint256 public guaranteedWorkingEpochDuration;

    uint256 internal _lastEpochId;
    address internal _activeSigner;

    EnumerableSet.AddressSet internal _validators;

    mapping(uint256 => DKGEpochData) internal _dkgEpochsData;
    mapping(address => ValidatorData) internal _validatorsData;

    modifier onlyActiveValidator() {
        _onlyActiveValidator();
        _;
    }

    modifier onlyStaking() {
        _onlyStaking();
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

        uint256 epochId = ++_lastEpochId;

        _activeSigner = msg.sender;

        _dkgEpochsData[epochId].epochSigner = msg.sender;
        _dkgEpochsData[epochId].epochStartTime = block.timestamp;

        _validators.add(msg.sender);
        _validatorsData[msg.sender] = ValidatorData(
            ValidationData(block.timestamp, epochId),
            ValidationData(type(uint256).max, 0)
        );
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _staking = IStaking(contractsRegistry.getStakingContract());
    }

    // solhint-disable-next-line ordering
    function addValidator(address _validatorToAdd) external onlyStaking {
        require(!isValidator(_validatorToAdd), "DKG: Validator already exists");

        ValidationData memory startValidationData = _createEpochEndUpdateValidators();

        _validatorsData[_validatorToAdd] = ValidatorData(startValidationData, ValidationData(type(uint256).max, 0));
        _validators.add(_validatorToAdd);

        emit NewValidatorAdded(
            _validatorToAdd,
            startValidationData.validationTime,
            startValidationData.validationEpoch
        );
    }

    function announceValidatorExit(address _validatorToExit) external onlyStaking returns (uint256) {
        require(isActiveValidator(_validatorToExit), "DKG: Validator is not active");
        require(
            _validatorsData[_validatorToExit].endValidationData.validationTime == type(uint256).max,
            "DKG: Exit of the validator has already been announced"
        );

        ValidationData memory endValidationData = _createEpochEndUpdateValidators();

        _validatorsData[_validatorToExit].endValidationData = endValidationData;

        emit ValidatorExitAnnounced(
            _validatorToExit,
            endValidationData.validationTime,
            endValidationData.validationEpoch
        );

        return endValidationData.validationTime;
    }

    function removeValidator(address _validatorToRemove) external onlyStaking {
        require(_validators.contains(_validatorToRemove), "DKG: Validator does not exist");
        require(
            _validatorsData[_validatorToRemove].endValidationData.validationTime <= block.timestamp,
            "DKG: Validator can't be removed yet"
        );

        _removeValidator(_validatorToRemove);
    }

    function voteSigner(
        uint256 _epochId,
        address _signerAddress,
        bytes memory _signature
    ) external onlyActiveValidator {
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

            emit SignerAddressUpdated(_epochId, _signerAddress);
        }
    }

    function updateAllValidators() public {
        address[] memory allCurrentValidators = _validators.values();

        for (uint256 i = 0; i < allCurrentValidators.length; ++i) {
            ValidatorData storage validatorData = _validatorsData[allCurrentValidators[i]];

            if (!isActiveValidator(allCurrentValidators[i])) {
                _updateValidationData(validatorData.startValidationData);
                _updateValidationData(validatorData.endValidationData);
            }

            if (validatorData.endValidationData.validationTime <= block.timestamp) {
                _removeValidator(allCurrentValidators[i]);
            }
        }
    }

    function getSignerAddress() external view returns (address) {
        return _activeSigner;
    }

    function getDKGEpochInfo(uint256 _epochId) external view returns (DKGEpochInfo memory) {
        DKGEpochData storage epochData = _dkgEpochsData[_epochId];

        return DKGEpochInfo(_epochId, epochData.epochStartTime, epochData.epochSigner, getEpochStatus(_epochId));
    }

    function getValidatorInfo(address _validator) external view returns (ValidatorInfo memory) {
        ValidatorData storage validatorData = _validatorsData[_validator];

        return ValidatorInfo(_validator, validatorData.startValidationData, validatorData.endValidationData);
    }

    function getSignerVotesCount(uint256 _epochId, address _signerAddr) external view returns (uint256) {
        return _dkgEpochsData[_epochId].signerVotesCount[_signerAddr];
    }

    function getValidatorVote(uint256 _epochId, address _validatorAddr) external view returns (address) {
        return _dkgEpochsData[_epochId].signerVotes[_validatorAddr];
    }

    function getActiveValidators() external view returns (address[] memory) {
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

    function getAllValidators() external view returns (address[] memory) {
        return _validators.values();
    }

    function getLastEpochId() external view returns (uint256) {
        return _lastEpochId;
    }

    function getAllValidatorsCount() external view returns (uint256) {
        return _validators.length();
    }

    function getActiveValidatorsCount() public view returns (uint256 _activeValidatorsCount) {
        uint256 allValidatorsCount = _validators.length();

        for (uint256 i = 0; i < allValidatorsCount; ++i) {
            if (isActiveValidator(_validators.at(i))) {
                _activeValidatorsCount++;
            }
        }
    }

    function getCurrentEpochStatus() public view returns (DKGEpochStatuses) {
        return getEpochStatus(getCurrentEpochId());
    }

    function getEpochStatus(uint256 _epochId) public view returns (DKGEpochStatuses) {
        if (_epochId > _lastEpochId) {
            return DKGEpochStatuses.NONE;
        }

        uint256 epochStartTime = _dkgEpochsData[_epochId].epochStartTime;

        if (isLastEpoch(_epochId) && epochStartTime > block.timestamp) {
            return DKGEpochStatuses.NOT_STARTED;
        }

        if (!isCurrentEpoch(_epochId)) {
            return DKGEpochStatuses.FINISHED;
        }

        if ((epochStartTime += updatesCollectionEpochDuration) >= block.timestamp) {
            return DKGEpochStatuses.UPDATES_COLLECTION;
        } else if ((epochStartTime += dkgGenerationEpochDuration) >= block.timestamp) {
            return DKGEpochStatuses.DKG_GENERATION;
        } else if ((epochStartTime += guaranteedWorkingEpochDuration) >= block.timestamp) {
            return DKGEpochStatuses.GUARANTEED_WORKING;
        } else {
            return DKGEpochStatuses.ACTIVE;
        }
    }

    function getUpdatedCollectionPeriodEndTime(uint256 _epochId) public view returns (uint256) {
        return _dkgEpochsData[_epochId].epochStartTime + updatesCollectionEpochDuration;
    }

    function getDKGPeriodEndTime(uint256 _epochId) public view returns (uint256) {
        return getUpdatedCollectionPeriodEndTime(_epochId) + dkgGenerationEpochDuration;
    }

    function getEpochEndTime(uint256 _epochId) public view returns (uint256) {
        return getDKGPeriodEndTime(_epochId) + guaranteedWorkingEpochDuration;
    }

    function getCurrentEpochId() public view returns (uint256 _currentEpochId) {
        _currentEpochId = _lastEpochId;

        if (_dkgEpochsData[_currentEpochId].epochStartTime > block.timestamp) {
            _currentEpochId--;
        }
    }

    function isCurrentEpoch(uint256 _epochId) public view returns (bool) {
        return getCurrentEpochId() == _epochId;
    }

    function isLastEpoch(uint256 _epochId) public view returns (bool) {
        return _lastEpochId == _epochId;
    }

    function isActiveValidator(address _validatorAddr) public view returns (bool) {
        ValidatorData memory validatorData = _validatorsData[_validatorAddr];

        bool startValidationTimeCheck = isDKGGenSuccessful(validatorData.startValidationData.validationEpoch) &&
            validatorData.startValidationData.validationTime <= block.timestamp;
        bool endValidationTimeCheck = !isDKGGenSuccessful(validatorData.endValidationData.validationEpoch) ||
            validatorData.endValidationData.validationTime > block.timestamp;

        return startValidationTimeCheck && endValidationTimeCheck;
    }

    function isValidator(address _validatorAddr) public view returns (bool) {
        return _validators.contains(_validatorAddr);
    }

    function isDKGGenSuccessful(uint256 _epochId) public view returns (bool) {
        return _dkgEpochsData[_epochId].epochSigner != address(0);
    }

    function _createEpoch() internal returns (ValidationData memory) {
        uint256 currentEpochId = getCurrentEpochId();
        DKGEpochStatuses currentEpochStatus = getEpochStatus(currentEpochId);

        uint256 epochId = currentEpochId;

        if (currentEpochStatus != DKGEpochStatuses.UPDATES_COLLECTION) {
            epochId = _lastEpochId;

            if (isLastEpoch(currentEpochId)) {
                uint256 nextEpochStartTime = block.timestamp;

                if (currentEpochStatus != DKGEpochStatuses.ACTIVE) {
                    nextEpochStartTime = getEpochEndTime(epochId);
                }

                epochId = ++_lastEpochId;

                _dkgEpochsData[epochId].epochStartTime = nextEpochStartTime;

                emit NewEpochCreated(epochId, nextEpochStartTime);
            }
        }

        return ValidationData(getDKGPeriodEndTime(epochId), epochId);
    }

    function _createEpochEndUpdateValidators() internal returns (ValidationData memory _validationData) {
        uint256 currentLastEpochId = _lastEpochId;

        _validationData = _createEpoch();

        if (_validationData.validationEpoch != currentLastEpochId) {
            updateAllValidators();
        }
    }

    function _updateValidationData(ValidationData storage _validationData) internal {
        if (_validationData.validationTime <= block.timestamp && !isDKGGenSuccessful(_validationData.validationEpoch)) {
            ValidationData memory validationData = _createEpoch();

            _validationData.validationTime = validationData.validationTime;
            _validationData.validationEpoch = validationData.validationEpoch;
        }
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
}
