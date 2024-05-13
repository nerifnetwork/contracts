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
     * @notice Struct containing data related to the deposit asset
     * @param tokenAddr The address of the token for the deposit asset
     * @param workflowExecutionDiscount The discount applied for workflow execution
     * @param totalAssetAmount The total deposited asset amount
     * @param isPermitable Indicates whether permit is allowed for this asset
     * @param isEnabled Indicates whether the asset is enabled for deposit
     */
    struct DepositAssetData {
        address tokenAddr;
        uint256 workflowExecutionDiscount;
        uint256 totalAssetAmount;
        bool isPermitable;
        bool isEnabled;
    }

    /**
     * @notice Struct containing information about a specific deposit asset
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param depositAssetData Data related to the deposit asset
     */
    struct DepositAssetInfo {
        string depositAssetKey;
        DepositAssetData depositAssetData;
    }

    /**
     * @notice Struct containing information about a user's deposit
     * @param userAddr The address of the user
     * @param userDepositedAmount The amount deposited by the user
     */
    struct UserDepositInfo {
        address userAddr;
        uint256 userDepositedAmount;
    }

    /**
     * @notice Struct containing information about a network withdrawal
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
     * @notice Struct containing information about a withdrawal
     * @param depositAssetKeys The keys identifying the deposit assets
     * @param withdrawAmounts The amounts of tokens to be withdrawn
     * @param nonce The nonce associated with the user's withdrawal
     * @param sigData Signature data used for withdrawal verification
     */
    struct WithdrawData {
        string[] depositAssetKeys;
        uint256[] withdrawAmounts;
        uint256 nonce;
        SigData sigData;
    }

    /**
     * @notice Struct containing signature data for withdrawal verification
     * @param sigExpirationTime The expiration time of the signature
     * @param v The 'v' component of the signature
     * @param r The 'r' component of the signature
     * @param s The 's' component of the signature
     */
    struct SigData {
        uint256 sigExpirationTime;
        uint8 v;
        bytes32 r;
        bytes32 s;
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
     * @notice Emitted when slashed tokens are added to the asset pool
     * @param depositAssetKey The key identifying the deposit asset associated with the slashed tokens
     * @param tokensAmount The amount of slashed tokens added
     */
    event SlashedTokensAdded(string indexed depositAssetKey, uint256 tokensAmount);

    /**
     * @notice Emitted when funds are withdrawn from the system
     * @param userAddr The address of the user who initiated the withdrawal
     * @param depositAssetKeys The keys identifying the deposit assets for the withdrawn funds
     * @param withdrawnAmounts The amounts of tokens withdrawn
     * @param nonce The nonce associated with the withdrawal
     */
    event FundsWithdrawn(
        address indexed userAddr,
        string[] depositAssetKeys,
        uint256[] withdrawnAmounts,
        uint256 nonce
    );

    /**
     * @notice Emitted when rewards are withdrawn from the system
     * @param userAddr The address of the user who initiated the withdrawal
     * @param depositAssetKeys The keys identifying the deposit assets for the withdrawn rewards
     * @param rewardsAmounts The amounts of rewards withdrawn
     * @param nonce The nonce associated with the withdrawal
     */
    event RewardsWithdrawn(
        address indexed userAddr,
        string[] depositAssetKeys,
        uint256[] rewardsAmounts,
        uint256 nonce
    );

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
    function addDepositAssets(DepositAssetInfo[] calldata _depositAssetInfoArr) external;

    /**
     * @notice Updates the workflow execution discount for a specific deposit asset
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _newWorkflowExecutionDiscount The new workflow execution discount to be set
     */
    function updateWorkflowExecutionDiscount(
        string calldata _depositAssetKey,
        uint256 _newWorkflowExecutionDiscount
    ) external;

    /**
     * @notice Updates the enabled status of a specific deposit asset
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _newEnabledStatus The new enabled status to be set
     */
    function updateDepositAssetEnabledStatus(string calldata _depositAssetKey, bool _newEnabledStatus) external;

    /**
     * @notice Adds slashed tokens to the asset pool
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _tokensAmount The amount of tokens to add
     */
    function addSlashedTokens(string calldata _depositAssetKey, uint256 _tokensAmount) external;

    /**
     * @notice Deposits a specified amount of tokens into the protocol.
     * This function can accept ETH if the deposit asset is the key from the `nativeDepositAssetKey()` function
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _recipientAddr The address to which the deposited tokens will be attributed
     * @param _depositAmount The amount of tokens to deposit
     */
    function deposit(string calldata _depositAssetKey, address _recipientAddr, uint256 _depositAmount) external payable;

    /**
     * @notice Deposits a specified amount of tokens into the protocol using permit for approval.
     * This function can accept only ERC20 tokens
     * @param _depositAssetKey The unique key identifying the deposit asset
     * @param _recipientAddr The address to which the deposited tokens will be attributed
     * @param _depositAmount The amount of tokens to deposit
     * @param _sigData The signature data for permit approval
     */
    function depositWithPermit(
        string calldata _depositAssetKey,
        address _recipientAddr,
        uint256 _depositAmount,
        SigData calldata _sigData
    ) external;

    /**
     * @notice Retrieves the key identifying the native deposit asset.
     * @return string The key identifying the native deposit asset.
     */
    function nativeDepositAssetKey() external view returns (string memory);

    /**
     * @notice Retrieves the key identifying the Nerif token deposit asset
     * @return string The key identifying the Nerif token deposit asset
     */
    function nerifTokenDepositAssetKey() external view returns (string memory);

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
     * @notice Retrieves the address of the token contract associated with a deposit asset
     * @param _depositAssetKey The key identifying the deposit asset
     * @return The address of the token contract
     */
    function getDepositAssetTokenAddr(string calldata _depositAssetKey) external view returns (address);

    /**
     * @notice Retrieves the total amount of tokens deposited for a specific deposit asset
     * @param _depositAssetKey The key identifying the deposit asset
     * @return The total amount of tokens deposited
     */
    function getTotalDepositAssetAmount(string calldata _depositAssetKey) external view returns (uint256);

    /**
     * @notice Checks if a user's nonce is used
     * @param _userAddr The address of the user
     * @param _nonce The nonce to check
     * @return True if the nonce is used, otherwise false
     */
    function isUserNonceUsed(address _userAddr, uint256 _nonce) external view returns (bool);

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
     * @notice Checks if a deposit asset allowed for permit logic
     * @param _depositAssetKey The key identifying the deposit asset
     * @return True if the permit logic is allowed, otherwise false
     */
    function isDepositAssetPermitable(string memory _depositAssetKey) external view returns (bool);
}
