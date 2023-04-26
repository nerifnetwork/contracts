// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IBridgeValidatorFeePool {
    function setValidatorFeeReceiver(address _validatorFeeReceiver) external;

    function setLimitPerToken(address _token, uint256 _limit) external;

    function setERC20Bridge(address _erc20Bridge) external;

    function collect(address _token) external;
}
