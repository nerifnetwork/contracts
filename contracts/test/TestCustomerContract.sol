// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

contract TestCustomerContract {
    bool public onOffButton;

    mapping(uint256 => uint256) public counter;
    mapping(uint256 => uint256) public lastTimeStamp;

    event Logger(uint256 timestamp, uint256 blocknbr, uint256 curcount);
    event Burn(uint256 amount);

    constructor() {
        onOffButton = true;
    }

    function setOnOffButton(bool newval) external {
        onOffButton = newval;
    }

    function perform(bytes calldata performData) external {
        uint256 baseGasAtm = gasleft();

        (uint256 id, uint256 interval, uint256 burnAmount) = abi.decode(performData, (uint256, uint256, uint256));

        // We highly recommend revalidating the upkeep in the performUpkeep function
        if ((block.timestamp - lastTimeStamp[id]) >= interval) {
            lastTimeStamp[id] = block.timestamp;
            counter[id] += 1;
            emit Logger(block.timestamp, block.number, counter[id]);

            if (burnAmount > 0) {
                emit Burn(burnAmount);
                uint256 gasLeftAmt = gasleft();
                burn(gasLeftAmt, gasLeftAmt - burnAmount + (baseGasAtm - gasLeftAmt) + 22000);
            }
        }
    }

    function check(bytes calldata checkData) external view returns (bool upkeepNeeded, bytes memory performData) {
        (uint256 id, uint256 interval, ) = abi.decode(checkData, (uint256, uint256, uint256));

        upkeepNeeded = false;
        if (onOffButton) {
            upkeepNeeded = (block.timestamp - lastTimeStamp[id]) >= interval;
        }

        if (upkeepNeeded) {
            performData = checkData;
        }
    }

    function burn(uint256 start, uint256 end) internal {
        while (gasleft() <= start && gasleft() > end) {}
    }
}
