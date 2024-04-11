// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IDKG {
    enum DKGEpochStatuses {
        NONE,
        NOT_STARTED,
        UPDATES_COLLECTION,
        DKG_GENERATION,
        GUARANTEED_WORKING,
        ACTIVE,
        FINISHED
    }

    struct DKGEpochInfo {
        uint256 epochId;
        uint256 epochStartTime;
        address epochSigner;
        DKGEpochStatuses epochStatus;
    }

    struct ValidatorInfo {
        address validator;
        ValidationData startValidationData;
        ValidationData endValidationData;
    }

    struct DKGEpochData {
        address epochSigner;
        uint256 epochStartTime;
        mapping(address => uint256) signerVotesCount;
        mapping(address => address) signerVotes;
    }

    struct ValidationData {
        uint256 validationTime;
        uint256 validationEpoch;
    }

    struct ValidatorData {
        ValidationData startValidationData;
        ValidationData endValidationData;
    }

    event NewEpochCreated(uint256 epochId, uint256 epochStartTime);
    event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationTime, uint256 startValidationEpoch);
    event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationTime, uint256 endValidationEpoch);
    event ValidatorRemoved(address indexed validatorAddr);
    event SignerVoted(uint256 generation, address validator, address collectiveSigner);
    event SignerAddressUpdated(uint256 generation, address signerAddress);

    function addValidator(address _validatorToAdd) external;

    function announceValidatorExit(address _validatorToExit) external returns (uint256);

    function removeValidator(address _validatorToRemove) external;

    function voteSigner(uint256 _epochId, address _signerAddress, bytes memory _signature) external;

    function updateAllValidators() external;

    function getSignerAddress() external view returns (address);

    function getDKGEpochInfo(uint256 _epochId) external view returns (DKGEpochInfo memory);

    function getValidatorInfo(address _validator) external view returns (ValidatorInfo memory);

    function getSignerVotesCount(uint256 _epochId, address _signerAddr) external view returns (uint256);

    function getValidatorVote(uint256 _epochId, address _validatorAddr) external view returns (address);

    function getActiveValidators() external view returns (address[] memory);

    function getAllValidators() external view returns (address[] memory);

    function getAllValidatorsCount() external view returns (uint256);

    function getActiveValidatorsCount() external view returns (uint256);

    function getCurrentEpochStatus() external view returns (DKGEpochStatuses);

    function getEpochStatus(uint256 _epochId) external view returns (DKGEpochStatuses);

    function getUpdatedCollectionPeriodEndTime(uint256 _epochId) external view returns (uint256);

    function getDKGPeriodEndTime(uint256 _epochId) external view returns (uint256);

    function getEpochEndTime(uint256 _epochId) external view returns (uint256);

    function getCurrentEpochId() external view returns (uint256 _currentEpochId);

    function isCurrentEpoch(uint256 _epochId) external view returns (bool);

    function isLastEpoch(uint256 _epochId) external view returns (bool);

    function isActiveValidator(address _validatorAddr) external view returns (bool);

    function isValidator(address _validatorAddr) external view returns (bool);
}
