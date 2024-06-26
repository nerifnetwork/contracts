// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

/**
 * @notice Interface for managing the Distributed Key Generation (DKG) process and validator operations
 */
interface IDKG {
    /**
     * @notice Enum defining the statuses of DKG epochs
     */
    enum DKGEpochStatuses {
        NONE,
        GUARANTEED_WORKING,
        ACTIVE,
        UPDATES_COLLECTION,
        DKG_GENERATION,
        FINISHED
    }

    /**
     * @notice Struct containing information about a DKG epoch
     * @param epochId The ID of the epoch
     * @param epochStartTime The start time of the epoch
     * @param dkgGenPeriodStartTime The start time of the DKG generation period within the epoch
     * @param epochSigner The address of the signer for the epoch
     * @param epochStatus The status of the epoch
     */
    struct DKGEpochInfo {
        uint256 epochId;
        uint256 epochStartTime;
        uint256 dkgGenPeriodStartTime;
        address epochSigner;
        DKGEpochStatuses epochStatus;
    }

    /**
     * @notice Struct containing information about a validator
     * @param validator The address of the validator
     * @param startValidationEpochId The epoch ID of the start of the validator's validation period
     * @param endValidationEpochId The epoch ID of the end of the validator's validation period
     */
    struct ValidatorInfo {
        address validator;
        uint256 startValidationEpochId;
        uint256 endValidationEpochId;
    }

    /**
     * @notice Struct containing data for a DKG epoch
     * @param epochSigner The address of the signer for the epoch
     * @param epochStartTime The start time of the epoch
     * @param dkgGenPeriodStartTime The start time of the DKG generation period within the epoch
     * @param signerVotesCount Mapping from signer addresses to the number of votes they have received
     * @param signerVotes Mapping from validator addresses to the collective signer they voted for
     */
    struct DKGEpochData {
        address epochSigner;
        uint256 epochStartTime;
        uint256 dkgGenPeriodStartTime;
        mapping(address => uint256) signerVotesCount;
        mapping(address => address) signerVotes;
    }

    /**
     * @notice Represents the data associated with a validator
     * @param startValidationEpochId The epoch ID indicating the start of the validator's validation period
     * @param endValidationEpochId The epoch ID indicating the end of the validator's validation period
     */
    struct ValidatorData {
        uint256 startValidationEpochId;
        uint256 endValidationEpochId;
    }

    /**
     * @notice Emitted when a new epoch is created
     * @param epochId The unique identifier of the epoch
     * @param epochStartTime The start time of the epoch
     */
    event NewEpochCreated(uint256 epochId, uint256 epochStartTime);

