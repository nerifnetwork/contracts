// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "../system/SlashingVoting.sol";

contract TestSlashingVoting is SlashingVoting {
    function getStaking() external view returns (address) {
        return address(_staking);
    }
}
