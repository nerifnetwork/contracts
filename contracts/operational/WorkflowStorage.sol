// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// WorkflowStatus is the list of workflow statuses.
enum WorkflowStatus {
    PENDING,
    ACTIVE,
    PAUSED,
    CANCELLED
}

// Workflow is the workflow model.
struct Workflow {
    uint256 id;
    address owner;
    bytes hash;
    WorkflowStatus status;
    bool isInternal;
    uint256 totalSpent;
}

// WorkflowStorage is the workflows storage contract.
contract WorkflowStorage is Ownable, Initializable {
    address public registry;
    mapping(uint256 => uint256) internal indexMap;
    mapping(address => uint64) internal workflowsPerAddress;
    Workflow[] internal workflows;

    // onlyRegistry permits transactions coming from the registry contract.
    modifier onlyRegistry() {
        require(registry == msg.sender, "WorkflowStorage: operation is not permitted");
        _;
    }

    function initialize(address _registry, Workflow[] memory _workflows) external virtual initializer {
        registry = _registry;

        for (uint256 i = 0; i < _workflows.length; i++) {
            workflows.push(_workflows[i]);
            indexMap[_workflows[i].id] = workflows.length;
        }
    }

    function setRegistry(address _registry) external onlyOwner {
        registry = _registry;
    }

    function mustAdd(Workflow memory workflow) external onlyRegistry {
        require(_add(workflow), "WorkflowStorage: failed to add workflow");
    }

    function mustUpdate(Workflow memory workflow) external onlyRegistry {
        require(_update(workflow), "WorkflowStorage: failed to update workflow");
    }

    function mustRemove(uint256 id) external onlyRegistry {
        require(_remove(id), "WorkflowStorage: failed to remove workflow");
    }

    function clear() external onlyOwner returns (bool) {
        for (uint256 i = 0; i < workflows.length; i++) {
            delete indexMap[workflows[i].id];
        }

        delete workflows;

        return true;
    }

    function size() external view returns (uint256) {
        return workflows.length;
    }

    function getWorkflows() external view returns (Workflow[] memory) {
        return workflows;
    }

    function getWorkflowsPerAddress(address owner) external view returns (uint256) {
        return workflowsPerAddress[owner];
    }

    function getWorkflow(uint256 id) external view returns (Workflow memory) {
        return workflows[indexMap[id] - 1];
    }

    function contains(uint256 id) public view returns (bool) {
        return indexMap[id] > 0;
    }

    function _add(Workflow memory _workflow) private onlyRegistry returns (bool) {
        if (contains(_workflow.id)) {
            return false;
        }

        workflows.push(_workflow);
        indexMap[_workflow.id] = workflows.length;
        workflowsPerAddress[_workflow.owner]++;

        _checkEntry(_workflow.id);

        return true;
    }

    function _update(Workflow memory _workflow) private onlyRegistry returns (bool) {
        if (!contains(_workflow.id)) {
            return false;
        }

        workflows[indexMap[_workflow.id] - 1] = _workflow;

        _checkEntry(_workflow.id);

        return true;
    }

    function _remove(uint256 id) private onlyRegistry returns (bool) {
        if (!contains(id)) {
            return false;
        }

        uint256 index = indexMap[id];
        Workflow memory existingWorkflow = workflows[index - 1];

        uint256 lastListIndex = workflows.length - 1;
        Workflow memory lastListWorkflow = workflows[lastListIndex];
        if (lastListIndex != index - 1) {
            indexMap[lastListWorkflow.id] = index;
            workflows[index - 1] = lastListWorkflow;
        }

        workflows.pop();
        delete indexMap[id];
        workflowsPerAddress[existingWorkflow.owner]--;

        _checkEntry(id);

        if (lastListWorkflow.id != id) {
            _checkEntry(lastListWorkflow.id);
        }

        return true;
    }

    function _checkEntry(uint256 id) private view {
        uint256 index = indexMap[id];
        assert(index <= workflows.length);

        if (contains(id)) {
            assert(index > 0);
        } else {
            assert(index == 0);
        }
    }
}
