// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";
import "@solarity/solidity-lib/libs/arrays/Paginator.sol";
import "@solarity/solidity-lib/libs/data-structures/StringSet.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/operational/IGateway.sol";
import "../interfaces/operational/IGatewayFactory.sol";
import "../interfaces/operational/IBillingManager.sol";
import "../interfaces/operational/IRegistry.sol";

contract Registry is IRegistry, Initializable, AbstractDependant {
    using EnumerableSet for EnumerableSet.AddressSet;
    using StringSet for StringSet.Set;

    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;
    string internal constant GATEWAY_PERFORM_FUNC_SIGNATURE = "perform(uint256,address,bytes)";

    IContractsRegistry internal _contractsRegistry;
    IGatewayFactory internal _gatewayFactory;
    IBillingManager internal _billingManager;

    uint16 public maxWorkflowsPerAccount;

    uint256[] internal _existingWorkflowIds;
    EnumerableSet.AddressSet internal _existingGatewayOwners;

    mapping(address => uint256) public workflowsPerAddress;

    mapping(uint256 => WorkflowData) internal _workflowsData;
    mapping(address => address) internal _gateways;

    modifier onlySigner() {
        _onlySigner();
        _;
    }

    modifier onlyExistingWorkflow(uint256 _id) {
        _onlyExistingWorkflow(_id);
        _;
    }

    function initialize(uint16 _maxWorkflowsPerAccount) external initializer {
        maxWorkflowsPerAccount = _maxWorkflowsPerAccount;
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _contractsRegistry = contractsRegistry;

        _gatewayFactory = IGatewayFactory(contractsRegistry.getGatewayFactoryContract());
        _billingManager = IBillingManager(contractsRegistry.getBillingManagerContract());
    }

    // solhint-disable-next-line ordering
    function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) external override onlySigner {
        maxWorkflowsPerAccount = _newMaxWorkflowsPerAccount;
    }

    function setGateway(address _gateway) external override {
        _setGateway(msg.sender, _gateway);
    }

    function deployAndSetGateway() external override returns (address) {
        return _deployAndSetGateway(msg.sender);
    }

    function updateWorkflowTotalSpent(
        string memory _depositAssetKey,
        uint256 _workflowId,
        uint256 _workflowExecutionAmount
    ) external override onlyExistingWorkflow(_workflowId) {
        require(msg.sender == address(_billingManager), "Registry: sender is not a billing manager");

        _workflowsData[_workflowId].depositAssetKeys.add(_depositAssetKey);
        _workflowsData[_workflowId].depositAssetsTotalSpent[_depositAssetKey] += _workflowExecutionAmount;
    }

    function registerWorkflows(RegisterWorkflowInfo[] calldata _registerWorkflowInfoArr) external override {
        bool isMainChain = _contractsRegistry.isMainChain();
        address signerAddr = _contractsRegistry.getSigner();

        if (isMainChain && maxWorkflowsPerAccount > 0) {
            require(
                workflowsPerAddress[msg.sender] + _registerWorkflowInfoArr.length <= maxWorkflowsPerAccount,
                "Registry: reached max workflows capacity"
            );
        }

        for (uint256 i = 0; i < _registerWorkflowInfoArr.length; i++) {
            RegisterWorkflowInfo calldata currentRegisterInfo = _registerWorkflowInfoArr[i];

            require(
                isMainChain ? currentRegisterInfo.workflowOwner == msg.sender : signerAddr == msg.sender,
                "Registry: not a sender or signer"
            );
            require(!isWorkflowRegistered(currentRegisterInfo.id), "Registry: workflow id already exists");

            if (currentRegisterInfo.requireGateway) {
                address currentGateway = _gateways[currentRegisterInfo.workflowOwner];

                if (currentRegisterInfo.deployGateway && currentGateway == address(0)) {
                    _deployAndSetGateway(currentRegisterInfo.workflowOwner);
                } else {
                    require(currentGateway != address(0), "Registry: gateway not found");
                }
            }

            _workflowsData[currentRegisterInfo.id].baseInfo = BaseWorkflowInfo(
                currentRegisterInfo.id,
                currentRegisterInfo.workflowOwner
            );
            workflowsPerAddress[currentRegisterInfo.workflowOwner]++;
            _existingWorkflowIds.push(currentRegisterInfo.id);

            emit WorkflowRegistered(currentRegisterInfo.workflowOwner, currentRegisterInfo.id);
        }
    }

    function perform(
        uint256 _workflowId,
        uint256 _gasAmount,
        bytes calldata _data,
        address _target
    ) external override onlySigner onlyExistingWorkflow(_workflowId) {
        require(address(this) != _target, "Registry: operation is not permitted");

        address existingGateway = _gateways[getWorkflowOwner(_workflowId)];

        require(existingGateway != address(0), "Registry: zero gateway address");

        bool success = _callWithExactGas(
            _gasAmount,
            existingGateway,
            abi.encodeWithSignature(GATEWAY_PERFORM_FUNC_SIGNATURE, _workflowId, _target, _data)
        );

        emit Performance(_workflowId, success);
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

    function getWorkflowsInfo(
        uint256 _offset,
        uint256 _limit
    ) external view override returns (WorkflowInfo[] memory _workflowsInfoArr) {
        uint256 to = Paginator.getTo(_existingWorkflowIds.length, _offset, _limit);

        _workflowsInfoArr = new WorkflowInfo[](to - _offset);

        for (uint256 i = _offset; i < to; i++) {
            _workflowsInfoArr[i - _offset] = getWorkflowInfo(_existingWorkflowIds[i]);
        }
    }

    function getBaseWorkflowInfo(uint256 _workflowId) public view override returns (BaseWorkflowInfo memory) {
        return _workflowsData[_workflowId].baseInfo;
    }

    function getWorkflowDepositAssetKeys(uint256 _workflowId) public view override returns (string[] memory) {
        return _workflowsData[_workflowId].depositAssetKeys.values();
    }

    function getWorkflowDepositAssetsInfo(
        uint256 _workflowId,
        string[] memory _depositAssetKeys
    ) public view override returns (DepositAssetInfo[] memory _depositAssetsArr) {
        _depositAssetsArr = new DepositAssetInfo[](_depositAssetKeys.length);

        for (uint256 i = 0; i < _depositAssetKeys.length; i++) {
            _depositAssetsArr[i] = DepositAssetInfo(
                _depositAssetKeys[i],
                _workflowsData[_workflowId].depositAssetsTotalSpent[_depositAssetKeys[i]]
            );
        }
    }

    function getWorkflowInfo(uint256 _workflowId) public view override returns (WorkflowInfo memory) {
        string[] memory depositAssetKeys = getWorkflowDepositAssetKeys(_workflowId);

        return
            WorkflowInfo(
                getBaseWorkflowInfo(_workflowId),
                depositAssetKeys,
                getWorkflowDepositAssetsInfo(_workflowId, depositAssetKeys)
            );
    }

    function getWorkflowOwner(uint256 _id) public view override returns (address) {
        return _workflowsData[_id].baseInfo.owner;
    }

    function isWorkflowRegistered(uint256 _id) public view override returns (bool) {
        return _workflowsData[_id].baseInfo.owner != address(0);
    }

    function _deployAndSetGateway(address _gatewayOwner) internal returns (address) {
        address newGatewayAddr = _gatewayFactory.deployGateway(_gatewayOwner);

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

    function _onlySigner() internal view {
        require(_contractsRegistry.getSigner() == msg.sender, "Registry: Not a signer");
    }

    function _onlyExistingWorkflow(uint256 _id) internal view {
        require(isWorkflowRegistered(_id), "Registry: workflow does not exist");
    }
}
