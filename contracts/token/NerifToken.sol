// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";

import "../interfaces/core/IContractsRegistry.sol";
import "../interfaces/token/INerifToken.sol";

contract NerifToken is INerifToken, ERC20PermitUpgradeable, OwnableUpgradeable, AbstractDependant {
    address internal _tokensVestingAddr;

    modifier onlyTokensVesting() {
        _onlyTokensVesting();
        _;
    }

    function initialize(uint256 _initTokensAmount, string memory _name, string memory _symbol) external initializer {
        __Ownable_init();
        __ERC20_init(_name, _symbol);
        __ERC20Permit_init(_name);

        _mint(msg.sender, _initTokensAmount);
    }

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _tokensVestingAddr = contractsRegistry.getTokensVestingContract();
    }

    // solhint-disable-next-line ordering
    function ownerMint(address _recipient, uint256 _tokensAmount) external override onlyOwner {
        _mint(_recipient, _tokensAmount);
    }

    function vestingMint(address _recipient, uint256 _tokensAmount) external override onlyTokensVesting {
        _mint(_recipient, _tokensAmount);
    }

    function _onlyTokensVesting() internal view {
        require(_tokensVestingAddr == msg.sender, "NerifToken: Not a tokens vesting");
    }
}
