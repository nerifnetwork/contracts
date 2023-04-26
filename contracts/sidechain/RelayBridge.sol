// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Globals.sol";
import "./SignerStorage.sol";
import "../interfaces/IBridgeApp.sol";
import "../interfaces/IBridgeMediator.sol";
import "../interfaces/IBridgeValidatorFeePool.sol";

contract RelayBridge is Initializable {
    mapping(bytes32 => bytes) public sentData;

    mapping(bytes32 => bool) public sent;
    mapping(bytes32 => bool) public executed;
    mapping(bytes32 => bool) public reverted;
    mapping(bytes32 => bool) public failed;

    address[] public leaderHistory;

    SignerStorage public signerStorage;
    IBridgeValidatorFeePool public bridgeValidatorFeePool;

    uint256 public nonce;

    event Sent(
        bytes32 hash,
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        bytes data,
        uint256 gasLimit,
        uint256 nonce,
        uint256 validatorFee
    );
    event FailedSend(
        bytes32 hash,
        address appContract,
        uint256 sourceChain,
        uint256 destinationChain,
        bytes data,
        uint256 gasLimit,
        uint256 nonce,
        uint256 validatorFee
    );
    event Reverted(bytes32 hash, uint256 sourceChain, uint256 destinationChain);
    event Executed(bytes32 hash, uint256 sourceChain, uint256 destinationChain);

    modifier onlySigner() {
        require(signerStorage.getAddress() == msg.sender, "SignerOwnable: only signer");
        _;
    }

    function initialize(address _signerStorage, address payable _bridgeValidatorFeePool) external initializer {
        signerStorage = SignerStorage(_signerStorage);
        bridgeValidatorFeePool = IBridgeValidatorFeePool(_bridgeValidatorFeePool);
    }

    function send(
        uint256 _destinationChain,
        uint256 _gasLimit,
        bytes memory _data
    ) external payable {
        bytes32 hash = dataHash(msg.sender, block.chainid, _destinationChain, _gasLimit, _data, nonce);
        require(sentData[hash].length == 0, "RelayBridge: data already send");

        sent[hash] = true;
        sentData[hash] = _data;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = address(bridgeValidatorFeePool).call{value: msg.value, gas: 21000}("");
        require(success, "RelayBridge: transfer value failed");

        emit Sent(hash, msg.sender, block.chainid, _destinationChain, _data, _gasLimit, nonce, msg.value);

        nonce++;
    }

    function failedSend(
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        bytes memory _data,
        uint256 _gasLimit,
        uint256 _nonce,
        uint256 _validatorFee
    ) external onlySigner {
        bytes32 hash = dataHash(_appContract, _sourceChain, _destinationChain, _gasLimit, _data, _nonce);
        require(!failed[hash], "RelayBridge: data already failed");

        failed[hash] = true;
        emit FailedSend(hash, _appContract, _sourceChain, _destinationChain, _data, _gasLimit, _nonce, _validatorFee);
    }

    function revertSend(
        address _appContract,
        uint256 _destinationChain,
        uint256 _gasLimit,
        bytes memory _data,
        uint256 _nonce,
        address _leader
    ) external onlySigner {
        bytes32 hash = dataHash(_appContract, block.chainid, _destinationChain, _gasLimit, _data, _nonce);
        require(sent[hash], "RelayBridge: data never sent");
        require(!reverted[hash], "RelayBridge: data already reverted");

        reverted[hash] = true;
        leaderHistory.push(_leader);

        IBridgeApp(_appContract).revertSend(_destinationChain, _data);

        emit Reverted(hash, block.chainid, _destinationChain);
    }

    function execute(
        address _appContract,
        uint256 _sourceChain,
        uint256 _gasLimit,
        bytes memory _data,
        uint256 _nonce,
        address _leader
    ) external onlySigner {
        bytes32 hash = dataHash(_appContract, _sourceChain, block.chainid, _gasLimit, _data, _nonce);
        require(!executed[hash], "RelayBridge: data already executed");

        executed[hash] = true;
        leaderHistory.push(_leader);

        IBridgeApp(_appContract).execute(_sourceChain, _data);

        emit Executed(hash, _sourceChain, block.chainid);
    }

    function leaderHistoryLength() external view returns (uint256) {
        return leaderHistory.length;
    }

    function dataHash(
        address _appContract,
        uint256 _sourceChain,
        uint256 _destinationChain,
        uint256 _gasLimit,
        bytes memory _data,
        uint256 _nonce
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(_appContract, _sourceChain, _destinationChain, _gasLimit, _data, _nonce));
    }
}
