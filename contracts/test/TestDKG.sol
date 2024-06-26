// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../system/DKG.sol";

contract TestDKG is DKG {
    function setSigner(uint256 _epochId, address _signer) external {
        _activeSigner = _signer;
        _dkgEpochsData[_epochId].epochSigner = _signer;
    }

    function setSlashedValidator(address _validatorAddr, bool _newValue) external {
        _slashedValidators[_validatorAddr] = _newValue;
    }

    function getStaking() external view returns (address) {
        return address(_staking);
    }
}
