// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../interfaces/SignerOwnable.sol";
import "../interfaces/IGateway.sol";
import "../interfaces/IGatewayFactory.sol";
import "../interfaces/IBillingManager.sol";
import "../interfaces/IRegistry.sol";

contract Registry is IRegistry, Initializable, SignerOwnable {
    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;
    string internal constant GATEWAY_PERFORM_FUNC_SIGNATURE = "perform(uint256,address,bytes)";

    IGatewayFactory public gatewayFactory;
    IBillingManager public billingManager;

    bool public isMainChain;
    uint16 public maxWorkflowsPerAccount;

    mapping(address => uint256) public workflowsPerAddress;

    mapping(uint256 => Workflow) internal _workflowsInfo;
    mapping(address => address) internal _gateways;

    modifier onlyMainchain() {
        _onlyMainChain();
        _;
    }

    modifier onlyExistingWorkflow(uint256 _id) {
        _onlyExistingWorkflow(_id);
        _;
    }

    modifier onlyWorkflowOwner(uint256 _id) {
        _onlyWorkflowOwner(_id);
        _;
    }

    function initialize(
        bool _isMainChain,
        address _signerGetterAddress,
        address _gatewayFactoryAddr,
        address _billingManagerAddr,
        uint16 _maxWorkflowsPerAccount
    ) external initializer {
        isMainChain = _isMainChain;

        gatewayFactory = IGatewayFactory(_gatewayFactoryAddr);
        billingManager = IBillingManager(_billingManagerAddr);

        _setSignerGetter(_signerGetterAddress);

        maxWorkflowsPerAccount = _maxWorkflowsPerAccount;
    }

    function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) external override onlySigner {
        maxWorkflowsPerAccount = _newMaxWorkflowsPerAccount;
    }

    function setGatewayFactory(address _newGatewayFactory) external override onlySigner {
        gatewayFactory = IGatewayFactory(_newGatewayFactory);
    }

    function setGateway(address _gateway) external override {
        _gateways[msg.sender] = _gateway;

        emit GatewaySet(msg.sender, _gateway);
    }

    function deployAndSetGateway() external override returns (address) {
        address newGatewayAddr = gatewayFactory.deployGateway(msg.sender);

        _gateways[msg.sender] = newGatewayAddr;

        emit GatewaySet(msg.sender, newGatewayAddr);

        return newGatewayAddr;
    }

    function updateWorkflowTotalSpent(uint256 _workflowId, uint256 _workflowExecutionAmount)
        external
        override
        onlyExistingWorkflow(_workflowId)
    {
        require(msg.sender == address(billingManager), "Registry: sender is not a billing manager");

        _workflowsInfo[_workflowId].totalSpent += _workflowExecutionAmount;
    }

    function pauseWorkflow(uint256 _id)
        external
        override
        onlyMainchain
        onlyExistingWorkflow(_id)
        onlyWorkflowOwner(_id)
    {
        _onlyRequiredStatus(_id, WorkflowStatus.ACTIVE, true);
        _updateWorkflowStatus(_id, WorkflowStatus.PAUSED);
    }

    function resumeWorkflow(uint256 _id)
        external
        override
        onlyMainchain
        onlyExistingWorkflow(_id)
        onlyWorkflowOwner(_id)
    {
        _onlyRequiredStatus(_id, WorkflowStatus.PAUSED, true);
        _updateWorkflowStatus(_id, WorkflowStatus.ACTIVE);
    }

    function registerWorkflow(
        uint256 _id,
        address _workflowOwner,
        bytes calldata _hash,
        bool _requireGateway
    ) external override {
        require(
            isMainChain ? _workflowOwner == msg.sender : signerGetter.getSignerAddress() == msg.sender,
            "Registry: not a sender or signer"
        );

        if (_requireGateway) {
            require(_gateways[_workflowOwner] != address(0), "Registry: gateway not found");
        }

        require(_workflowsInfo[_id].status == WorkflowStatus.NONE, "Registry: workflow id is already exists");

        if (isMainChain && maxWorkflowsPerAccount > 0) {
            require(
                workflowsPerAddress[msg.sender] < maxWorkflowsPerAccount,
                "Registry: reached max workflows capacity"
            );
        }

        _workflowsInfo[_id] = Workflow(
            _id,
            _workflowOwner,
            _hash,
            isMainChain ? WorkflowStatus.PENDING : WorkflowStatus.ACTIVE,
            0
        );
        workflowsPerAddress[_workflowOwner]++;

        emit WorkflowRegistered(_workflowOwner, _id, _hash);
    }

    function activateWorkflow(uint256 _id) external override onlyMainchain onlySigner onlyExistingWorkflow(_id) {
        _onlyRequiredStatus(_id, WorkflowStatus.PENDING, true);
        _updateWorkflowStatus(_id, WorkflowStatus.ACTIVE);
    }

    function cancelWorkflow(uint256 _id) external override onlyExistingWorkflow(_id) {
        require(
            isMainChain ? getWorkflowOwner(_id) == msg.sender : signerGetter.getSignerAddress() == msg.sender,
            "Registry: not a workflow owner or signer"
        );

        _onlyRequiredStatus(_id, WorkflowStatus.CANCELLED, false);
        _updateWorkflowStatus(_id, WorkflowStatus.CANCELLED);
    }

    function perform(
        uint256 _workflowId,
        uint256 _workflowExecutionId,
        uint256 _gasAmount,
        bytes calldata _data,
        address _target
    ) external override onlySigner {
        _onlyRequiredStatus(_workflowId, WorkflowStatus.ACTIVE, true);

        require(
            billingManager.getExecutionWorkflowId(_workflowExecutionId) == _workflowId &&
                billingManager.getWorkflowExecutionStatus(_workflowExecutionId) ==
                IBillingManager.WorkflowExecutionStatus.PENDING,
            "Registry: invalid workflow execution ID"
        );
        require(address(this) != _target, "Registry: operation is not permitted");

        address existingGateway = _gateways[getWorkflowOwner(_workflowId)];

        require(existingGateway != address(0), "Registry: zero gateway address");

        bool success = _callWithExactGas(
            _gasAmount,
            existingGateway,
            abi.encodeWithSignature(GATEWAY_PERFORM_FUNC_SIGNATURE, _workflowId, _target, _data)
        );

        emit Performance(_workflowId, _workflowExecutionId, success);
    }

    function getGateway(address _userAddr) external view returns (address) {
        return _gateways[_userAddr];
    }

    function getWorkflow(uint256 _id) external view returns (Workflow memory) {
        return _workflowsInfo[_id];
    }

    function getWorkflowOwner(uint256 _id) public view returns (address) {
        return _workflowsInfo[_id].owner;
    }

    function getWorkflowStatus(uint256 _id) public view returns (WorkflowStatus) {
        return _workflowsInfo[_id].status;
    }

    function isWorkflowExist(uint256 _id) public view returns (bool) {
        return _workflowsInfo[_id].status != WorkflowStatus.NONE;
    }

    function _updateWorkflowStatus(uint256 _id, WorkflowStatus _newStatus) internal {
        _workflowsInfo[_id].status = _newStatus;

        emit WorkflowStatusChanged(_id, _newStatus);
    }

    function _callWithExactGas(
        uint256 _gasAmount,
        address _target,
        bytes memory _data
    ) internal returns (bool _success) {
        assembly {
            let g := gas()

            // Compute g -= PERFORM_GAS_CUSHION and check for underflow
            if lt(g, PERFORM_GAS_CUSHION) {
                revert(0, 0)
            }

            g := sub(g, PERFORM_GAS_CUSHION)

            // if g - g//64 <= gasAmount, revert
            // (we subtract g//64 because of EIP-150)
            if iszero(gt(sub(g, div(g, 64)), _gasAmount)) {
                revert(0, 0)
            }

            // solidity calls check that a contract actually exists at the destination, so we do the same
            if iszero(extcodesize(_target)) {
                revert(0, 0)
            }

            // call and return whether we succeeded. ignore return data
            _success := call(_gasAmount, _target, 0, add(_data, 0x20), mload(_data), 0, 0)
        }
    }

    function _onlyMainChain() internal view {
        require(isMainChain, "Registry: not a main chain");
    }

    function _onlyRequiredStatus(
        uint256 _id,
        WorkflowStatus _statusToCheck,
        bool _isEqual
    ) internal view {
        require((getWorkflowStatus(_id) == _statusToCheck) == _isEqual, "Registry: invalid workflow status");
    }

    function _onlyExistingWorkflow(uint256 _id) internal view {
        require(isWorkflowExist(_id), "Registry: workflow does not exist");
    }

    function _onlyWorkflowOwner(uint256 _id) internal view {
        require(getWorkflowOwner(_id) == msg.sender, "Registry: sender is not the workflow owner");
    }
}
