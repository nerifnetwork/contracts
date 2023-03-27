// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

contract TestCustomerContract {
    bool public onOffButton;

    mapping (uint256 => uint) public counter;
    mapping (uint256 => uint) public lastTimeStamp;

    event Logger(uint timestamp, uint blocknbr, uint curcount);
    event Burn(uint256 amount);

    constructor() {
        onOffButton = true;
    }

    function setOnOffButton(bool newval) external{
        onOffButton = newval;
    }

    function check(bytes calldata checkData) external view returns (bool upkeepNeeded, bytes memory performData) {
        (
            uint256 id,
            uint interval,
            uint256 burnAmount
        ) = abi.decode(checkData, (uint256, uint, uint256));

        upkeepNeeded = false;
        if (onOffButton){
            upkeepNeeded = (block.timestamp - lastTimeStamp[id]) >= interval;
        }

        if (upkeepNeeded) {
            performData = checkData;
        }
    }

    function perform(bytes calldata performData) external {
        uint256 base_gas_atm = gasleft();

        (
            uint256 id,
            uint interval,
            uint256 burnAmount
        ) = abi.decode(performData, (uint256, uint, uint256));

        // We highly recommend revalidating the upkeep in the performUpkeep function
        if ((block.timestamp - lastTimeStamp[id]) >= interval) {
            lastTimeStamp[id] = block.timestamp;
            counter[id] += 1;
            emit Logger(block.timestamp,block.number, counter[id]);

            if (burnAmount > 0) {
                emit Burn(burnAmount);
                uint256 gas_left_amt = gasleft();
                burn(gas_left_amt, gas_left_amt - burnAmount + (base_gas_atm - gas_left_amt) + 22000);
            }
        }
    }

    function burn(uint256 start, uint256 end) internal {
        while (gasleft() <= start && gasleft() > end) {}
    }
}