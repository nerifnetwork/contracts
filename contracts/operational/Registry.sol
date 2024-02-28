// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

import "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import "../interfaces/SignerOwnable.sol";
import "../interfaces/IGateway.sol";
import "../interfaces/IGatewayFactory.sol";
import "../interfaces/IBillingManager.sol";
import "../interfaces/IRegistry.sol";

contract Registry is IRegistry, Initializable, SignerOwnable, UUPSUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;
    string internal constant GATEWAY_PERFORM_FUNC_SIGNATURE = "perform(uint256,address,bytes)";

    IGatewayFactory public gatewayFactory;
    IBillingManager public billingManager;

    bool public isMainChain;
    uint16 public maxWorkflowsPerAccount;

    uint256[] internal _existingWorkflowIds;
    EnumerableSet.AddressSet internal _existingGatewayOwners;

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
        _setGateway(msg.sender, _gateway);
    }

    function deployAndSetGateway() external override returns (address) {
        return _deployAndSetGateway(msg.sender);
    }

    function updateWorkflowTotalSpent(
        uint256 _workflowId,
        uint256 _workflowExecutionAmount
    ) external override onlyExistingWorkflow(_workflowId) {
        require(msg.sender == address(billingManager), "Registry: sender is not a billing manager");

        _workflowsInfo[_workflowId].totalSpent += _workflowExecutionAmount;
    }

    function pauseWorkflows(uint256[] calldata _workflowIds) external override onlyMainchain {
        for (uint256 i = 0; i < _workflowIds.length; i++) {
            _onlyWorkflowOwner(_workflowIds[i]);
            _onlyRequiredStatus(_workflowIds[i], WorkflowStatus.ACTIVE, true);
            _updateWorkflowStatus(_workflowIds[i], WorkflowStatus.PAUSED);
        }
    }

    function resumeWorkflows(uint256[] calldata _workflowIds) external override onlyMainchain {
        for (uint256 i = 0; i < _workflowIds.length; i++) {
            _onlyWorkflowOwner(_workflowIds[i]);
            _onlyRequiredStatus(_workflowIds[i], WorkflowStatus.PAUSED, true);
            _updateWorkflowStatus(_workflowIds[i], WorkflowStatus.ACTIVE);
        }
    }

    function registerWorkflows(RegisterWorkflowInfo[] calldata _registerWorkflowInfoArr) external override {
        bool isMainChainLocal = isMainChain;

        if (isMainChainLocal && maxWorkflowsPerAccount > 0) {
            require(
                workflowsPerAddress[msg.sender] + _registerWorkflowInfoArr.length <= maxWorkflowsPerAccount,
                "Registry: reached max workflows capacity"
            );
        }

        for (uint256 i = 0; i < _registerWorkflowInfoArr.length; i++) {
            RegisterWorkflowInfo calldata currentRegisterInfo = _registerWorkflowInfoArr[i];

            require(
                isMainChainLocal
                    ? currentRegisterInfo.workflowOwner == msg.sender
                    : signerGetter.getSignerAddress() == msg.sender,
                "Registry: not a sender or signer"
            );

            if (currentRegisterInfo.requireGateway) {
                address currentGateway = _gateways[currentRegisterInfo.workflowOwner];

                if (currentRegisterInfo.deployGateway && currentGateway == address(0)) {
                    _deployAndSetGateway(currentRegisterInfo.workflowOwner);
                } else {
                    require(currentGateway != address(0), "Registry: gateway not found");
                }
            }

            require(
                _workflowsInfo[currentRegisterInfo.id].status == WorkflowStatus.NONE,
                "Registry: workflow id is already exists"
            );

            _workflowsInfo[currentRegisterInfo.id] = Workflow(
                currentRegisterInfo.id,
                currentRegisterInfo.workflowOwner,
                currentRegisterInfo.hash,
                isMainChainLocal ? WorkflowStatus.PENDING : WorkflowStatus.ACTIVE,
                0
            );
            workflowsPerAddress[currentRegisterInfo.workflowOwner]++;
            _existingWorkflowIds.push(currentRegisterInfo.id);

            emit WorkflowRegistered(
                currentRegisterInfo.workflowOwner,
                currentRegisterInfo.id,
                currentRegisterInfo.hash
            );
        }
    }

    function activateWorkflows(uint256[] calldata _workflowIds) external override onlyMainchain onlySigner {
        for (uint256 i = 0; i < _workflowIds.length; i++) {
            _onlyRequiredStatus(_workflowIds[i], WorkflowStatus.PENDING, true);
            _updateWorkflowStatus(_workflowIds[i], WorkflowStatus.ACTIVE);
        }
    }

    function cancelWorkflows(uint256[] calldata _workflowIds) external override {
        bool isMainChainLocal = isMainChain;
        address signerAddr = signerGetter.getSignerAddress();

        for (uint256 i = 0; i < _workflowIds.length; i++) {
            _onlyExistingWorkflow(_workflowIds[i]);

            address workflowOwner = getWorkflowOwner(_workflowIds[i]);

            require(
                isMainChainLocal ? workflowOwner == msg.sender : signerAddr == msg.sender,
                "Registry: not a workflow owner or signer"
            );

            _onlyRequiredStatus(_workflowIds[i], WorkflowStatus.CANCELLED, false);
            _updateWorkflowStatus(_workflowIds[i], WorkflowStatus.CANCELLED);

            workflowsPerAddress[workflowOwner]--;
        }
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

    function getTotalGatewaysCount() external view override returns (uint256) {
        return _existingGatewayOwners.length();
    }

    function getGateway(address _userAddr) external view override returns (address) {
        return _gateways[_userAddr];
    }

    function getGatewaysInfo(
        uint256 _offset,
        uint256 _limit
    ) external view override returns (GatewayInfo[] memory _gatewaysInfoArr) {
        uint256 to = Paginator.getTo(_existingGatewayOwners.length(), _offset, _limit);

        _gatewaysInfoArr = new GatewayInfo[](to - _offset);

        for (uint256 i = _offset; i < to; i++) {
            address gatewayOwner = _existingGatewayOwners.at(i);

            _gatewaysInfoArr[i - _offset] = GatewayInfo(gatewayOwner, _gateways[gatewayOwner]);
        }
    }

    function getTotalWorkflowsCount() external view override returns (uint256) {
        return _existingWorkflowIds.length;
    }

    function getWorkflow(uint256 _id) external view override returns (Workflow memory) {
        return _workflowsInfo[_id];
    }

    function getWorkflows(
        uint256 _offset,
        uint256 _limit
    ) external view override returns (Workflow[] memory _workflowsArr) {
        uint256 to = Paginator.getTo(_existingWorkflowIds.length, _offset, _limit);

        _workflowsArr = new Workflow[](to - _offset);

        for (uint256 i = _offset; i < to; i++) {
            _workflowsArr[i - _offset] = _workflowsInfo[_existingWorkflowIds[i]];
        }
    }

    function getWorkflowOwner(uint256 _id) public view override returns (address) {
        return _workflowsInfo[_id].owner;
    }

    function getWorkflowStatus(uint256 _id) public view override returns (WorkflowStatus) {
        return _workflowsInfo[_id].status;
    }

    function isWorkflowExist(uint256 _id) public view override returns (bool) {
        return _workflowsInfo[_id].status != WorkflowStatus.NONE;
    }

    function _authorizeUpgrade(address) internal virtual override onlySigner {}

    function _updateWorkflowStatus(uint256 _id, WorkflowStatus _newStatus) internal {
        _workflowsInfo[_id].status = _newStatus;

        emit WorkflowStatusChanged(_id, _newStatus);
    }

    function _deployAndSetGateway(address _gatewayOwner) internal returns (address) {
        address newGatewayAddr = gatewayFactory.deployGateway(_gatewayOwner);

        _setGateway(_gatewayOwner, newGatewayAddr);

        return newGatewayAddr;
    }

    function _setGateway(address _gatewayOwner, address _gateway) internal {
        if (_gateway != address(0)) {
            _existingGatewayOwners.add(_gatewayOwner);
        } else {
            _existingGatewayOwners.remove(_gatewayOwner);
        }

        _gateways[_gatewayOwner] = _gateway;

        emit GatewaySet(_gatewayOwner, _gateway);
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

    function _onlyRequiredStatus(uint256 _id, WorkflowStatus _statusToCheck, bool _isEqual) internal view {
        require((getWorkflowStatus(_id) == _statusToCheck) == _isEqual, "Registry: invalid workflow status");
    }

    function _onlyExistingWorkflow(uint256 _id) internal view {
        require(isWorkflowExist(_id), "Registry: workflow does not exist");
    }

    function _onlyWorkflowOwner(uint256 _id) internal view {
        require(getWorkflowOwner(_id) == msg.sender, "Registry: sender is not the workflow owner");
    }
}
