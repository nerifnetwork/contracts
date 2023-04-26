// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface SignerGetter {
    function getSignerAddress() external view returns (address);
}

abstract contract SignerOwnable {
    SignerGetter public signerGetter;

    modifier onlySigner() {
        require(signerGetter.getSignerAddress() == msg.sender, "SignerOwnable: only signer");
        _;
    }

    function _setSignerGetter(address _signerGetterAddress) internal virtual {
        signerGetter = SignerGetter(_signerGetterAddress);
    }
}
