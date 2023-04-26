// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./ValidatorOwnable.sol";
import "./ContractKeys.sol";

contract EventRegistry is ValidatorOwnable, Initializable, ContractKeys {
    enum EventType {
        SEND,
        REVERT_SEND,
        SET_SIGNER
    }

    mapping(bytes32 => bool) public registeredEvents;
    mapping(bytes32 => EventType) public eventType;

    event EventRegistered(
        bytes32 _hash,
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes _data,
        uint256 _validatorFee,
        EventType _eventType
    );

    function initialize(address _validatorGetterAddress) external initializer {
        _setValidatorGetter(_validatorGetterAddress);
    }

    function registerEvent(
        bytes32 _hash,
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _data,
        uint256 _validatorFee,
        EventType _eventType
    ) external onlyValidator {
        bytes32 key = this.eventKey(
            _hash,
            _appContract,
            _sourceChain,
            _destinationChain,
            _data,
            _validatorFee,
            _eventType
        );
        require(registeredEvents[key] == false, "EventRegistry: event is already registered");

        registeredEvents[key] = true;
        eventType[key] = _eventType;

        emit EventRegistered(_hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType);
    }

    function eventKey(
        bytes32 _hash,
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _data,
        uint256 _validatorFee,
        EventType _eventType
    ) external pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(_hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType)
            );
    }

    function getEventType(bytes32 _eventKey) public view returns (EventType) {
        require(registeredEvents[_eventKey] == true, "EventRegistry: event is not registered");

        return eventType[_eventKey];
    }
}
