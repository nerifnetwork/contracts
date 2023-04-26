// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract SignerStorage is Initializable {
    address public signer;

    event SignerUpdated(address signer);

    modifier onlySigner() {
        require(this.getAddress() == msg.sender, "SignerOwnable: only signer");
        _;
    }

    function initialize(address _signer) external initializer {
        signer = _signer;
    }

    function setAddress(address _newSigner) public payable {
        require(signer == msg.sender, "SignerStorage: only signer");
        signer = _newSigner;

        uint256 _value = msg.value;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _newSigner.call{value: _value, gas: 21000}("");
        require(success, "SignerStorage: transfer value failed");

        emit SignerUpdated(_newSigner);
    }

    function getAddress() public view returns (address) {
        return signer;
    }
}
