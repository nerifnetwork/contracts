import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry, NerifToken, SignerStorage, TokensVesting, Vesting } from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setTime } from '../helpers/block-helper';

describe('TokensVesting', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;
  let tokensVesting: TokensVesting;
  let nerifToken: NerifToken;
  let signerStorage: SignerStorage;

  const nerifTokenName = 'Nerif Token';
  const nerifTokenSymbol = 'NRF';

  const tokensAmount = wei('1000');
  const defaultPeriodDuration = wei('3600', 0);

  before(async () => {
    [OWNER, FIRST, SIGNER] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const TokensVestingFactory = await ethers.getContractFactory('TokensVesting');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const nerifTokenImpl = await NerifTokenFactory.deploy();

    tokensVesting = await TokensVestingFactory.deploy();
    signerStorage = await SignerStorageFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);

    await contractsRegistry.addContract(await contractsRegistry.TOKENS_VESTING_NAME(), tokensVesting.address);

    nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());

    await nerifToken.initialize(tokensAmount, nerifTokenName, nerifTokenSymbol);
    await tokensVesting.initialize();
    await signerStorage.initialize(SIGNER.address);

    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await tokensVesting.owner()).to.be.eq(OWNER.address);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(tokensVesting.connect(FIRST).initialize()).to.be.revertedWith(reason);
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

      const tx = await tokensVesting.createSchedule(baseSchedule);

      await expect(tx).to.emit(tokensVesting, 'ScheduleCreated').withArgs(expectedScheduleId);

      const storedScheduleData = (await tokensVesting.getSchedule(expectedScheduleId)).scheduleData;

      expect(storedScheduleData.secondsInPeriod).to.be.eq(baseSchedule.secondsInPeriod);
      expect(storedScheduleData.durationInPeriods).to.be.eq(baseSchedule.durationInPeriods);
      expect(storedScheduleData.cliffInPeriods).to.be.eq(baseSchedule.cliffInPeriods);
    });

    it('should correctly create several schedules', async () => {
      const expectedScheduleId1 = 1;
      const expectedScheduleId2 = 2;

      await tokensVesting.createSchedule(baseSchedule);

      let storedScheduleData = (await tokensVesting.getSchedule(expectedScheduleId1)).scheduleData;

      expect(storedScheduleData.secondsInPeriod).to.be.eq(baseSchedule.secondsInPeriod);
      expect(storedScheduleData.durationInPeriods).to.be.eq(baseSchedule.durationInPeriods);
      expect(storedScheduleData.cliffInPeriods).to.be.eq(baseSchedule.cliffInPeriods);

      baseSchedule.durationInPeriods = 48;
      baseSchedule.cliffInPeriods = 2;

      await tokensVesting.createSchedule(baseSchedule);

      storedScheduleData = (await tokensVesting.getSchedule(expectedScheduleId2)).scheduleData;

      expect(storedScheduleData.secondsInPeriod).to.be.eq(baseSchedule.secondsInPeriod);
      expect(storedScheduleData.durationInPeriods).to.be.eq(baseSchedule.durationInPeriods);
      expect(storedScheduleData.cliffInPeriods).to.be.eq(baseSchedule.cliffInPeriods);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(tokensVesting.connect(FIRST).createSchedule(baseSchedule)).to.be.revertedWith(reason);
      await expect(tokensVesting.connect(SIGNER).createSchedule(baseSchedule)).to.be.revertedWith(reason);
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

      await tokensVesting.createSchedule(baseSchedule);
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

      const tx = await tokensVesting.createVesting(vestingData);

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

      await tokensVesting.createVesting(vestingData);

      let vestingInfo = await tokensVesting.getVesting(expectedVestingId1);

      expect(vestingInfo.beneficiary).to.be.eq(FIRST.address);
      expect(vestingInfo.vestingToken).to.be.eq(nerifToken.address);
      expect(vestingInfo.vestingAmount).to.be.eq(baseTokensAmount);
      expect(vestingInfo.scheduleId).to.be.eq(baseScheduleId);

      vestingData.beneficiary = OWNER.address;
      vestingData.vestingAmount = baseTokensAmount.mul(2);

      await tokensVesting.createVesting(vestingData);

      vestingInfo = await tokensVesting.getVesting(expectedVestingId2);

      expect(vestingInfo.beneficiary).to.be.eq(OWNER.address);
      expect(vestingInfo.vestingToken).to.be.eq(nerifToken.address);
      expect(vestingInfo.vestingAmount).to.be.eq(baseTokensAmount.mul(2));
      expect(vestingInfo.scheduleId).to.be.eq(baseScheduleId);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(tokensVesting.connect(FIRST).createVesting(vestingData)).to.be.revertedWith(reason);
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

      await tokensVesting.createSchedule(baseSchedule);
      await setTime(currentTime.toNumber());

      vestingData = {
        beneficiary: FIRST.address,
        vestingToken: nerifToken.address,
        scheduleId: baseScheduleId,
        vestingStartTime: baseStartTime,
        paidAmount: 0,
        vestingAmount: baseTokensAmount,
      };

      await tokensVesting.createVesting(vestingData);
    });

    it('should correctly count withdrawable tokens amount', async () => {
      const expectedAmountPerPeriod = baseTokensAmount.div(await baseSchedule.durationInPeriods);

      expect(await tokensVesting.getWithdrawableAmount(baseVestingId)).to.be.eq(0);

      let periodsToMove = 4;

      await setTime(baseStartTime.add(defaultPeriodDuration.mul(periodsToMove).add(1)).toNumber());

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
