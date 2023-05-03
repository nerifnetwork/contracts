// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ValidatorGetter {
    function isValidatorActive(address _sender) external view returns (bool);
}

abstract contract ValidatorOwnable {
    ValidatorGetter public validatorGetter;

    modifier onlyValidator() {
        require(validatorGetter.isValidatorActive(msg.sender), "ValidatorOwnable: only validator");
        _;
    }

    function _setValidatorGetter(address _validatorGetterAddress) internal virtual {
        validatorGetter = ValidatorGetter(_validatorGetterAddress);
    }
}
