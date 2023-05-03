import { ethers } from 'hardhat';
import { expect } from 'chai';
import { BigNumber } from 'ethers';
import { deployMainchain } from '../utils/deploy';

describe('RewardDistributionPool', function () {
  it('should collect rewards', async function () {
    const [signer, v2, v3, v4] = await ethers.getSigners();
    const totalReward = ethers.utils.parseEther('0.1');
    const minimalStake = ethers.utils.parseEther('100');

    const { rewardDistributionPool, staking } = await deployMainchain();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);
    const staking4 = await ethers.getContractAt('Staking', staking.address, v4);

    const rewardDistributionPool2 = await ethers.getContractAt(
      'RewardDistributionPool',
      rewardDistributionPool.address,
      v2
    );
    const rewardDistributionPool3 = await ethers.getContractAt(
      'RewardDistributionPool',
      rewardDistributionPool.address,
      v3
    );

    const rewardDistributionPool4 = await ethers.getContractAt(
      'RewardDistributionPool',
      rewardDistributionPool.address,
      v4
    );

    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await expect(rewardDistributionPool2.collectRewards()).to.be.revertedWith(
      'RewardDistributionPool: reward must be greater than 0'
    );
    await expect(rewardDistributionPool3.collectRewards()).to.be.revertedWith(
      'RewardDistributionPool: reward must be greater than 0'
    );

    await expect(rewardDistributionPool.distributeRewards()).to.be.revertedWith(
      'RewardDistributionPool: amount must be greater than 0'
    );

    await (
      await signer.sendTransaction({
        to: rewardDistributionPool.address,
        value: totalReward,
      })
    ).wait();

    await rewardDistributionPool.distributeRewards();

    let validator3BalanceBefore = await ethers.provider.getBalance(v3.address);
    let validator2BalanceBefore = await ethers.provider.getBalance(v2.address);

    let reward = totalReward.div(BigNumber.from('2'));

    await expect(rewardDistributionPool3.collectRewards())
      .to.emit(rewardDistributionPool3, 'CollectRewards')
      .withArgs(v3.address, reward);
    await expect(rewardDistributionPool2.collectRewards())
      .to.emit(rewardDistributionPool2, 'CollectRewards')
      .withArgs(v2.address, reward);

    let validator3BalanceAfte = await ethers.provider.getBalance(v3.address);
    let validator2BalanceAfte = await ethers.provider.getBalance(v2.address);

    let txFee1 = validator3BalanceBefore.add(reward).sub(validator3BalanceAfte);
    let txFee2 = validator2BalanceBefore.add(reward).sub(validator2BalanceAfte);

    expect(validator3BalanceAfte).to.equal(validator3BalanceBefore.add(reward).sub(txFee1));
    expect(validator2BalanceAfte).to.equal(validator2BalanceBefore.add(reward).sub(txFee2));

    await expect(rewardDistributionPool2.collectRewards()).to.be.revertedWith(
      'RewardDistributionPool: reward must be greater than 0'
    );
    await expect(rewardDistributionPool3.collectRewards()).to.be.revertedWith(
      'RewardDistributionPool: reward must be greater than 0'
    );

    await staking4.stake({ value: minimalStake });

    await (
      await signer.sendTransaction({
        to: rewardDistributionPool.address,
        value: totalReward,
      })
    ).wait();

    await rewardDistributionPool.distributeRewards();

    let validator4BalanceBefore = await ethers.provider.getBalance(v4.address);
    validator3BalanceBefore = await ethers.provider.getBalance(v3.address);
    validator2BalanceBefore = await ethers.provider.getBalance(v2.address);

    reward = BigNumber.from('33333333333333300');

    await expect(rewardDistributionPool4.collectRewards())
      .to.emit(rewardDistributionPool4, 'CollectRewards')
      .withArgs(v4.address, reward);
    await expect(rewardDistributionPool3.collectRewards())
      .to.emit(rewardDistributionPool3, 'CollectRewards')
      .withArgs(v3.address, reward);
    await expect(rewardDistributionPool2.collectRewards())
      .to.emit(rewardDistributionPool2, 'CollectRewards')
      .withArgs(v2.address, reward);

    let validator4BalanceAfte = await ethers.provider.getBalance(v4.address);
    validator3BalanceAfte = await ethers.provider.getBalance(v3.address);
    validator2BalanceAfte = await ethers.provider.getBalance(v2.address);

    let txFee3 = validator4BalanceBefore.add(reward).sub(validator4BalanceAfte);
    txFee1 = validator3BalanceBefore.add(reward).sub(validator3BalanceAfte);
    txFee2 = validator2BalanceBefore.add(reward).sub(validator2BalanceAfte);

    expect(validator4BalanceAfte).to.equal(validator4BalanceBefore.add(reward).sub(txFee3));
    expect(validator3BalanceAfte).to.equal(validator3BalanceBefore.add(reward).sub(txFee1));
    expect(validator2BalanceAfte).to.equal(validator2BalanceBefore.add(reward).sub(txFee2));

    await rewardDistributionPool.distributeRewards();

    await (
      await signer.sendTransaction({
        to: rewardDistributionPool.address,
        value: totalReward,
      })
    ).wait();

    reward = BigNumber.from('33333333333333300');

    await rewardDistributionPool.distributeRewards();

    const validator4StakeBefore = await staking.getStake(v4.address);
    const validator3StakeBefore = await staking.getStake(v3.address);
    const validator2StakeBefore = await staking.getStake(v2.address);

    await expect(rewardDistributionPool4.reinvestRewards())
      .to.emit(rewardDistributionPool4, 'CollectRewards')
      .withArgs(v4.address, reward);
    await expect(rewardDistributionPool3.reinvestRewards())
      .to.emit(rewardDistributionPool3, 'CollectRewards')
      .withArgs(v3.address, reward);
    await expect(rewardDistributionPool2.reinvestRewards())
      .to.emit(rewardDistributionPool2, 'CollectRewards')
      .withArgs(v2.address, reward);

    const validator4StakeAfter = await staking.getStake(v4.address);
    const validator3StakeAfter = await staking.getStake(v3.address);
    const validator2StakeAfter = await staking.getStake(v2.address);

    expect(validator4StakeAfter).to.equal(validator4StakeBefore.add(reward));
    expect(validator3StakeAfter).to.equal(validator3StakeBefore.add(reward));
    expect(validator2StakeAfter).to.equal(validator2StakeBefore.add(reward));
  });
});
