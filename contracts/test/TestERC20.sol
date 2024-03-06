// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract TestERC20 is ERC20Permit {
    constructor(
        string memory _name,
        string memory _symbol,
        uint256 _initMintAmount
    ) ERC20Permit(_name) ERC20(_name, _symbol) {
        _mint(msg.sender, _initMintAmount);
    }

    function mint(address _to, uint256 _mintAmount) external {
        _mint(_to, _mintAmount);
    }

    function getTimestamp() external view returns (uint256) {
        return block.timestamp;
    }

    function getChainId() external view returns (uint256) {
        return block.chainid;
    }
}
