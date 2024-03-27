// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@solarity/solidity-lib/libs/data-structures/StringSet.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title IBillingManager
 * @notice Interface for the Billing Manager contract, responsible for managing user funds,
 * workflow executions funds locking, and network rewards
 */
interface IBillingManager {
    /**
     * @dev Struct containing data related to the deposit asset
     * @param tokenAddr The address of the token for the deposit asset
     * @param workflowExecutionDiscount The discount applied for workflow execution
     * @param networkRewards The rewards earned from network activities
     * @param isPermitable Indicates whether permit is allowed for this asset
     * @param isEnabled Indicates whether the asset is enabled for deposit
     */
    struct DepositAssetData {
        address tokenAddr;
        uint256 workflowExecutionDiscount;
        uint256 networkRewards;
        bool isPermitable;
        bool isEnabled;
    }

    /**
     * @dev Struct containing information about a specific deposit asset
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param depositAssetData Data related to the deposit asset
     */
    struct DepositAssetInfo {
        string depositAssetKey;
        DepositAssetData depositAssetData;
    }

    /**
     * @dev Struct containing user data related to deposits
     * @param withdrawNonce The nonce for withdrawals, used to prevent replay attacks
     * @param depositAssetKeys Set containing keys of all deposit assets associated with the user
     * @param userDepositsAmount Mapping from deposit asset keys to the amount deposited by the user
     */
    struct UserData {
        uint256 withdrawNonce;
        StringSet.Set depositAssetKeys;
        mapping(string => uint256) userDepositsAmount;
    }

    /**
     * @dev Struct containing information about a user's deposit
     * @param userAddr The address of the user
     * @param userDepositedAmount The amount deposited by the user
     */
    struct UserDepositInfo {
        address userAddr;
        uint256 userDepositedAmount;
    }

    /**
     * @dev Struct containing information about a network withdrawal
     * @param depositAssetKey The key identifying the deposit asset
     * @param userAddr The address of the user from whom the tokens will be withdrawn
     * @param amountToWithdraw The amount of tokens to be withdrawn
     */
    struct NetworkWithdrawInfo {
        string depositAssetKey;
        address userAddr;
        uint256 amountToWithdraw;
    }

    /**
     * @notice Emitted when a new deposit asset is added
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param depositAssetData Data related to the added deposit asset
     */
    event DepositAssetAdded(string depositAssetKey, DepositAssetData depositAssetData);

    /**
     * @notice Emitted when the workflow execution discount for a deposit asset is updated
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param newWorkflowExecutionDiscount The new workflow execution discount
     * @param prevWorkflowExecutionDiscount The previous workflow execution discount
     */
    event WorkflowExecutionDiscountUpdated(
        string indexed depositAssetKey,
        uint256 newWorkflowExecutionDiscount,
        uint256 prevWorkflowExecutionDiscount
    );

    /**
     * @notice Emitted when the enabled status of a deposit asset is updated
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param newEnabledStatus The new enabled status of the deposit asset
     */
    event DepositAssetEnabledStatusUpdated(string indexed depositAssetKey, bool newEnabledStatus);

    /**
     * @notice Emitted when an asset is deposited
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param tokensSender The address that sent the tokens for deposit
     * @param depositRecipient The address that received the deposit
     * @param depositAmount The amount of tokens deposited
     */
    event AssetDeposited(
        string indexed depositAssetKey,
        address tokensSender,
        address depositRecipient,
        uint256 depositAmount
    );

    /**
     * @notice Event emitted when funds are withdrawn from a user's balance
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param userAddr The address of the user
     * @param withdrawnAmount The amount withdrawn from the user's balance
     */
    event UserFundsWithdrawn(string indexed depositAssetKey, address indexed userAddr, uint256 withdrawnAmount);

    /**
     * @notice Event emitted when network rewards are withdrawn
     * @param recipientAddr The address of the recipient
     * @param rewardsAmount The amount of rewards withdrawn
     */
    event RewardsWithdrawn(string indexed depositAssetKey, address indexed recipientAddr, uint256 rewardsAmount);

    /**
     * @notice Emitted when a network withdrawal is completed
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param userAddr The address of the user from whom the tokens were withdrawn
     * @param amountToWithdraw The amount of tokens withdrawn
     */
    event NetworkWithdrawCompleted(string indexed depositAssetKey, address indexed userAddr, uint256 amountToWithdraw);

    /**
     * @notice Adds multiple deposit assets with their corresponding information
     * @param _depositAssetInfoArr An array containing DepositAssetInfo structs for the deposit assets to be added
     */
    function addDepositAssets(DepositAssetInfo[] memory _depositAssetInfoArr) external;

    /**
     * @notice Updates the workflow execution discount for a specific deposit asset
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _newWorkflowExecutionDiscount The new workflow execution discount to be set
     */
    function updateWorkflowExecutionDiscount(
        string memory _depositAssetKey,
        uint256 _newWorkflowExecutionDiscount
    ) external;

    /**
     * @notice Updates the enabled status of a specific deposit asset
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _newEnabledStatus The new enabled status to be set
     */
    function updateDepositAssetEnabledStatus(string memory _depositAssetKey, bool _newEnabledStatus) external;

    /**
     * @notice Allows the network to withdraw tokens on behalf of users
     * @param _networkWithdrawArr An array containing NetworkWithdrawInfo structs for each withdrawal
     */
    function networkWithdraw(NetworkWithdrawInfo[] calldata _networkWithdrawArr) external;

