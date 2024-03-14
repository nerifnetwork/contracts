import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractRegistry, NerifToken, SignerStorage, TokensVesting, Vesting } from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setTime } from '../helpers/block-helper';

describe('TokensVesting', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let tokensVesting: TokensVesting;
  let nerifToken: NerifToken;
  let contractsRegistry: ContractRegistry;
  let signerStorage: SignerStorage;

  const nerifTokenName = 'Nerif Token';
  const nerifTokenSymbol = 'NRF';

  const tokensAmount = wei('1000');
  const defaultPeriodDuration = wei('3600', 0);

  before(async () => {
    [OWNER, FIRST, SIGNER] = await ethers.getSigners();

    const TokensVestingFactory = await ethers.getContractFactory('TokensVesting');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractRegistry');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');

    tokensVesting = await TokensVestingFactory.deploy();
    nerifToken = await NerifTokenFactory.deploy();
    contractsRegistry = await ContractsRegistryFactory.deploy();
    signerStorage = await SignerStorageFactory.deploy();

    await tokensVesting.initialize(signerStorage.address, contractsRegistry.address);
    await nerifToken.initialize(contractsRegistry.address, tokensAmount, nerifTokenName, nerifTokenSymbol);
    await signerStorage.initialize(SIGNER.address);
    await contractsRegistry.initialize(signerStorage.address);

    await contractsRegistry.connect(SIGNER).setContract(await nerifToken.TOKENS_VESTING_KEY(), tokensVesting.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await tokensVesting.contractsRegistry()).to.be.eq(contractsRegistry.address);
      expect(await tokensVesting.signerGetter()).to.be.eq(signerStorage.address);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(tokensVesting.initialize(signerStorage.address, contractsRegistry.address)).to.be.revertedWith(
        reason
      );
    });
  });

  describe('createSchedule', () => {
    let baseSchedule: Vesting.BaseScheduleStruct;

    beforeEach('setup', async () => {
      baseSchedule = {
        secondsInPeriod: defaultPeriodDuration,
        durationInPeriods: 24,
        cliffInPeriods: 1,
      };
    });

    it('should correctly create schedule', async () => {
      const expectedScheduleId = 1;

      const tx = await tokensVesting.connect(SIGNER).createSchedule(baseSchedule);

      await expect(tx).to.emit(tokensVesting, 'ScheduleCreated').withArgs(expectedScheduleId);

      const storedScheduleData = (await tokensVesting.getSchedule(expectedScheduleId)).scheduleData;

      expect(storedScheduleData.secondsInPeriod).to.be.eq(baseSchedule.secondsInPeriod);
      expect(storedScheduleData.durationInPeriods).to.be.eq(baseSchedule.durationInPeriods);
      expect(storedScheduleData.cliffInPeriods).to.be.eq(baseSchedule.cliffInPeriods);
    });

    it('should correctly create several schedules', async () => {
      const expectedScheduleId1 = 1;
      const expectedScheduleId2 = 2;

      await tokensVesting.connect(SIGNER).createSchedule(baseSchedule);

      let storedScheduleData = (await tokensVesting.getSchedule(expectedScheduleId1)).scheduleData;

      expect(storedScheduleData.secondsInPeriod).to.be.eq(baseSchedule.secondsInPeriod);
      expect(storedScheduleData.durationInPeriods).to.be.eq(baseSchedule.durationInPeriods);
      expect(storedScheduleData.cliffInPeriods).to.be.eq(baseSchedule.cliffInPeriods);

      baseSchedule.durationInPeriods = 48;
      baseSchedule.cliffInPeriods = 2;

      await tokensVesting.connect(SIGNER).createSchedule(baseSchedule);

      storedScheduleData = (await tokensVesting.getSchedule(expectedScheduleId2)).scheduleData;

      expect(storedScheduleData.secondsInPeriod).to.be.eq(baseSchedule.secondsInPeriod);
      expect(storedScheduleData.durationInPeriods).to.be.eq(baseSchedule.durationInPeriods);
      expect(storedScheduleData.cliffInPeriods).to.be.eq(baseSchedule.cliffInPeriods);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(tokensVesting.createSchedule(baseSchedule)).to.be.revertedWith(reason);
    });
  });

  describe('createVesting', () => {
    const currentTime = wei('10000', 1);
    const baseScheduleId = 1;
    const baseStartTime = currentTime.add('100');
    const baseTokensAmount = wei('10000');

    let baseSchedule: Vesting.BaseScheduleStruct;
    let vestingData: Vesting.VestingDataStruct;

    beforeEach('setup', async () => {
      baseSchedule = {
        secondsInPeriod: defaultPeriodDuration,
        durationInPeriods: 24,
        cliffInPeriods: 1,
      };

      await tokensVesting.connect(SIGNER).createSchedule(baseSchedule);
      await setTime(currentTime.toNumber());

      vestingData = {
        beneficiary: FIRST.address,
        vestingToken: nerifToken.address,
        scheduleId: baseScheduleId,
        vestingStartTime: baseStartTime,
        paidAmount: 0,
        vestingAmount: baseTokensAmount,
      };
    });

    it('should correctly create vesting', async () => {
      const expectedVestingId = 1;

      const tx = await tokensVesting.connect(SIGNER).createVesting(vestingData);

      await expect(tx)
        .to.emit(tokensVesting, 'VestingCreated')
        .withArgs(expectedVestingId, FIRST.address, nerifToken.address);

      const vestingInfo = await tokensVesting.getVesting(expectedVestingId);

      expect(vestingInfo.beneficiary).to.be.eq(FIRST.address);
      expect(vestingInfo.vestingToken).to.be.eq(nerifToken.address);
      expect(vestingInfo.vestingAmount).to.be.eq(baseTokensAmount);
      expect(vestingInfo.scheduleId).to.be.eq(baseScheduleId);
    });

    it('should correctly create several vestings', async () => {
      const expectedVestingId1 = 1;
      const expectedVestingId2 = 2;

      await tokensVesting.connect(SIGNER).createVesting(vestingData);

      let vestingInfo = await tokensVesting.getVesting(expectedVestingId1);

      expect(vestingInfo.beneficiary).to.be.eq(FIRST.address);
      expect(vestingInfo.vestingToken).to.be.eq(nerifToken.address);
      expect(vestingInfo.vestingAmount).to.be.eq(baseTokensAmount);
      expect(vestingInfo.scheduleId).to.be.eq(baseScheduleId);

      vestingData.beneficiary = OWNER.address;
      vestingData.vestingAmount = baseTokensAmount.mul(2);

      await tokensVesting.connect(SIGNER).createVesting(vestingData);

      vestingInfo = await tokensVesting.getVesting(expectedVestingId2);

      expect(vestingInfo.beneficiary).to.be.eq(OWNER.address);
      expect(vestingInfo.vestingToken).to.be.eq(nerifToken.address);
      expect(vestingInfo.vestingAmount).to.be.eq(baseTokensAmount.mul(2));
      expect(vestingInfo.scheduleId).to.be.eq(baseScheduleId);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(tokensVesting.createVesting(vestingData)).to.be.revertedWith(reason);
    });
  });

  describe('withdrawFromVesting', () => {
    const currentTime = wei('10000', 1);
    const baseScheduleId = 1;
    const baseVestingId = 1;
    const baseStartTime = currentTime.add('100');
    const baseTokensAmount = wei('12000');

    let baseSchedule: Vesting.BaseScheduleStruct;
    let vestingData: Vesting.VestingDataStruct;

    beforeEach('setup', async () => {
      baseSchedule = {
        secondsInPeriod: defaultPeriodDuration,
        durationInPeriods: 24,
        cliffInPeriods: 4,
      };

      await tokensVesting.connect(SIGNER).createSchedule(baseSchedule);
      await setTime(currentTime.toNumber());

      vestingData = {
        beneficiary: FIRST.address,
        vestingToken: nerifToken.address,
        scheduleId: baseScheduleId,
        vestingStartTime: baseStartTime,
        paidAmount: 0,
        vestingAmount: baseTokensAmount,
      };

      await tokensVesting.connect(SIGNER).createVesting(vestingData);
    });

    it('should correctly count withdrawable tokens amount', async () => {
      const expectedAmountPerPeriod = baseTokensAmount.div(await baseSchedule.durationInPeriods);

      expect(await tokensVesting.getWithdrawableAmount(baseVestingId)).to.be.eq(0);

      let periodsToMove = 5;

      await setTime(baseStartTime.add(defaultPeriodDuration.mul(periodsToMove)).toNumber());

      expect(await tokensVesting.getWithdrawableAmount(baseVestingId)).to.be.closeTo(
        expectedAmountPerPeriod.mul(periodsToMove),
        10
      );

      periodsToMove = 10;

      await setTime(baseStartTime.add(defaultPeriodDuration.mul(periodsToMove)).toNumber());

      expect(await tokensVesting.getWithdrawableAmount(baseVestingId)).to.be.closeTo(
        expectedAmountPerPeriod.mul(periodsToMove),
        10
      );

      periodsToMove = 25;

      await setTime(baseStartTime.add(defaultPeriodDuration.mul(periodsToMove)).toNumber());

      expect(await tokensVesting.getWithdrawableAmount(baseVestingId)).to.be.eq(baseTokensAmount);
    });

    it('should withdraw tokens from vesting', async () => {
      const expectedAmountPerPeriod = baseTokensAmount.div(await baseSchedule.durationInPeriods);

      expect(await tokensVesting.getWithdrawableAmount(baseVestingId)).to.be.eq(0);

      let periodsToMove = 5;

      await setTime(baseStartTime.add(defaultPeriodDuration.mul(periodsToMove)).toNumber());

      let withdrawableAmount = await tokensVesting.getWithdrawableAmount(baseVestingId);

      expect(withdrawableAmount).to.be.closeTo(expectedAmountPerPeriod.mul(periodsToMove), 10);

      let tx = await tokensVesting.connect(FIRST).withdrawFromVesting(baseVestingId);

      await expect(tx).to.emit(tokensVesting, 'WithdrawnFromVesting').withArgs(baseVestingId, withdrawableAmount);
      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(withdrawableAmount);

      periodsToMove = 10;

      await setTime(baseStartTime.add(defaultPeriodDuration.mul(periodsToMove)).toNumber());

      let vestedAmount = withdrawableAmount;
      withdrawableAmount = await tokensVesting.getWithdrawableAmount(baseVestingId);

      expect(withdrawableAmount).to.be.closeTo(expectedAmountPerPeriod.mul(periodsToMove).sub(vestedAmount), 10);
      expect(await tokensVesting.getVestedAmount(baseVestingId)).to.be.eq(vestedAmount.add(withdrawableAmount));

      tx = await tokensVesting.connect(FIRST).withdrawFromVesting(baseVestingId);

      await expect(tx).to.emit(tokensVesting, 'WithdrawnFromVesting').withArgs(baseVestingId, withdrawableAmount);
      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(withdrawableAmount.add(vestedAmount));

      periodsToMove = 25;

      await setTime(baseStartTime.add(defaultPeriodDuration.mul(periodsToMove)).toNumber());

      vestedAmount = vestedAmount.add(withdrawableAmount);
      withdrawableAmount = await tokensVesting.getWithdrawableAmount(baseVestingId);

      expect(withdrawableAmount).to.be.eq(baseTokensAmount.sub(vestedAmount));
      expect(await tokensVesting.getVestedAmount(baseVestingId)).to.be.eq(baseTokensAmount);

      tx = await tokensVesting.connect(FIRST).withdrawFromVesting(baseVestingId);

      await expect(tx).to.emit(tokensVesting, 'WithdrawnFromVesting').withArgs(baseVestingId, withdrawableAmount);
      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(withdrawableAmount.add(vestedAmount));
    });
  });
});
