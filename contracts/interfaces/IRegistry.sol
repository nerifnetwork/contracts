// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

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

// IRegistry represents the public behaviour of the registry contract.
interface IRegistry {
    // getWorkflow returns the workflow by its ID.
    function getWorkflow(uint256 id) external view returns (Workflow memory workflow);
}
