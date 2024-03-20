// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@solarity/solidity-lib/contracts-registry/AbstractDependant.sol";

import "../interfaces/core/IContractsRegistry.sol";

import "./Staking.sol";
import "../common/Globals.sol";

contract RewardDistributionPool is AbstractDependant {
    struct RewardPosition {
        uint256 balance;
        uint256 lastRewardPoints;
    }

    Staking internal _staking;

    uint256 public collectedRewards;
    uint256 public totalRewardPoints;
    uint256 public providedStake;

    mapping(address => RewardPosition) public rewardPositions;

    event CollectRewards(address validator, uint256 amount);

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function setDependencies(address _contractsRegistryAddr, bytes memory) public override dependant {
        IContractsRegistry contractsRegistry = IContractsRegistry(_contractsRegistryAddr);

        _staking = Staking(contractsRegistry.getStakingContract());
    }

    // solhint-disable-next-line ordering
    function distributeRewards() external {
        uint256 amount = address(this).balance;

        require(amount > 0, "RewardDistributionPool: amount must be greater than 0");

        _updateLastRewardPoints();

        providedStake = _staking.totalStake();
        totalRewardPoints += (amount * BASE_DIVISOR) / providedStake;
        collectedRewards += amount;
    }

    function collectRewards() external {
        claimRewards();
        _sendRewards(msg.sender);
    }

    function reinvestRewards() external {
        claimRewards();

        uint256 reward = rewardPositions[msg.sender].balance;

        _sendRewards(address(_staking));
        _staking.addRewardsToStake(msg.sender, reward);
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

        return (_staking.getStake(msg.sender) * newRewardPoints) / BASE_DIVISOR;
    }

    function _updateLastRewardPoints() private {
        address[] memory validators = _staking.getValidators();

        for (uint256 i = 0; i < validators.length; i++) {
            rewardPositions[validators[i]].lastRewardPoints = totalRewardPoints;
        }
    }

    function _sendRewards(address _receiver) private {
        uint256 reward = rewardPositions[msg.sender].balance;

        require(reward > 0, "RewardDistributionPool: reward must be greater than 0");

        rewardPositions[msg.sender].balance -= reward;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _receiver.call{value: reward, gas: 21000}("");
        require(success, "RewardDistributionPool: transfer failed");

        emit CollectRewards(msg.sender, reward);
    }

    function _createPath(address _from, address _to) private pure returns (address[] memory) {
        address[] memory path = new address[](2);
        path[0] = _from;
        path[1] = _to;

        return path;
    }
}
