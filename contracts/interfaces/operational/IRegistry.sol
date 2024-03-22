// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@solarity/solidity-lib/libs/data-structures/StringSet.sol";

/**
 * @title IRegistry
 * @notice Interface for the Registry contract, responsible for managing workflows, gateways, and configurations
 */
interface IRegistry {
    /**
     * @dev Struct containing information about a deposit asset
     * @param depositAssetKey The unique key identifying the deposit asset
     * @param depositAssetTotalSpent The total amount spent on the deposit asset
     */
    struct DepositAssetInfo {
        string depositAssetKey;
        uint256 depositAssetTotalSpent;
    }

    /**
     * @dev Struct containing basic information about a workflow
     * @param id The unique identifier of the workflow
     * @param owner The address of the owner of the workflow
     */
    struct BaseWorkflowInfo {
        uint256 id;
        address owner;
    }

    /**
     * @dev Struct containing data about a workflow, including deposit asset information
     * @param baseInfo Basic information about the workflow
     * @param depositAssetKeys Set containing keys of all deposit assets associated with the workflow
     * @param depositAssetsTotalSpent Mapping from deposit asset keys to total
     * amounts spent on those assets within the workflow
     */
    struct WorkflowData {
        BaseWorkflowInfo baseInfo;
        StringSet.Set depositAssetKeys;
        mapping(string => uint256) depositAssetsTotalSpent;
    }

    /**
     * @dev Struct containing information about a workflow, including deposit asset details
     * @param baseInfo Basic information about the workflow
     * @param depositAssetKeys Array of keys identifying deposit assets associated with the workflow
     * @param depositAssetsInfo Array containing information about the deposit assets associated with the workflow
     */
    struct WorkflowInfo {
        BaseWorkflowInfo baseInfo;
        string[] depositAssetKeys;
        DepositAssetInfo[] depositAssetsInfo;
    }

    /**
     * @notice Struct containing information about a gateway
     * @param gatewayOwner The address of the owner of the gateway
     * @param gateway The address of the gateway
     */
    struct GatewayInfo {
        address gatewayOwner;
        address gateway;
    }

    /**
     * @notice Struct containing information required to register a new workflow
     * @param id The unique identifier of the workflow
     * @param workflowOwner The address of the owner of the workflow
     * @param requireGateway The flag indicating whether the workflow requires a gateway
     */
    struct RegisterWorkflowInfo {
        uint256 id;
        address workflowOwner;
        bool requireGateway;
        bool deployGateway;
    }

    /**
     * @notice Event emitted when a new gateway is set
     * @param owner The address of the owner setting the gateway
     * @param gateway The address of the newly set gateway
     */
    event GatewaySet(address owner, address gateway);

    /**
     * @notice Event emitted when a workflow is registered
     * @param owner The address of the owner registering the workflow
     * @param id The ID of the registered workflow
     */
    event WorkflowRegistered(address owner, uint256 id);

    /**
     * @notice Event emitted upon completion of a workflow execution
     * @param workflowId The ID of the executed workflow
     * @param workflowExecutionId The ID of the specific execution within the workflow
     * @param success The boolean indicating whether the execution was successful
     */
    event Performance(uint256 workflowId, uint256 workflowExecutionId, bool success);

