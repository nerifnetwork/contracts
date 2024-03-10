// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "../interfaces/system/IStaking.sol";
import "../system/SlashingVoting.sol";

contract TestSlashingVoting {
    mapping(address => mapping(SlashingReason => bool)) internal _testBans;

    function slash(address _stakingAddr, address _validator) external {
        IStaking(_stakingAddr).slash(_validator);
    }

    function isBannedByReason(address _validator, SlashingReason _reason) public view returns (bool) {
        return _testBans[_validator][_reason];
    }
}
