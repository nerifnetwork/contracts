// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IERC20MintableBurnable.sol";

contract MockMintableBurnableToken is ERC20, IERC20MintableBurnable, ERC20Burnable, Ownable {
    constructor(string memory _name, string memory _symbol)
        ERC20(_name, _symbol)
    // solhint-disable-next-line no-empty-blocks
    {

    }

    function mint(address _to, uint256 _amount) public onlyOwner {
        _mint(_to, _amount);
    }

    function burnFrom(address _account, uint256 _amountt) public override(IERC20MintableBurnable, ERC20Burnable) {
        ERC20Burnable.burnFrom(_account, _amountt);
    }
}
