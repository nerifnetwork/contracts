// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

enum WorkflowStatus {
    PENDING,
    ACTIVE,
    PAUSED,
    CANCELLED
}

struct Workflow {
    uint256 id;
    address owner;
    bytes hash;
    WorkflowStatus status;
    uint256 totalSpent;
}

abstract contract RegistryWorkflow {
    mapping(uint256 => uint256) private indexMap;
    Workflow[] private workflows;
    mapping(address => uint64) private perAddress;

    modifier onlyExistingWorkflow(uint256 id) {
        Workflow memory workflow = getWorkflow(id);
        require(workflow.id > 0, "Registry: workflow does not exist");
        _;
    }

    modifier onlyWorkflowOwner(uint256 id) {
        Workflow memory workflow = getWorkflow(id);
        require(workflow.owner == msg.sender, "Registry: operation not permitted");
        _;
    }

    function getWorkflow(uint256 id) public view returns (Workflow memory) {
        require(_hasWorkflow(id), "Registry: workflow does not exist");
        return workflows[indexMap[id] - 1];
    }

    function getWorkflows() public view returns (Workflow[] memory) {
        return workflows;
    }

    function _addWorkflow(Workflow memory _workflow) internal returns (bool) {
        if (_hasWorkflow(_workflow.id)) {
            return false;
        }

        workflows.push(_workflow);
        indexMap[_workflow.id] = workflows.length;
        perAddress[_workflow.owner]++;

        _checkWorkflowEntry(_workflow.id);

        return true;
    }

    function _updateWorkflow(Workflow memory _workflow) internal returns (bool) {
        if (!_hasWorkflow(_workflow.id)) {
            return false;
        }

        workflows[indexMap[_workflow.id] - 1] = _workflow;

        _checkWorkflowEntry(_workflow.id);

        return true;
    }

    function _workflowsPerAddress(address owner) internal view returns (uint256) {
        return perAddress[owner];
    }

    function _checkWorkflowEntry(uint256 id) private view {
        uint256 index = indexMap[id];
        assert(index <= workflows.length);

        if (_hasWorkflow(id)) {
            assert(index > 0);
        } else {
            assert(index == 0);
        }
    }

    function _hasWorkflow(uint256 id) private view returns (bool) {
        return indexMap[id] > 0;
    }
}