    /**
     * @notice Sets the maximum number of workflows allowed per account
     * @param _newMaxWorkflowsPerAccount The new maximum number of workflows per account to be set
     */
    function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) external;

    /**
     * @notice Sets the address of the gateway contract for the msg.sender address
     * @param _gateway The address of the gateway contract to be set
     */
    function setGateway(address _gateway) external;

    /**
     * @notice Deploys a new gateway contract and sets it as the current gateway for the msg.sender address
     * @return The address of the newly deployed gateway contract
     */
    function deployAndSetGateway() external returns (address);

    /**
     * @notice Updates the total amount spent on a workflow execution
     * @param _workflowId The ID of the workflow
     * @param _workflowExecutionAmount The amount spent on the workflow execution
     */
    function updateWorkflowTotalSpent(
        string memory _depositAssetKey,
        uint256 _workflowId,
        uint256 _workflowExecutionAmount
    ) external;

    /**
     * @notice Registers new workflows with the provided information
     * @param _registerWorkflowInfoArr The array containing information to register new workflows
     */
    function registerWorkflows(RegisterWorkflowInfo[] calldata _registerWorkflowInfoArr) external;

    /**
     * @notice Performs an action on a workflow execution
     * @param _workflowId The ID of the workflow
     * @param _workflowExecutionId The ID of the specific execution within the workflow
     * @param _gasAmount The gas amount for the transaction
     * @param _data The transaction data
     * @param _target The address of the target contract
     */
    function perform(
        uint256 _workflowId,
        uint256 _workflowExecutionId,
        uint256 _gasAmount,
        bytes calldata _data,
        address _target
    ) external;

    /**
     * @notice Retrieves the total count of gateways
     * @return The total count of gateways
     */
    function getTotalGatewaysCount() external view returns (uint256);

    /**
     * @notice Retrieves the gateway associated with the given user address
     * @param _userAddr The address of the user to retrieve the gateway for
     * @return The gateway address associated with the given user address
     */
    function getGateway(address _userAddr) external view returns (address);

    /**
     * @notice Retrieves an array of gateway information within the specified range
     * @param _offset The starting index of gateways to retrieve
     * @param _limit The maximum number of gateways to retrieve
     * @return _gatewaysInfoArr The array containing gateway information
     */
    function getGatewaysInfo(
        uint256 _offset,
        uint256 _limit
    ) external view returns (GatewayInfo[] memory _gatewaysInfoArr);

    /**
     * @notice Retrieves the total count of workflows
     * @return The total count of workflows
     */
    function getTotalWorkflowsCount() external view returns (uint256);

    /**
     * @notice Retrieves an array of workflows within the specified range
     * @param _offset The starting index of workflows to retrieve
     * @param _limit The maximum number of workflows to retrieve
     * @return _workflowsInfoArr The array containing workflow information
     */
    function getWorkflowsInfo(
        uint256 _offset,
        uint256 _limit
    ) external view returns (WorkflowInfo[] memory _workflowsInfoArr);

    /**
     * @notice Retrieves basic information about a workflow
     * @param _workflowId The unique identifier of the workflow
     * @return The basic information about the workflow
     */
    function getBaseWorkflowInfo(uint256 _workflowId) external view returns (BaseWorkflowInfo memory);

    /**
     * @notice Retrieves the deposit asset keys associated with a workflow
     * @param _workflowId The unique identifier of the workflow
     * @return An array containing the deposit asset keys associated with the workflow
     */
    function getWorkflowDepositAssetKeys(uint256 _workflowId) external view returns (string[] memory);

    /**
     * @notice Retrieves information about specific deposit assets associated with a workflow
     * @param _workflowId The unique identifier of the workflow
     * @param _depositAssetKeys An array containing the keys of deposit assets for which to retrieve information
     * @return An array containing information about the specified deposit assets
     */
    function getWorkflowDepositAssetsInfo(
        uint256 _workflowId,
        string[] memory _depositAssetKeys
    ) external view returns (DepositAssetInfo[] memory);

    /**
     * @notice Retrieves detailed information about a workflow, including deposit asset details
     * @param _workflowId The unique identifier of the workflow
     * @return Information about the workflow, including deposit asset details
     */
    function getWorkflowInfo(uint256 _workflowId) external view returns (WorkflowInfo memory);

    /**
     * @notice Retrieves the owner address of the workflow associated with the given ID
     * @param _id The ID of the workflow to retrieve the owner address for
     * @return The owner address of the workflow associated with the given ID
     */
    function getWorkflowOwner(uint256 _id) external view returns (address);

    /**
     * @notice Checks whether a workflow with the given ID registered
     * @param _id The ID of the workflow to check for existence
     * @return True if a workflow with the given ID registered, otherwise false
     */
    function isWorkflowRegistered(uint256 _id) external view returns (bool);
}
