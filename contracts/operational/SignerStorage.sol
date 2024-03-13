// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../interfaces/ISignerAddress.sol";
import "../interfaces/SignerOwnable.sol";

// SignerStorage is the simple collective signer address storage contract.
// This is used on sidechains only. Mainchain uses DKG contract instead.
contract SignerStorage is Initializable, ISignerAddress, SignerOwnable {
    address private signer;

    event SignerUpdated(address signer);

    function initialize(address _signer) external initializer {
        _setSignerGetter(address(this));
        signer = _signer;
    }

    // setAddress updates signer address
    function setAddress(address _newSigner) public payable onlySigner {
        signer = _newSigner;

        uint256 _value = msg.value;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _newSigner.call{value: _value, gas: 21000}("");
        require(success, "SignerStorage: transfer value failed");

        emit SignerUpdated(_newSigner);
    }

    // getSignerAddress implements ISignerAddress interface
    function getSignerAddress() public view returns (address) {
        return signer;
    }
}
