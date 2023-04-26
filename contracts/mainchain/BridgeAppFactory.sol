// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./BridgeApp.sol";

contract BridgeAppFactory {
    address[] public apps;

    event BridgeAppCreated(address contractAddress, address owner);

    function createApp() public returns (BridgeApp) {
        BridgeApp bridgeApp = new BridgeApp(msg.sender);
        apps.push(address(bridgeApp));

        emit BridgeAppCreated(address(bridgeApp), msg.sender);

        return bridgeApp;
    }
}
