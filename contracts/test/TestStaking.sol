// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "../system/Staking.sol";

contract TestStaking is Staking {
    function dkgAddress() external view returns (address) {
        return address(_dkg);
    }
}