    /**
     * @notice Deposits a specified amount of tokens into the protocol
     * @dev This function can accept ETH if the deposit asset is the Native
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _recipientAddr The address to which the deposited tokens will be attributed
     * @param _depositAmount The amount of tokens to deposit
     */
    function deposit(string memory _depositAssetKey, address _recipientAddr, uint256 _depositAmount) external payable;

    /**
     * @notice Deposits a specified amount of tokens into the protocol using permit for approval
     * @dev This function can accept only ERC20 tokens
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _recipientAddr The address to which the deposited tokens will be attributed
     * @param _depositAmount The amount of tokens to deposit
     * @param _sigExpirationTime The expiration time for the permit signature
     * @param _v The recovery byte of the permit signature
     * @param _r The R part of the permit signature
     * @param _s The S part of the permit signature
     */
    function depositWithPermit(
        string memory _depositAssetKey,
        address _recipientAddr,
        uint256 _depositAmount,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external;

    /**
     * @notice Initiates a withdrawal of funds
     * @param _depositAssetKey The key identifying the deposit asset
     * @param _amountToWithdraw The amount of tokens to withdraw
     * @param _sigExpirationTime The expiration time for the withdrawal signature
     * @param _v The recovery byte of the withdrawal signature
     * @param _r The R part of the withdrawal signature
     * @param _s The S part of the withdrawal signature
     */
    function withdrawFunds(
        string memory _depositAssetKey,
        uint256 _amountToWithdraw,
        uint256 _sigExpirationTime,
        uint8 _v,
        bytes32 _r,
        bytes32 _s
    ) external;

    /**
     * @notice Withdraws network rewards
     * @param _depositAssetKey The unique key identifying the deposit asset
     */
    function withdrawNetworkRewards(string memory _depositAssetKey) external;

    /**
     * @notice Retrieves the key identifying the native deposit asset
     * @return string The key identifying the native deposit asset
     */
    function nativeDepositAssetKey() external view returns (string memory);

    /**
     * @notice Retrieves the total count of registered users
     * @return uint256 The total count of registered users
     */
    function getTotalUsersCount() external view returns (uint256);

    /**
     * @notice Retrieves an array of supported deposit asset keys
     * @return An array containing the keys of supported deposit assets
     */
    function getSupportedDepositAssetKeys() external view returns (string[] memory);

    /**
     * @notice Retrieves information about specified deposit assets
     * @param _depositAssetKeysArr An array containing the keys of deposit assets to retrieve information for
     * @return An array containing information about the specified deposit assets
     */
    function getDepositAssetsInfo(
        string[] memory _depositAssetKeysArr
    ) external view returns (DepositAssetInfo[] memory);

    /**
     * @notice Retrieves an array of existing user addresses within the specified range
     * @param _offset The starting index of users to retrieve
     * @param _limit The maximum number of users to retrieve
     * @return An array containing user addresses
     */
    function getExistingUsers(uint256 _offset, uint256 _limit) external view returns (address[] memory);

    /**
     * @notice Retrieves an array of user funds information within the specified range
     * @param _offset The starting index of user funds information to retrieve
     * @param _limit The maximum number of user funds information to retrieve
     * @return _usersInfoArr An array containing information about user funds
     */
    function getUsersDepositInfo(
        string memory _depositAssetKey,
        uint256 _offset,
        uint256 _limit
    ) external view returns (UserDepositInfo[] memory _usersInfoArr);

    /**
     * @notice Retrieves the withdrawal nonce for a specific user
     * @param _userAddr The address of the user
     * @return The withdrawal nonce for the user
     */
    function getUserWithdrawNonce(address _userAddr) external view returns (uint256);

    /**
     * @notice Retrieves the deposit asset keys associated with a specific user
     * @param _userAddr The address of the user for whom to retrieve deposit asset keys
     * @return An array containing the deposit asset keys associated with the specified user
     */
    function getUserDepositAssetKeys(address _userAddr) external view returns (string[] memory);

    /**
     * @notice Retrieves the network rewards for a specific deposit asset
     * @param _depositAssetKey The key identifying the deposit asset for which to retrieve network rewards
     * @return uint256 The amount of network rewards associated with the specified deposit asset
     */
    function getNetworkRewards(string memory _depositAssetKey) external view returns (uint256);

    /**
     * @notice Retrieves the funds information of a user
     * @param _userAddr The address of the user
     * @return Information about the user's funds
     */
    function getUserDepositInfo(
        address _userAddr,
        string memory _depositAssetKey
    ) external view returns (UserDepositInfo memory);

    /**
     * @notice Checks if a deposit asset is supported
     * @param _depositAssetKey The key identifying the deposit asset
     * @return True if the deposit asset is supported, otherwise false
     */
    function isDepositAssetSupported(string memory _depositAssetKey) external view returns (bool);

    /**
     * @notice Checks if a deposit asset is native to the protocol
     * @param _depositAssetKey The key identifying the deposit asset
     * @return True if the deposit asset is native, otherwise false
     */
    function isNativeDepositAsset(string memory _depositAssetKey) external view returns (bool);

    /**
     * @notice Checks if a deposit asset is enabled
     * @param _depositAssetKey The key identifying the deposit asset
     * @return True if the deposit asset is enabled, otherwise false
     */
    function isDepositAssetEnabled(string memory _depositAssetKey) external view returns (bool);

    /**
     * @notice Checks if a deposit asset allows permitting
     * @param _depositAssetKey The key identifying the deposit asset
     * @return True if permitting is allowed for the deposit asset, otherwise false
     */
    function isDepositAssetPermitable(string memory _depositAssetKey) external view returns (bool);
}
