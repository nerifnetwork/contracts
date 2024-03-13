// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";

import "./ContractKeys.sol";
import "../common/ContractRegistry.sol";
import "../interfaces/IMintableToken.sol";

contract NerifToken is IMintableToken, ERC20PermitUpgradeable, ContractKeys {
    ContractRegistry public contractsRegistry;

    modifier onlyTokensVesting() {
        _onlyTokensVesting();
        _;
    }

    function initialize(address _contractsRegistry, string memory _name, string memory _symbol) external initializer {
        __ERC20_init(_name, _symbol);
        __ERC20Permit_init(_name);

        contractsRegistry = ContractRegistry(_contractsRegistry);
    }

    function mint(address _recipient, uint256 _tokensAmount) external override onlyTokensVesting {
        _mint(_recipient, _tokensAmount);
    }

    function _onlyTokensVesting() internal view {
        require(contractsRegistry.getContract(TOKENS_VESTING_KEY) == msg.sender, "NerifToken: Not a tokens vesting");
    }
}
