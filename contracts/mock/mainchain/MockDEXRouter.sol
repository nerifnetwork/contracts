// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract MockDEXRouter is Initializable {
    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function swapExactTokensForETH(
        uint256 _amountIn,
        uint256 _amountOutMin,
        address[] calldata _path,
        address to,
        uint256 _deadline
    ) external returns (uint256[] memory amounts) {
        require(address(this).balance >= _amountIn, "MockDEXRouter: not enough balance");

        IERC20(_path[0]).transferFrom(msg.sender, address(this), _amountIn);

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = to.call{value: _amountIn, gas: 21000}("");
        require(success, "MockDEXRouter: transfer failed");

        amounts = new uint256[](2);
        amounts[0] = _amountOutMin;
        amounts[1] = _deadline;

        return amounts;
    }

    function wcfn() external pure returns (address) {
        return address(0);
    }

    function getAmountsOut(uint256 _amountIn, address[] calldata path)
        external
        pure
        returns (uint256[] memory amounts)
    {
        amounts = new uint256[](2);
        amounts[0] = _amountIn;
        amounts[1] = path.length;

        return amounts;
    }
}
