// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../interfaces/IDEXV2Router02.sol";
import "./ContractRegistry.sol";
import "./Staking.sol";
import "./ContractKeys.sol";
import "./SignerOwnable.sol";
import "../sidechain/Globals.sol";

contract ValidatorRewardDistributionPool is Initializable, ContractKeys, SignerOwnable {
    struct RewardPosition {
        uint256 balance;
        uint256 lastRewardPoints;
    }

    ContractRegistry public contractRegistry;
    IDEXV2Router02 public dexRouter;

    uint256 public collectedRewards;
    uint256 public totalRewardPoints;
    uint256 public providedStake;

    mapping(address => RewardPosition) public rewardPositions;

    event CollectRewards(address validator, uint256 amount);
    event RouterUpdated(address router);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address _contractRegistry,
        address _router,
        address _signerGetterAddress
    ) external initializer {
        _setSignerGetter(_signerGetterAddress);
        setRouter(_router);
        contractRegistry = ContractRegistry(_contractRegistry);
    }

    function setRouter(address _router) public onlySigner {
        dexRouter = IDEXV2Router02(_router);

        emit RouterUpdated(_router);
    }

    function distributeRewards() public {
        uint256 amount = address(this).balance;

        require(amount > 0, "ValidatorRewardDistributionPool: amount must be greater than 0");

        _updateLastRewardPoints();

        providedStake = _stakingContract().getTotalStake();
        totalRewardPoints += (amount * BASE_DIVISOR) / providedStake;
        collectedRewards += amount;
    }

    function distribute(address _token) public {
        // solhint-disable-next-line not-rely-on-time
        uint256 deadline = block.timestamp + 1;
        uint256 amountIn = IERC20(_token).balanceOf(address(this));

        address[] memory path = _createPath(_token, dexRouter.wcfn());
        uint256[] memory amountsOut = dexRouter.getAmountsOut(amountIn, path);

        IERC20(_token).approve(address(dexRouter), amountIn);

        dexRouter.swapExactTokensForETH(amountIn, amountsOut[amountsOut.length - 1], path, address(this), deadline);
    }

    function collectRewards() public {
        claimRewards();
        _sendRewards(msg.sender);
    }

    function reinvestRewards() public {
        claimRewards();

        uint256 reward = rewardPositions[msg.sender].balance;

        _sendRewards(address(_stakingContract()));
        _stakingContract().addRewardsToStake(msg.sender, reward);
    }

    function claimRewards() public {
        uint256 rewardsOwingAmount = rewardsOwing();
        if (rewardsOwingAmount > 0) {
            collectedRewards -= rewardsOwingAmount;
            rewardPositions[msg.sender].balance += rewardsOwingAmount;
            rewardPositions[msg.sender].lastRewardPoints = totalRewardPoints;
        }
    }

    function rewardsOwing() public view returns (uint256) {
        uint256 newRewardPoints = totalRewardPoints - rewardPositions[msg.sender].lastRewardPoints;

        return (_stakingContract().getStake(msg.sender) * newRewardPoints) / BASE_DIVISOR;
    }

    function _updateLastRewardPoints() private {
        address[] memory validators = _stakingContract().getValidators();
        for (uint256 i = 0; i < validators.length; i++) {
            rewardPositions[validators[i]].lastRewardPoints = totalRewardPoints;
        }
    }

    function _sendRewards(address _receiver) private {
        uint256 reward = rewardPositions[msg.sender].balance;

        require(reward > 0, "ValidatorRewardDistributionPool: reward must be greater than 0");

        rewardPositions[msg.sender].balance -= reward;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _receiver.call{value: reward, gas: 21000}("");
        require(success, "ValidatorRewardDistributionPool: transfer failed");

        emit CollectRewards(msg.sender, reward);
    }

    function _stakingContract() private view returns (Staking) {
        return Staking(payable(contractRegistry.getContract(STAKING_KEY)));
    }

    function _createPath(address _from, address _to) private pure returns (address[] memory) {
        address[] memory path = new address[](2);
        path[0] = _from;
        path[1] = _to;

        return path;
    }
}
