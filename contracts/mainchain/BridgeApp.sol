// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract BridgeApp is Ownable {
    mapping(uint256 => address) public contractAddresses;

    // mediator address, if 0x000 - no mediator
    address public mediatorAddress;

    event ContractAddressUpdated(uint256 chainId, address contractAddress);
    event UpdatedMediatorAddress(address mediatorAddress);

    constructor(address _owner) {
        _transferOwnership(_owner);
    }

    function setContractAddress(uint256 _chainId, address _contractAddress) public onlyOwner {
        contractAddresses[_chainId] = _contractAddress;
        emit ContractAddressUpdated(_chainId, _contractAddress);
    }

    function setMediator(address _mediatorAddress) public onlyOwner {
        mediatorAddress = _mediatorAddress;
        emit UpdatedMediatorAddress(mediatorAddress);
    }
}
