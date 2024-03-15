// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";

import "../common/ContractRegistry.sol";
import "../system/ContractKeys.sol";
import "../interfaces/token/INerifToken.sol";

contract NerifToken is INerifToken, ERC20PermitUpgradeable, OwnableUpgradeable, ContractKeys {
    ContractRegistry public contractsRegistry;

    modifier onlyTokensVesting() {
        _onlyTokensVesting();
        _;
    }

    function initialize(
        address _contractsRegistry,
        uint256 _initTokensAmount,
        string memory _name,
        string memory _symbol
    ) external initializer {
        __Ownable_init();
        __ERC20_init(_name, _symbol);
        __ERC20Permit_init(_name);

        contractsRegistry = ContractRegistry(_contractsRegistry);

        _mint(msg.sender, _initTokensAmount);
    }

    function ownerMint(address _recipient, uint256 _tokensAmount) external override onlyOwner {
        _mint(_recipient, _tokensAmount);
    }

    function vestingMint(address _recipient, uint256 _tokensAmount) external override onlyTokensVesting {
        _mint(_recipient, _tokensAmount);
    }

    function _onlyTokensVesting() internal view {
        require(contractsRegistry.getContract(TOKENS_VESTING_KEY) == msg.sender, "NerifToken: Not a tokens vesting");
    }
}
