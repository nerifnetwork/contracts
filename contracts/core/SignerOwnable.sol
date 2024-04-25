// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/core/ISignerAddress.sol";

abstract contract SignerOwnable {
    ISignerAddress public signerGetter;

    modifier onlySigner() {
        require(signerGetter.getSignerAddress() == msg.sender, "SignerOwnable: only signer");
        _;
    }

    function _setSignerGetter(address _signerGetterAddress) internal virtual {
        signerGetter = ISignerAddress(_signerGetterAddress);
    }
}
