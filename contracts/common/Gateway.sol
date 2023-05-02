// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IRegistry.sol";
import "../interfaces/IGateway.sol";

contract Gateway is Initializable, Ownable, IGateway {
    struct Config {
        uint256[] knownWorkflows;
        address[] knownWorkflowOwners;
        address[] knownCustomerContracts;
    }

    uint256 internal constant PERFORM_GAS_CUSHION = 5_000;

    IRegistry public registry;
    Config internal config;

    // onlyRegistry permits transactions coming from the registry contract.
    modifier onlyRegistry() {
        require(address(registry) == msg.sender, "Gateway: operation is not permitted");
        _;
    }

    // onlyAllowedWorkflow allows transactions passed through checks based on config.
    modifier onlyAllowedWorkflow(uint256 id, address target) {
        Workflow memory workflow = registry.getWorkflow(id);

        // Gateway owner's workflows can be executed anytime
        bool allowed = super.owner() == workflow.owner;

        // Check if the given workflow owner is known.
        if (!allowed) {
            for (uint256 i = 0; i < config.knownWorkflowOwners.length; i++) {
                if (config.knownWorkflowOwners[i] == workflow.owner) {
                    allowed = true;
                    break;
                }
            }
        }

        // Check if the given workflow is known.
        if (!allowed) {
            for (uint256 i = 0; i < config.knownWorkflows.length; i++) {
                if (config.knownWorkflows[i] == id) {
                    allowed = true;
                    break;
                }
            }
        }

        // Check if the given customer contract is allowed.
        if (!allowed) {
            for (uint256 i = 0; i < config.knownCustomerContracts.length; i++) {
                if (config.knownCustomerContracts[i] == target) {
                    allowed = true;
                    break;
                }
            }
        }

        require(allowed, "Gateway: operation is not permitted");
        _;
    }

    function initialize(address _registry) external initializer {
        registry = IRegistry(_registry);
    }

    function perform(
        uint256 id,
        address target,
        bytes calldata payload
    ) external onlyRegistry onlyAllowedWorkflow(id, target) {
        (bool success, ) = target.call(payload);
        require(success, "Gateway: failed to execute customer contract");
    }

    // setConfig sets the given configuration
    function setConfig(Config calldata _config) external onlyOwner {
        config = _config;
    }

    // getKnownWorkflows returns the list of known workflows
    function getConfig() external view returns (Config memory) {
        return config;
    }
}
