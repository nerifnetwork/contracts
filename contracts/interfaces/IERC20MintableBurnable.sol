// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IERC20MintableBurnable {
    function mint(address _to, uint256 _amount) external;

    function burnFrom(address _account, uint256 _amount) external;
}
