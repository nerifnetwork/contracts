// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IBridgeMediator {
    function mediate(
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _data
    ) external view returns (bytes memory);
}
