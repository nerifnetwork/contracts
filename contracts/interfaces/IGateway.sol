// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// IGateway represents the customer gateway contract behaviour.
interface IGateway {
    // perform is the entrypoint function of all customer contracts.
    // This function accepts the workflow ID and the end customer contract
    // execution payload.
    // The function checks that the given transaction is permitted and can be forwarded
    // next to the end customer contract.
    function perform(
        uint256 id,
        address target,
        bytes calldata payload
    ) external;
}
