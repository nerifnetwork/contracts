// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface INerifToken {
    function ownerMint(address _recipient, uint256 _tokensAmount) external;

    function vestingMint(address _recipient, uint256 _tokensAmount) external;
}
