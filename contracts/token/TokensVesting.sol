// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "@solarity/solidity-lib/finance/vesting/Vesting.sol";

import "../interfaces/token/INerifToken.sol";
import "../interfaces/token/ITokensVesting.sol";

contract TokensVesting is ITokensVesting, Vesting, OwnableUpgradeable {
    function initialize() external initializer {
        __Ownable_init();
    }

    function createSchedule(BaseSchedule memory _baseSchedule) external override onlyOwner returns (uint256) {
        return _createSchedule(_baseSchedule);
    }

    function createVesting(VestingData memory _vesting) external override onlyOwner returns (uint256) {
        return _createVesting(_vesting);
    }

    function _releaseTokens(address _token, address _beneficiary, uint256 _amountToPay) internal override {
        INerifToken(_token).vestingMint(_beneficiary, _amountToPay);
    }
}
