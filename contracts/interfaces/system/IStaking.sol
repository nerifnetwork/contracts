// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

/**
 * @notice Interface for managing staking and withdrawals
 */
interface IStaking {
    /**
     * @dev Struct containing information about a user's stake
     * @param userAddr The address of the user
     * @param userStakedAmount The amount staked by the user
     */
    struct UserStakeInfo {
        address userAddr;
        uint256 userStakedAmount;
    }

    /**
     * @dev Struct containing information about a withdrawal announcement
     * @param tokensAmount The amount of tokens to be withdrawn
     */
    struct WithdrawalAnnouncement {
        uint256 tokensAmount;
        uint256 withdrawalEpochId;
    }

    /**
     * @notice Emitted when the minimal stake requirement is updated
     * @param minimalStake The new minimal stake requirement
     */
    event MinimalStakeUpdated(uint256 minimalStake);

    /**
     * @notice Emitted when tokens are staked
     * @param tokensSender The address of the sender of the tokens
     * @param tokensRecipient The address of the recipient of the tokens
     * @param stakeAmount The amount of tokens staked
     */
    event TokensStaked(address indexed tokensSender, address indexed tokensRecipient, uint256 stakeAmount);

    /**
     * @notice Emitted when a withdrawal is announced
     * @param userAddr The address of the user announcing the withdrawal
     * @param tokensAmount The amount of tokens being withdrawn
     * @param withdrawalEpochId The epoch ID when the withdrawal will be allowed
     */
    event WithdrawalAnnounced(address indexed userAddr, uint256 tokensAmount, uint256 withdrawalEpochId);

    /**
     * @notice Emitted when tokens are withdrawn
     * @param userAddr The address of the user withdrawing tokens
     * @param tokensAmount The amount of tokens withdrawn
     */
    event TokensWithdrawn(address indexed userAddr, uint256 tokensAmount);

    /**
     * @notice Sets the minimal stake requirement
     * @param _minimalStake The new minimal stake requirement
     */
    function setMinimalStake(uint256 _minimalStake) external;

    /**
     * @notice Updates the whitelist status of users
     * @param _usersToUpdate The list of users to update
     * @param _isAdding Whether to add or remove users from the whitelist
     */
    function updateWhitelistedUsers(address[] calldata _usersToUpdate, bool _isAdding) external;

    /**
     * @notice Slashes a validator
     * @param _validator The address of the validator to be slashed
     */
    function slashValidator(address _validator) external;

    /**
     * @notice Announces a withdrawal
     */
    function announceWithdrawal() external;

    /**
     * @notice Withdraws tokens
     */
    function withdraw() external;

    /**
     * @notice Stakes tokens
     * @param _stakeAmount The amount of tokens to stake
     */
    function stake(uint256 _stakeAmount) external;

    /**
     * @notice Stakes tokens with permit
     * @param _stakeAmount The amount of tokens to stake
     * @param _sigExpirationTime The expiration time of the permit signature
     * @param _v The V part of the permit signature
     * @param _r The R part of the permit signature
     * @param _s The S part of the permit signature
     */
    function stakeWithPermit(
        uint256 _stakeAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external;

    /**
     * @notice Retrieves the minimal stake requirement
     * @return uint256 The minimal stake requirement
     */
    function minimalStake() external view returns (uint256);

    /**
     * @notice Retrieves the total stake across all validators
     * @return uint256 The total stake
     */
    function totalStake() external view returns (uint256);

    /**
     * @notice Retrieves stake information for users
     * @param _offset The offset for pagination
     * @param _limit The maximum number of results to retrieve
     * @return The array containing stake information for users
     */
    function getUsersStakeInfo(uint256 _offset, uint256 _limit) external view returns (UserStakeInfo[] memory);

    /**
     * @notice Retrieves the stake of a validator
     * @param _validator The address of the validator
     * @return The stake amount
     */
    function getStake(address _validator) external view returns (uint256);

    /**
     * @notice Retrieves withdrawal announcement information for a user
     * @param _userAddr The address of the user
     * @return The withdrawal announcement information
     */
    function getWithdrawalAnnouncement(address _userAddr) external view returns (WithdrawalAnnouncement memory);

    /**
     * @notice Checks if a user is whitelisted
     * @param _userAddr The address of the user
     * @return True if the user is whitelisted, false otherwise
     */
    function isUserWhitelisted(address _userAddr) external view returns (bool);

    /**
     * @notice Checks if a user has announced withdrawal
     * @param _userAddr The address of the user
     * @return True if the user has announced withdrawal, false otherwise
     */
    function hasWithdrawalAnnouncement(address _userAddr) external view returns (bool);
}
