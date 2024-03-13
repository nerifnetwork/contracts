// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IMintableToken {
    function mint(address _recipient, uint256 _tokensAmount) external;
}
