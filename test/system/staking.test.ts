import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import {
  ContractRegistry,
  DKG,
  RewardDistributionPool,
  SignerStorage,
  Staking,
  TestERC20,
  TestSlashingVoting,
} from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setNextBlockTime, setTime } from '../helpers/block-helper';

const { signPermit } = require('../helpers/signatures.js');

describe('Staking', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let staking: Staking;
  let dkg: DKG;
  let rewardsDistributionPool: RewardDistributionPool;
  let contractsRegistry: ContractRegistry;
  let slashingVoting: TestSlashingVoting;
  let signerStorage: SignerStorage;
  let stakeToken: TestERC20;

  const OWNER_PK: string = 'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';

  const tokensAmount = wei('1000');
  const defMinimalStake = wei('100');
  const defWithdrawalPeriod = wei('3600', 1);
  const deadlinePeriod = wei('3600', 1);

  function createPermitSig(
    tokenName: string,
    tokenAddr: string,
    deadline: string | number = 0,
    value: string | number = tokensAmount.toString(),
    ownerAddr: string = OWNER.address,
    nonce: string | number = 0
  ) {
    const domain = {
      name: tokenName,
      version: '1',
      verifyingContract: tokenAddr,
      chainId: '1',
    };

    const message = {
      owner: ownerAddr,
      spender: staking.address,
      value: value,
      nonce: nonce,
      deadline: deadline,
    };

    const buffer = Buffer.from(OWNER_PK, 'hex');

    return signPermit(domain, message, buffer);
  }

  before(async () => {
    [OWNER, FIRST, SECOND, SIGNER] = await ethers.getSigners();

    const StakingFactory = await ethers.getContractFactory('Staking');
    const DKGFactory = await ethers.getContractFactory('DKG');
    const RewardDistributionPoolFactory = await ethers.getContractFactory('RewardDistributionPool');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractRegistry');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const TestERC20Factory = await ethers.getContractFactory('TestERC20');
    const TestSlashingVotingFactory = await ethers.getContractFactory('TestSlashingVoting');

    staking = await StakingFactory.deploy();
    dkg = await DKGFactory.deploy();
    rewardsDistributionPool = await RewardDistributionPoolFactory.deploy();
    contractsRegistry = await ContractsRegistryFactory.deploy();
    slashingVoting = await TestSlashingVotingFactory.deploy();
    signerStorage = await SignerStorageFactory.deploy();
    stakeToken = await TestERC20Factory.deploy('Stake Token', 'ST', tokensAmount);

    await staking.initialize(
      signerStorage.address,
      contractsRegistry.address,
      stakeToken.address,
      defMinimalStake,
      defWithdrawalPeriod
    );
    await signerStorage.initialize(SIGNER.address);
    await dkg.initialize(contractsRegistry.address, deadlinePeriod);
    await contractsRegistry.initialize(signerStorage.address);

    await contractsRegistry.connect(SIGNER).setContract(await dkg.SLASHING_VOTING_KEY(), slashingVoting.address);
    await contractsRegistry.connect(SIGNER).setContract(await dkg.DKG_KEY(), dkg.address);
    await contractsRegistry.connect(SIGNER).setContract(await staking.STAKING_KEY(), staking.address);
    await contractsRegistry
      .connect(SIGNER)
      .setContract(await rewardsDistributionPool.REWARD_DISTRIBUTION_POOL_KEY(), rewardsDistributionPool.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await staking.contractRegistry()).to.be.eq(contractsRegistry.address);
      expect(await staking.stakeToken()).to.be.eq(stakeToken.address);
      expect(await staking.signerGetter()).to.be.eq(signerStorage.address);

      expect(await staking.minimalStake()).to.be.eq(defMinimalStake);
      expect(await staking.withdrawalPeriod()).to.be.eq(defWithdrawalPeriod);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(
        staking.initialize(
          signerStorage.address,
          contractsRegistry.address,
          stakeToken.address,
          defMinimalStake,
          defWithdrawalPeriod
        )
      ).to.be.revertedWith(reason);
    });
  });

  describe('setMinimalStake', () => {
    const newMinimalStake = wei('50');

    it('should correctly set new minimal stake amount', async () => {
      const tx = await staking.connect(SIGNER).setMinimalStake(newMinimalStake);

      await expect(tx).to.emit(staking, 'MinimalStakeUpdated').withArgs(newMinimalStake);

      expect(await staking.minimalStake()).to.be.eq(newMinimalStake);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(staking.setMinimalStake(newMinimalStake)).to.be.revertedWith(reason);
    });
  });

  describe('setWithdrawalPeriod', () => {
    const newWithdrawalPeriod = wei('36000', 1);

    it('should correctly set new withdrawal period', async () => {
      const tx = await staking.connect(SIGNER).setWithdrawalPeriod(newWithdrawalPeriod);

      await expect(tx).to.emit(staking, 'WithdrawalPeriodUpdated').withArgs(newWithdrawalPeriod);

      expect(await staking.withdrawalPeriod()).to.be.eq(newWithdrawalPeriod);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(staking.setWithdrawalPeriod(newWithdrawalPeriod)).to.be.revertedWith(reason);
    });
  });

  describe('slash', () => {
    const stakeAmount = wei('200');

    beforeEach('setup', async () => {
      await stakeToken.approve(staking.address, stakeAmount);
      await staking.stake(stakeAmount);

      expect(await staking.getStake(OWNER.address)).to.be.eq(stakeAmount);
      expect(await staking.getValidators()).to.be.deep.eq([OWNER.address]);
    });

    it('should correctly slash validator', async () => {
      expect(await staking.getValidatorStatus(OWNER.address)).to.be.eq(1);

      await slashingVoting.slash(staking.address, OWNER.address);

      expect(await staking.getValidatorStatus(OWNER.address)).to.be.eq(2);
      expect(await staking.isValidatorSlashed(OWNER.address)).to.be.eq(true);
      expect(await staking.getValidators()).to.be.deep.eq([]);
    });

    it('should get exception if not a slashing voting try to call this function', async () => {
      const reason = 'Staking: not a slashing voting';

      await expect(staking.slash(OWNER.address)).to.be.revertedWith(reason);
    });

    it('should get exception if the validator is not in the validators list', async () => {
      const reason = 'Staking: Failed to update validator status';

      await expect(slashingVoting.slash(staking.address, FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('announceWithdrawal', () => {
    const stakeAmount = wei('150');
    const announceAmount = wei('30');
    const announceTime = wei('100000', 1);

    beforeEach('setup', async () => {
      await stakeToken.approve(staking.address, stakeAmount);
      await staking.stake(stakeAmount);
    });

    it('should correctly announce withdrawal without removing from validators list', async () => {
      await setNextBlockTime(announceTime.toNumber());
      await staking.announceWithdrawal(announceAmount);

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(OWNER.address);

      expect(withdrawalInfo.time).to.be.eq(announceTime);
      expect(withdrawalInfo.amount).to.be.eq(announceAmount);

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(true);
    });

    it('should correctly announce withdrawal with removing from the validators list', async () => {
      await setNextBlockTime(announceTime.toNumber());
      await staking.announceWithdrawal(announceAmount.mul(2));

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(OWNER.address);

      expect(withdrawalInfo.time).to.be.eq(announceTime);
      expect(withdrawalInfo.amount).to.be.eq(announceAmount.mul(2));

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(false);
      expect(await staking.getValidatorStatus(OWNER.address)).to.be.eq(0);
    });

    it('should correctly announce withdrawal several times', async () => {
      await setNextBlockTime(announceTime.toNumber());
      await staking.announceWithdrawal(announceAmount.mul(2));

      let withdrawalInfo = await staking.getWithdrawalAnnouncement(OWNER.address);

      expect(withdrawalInfo.time).to.be.eq(announceTime);
      expect(withdrawalInfo.amount).to.be.eq(announceAmount.mul(2));

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(false);
      expect(await staking.getValidatorStatus(OWNER.address)).to.be.eq(0);

      await setNextBlockTime(announceTime.add(100).toNumber());
      await staking.announceWithdrawal(announceAmount.mul(3));

      withdrawalInfo = await staking.getWithdrawalAnnouncement(OWNER.address);

      expect(withdrawalInfo.time).to.be.eq(announceTime.add(100));
      expect(withdrawalInfo.amount).to.be.eq(announceAmount.mul(3));

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(false);
      expect(await staking.getValidatorStatus(OWNER.address)).to.be.eq(0);
    });

    it('should get exception if pass zero amount', async () => {
      const reason = 'Staking: amount must be greater than zero';

      await expect(staking.announceWithdrawal(0)).to.be.revertedWith(reason);
    });

    it('should get exception if pass amount that higher than the stake', async () => {
      const reason = 'Staking: amount must be <= to stake';

      await expect(staking.announceWithdrawal(stakeAmount.mul(2))).to.be.revertedWith(reason);
    });

    it('should get exception if try to call from the slashed validator addr', async () => {
      const reason = 'Staking: validator is slashed';

      await slashingVoting.slash(staking.address, OWNER.address);

      await expect(staking.announceWithdrawal(stakeAmount.mul(2))).to.be.revertedWith(reason);
    });
  });

  describe('revokeWithdrawal', () => {
    const stakeAmount = wei('150');
    const announceAmount = wei('30');
    const announceTime = wei('100000', 1);

    beforeEach('setup', async () => {
      await stakeToken.approve(staking.address, stakeAmount);
      await staking.stake(stakeAmount);

      await setNextBlockTime(announceTime.toNumber());
      await staking.announceWithdrawal(announceAmount);
    });

    it('should correctly revoke withdrawal', async () => {
      await staking.announceWithdrawal(announceAmount.mul(2));

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(false);

      await staking.revokeWithdrawal();

      let withdrawalInfo = await staking.getWithdrawalAnnouncement(OWNER.address);

      expect(withdrawalInfo.time).to.be.eq(0);
      expect(withdrawalInfo.amount).to.be.eq(0);

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(true);
    });

    it('should correctly revoke withdrawal without adding to the validators list', async () => {
      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(true);

      await staking.revokeWithdrawal();

      let withdrawalInfo = await staking.getWithdrawalAnnouncement(OWNER.address);

      expect(withdrawalInfo.time).to.be.eq(0);
      expect(withdrawalInfo.amount).to.be.eq(0);

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(true);
    });

    it('should get exception if try to revoke withdrawal without announcement', async () => {
      const reason = 'Staking: user does not have withdrawal announcement';

      await expect(staking.connect(FIRST).revokeWithdrawal()).to.be.revertedWith(reason);
    });

    it('should get exception if the validator has slashed status', async () => {
      const reason = 'Staking: validator is slashed';

      await slashingVoting.slash(staking.address, OWNER.address);

      await expect(staking.revokeWithdrawal()).to.be.revertedWith(reason);
    });
  });

  describe('withdraw', () => {
    const stakeAmount = wei('150');
    const announceAmount = wei('30');
    const announceTime = wei('100000', 1);

    beforeEach('setup', async () => {
      await stakeToken.approve(staking.address, stakeAmount);
      await staking.stake(stakeAmount);

      await setNextBlockTime(announceTime.toNumber());
      await staking.announceWithdrawal(announceAmount);
    });

    it('should correctly withdraw stake tokens', async () => {
      await setNextBlockTime(announceTime.add(defWithdrawalPeriod).add(10).toNumber());
      const tx = await staking.withdraw();

      await expect(tx).to.emit(staking, 'TokensWithdrawn').withArgs(OWNER.address, announceAmount);

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(OWNER.address);

      expect(withdrawalInfo.time).to.be.eq(0);
      expect(withdrawalInfo.amount).to.be.eq(0);

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(true);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.sub(announceAmount));
      expect(await staking.getStake(OWNER.address)).to.be.eq(stakeAmount.sub(announceAmount));

      expect(await stakeToken.balanceOf(OWNER.address)).to.be.eq(tokensAmount.sub(stakeAmount).add(announceAmount));
      expect(await stakeToken.balanceOf(staking.address)).to.be.eq(stakeAmount.sub(announceAmount));
    });

    it('should get exception if try to withdraw without announcement', async () => {
      const reason = 'Staking: user does not have withdrawal announcement';

      await expect(staking.connect(FIRST).withdraw()).to.be.revertedWith(reason);
    });

    it('should get exception if the withdrawal period not passed', async () => {
      const reason = 'Staking: withdrawal period not passed';

      await expect(staking.withdraw()).to.be.revertedWith(reason);
    });

    it('should get exception if the validator has slashed status', async () => {
      const reason = 'Staking: validator is slashed';

      await slashingVoting.slash(staking.address, OWNER.address);

      await expect(staking.withdraw()).to.be.revertedWith(reason);
    });
  });

  describe('stake', () => {
    const stakeAmount = wei('150');

    beforeEach('setup', async () => {
      await stakeToken.approve(staking.address, tokensAmount);
    });

    it('should correctly stake tokens and add address to the validators list', async () => {
      const tx = await staking.stake(stakeAmount);

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(OWNER.address, OWNER.address, stakeAmount);

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(true);

      expect(await staking.totalStake()).to.be.eq(stakeAmount);
      expect(await staking.getStake(OWNER.address)).to.be.eq(stakeAmount);

      expect(await stakeToken.balanceOf(OWNER.address)).to.be.eq(tokensAmount.sub(stakeAmount));
      expect(await stakeToken.balanceOf(staking.address)).to.be.eq(stakeAmount);

      await staking.stake(stakeAmount);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.mul(2));
      expect(await staking.getStake(OWNER.address)).to.be.eq(stakeAmount.mul(2));
    });

    it('should correctly stake tokens and does not add validator to the validators list', async () => {
      const tx = await staking.stake(stakeAmount.div(2));

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(OWNER.address, OWNER.address, stakeAmount.div(2));

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(false);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.div(2));
      expect(await staking.getStake(OWNER.address)).to.be.eq(stakeAmount.div(2));
    });

    it('should correctly stake tokens without updating validators list', async () => {
      const tx = await staking.stake(stakeAmount.div(2));

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(OWNER.address, OWNER.address, stakeAmount.div(2));

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(false);
    });

    it('should get exception if try to stake zero amount', async () => {
      const reason = 'Staking: Zero stake amount';

      await expect(staking.stake(0)).to.be.revertedWith(reason);
    });

    it('should get exception if try to stake amount that higher than balance', async () => {
      const reason = 'Staking: Not enough tokens to stake';

      await expect(staking.stake(tokensAmount.add(1))).to.be.revertedWith(reason);
    });

    it('should get exception if the slashed validator try to stake tokens', async () => {
      const reason = 'Staking: validator is slashed';

      await staking.stake(stakeAmount);

      await slashingVoting.slash(staking.address, OWNER.address);

      await expect(staking.stake(stakeAmount)).to.be.revertedWith(reason);
    });
  });

  describe('stakeWithPermit', () => {
    const stakeAmount = wei('150');
    const currentTime = wei('100000', 1);
    const deadline = currentTime.add(1000);

    beforeEach('setup', async () => {
      await setTime(currentTime.toNumber());
    });

    it('should correctly stake tokens with permit', async () => {
      const sig = createPermitSig(
        await stakeToken.name(),
        stakeToken.address,
        deadline.toString(),
        stakeAmount.toString()
      );

      const tx = await staking.stakeWithPermit(stakeAmount, deadline, sig.v, sig.r, sig.s);

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(OWNER.address, OWNER.address, stakeAmount);

      expect(await staking.isValidatorActive(OWNER.address)).to.be.eq(true);

      expect(await staking.totalStake()).to.be.eq(stakeAmount);
      expect(await staking.getStake(OWNER.address)).to.be.eq(stakeAmount);
    });

    it('should get exception if the slashed validator try to stake tokens', async () => {
      const reason = 'Staking: validator is slashed';

      await stakeToken.approve(staking.address, stakeAmount);
      await staking.stake(stakeAmount);

      await slashingVoting.slash(staking.address, OWNER.address);

      const sig = createPermitSig(
        await stakeToken.name(),
        stakeToken.address,
        deadline.toString(),
        stakeAmount.toString()
      );

      await expect(staking.stakeWithPermit(stakeAmount, deadline, sig.v, sig.r, sig.s)).to.be.revertedWith(reason);
    });
  });

  describe('getters', () => {
    const stakeAmount = wei('150');

    it('should correctly return validators info', async () => {
      const expectedValidatorsArr = [OWNER, FIRST, SECOND];
      const expectedValidatorsStatuses = [1, 1, 1];

      for (let i = 0; i < expectedValidatorsArr.length; i++) {
        await stakeToken.mint(expectedValidatorsArr[i].address, stakeAmount.mul(i + 1));
        await stakeToken.connect(expectedValidatorsArr[i]).approve(staking.address, stakeAmount.mul(i + 1));
        await staking.connect(expectedValidatorsArr[i]).stake(stakeAmount.mul(i + 1));
      }

      expect(await staking.getValidatorsCount()).to.be.eq(expectedValidatorsArr.length);

      const validatorsInfo = await staking.getValidatorsInfo(0, 10);

      for (let i = 0; i < validatorsInfo.length; i++) {
        expect(validatorsInfo[i].validatorAddr).to.be.eq(expectedValidatorsArr[i].address);
        expect(validatorsInfo[i].validatorData.stake).to.be.eq(stakeAmount.mul(i + 1));
        expect(validatorsInfo[i].validatorData.status).to.be.eq(expectedValidatorsStatuses[i]);
      }
    });
  });
});
