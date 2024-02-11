// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

contract TestTarget {
    uint256 public number;

    function setNumber(uint256 _newNumber) external {
        number = _newNumber;
    }
}
