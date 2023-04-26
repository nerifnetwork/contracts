// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IBridgeApp {
    function execute(uint256 _sourceChain, bytes memory _data) external;

    function revertSend(uint256 _destinationChain, bytes memory _data) external;

    // address of BridgeApp in CFN blockchain
    function bridgeAppAddress() external view returns (address);
}
