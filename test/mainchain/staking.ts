import hre from 'hardhat';
import { ethers } from 'hardhat';
import { expect } from 'chai';
import { ValidatorStatusActive, ValidatorStatusInactive } from '../utils/helpers';
import { deploySystem } from '../utils/deploy';
import { hexValue } from 'ethers/lib/utils';

describe('Staking', function () {
  it('should change minimal stake', async function () {
    const newMinimalStake = ethers.utils.parseEther('2');
    const { staking } = await deploySystem();

    await expect(staking.setMinimalStake(newMinimalStake))
      .to.emit(staking, 'MinimalStakeUpdated')
      .withArgs(newMinimalStake);

    expect(await staking.minimalStake()).to.equal(newMinimalStake);
  });

  it('should set withdrawal period', async function () {
    const newWithdrawalPeriod = 2;
    const { staking } = await deploySystem();

    await expect(staking.setWithdrawalPeriod(newWithdrawalPeriod))
      .to.emit(staking, 'WithdrawalPeriodUpdated')
      .withArgs(newWithdrawalPeriod);

    expect(await staking.withdrawalPeriod()).to.equal(newWithdrawalPeriod);
  });

  it('should verify error at zero stake amount', async function () {
    const { staking } = await deploySystem();

    await expect(staking.stake({ value: 0 })).to.be.revertedWith('Staking: amount must be greater than zero');
  });

  it('should verify staking more than minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const additionalStake = ethers.utils.parseEther('5');
    const { staking, minimalStake } = await deploySystem();
    await staking.stake({ value: minimalStake.add(additionalStake) });

    const { validator, stake, status } = await staking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(minimalStake.add(additionalStake));
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check if it is possible to stake several times', async function () {
    const [owner] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    await staking.stake({ value: minimalStake });
    await staking.stake({ value: minimalStake });

    const { validator, stake, status } = await staking.stakes(owner.address);
    expect(validator).to.equal(owner.address);
    expect(stake).to.equal(minimalStake.mul(2));
    expect(status).to.equal(ValidatorStatusActive);
  });

  it('should check make validator active if new stake is more than minimum amount', async function () {
    const [owner] = await ethers.getSigners();
    let { staking, minimalStake } = await deploySystem();

    await staking.stake({ value: minimalStake.div(2) });
    var stake1 = await staking.stakes(owner.address);

    expect(stake1.status).to.equal(ValidatorStatusInactive);
    expect(stake1.stake).to.equal(minimalStake.div(2));

    await staking.stake({ value: minimalStake.div(2) });
    let stake2 = await staking.stakes(owner.address);

    expect(stake2.stake).to.equal(minimalStake);
    expect(stake2.validator).to.equal(owner.address);
    expect(stake2.status).to.equal(ValidatorStatusActive);
  });

  it('should check if only slashing voting can slash', async function () {
    const [v1, v2] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });

    expect(await staking.getValidators()).to.deep.equal([v1.address]);

    await expect(staking2.slash(v1.address)).to.be.revertedWith('Staking: not a slashing voting');
  });

  it('should check if validatorCount is decremented after slashing', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const nonce = ethers.utils.arrayify(0);
    const reason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');

    const { staking, addressStorage, slashingVoting, minimalStake } = await deploySystem();
    await slashingVoting.setSlashingEpochs(2);
    await slashingVoting.setSlashingThresold(2);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await slashingVoting.voteWithReason(v3.address, reason, nonce);
    await slashingVoting2.voteWithReason(v3.address, reason, nonce);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(await slashingVoting.epochPeriod())
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonce);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonce);

    blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(await slashingVoting.epochPeriod())
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    expect(await staking.isValidatorSlashed(v3.address)).to.equal(true);
    expect(await addressStorage.size()).to.be.equal(2);
  });

  it('should check if only validator with stake can announceWithdrawal', async function () {
    const [, v2] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });

    await expect(staking2.announceWithdrawal(minimalStake)).to.be.revertedWith('Staking: amount must be <= to stake');
  });

  it('should check if amount more than stake', async function () {
    const { staking, minimalStake } = await deploySystem();

    await staking.stake({ value: minimalStake });
    await expect(staking.announceWithdrawal(minimalStake.add(ethers.utils.parseEther('1')))).to.be.revertedWith(
      'Staking: amount must be <= to stake'
    );
  });

  it('should check if amount must be less or equal to stake', async function () {
    const [owner] = await ethers.getSigners();
    const { staking, minimalStake } = await deploySystem();

    await staking.stake({ value: minimalStake });
    await staking.announceWithdrawal(minimalStake);

    const { amount } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(minimalStake);
  });

  it('should check if only validator with stake can withdraw', async function () {
    const [, v2] = await ethers.getSigners();
    const { staking, minimalStake, stakeWithdrawalPeriod } = await deploySystem();

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking.stake({ value: minimalStake });
    await staking.announceWithdrawal(minimalStake);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);

    await expect(staking2.withdraw()).to.be.revertedWith('Staking: amount must be greater than zero');
  });

  it('should check if slashed validator can not announce withdraw', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const nonce = ethers.utils.arrayify(0);
    const reason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');
    const { staking, slashingVoting, minimalStake } = await deploySystem();
    await slashingVoting.setSlashingEpochs(2);
    await slashingVoting.setSlashingThresold(2);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    await slashingVoting.voteWithReason(v3.address, reason, nonce);
    await slashingVoting2.voteWithReason(v3.address, reason, nonce);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(await slashingVoting.epochPeriod())
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonce);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonce);

    expect(await staking.isValidatorSlashed(v3.address)).to.equal(true);

    await expect(staking3.announceWithdrawal(minimalStake)).to.be.revertedWith('Staking: validator is slashed');
  });

  it('should check if slashed validator can not withdraw', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const nonce = ethers.utils.arrayify(0);
    const firstreason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');

    const { staking, slashingVoting, minimalStake } = await deploySystem();
    await slashingVoting.setSlashingEpochs(2);
    await slashingVoting.setSlashingThresold(2);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });
    await staking3.stake({ value: minimalStake });

    expect(await staking.getValidatorsCount()).to.equal(3);
    const validatorsStake = await staking.listValidators(0, 3);

    expect(validatorsStake.length).to.equal(3);
    expect(validatorsStake[0].stake).to.equal(minimalStake);
    expect(validatorsStake[1].stake).to.equal(minimalStake);
    expect(validatorsStake[2].stake).to.equal(minimalStake);

    await slashingVoting.voteWithReason(v3.address, firstreason, nonce);
    await slashingVoting2.voteWithReason(v3.address, firstreason, nonce);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(await slashingVoting.epochPeriod())
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonce);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonce);

    expect(await staking.isValidatorSlashed(v3.address)).to.equal(true);
    await expect(staking3.withdraw()).to.be.revertedWith('Staking: validator is slashed');
  });

  it('should be able to withdraw even if no longer a validator', async function () {
    const [owner] = await ethers.getSigners();
    const minimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const firstWithdrawal = ethers.utils.parseEther('3');
    const secondWithdrawal = ethers.utils.parseEther('2');

    const { staking, stakeWithdrawalPeriod } = await deploySystem({ minimalStake });

    await staking.stake({ value: value });

    await staking.announceWithdrawal(firstWithdrawal);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();

    await staking.announceWithdrawal(secondWithdrawal);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();

    const { stake } = await staking.stakes(owner.address);

    expect(stake).to.equal(0);
  });

  it('should check if withdrawalPeriod not passed', async function () {
    const { staking, minimalStake } = await deploySystem();
    await staking.setWithdrawalPeriod(1000000);
    await staking.stake({ value: minimalStake });
    await staking.announceWithdrawal(minimalStake);

    await expect(staking.withdraw()).to.be.revertedWith('Staking: withdrawal period not passed');
  });

  it('should withdraw', async function () {
    const [owner] = await ethers.getSigners();
    const { staking, minimalStake, stakeWithdrawalPeriod, addressStorage } = await deploySystem();

    await staking.stake({ value: minimalStake });
    expect(await addressStorage.contains(owner.address)).to.equal(true);

    await staking.announceWithdrawal(minimalStake);

    expect(await addressStorage.contains(owner.address)).to.equal(false);

    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();

    const { amount, time } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });

  it('should revoke withdrawal', async function () {
    const [owner] = await ethers.getSigners();
    const { staking, minimalStake, addressStorage } = await deploySystem();

    await staking.stake({ value: minimalStake });

    await staking.announceWithdrawal(minimalStake);
    expect(await addressStorage.contains(owner.address)).to.equal(false);

    await staking.revokeWithdrawal();
    expect(await addressStorage.contains(owner.address)).to.equal(true);

    const { amount, time } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });

  it('should check if DKG validators are being updated after staking and slashing', async function () {
    const [owner, v1, v2, v3] = await ethers.getSigners();
    const minimalStake = ethers.utils.parseEther('3');
    const stake = ethers.utils.parseEther('5');
    const withdrawStake = ethers.utils.parseEther('10');

    const { staking, dkg, stakeWithdrawalPeriod } = await deploySystem({ minimalStake });
    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    expect((await dkg.getCurrentValidators()).length).to.equal(0);

    await staking1.stake({ value: stake });
    await staking1.stake({ value: stake });
    await staking2.stake({ value: stake });
    await staking3.stake({ value: stake });

    expect((await dkg.getCurrentValidators()).length).to.equal(3);
    await staking1.announceWithdrawal(withdrawStake);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking1.withdraw();
    expect((await dkg.getCurrentValidators()).length).to.equal(2);

    const { amount, time } = await staking.withdrawalAnnouncements(owner.address);

    expect(amount).to.equal(0);
    expect(time).to.equal(0);
  });

  it('should user can able to stake with the difference of the current stake and minimal stake', async function () {
    const [owner] = await ethers.getSigners();
    const minimalStake = ethers.utils.parseEther('3');
    const value = ethers.utils.parseEther('5');
    const withdrawal = ethers.utils.parseEther('3');
    const secondStake = ethers.utils.parseEther('2');

    const { staking, stakeWithdrawalPeriod } = await deploySystem({ minimalStake });

    await staking.stake({ value: value });
    await staking.announceWithdrawal(withdrawal);
    await hre.network.provider.send('hardhat_mine', [stakeWithdrawalPeriod.toHexString(), '0x1']);
    await staking.withdraw();
    await staking.stake({ value: secondStake });

    const { status } = await staking.stakes(owner.address);
    expect(status).to.equal(1);
  });

  it('should check if slashed validator can stake', async function () {
    const [, v2, v3] = await ethers.getSigners();
    const value = ethers.utils.parseEther('5');
    const nonce = ethers.utils.arrayify(0);
    const reason: number = 0;
    const secondReason: number = 1;
    const hre = require('hardhat');
    const { staking, slashingVoting } = await deploySystem();
    await slashingVoting.setSlashingEpochs(2);
    await slashingVoting.setSlashingThresold(2);

    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);
    const slashingVoting2 = await ethers.getContractAt('SlashingVoting', slashingVoting.address, v2);
    const staking3 = await ethers.getContractAt('Staking', staking.address, v3);

    await staking.stake({ value: value });
    await staking2.stake({ value: value });
    await staking3.stake({ value: value });

    await slashingVoting.voteWithReason(v3.address, reason, nonce);
    await slashingVoting2.voteWithReason(v3.address, reason, nonce);

    var blocksStep = (await slashingVoting.currentEpoch())
      .add(1)
      .mul(await slashingVoting.epochPeriod())
      .sub(await ethers.provider.getBlockNumber());
    await hre.network.provider.send('hardhat_mine', [hexValue(blocksStep)]);

    await slashingVoting.voteWithReason(v3.address, secondReason, nonce);
    await slashingVoting2.voteWithReason(v3.address, secondReason, nonce);

    expect(await staking.isValidatorSlashed(v3.address)).to.equal(true);
    await expect(staking3.stake({ value })).to.be.revertedWith('Staking: validator is slashed');
  });
});
