// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../interfaces/IGateway.sol";
import "../operational/RegistryWorkflow.sol";

contract TestRegistry is RegistryWorkflow {
    function callPerform(
        address _gatewayAddr,
        uint256 _id,
        address _target,
        bytes calldata _payload
    ) external {
        IGateway(_gatewayAddr).perform(_id, _target, _payload);
    }

    function addWorkflow(Workflow memory _workflow) external returns (bool) {
        return _addWorkflow(_workflow);
    }

    function updateWorkflow(Workflow memory _workflow) external returns (bool) {
        return _updateWorkflow(_workflow);
    }
}
