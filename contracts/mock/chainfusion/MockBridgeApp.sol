// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../../interfaces/IBridgeApp.sol";
import "../../sidechain/RelayBridge.sol";

contract MockBridgeApp is IBridgeApp, Initializable {
    string public value;
    RelayBridge public relayBridge;

    event Reverted(bytes32 hash);

    function initialize(address _relayBridgeAddress) external initializer {
        relayBridge = RelayBridge(_relayBridgeAddress);
    }

    function send(
        uint256 _destinationChain,
        uint256 _gasLimit,
        bytes memory _data
    ) public {
        relayBridge.send(_destinationChain, _gasLimit, _data);
    }

    function execute(uint256, bytes memory _data) public {
        value = abi.decode(_data, (string));
    }

    function revertSend(uint256, bytes memory _data) public {
        emit Reverted(keccak256(_data));
    }

    function bridgeAppAddress() public pure returns (address) {
        return address(0);
    }
}
