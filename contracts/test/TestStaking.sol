// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/interfaces/IERC20.sol";

import "../system/Staking.sol";

contract TestStaking is Staking {
    function setStakeToken(address _newToken) external {
        stakeToken = IERC20(_newToken);
    }

    function dkgAddress() external view returns (address) {
        return address(_dkg);
    }
}