    /**
     * @notice Emitted when a new validator is added
     * @param validatorAddr The address of the newly added validator
     * @param startValidationEpoch The epoch when the validation period starts for the validator
     */
    event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationEpoch);

    /**
     * @notice Emitted when a validator exit is announced
     * @param validatorAddr The address of the validator announcing the exit
     * @param endValidationEpoch The epoch when the validation period ends for the exiting validator
     */
    event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationEpoch);

    /**
     * @notice Emitted when a validator is removed
     * @param validatorAddr The address of the removed validator
     */
    event ValidatorRemoved(address indexed validatorAddr);

    /**
     * @notice Emitted when a validator is slashed
     * @param validatorAddr The address of the slashed validator
     */
    event ValidatorSlashed(address indexed validatorAddr);

    /**
     * @notice Emitted when a signer vote is cast
     * @param epochId The epoch id of the vote
     * @param validator The address of the validator casting the vote
     * @param collectiveSigner The address of the collective signer voted for
     */
    event SignerVoted(uint256 epochId, address validator, address collectiveSigner);

    /**
     * @notice Emitted when a signer address is updated
     * @param epochId The epoch id of the update
     * @param signerAddress The updated signer address
     */
    event SignerAddressUpdated(uint256 epochId, address signerAddress);

    /**
     * @notice Adds a new validator
     * @param _validatorToAdd The address of the validator to add
     */
    function addValidator(address _validatorToAdd) external;

    /**
     * @notice Announces the exit of a validator
     * @param _validatorToExit The address of the validator to exit
     * @return uint256 The end validation time
     */
    function announceValidatorExit(address _validatorToExit) external returns (uint256);

    /**
     * @notice Creates a slashing proposal
     * @return The DKG generation period start time
     */
    function createProposal() external returns (uint256);

    /**
     * @notice Slashes a validator
     * @param _validatorAddr The address of the validator to be slashed
     */
    function slashValidator(address _validatorAddr) external;

    /**
     * @notice Casts a vote for a signer
     * @param _epochId The ID of the epoch
     * @param _signerAddress The address of the signer
     * @param _signature The signature of the signer
     */
    function voteSigner(uint256 _epochId, address _signerAddress, bytes memory _signature) external;

    /**
     * @notice Updates all validators
     */
    function updateAllValidators() external;

    /**
     * @notice Retrieves the signer address
     * @return The signer address
     */
    function getSignerAddress() external view returns (address);

    /**
     * @notice Retrieves the ID of the active epoch
     * @return The ID of the active epoch
     */
    function getActiveEpochId() external view returns (uint256);

    /**
     * @notice Retrieves information about a DKG epoch
     * @param _epochId The ID of the epoch
     * @return Information about the DKG epoch
     */
    function getDKGEpochInfo(uint256 _epochId) external view returns (DKGEpochInfo memory);

    /**
     * @notice Retrieves information about a validator
     * @param _validator The address of the validator
     * @return Information about the validator
     */
    function getValidatorInfo(address _validator) external view returns (ValidatorInfo memory);

    /**
     * @notice Retrieves the vote count for a signer
     * @param _epochId The ID of the epoch
     * @param _signerAddr The address of the signer
     * @return The vote count for the signer
     */
    function getSignerVotesCount(uint256 _epochId, address _signerAddr) external view returns (uint256);

    /**
     * @notice Retrieves the vote of a validator
     * @param _epochId The ID of the epoch
     * @param _validatorAddr The address of the validator
     * @return The vote of the validator
     */
    function getValidatorVote(uint256 _epochId, address _validatorAddr) external view returns (address);

    /**
     * @notice Retrieves all validators
     * @return An array containing the addresses of all validators
     */
    function getAllValidators() external view returns (address[] memory);

    /**
     * @notice Retrieves active validators
     * @return An array containing the addresses of active validators
     */
    function getActiveValidators() external view returns (address[] memory);

    /**
     * @notice Retrieves the count of all validators
     * @return The count of all validators
     */
    function getAllValidatorsCount() external view returns (uint256);

    /**
     * @notice Retrieves the count of active validators
     * @return The count of active validators
     */
    function getActiveValidatorsCount() external view returns (uint256);

    /**
     * @notice Retrieves the current epoch status
     * @return The current epoch status
     */
    function getActiveEpochStatus() external view returns (DKGEpochStatuses);

    /**
     * @notice Retrieves the status of a specific epoch
     * @param _epochId The ID of the epoch
     * @return The status of the epoch
     */
    function getEpochStatus(uint256 _epochId) external view returns (DKGEpochStatuses);

    /**
     * @notice Retrieves the end time for the DKG period of an epoch
     * @param _epochId The ID of the epoch
     * @return The end time for the DKG period
     */
    function getDKGPeriodEndTime(uint256 _epochId) external view returns (uint256);

    /**
     * @notice Checks if a validator is active
     * @param _validatorAddr The address of the validator
     * @return True if the validator is active, otherwise false
     */
    function isActiveValidator(address _validatorAddr) external view returns (bool);

    /**
     * @notice Checks if an address is a validator
     * @param _validatorAddr The address to check
     * @return True if the address is a validator, otherwise false
     */
    function isValidator(address _validatorAddr) external view returns (bool);

    /**
     * @notice Checks if a validator has been slashed
     * @param _validatorAddr The address of the validator
     * @return True if the validator has been slashed, otherwise false
     */
    function isValidatorSlashed(address _validatorAddr) external view returns (bool);
}
