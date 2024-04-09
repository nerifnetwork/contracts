// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../system/DKG.sol";

contract TestDKG is DKG {
    function getStaking() external view returns (address) {
        return address(_staking);
    }
}
