// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "../token/NerifToken.sol";

contract TestNerifToken is NerifToken {
    function mint(address _to, uint256 _mintAmount) external {
        _mint(_to, _mintAmount);
    }

    function getTimestamp() external view returns (uint256) {
        return block.timestamp;
    }

    function getChainId() external view returns (uint256) {
        return block.chainid;
    }

    function tokensVestingAddress() external view returns (address) {
        return _tokensVestingAddr;
    }
}
