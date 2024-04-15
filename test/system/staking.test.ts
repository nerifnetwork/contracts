import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry, TestDKG, NerifToken, TestStaking } from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setNextBlockTime, setTime } from '../helpers/block-helper';
import { BigNumber } from 'ethers';

const { signPermit } = require('../helpers/signatures.js');

describe('Staking', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;
  let staking: TestStaking;
  let dkg: TestDKG;
  let nerifToken: NerifToken;

  const FIRST_PK: string = '59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d';

  const tokensAmount = wei('1000');
  const defMinimalStake = wei('100');

  const defUpdateCollectionsEpochDuration = wei('7200', 0);
  const defDKGGenerationEpochDuration = wei('600', 0);
  const defGuaranteedWorkingEpochDuration = wei('14400', 0);

  const startTime = wei('20000', 0);
  const startEpochId = wei('1', 0);

  function getEndDKGPeriodTime(startEpochTime: BigNumber) {
    return startEpochTime.add(defUpdateCollectionsEpochDuration).add(defDKGGenerationEpochDuration);
  }

  function createPermitSig(
    tokenName: string,
    tokenAddr: string,
    deadline: string | number = 0,
    value: string | number = tokensAmount.toString(),
    ownerAddr: string = FIRST.address,
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

    const buffer = Buffer.from(FIRST_PK, 'hex');

    return signPermit(domain, message, buffer);
  }

  before(async () => {
    [OWNER, FIRST, SECOND] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const StakingFactory = await ethers.getContractFactory('TestStaking');
    const DKGFactory = await ethers.getContractFactory('TestDKG');
    const RewardDistributionPoolFactory = await ethers.getContractFactory('RewardDistributionPool');
    const TokensVestingFactory = await ethers.getContractFactory('TokensVesting');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const stakingImpl = await StakingFactory.deploy();
    const dkgImpl = await DKGFactory.deploy();
    const rewardsDistributionPoolImpl = await RewardDistributionPoolFactory.deploy();
    const nerifTokenImpl = await NerifTokenFactory.deploy();

    const tokensVesting = await TokensVestingFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.DKG_NAME(), dkgImpl.address);
    await contractsRegistry.addProxyContract(await contractsRegistry.STAKING_NAME(), stakingImpl.address);
    await contractsRegistry.addProxyContract(
      await contractsRegistry.REWARDS_DISTRIBUTION_POOL_NAME(),
      rewardsDistributionPoolImpl.address
    );
    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);

    staking = StakingFactory.attach(await contractsRegistry.getStakingContract());
    dkg = DKGFactory.attach(await contractsRegistry.getDKGContract());
    nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());

    await contractsRegistry.addContract(await contractsRegistry.TOKENS_VESTING_NAME(), tokensVesting.address);
    await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), dkg.address);

    await nerifToken.initialize(tokensAmount, 'NERIF', 'NERIF');
    await staking.initialize(nerifToken.address, defMinimalStake, [OWNER.address]);

    await setNextBlockTime(startTime.toNumber());

    await dkg.initialize(
      defUpdateCollectionsEpochDuration,
      defDKGGenerationEpochDuration,
      defGuaranteedWorkingEpochDuration
    );

    await contractsRegistry.injectDependencies(await contractsRegistry.DKG_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.STAKING_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.REWARDS_DISTRIBUTION_POOL_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await staking.stakeToken()).to.be.eq(nerifToken.address);

      expect(await staking.minimalStake()).to.be.eq(defMinimalStake);
      expect(await staking.getWhitelistedUsers()).to.be.deep.eq([OWNER.address]);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(staking.initialize(nerifToken.address, defMinimalStake, [])).to.be.revertedWith(reason);
    });
  });

  describe('setDependencies', () => {
    it('should correctly update dependencies', async () => {
      expect(await staking.dkgAddress()).to.be.eq(dkg.address);

      await contractsRegistry.addContract(await contractsRegistry.DKG_NAME(), FIRST.address);
      await contractsRegistry.injectDependencies(await contractsRegistry.STAKING_NAME());

      expect(await staking.dkgAddress()).to.be.eq(FIRST.address);
    });

    it('should get exception if not a contracts registry try to call this function', async () => {
      const reason = 'Dependant: not an injector';

      await expect(staking.setDependencies(contractsRegistry.address, '0x')).to.be.revertedWith(reason);
    });
  });

  describe('setMinimalStake', () => {
    const newMinimalStake = wei('50');

    it('should correctly set new minimal stake amount', async () => {
      const tx = await staking.setMinimalStake(newMinimalStake);

      await expect(tx).to.emit(staking, 'MinimalStakeUpdated').withArgs(newMinimalStake);

      expect(await staking.minimalStake()).to.be.eq(newMinimalStake);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'Staking: Not a signer';

      await expect(staking.connect(FIRST).setMinimalStake(newMinimalStake)).to.be.revertedWith(reason);
    });
  });

  describe('updateWhitelistedUsers', () => {
    it('should correctly add users to whitelist', async () => {
      const usersToAdd = [FIRST.address, SECOND.address];

      expect(await staking.getWhitelistedUsers()).to.be.deep.eq([OWNER.address]);

      await staking.updateWhitelistedUsers(usersToAdd, true);

      expect(await staking.getWhitelistedUsers()).to.be.deep.eq([OWNER.address, ...usersToAdd]);
    });

    it('should correctly remove users from whitelist', async () => {
      const expectedWhitelistedUsersArr = [OWNER.address, SECOND.address];

      await staking.updateWhitelistedUsers([FIRST.address, SECOND.address], true);

      await staking.updateWhitelistedUsers([FIRST.address], false);

      expect(await staking.getWhitelistedUsers()).to.be.deep.eq(expectedWhitelistedUsersArr);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'Staking: Not a signer';

      await expect(staking.connect(FIRST).updateWhitelistedUsers([FIRST.address], true)).to.be.revertedWith(reason);
    });
  });

  describe('stake', () => {
    const stakeAmount = wei('150');

    beforeEach('setup', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, tokensAmount);
    });

    it('should correctly stake tokens and add address to the validators list', async () => {
      const tx = await staking.connect(FIRST).stake(stakeAmount);

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isLastEpoch(startEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(startEpochId)).to.be.eq(true);

      expect(await staking.totalStake()).to.be.eq(stakeAmount);
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount);

      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(tokensAmount.sub(stakeAmount));
      expect(await nerifToken.balanceOf(staking.address)).to.be.eq(stakeAmount);

      await staking.connect(FIRST).stake(stakeAmount);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.mul(2));
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount.mul(2));

      expect(await dkg.isLastEpoch(startEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(startEpochId)).to.be.eq(true);
    });

    it('should correctly stake tokens and not add validator to the validators list', async () => {
      const tx = await staking.connect(FIRST).stake(stakeAmount.div(2));

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount.div(2));

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.div(2));
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount.div(2));
    });

    it('should correctly stake tokens several times with adding to the validators', async () => {
      let tx = await staking.connect(FIRST).stake(stakeAmount.div(2));

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount.div(2));

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      tx = await staking.connect(FIRST).stake(stakeAmount.div(2));

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount.div(2));

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);
    });

    it('should get exception if try to stake zero amount', async () => {
      const reason = 'Staking: Zero stake amount';

      await expect(staking.connect(FIRST).stake(0)).to.be.revertedWith(reason);
    });

    it('should get exception if try to stake amount that higher than balance', async () => {
      const reason = 'Staking: Not enough tokens to stake';

      await expect(staking.connect(FIRST).stake(tokensAmount.add(1))).to.be.revertedWith(reason);
    });

    it('should get exception if not whitelisted validator try to stake', async () => {
      const reason = 'Staking: Not a whitelisted user';

      await expect(staking.connect(SECOND).stake(stakeAmount)).to.be.revertedWith(reason);
    });
  });

  describe('stakeWithPermit', () => {
    const stakeAmount = wei('150');
    const deadline = startTime.add('1000');

    beforeEach('setup', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address], true);
    });

    it('should correctly stake tokens with permit', async () => {
      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        stakeAmount.toString()
      );

      const tx = await staking.connect(FIRST).stakeWithPermit(stakeAmount, deadline, sig.v, sig.r, sig.s);

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);

      expect(await staking.totalStake()).to.be.eq(stakeAmount);
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount);
    });

    it('should get exception if not whitelisted validator try to stake', async () => {
      const reason = 'Staking: Not a whitelisted user';

      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        stakeAmount.toString()
      );

      await expect(
        staking.connect(SECOND).stakeWithPermit(stakeAmount, deadline, sig.v, sig.r, sig.s)
      ).to.be.revertedWith(reason);
    });
  });

  describe('announceWithdrawal', () => {
    const stakeAmount = wei('150');
    const announceAmount = wei('30');
    let activePeriodStartTime: BigNumber;
    let announceTime: BigNumber;

    beforeEach('setup', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, stakeAmount);
      await staking.connect(FIRST).stake(stakeAmount);

      activePeriodStartTime = (await dkg.getEpochEndTime(startEpochId)).add('100');
      announceTime = activePeriodStartTime.add('100');

      await setTime(activePeriodStartTime.toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);
    });

    it('should correctly announce withdrawal without removing from validators list', async () => {
      await setNextBlockTime(announceTime.toNumber());
      const tx = await staking.connect(FIRST).announceWithdrawal(announceAmount);

      await expect(tx).emit(staking, 'WithdrawalAnnounced').withArgs(FIRST.address, announceAmount, announceTime);

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(FIRST.address);

      expect(withdrawalInfo.withdrawalTime).to.be.eq(announceTime);
      expect(withdrawalInfo.tokensAmount).to.be.eq(announceAmount);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);
    });

    it('should correctly announce withdrawal with new epoch creation', async () => {
      const newEpochId = startEpochId.add('1');
      const expectedWithdrawalTime = getEndDKGPeriodTime(announceTime);

      await setNextBlockTime(announceTime.toNumber());
      const tx = await staking.connect(FIRST).announceWithdrawal(announceAmount.mul(2));

      await expect(tx)
        .emit(staking, 'WithdrawalAnnounced')
        .withArgs(FIRST.address, announceAmount.mul(2), expectedWithdrawalTime);

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(FIRST.address);

      expect(withdrawalInfo.withdrawalTime).to.be.eq(expectedWithdrawalTime);
      expect(withdrawalInfo.tokensAmount).to.be.eq(announceAmount.mul(2));

      expect(await dkg.isLastEpoch(newEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(newEpochId)).to.be.eq(true);

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(2);
    });

    it('should get exception if pass zero amount', async () => {
      const reason = 'Staking: Amount must be greater than zero';

      await expect(staking.connect(FIRST).announceWithdrawal(0)).to.be.revertedWith(reason);
    });

    it('should get exception if pass amount that higher than the stake', async () => {
      const reason = 'Staking: Not enough stake';

      await expect(staking.connect(FIRST).announceWithdrawal(stakeAmount.mul(2))).to.be.revertedWith(reason);
    });

    it('should get exception if the validator already has withdrawal announcement', async () => {
      const reason = 'Staking: User already has withdrawal announcement';

      await staking.connect(FIRST).announceWithdrawal(announceAmount);

      await expect(staking.connect(FIRST).announceWithdrawal(announceAmount)).to.be.revertedWith(reason);
    });
  });

  describe('withdraw', () => {
    const stakeAmount = wei('150');
    const announceAmount = wei('30');

    let activePeriodStartTime: BigNumber;
    let announceTime: BigNumber;

    beforeEach('setup', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, stakeAmount);
      await staking.connect(FIRST).stake(stakeAmount);

      activePeriodStartTime = (await dkg.getEpochEndTime(startEpochId)).add('100');
      announceTime = activePeriodStartTime.add('100');

      await setTime(activePeriodStartTime.toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);

      await setNextBlockTime(announceTime.toNumber());
    });

    it('should correctly withdraw stake tokens with removing from validators list', async () => {
      const withdrawalTime = getEndDKGPeriodTime(announceTime);

      await staking.connect(FIRST).announceWithdrawal(announceAmount.mul(2));

      expect((await staking.getWithdrawalAnnouncement(FIRST.address)).withdrawalTime).to.be.eq(withdrawalTime);

      await setNextBlockTime(withdrawalTime.add('100').toNumber());
      const tx = await staking.connect(FIRST).withdraw();

      await expect(tx).to.emit(staking, 'TokensWithdrawn').withArgs(FIRST.address, announceAmount.mul(2));

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(FIRST.address);

      expect(withdrawalInfo.withdrawalTime).to.be.eq(0);
      expect(withdrawalInfo.tokensAmount).to.be.eq(0);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.sub(announceAmount.mul(2)));
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount.sub(announceAmount.mul(2)));

      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(
        tokensAmount.sub(stakeAmount).add(announceAmount.mul(2))
      );
      expect(await nerifToken.balanceOf(staking.address)).to.be.eq(stakeAmount.sub(announceAmount.mul(2)));
    });

    it('should correctly withdraw tokens without removing from the validators list', async () => {
      await staking.connect(FIRST).announceWithdrawal(announceAmount);

      expect((await staking.getWithdrawalAnnouncement(FIRST.address)).withdrawalTime).to.be.eq(announceTime);

      await setNextBlockTime(announceTime.add('100').toNumber());
      const tx = await staking.connect(FIRST).withdraw();

      await expect(tx).to.emit(staking, 'TokensWithdrawn').withArgs(FIRST.address, announceAmount);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);
    });

    it('should correctly withdraw tokens with stake < min and validator not in the validators list', async () => {
      const withdrawalTime = getEndDKGPeriodTime(announceTime);

      await staking.connect(FIRST).announceWithdrawal(announceAmount.mul(2));

      await setNextBlockTime(withdrawalTime.add('100').toNumber());
      await dkg.setSigner(startEpochId.add('1'), OWNER.address);

      await dkg.updateAllValidators();

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      await setNextBlockTime(withdrawalTime.add('110').toNumber());
      const tx = await staking.connect(FIRST).withdraw();

      await expect(tx).to.emit(staking, 'TokensWithdrawn').withArgs(FIRST.address, announceAmount.mul(2));

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(FIRST.address);

      expect(withdrawalInfo.withdrawalTime).to.be.eq(0);
      expect(withdrawalInfo.tokensAmount).to.be.eq(0);
    });

    it('should get exception if try to withdraw without announcement', async () => {
      const reason = 'Staking: User does not have withdrawal announcement';

      await expect(staking.connect(SECOND).withdraw()).to.be.revertedWith(reason);
    });

    it('should get exception if the withdrawal time has not come', async () => {
      await staking.connect(FIRST).announceWithdrawal(announceAmount.mul(2));

      const reason = 'Staking: The time for withdrawal has not come';

      await expect(staking.connect(FIRST).withdraw()).to.be.revertedWith(reason);
    });
  });

  describe('getters', () => {
    const stakeAmount = wei('150');

    it('should correctly return validators info', async () => {
      await staking.updateWhitelistedUsers([FIRST.address, SECOND.address], true);

      const expectedValidatorsArr = [OWNER, FIRST, SECOND];

      for (let i = 0; i < expectedValidatorsArr.length; i++) {
        await nerifToken.ownerMint(expectedValidatorsArr[i].address, stakeAmount.mul(i + 1));
        await nerifToken.connect(expectedValidatorsArr[i]).approve(staking.address, stakeAmount.mul(i + 1));
        await staking.connect(expectedValidatorsArr[i]).stake(stakeAmount.mul(i + 1));
      }
      expect(await dkg.getAllValidatorsCount()).to.be.eq(expectedValidatorsArr.length);

      const userStakesInfo = await staking.getUsersStakeInfo(0, 10);

      for (let i = 0; i < userStakesInfo.length; i++) {
        expect(userStakesInfo[i].userAddr).to.be.eq(expectedValidatorsArr[i].address);
        expect(userStakesInfo[i].userStakedAmount).to.be.eq(stakeAmount.mul(i + 1));
      }
    });
  });
});
