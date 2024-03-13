// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@solarity/solidity-lib/finance/vesting/Vesting.sol";

import "./ContractKeys.sol";
import "../interfaces/SignerOwnable.sol";
import "../interfaces/IMintableToken.sol";
import "../common/ContractRegistry.sol";

contract TokensVesting is Vesting, SignerOwnable, ContractKeys {
    ContractRegistry public contractsRegistry;

    function initialize(address _signerGetterAddress, address _contractsRegistry) external initializer {
        contractsRegistry = ContractRegistry(_contractsRegistry);

        _setSignerGetter(_signerGetterAddress);
    }

    function createSchedule(BaseSchedule memory _baseSchedule) external onlySigner returns (uint256) {
        return _createSchedule(_baseSchedule);
    }

    function createVesting(VestingData memory _vesting) external onlySigner returns (uint256) {
        return _createVesting(_vesting);
    }

    function _releaseTokens(address _token, address _beneficiary, uint256 _amountToPay) internal override {
        IMintableToken(_token).mint(_beneficiary, _amountToPay);
    }
}
