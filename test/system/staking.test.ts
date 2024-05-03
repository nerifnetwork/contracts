import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import {
  ContractsRegistry,
  TestDKG,
  NerifToken,
  TestStaking,
  BillingManager,
  IBillingManager,
} from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setNextBlockTime, setTime } from '../helpers/block-helper';
import { BigNumber } from 'ethers';

const { signPermit } = require('../helpers/signatures.js');

describe('Staking', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let SLASHING_VOTING: SignerWithAddress;
  let SOME_TOKEN: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;
  let billingManager: BillingManager;
  let staking: TestStaking;
  let dkg: TestDKG;
  let nerifToken: NerifToken;

  const FIRST_PK: string = '59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d';

  const tokensAmount = wei('1000');
  const defMinimalStake = wei('100');

  const msgToSign = 'verify';
  const defUpdateCollectionsEpochDuration = wei('7200', 0);
  const defDKGGenerationEpochDuration = wei('600', 0);
  const defGuaranteedWorkingEpochDuration = wei('14400', 0);

  const startTime = wei('30000', 0);
  const startEpochId = wei('1', 0);

  const nativeDepositAssetKey: string = 'NATIVE';
  const nerifTokenDepositAssetKey: string = 'NERIF';

  let nativeDepositAssetData: IBillingManager.DepositAssetDataStruct;
  let nerifTokenDepositAssetData: IBillingManager.DepositAssetDataStruct;

  function getDKGGenPeriodStartTime(startTime: BigNumber) {
    return startTime.add(defGuaranteedWorkingEpochDuration).add(defUpdateCollectionsEpochDuration);
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
    [OWNER, FIRST, SECOND, SLASHING_VOTING, SOME_TOKEN] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const StakingFactory = await ethers.getContractFactory('TestStaking');
    const DKGFactory = await ethers.getContractFactory('TestDKG');
    const TokensVestingFactory = await ethers.getContractFactory('TokensVesting');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');

    const RegistryFactory = await ethers.getContractFactory('Registry');
    const BillingManagerFactory = await ethers.getContractFactory('BillingManager');
    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const stakingImpl = await StakingFactory.deploy();
    const dkgImpl = await DKGFactory.deploy();
    const nerifTokenImpl = await NerifTokenFactory.deploy();

    const registryImpl = await RegistryFactory.deploy();
    const billingManagerImpl = await BillingManagerFactory.deploy();
    const gatewayImpl = await GatewayImplFactory.deploy();
    const gatewayFactoryImpl = await GatewayFactoryFactory.deploy();

    const tokensVesting = await TokensVestingFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.DKG_NAME(), dkgImpl.address);
    await contractsRegistry.addProxyContract(await contractsRegistry.STAKING_NAME(), stakingImpl.address);
    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);

    await contractsRegistry.addProxyContract(await contractsRegistry.REGISTRY_NAME(), registryImpl.address);
    await contractsRegistry.addProxyContract(
      await contractsRegistry.BILLING_MANAGER_NAME(),
      billingManagerImpl.address
    );
    await contractsRegistry.addProxyContract(
      await contractsRegistry.GATEWAY_FACTORY_NAME(),
      gatewayFactoryImpl.address
    );

    staking = StakingFactory.attach(await contractsRegistry.getStakingContract());
    dkg = DKGFactory.attach(await contractsRegistry.getDKGContract());
    nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());
    billingManager = BillingManagerFactory.attach(await contractsRegistry.getBillingManagerContract());
    const gatewayFactory = GatewayFactoryFactory.attach(await contractsRegistry.getGatewayFactoryContract());
    const registry = RegistryFactory.attach(await contractsRegistry.getRegistryContract());

    await contractsRegistry.addContract(await contractsRegistry.TOKENS_VESTING_NAME(), tokensVesting.address);
    await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), dkg.address);
    await contractsRegistry.addContract(await contractsRegistry.SLASHING_VOTING_NAME(), SLASHING_VOTING.address);

    nativeDepositAssetData = {
      tokenAddr: ethers.constants.AddressZero,
      workflowExecutionDiscount: 0,
      totalAssetAmount: 0,
      isPermitable: false,
      isEnabled: true,
    };
    nerifTokenDepositAssetData = {
      tokenAddr: nerifToken.address,
      workflowExecutionDiscount: 10,
      totalAssetAmount: 0,
      isPermitable: true,
      isEnabled: true,
    };

    await contractsRegistry.setIsMainChain(true);

    await nerifToken.initialize(tokensAmount, 'NERIF', 'NERIF');
    await staking.initialize(nerifToken.address, defMinimalStake, [OWNER.address]);

    await gatewayFactory.initialize(gatewayImpl.address);
    await registry.initialize(0);
    await billingManager.initialize(
      { depositAssetKey: nativeDepositAssetKey, depositAssetData: nativeDepositAssetData },
      { depositAssetKey: nerifTokenDepositAssetKey, depositAssetData: nerifTokenDepositAssetData }
    );

    await setNextBlockTime(startTime.toNumber());

    await dkg.initialize(
      defUpdateCollectionsEpochDuration,
      defDKGGenerationEpochDuration,
      defGuaranteedWorkingEpochDuration
    );

    await contractsRegistry.injectDependencies(await contractsRegistry.DKG_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.STAKING_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.BILLING_MANAGER_NAME());

    await nerifToken.approve(staking.address, tokensAmount);
    await staking.stake(defMinimalStake);

    await dkg.setSigner(startEpochId, OWNER.address);

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

  describe('slashValidator', () => {
    const stakeAmount = wei('150');

    it('should correctly slash validator', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, tokensAmount);
      await staking.connect(FIRST).stake(stakeAmount);

      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount);
      expect(await billingManager.getTotalDepositAssetAmount(nerifTokenDepositAssetKey)).to.be.eq(0);

      await staking.connect(SLASHING_VOTING).slashValidator(FIRST.address);

      expect(await billingManager.getTotalDepositAssetAmount(nerifTokenDepositAssetKey)).to.be.eq(stakeAmount);
      expect(await staking.getStake(FIRST.address)).to.be.eq(0);
    });

    it('should get exception if the stake token is not the NERIF token', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, tokensAmount);
      await staking.connect(FIRST).stake(stakeAmount);

      await staking.setStakeToken(SOME_TOKEN.address);

      const reason = 'Staking: Stake token not a NERIF token';

      await expect(staking.connect(SLASHING_VOTING).slashValidator(FIRST.address)).to.be.revertedWith(reason);
    });

    it('should get exception if not a SlashingVoting tries to call this function', async () => {
      const reason = 'Staking: Not a slashing voting address';

      await expect(staking.slashValidator(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('stake', () => {
    const stakeAmount = wei('75');

    beforeEach('setup', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, tokensAmount);
    });

    it('should correctly stake tokens and add address to the validators list', async () => {
      const tx = await staking.connect(FIRST).stake(stakeAmount);

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount);

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      let totalStake = stakeAmount.add(defMinimalStake);

      expect(await staking.totalStake()).to.be.eq(totalStake);
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount);

      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(tokensAmount.sub(stakeAmount));
      expect(await nerifToken.balanceOf(staking.address)).to.be.eq(totalStake);

      await staking.connect(FIRST).stake(stakeAmount);

      totalStake = totalStake.add(stakeAmount);

      expect(await staking.totalStake()).to.be.eq(totalStake);
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount.mul(2));

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.getActiveEpochId()).to.be.eq(startEpochId);
    });

    it('should correctly stake tokens and not add validator to the validators list', async () => {
      const tx = await staking.connect(FIRST).stake(stakeAmount.div(2));

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount.div(2));

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.div(2).add(defMinimalStake));
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount.div(2));
    });

    it('should correctly stake tokens several times with adding to the validators', async () => {
      let tx = await staking.connect(FIRST).stake(stakeAmount);

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount);

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      tx = await staking.connect(FIRST).stake(stakeAmount);

      await expect(tx).to.emit(staking, 'TokensStaked').withArgs(FIRST.address, FIRST.address, stakeAmount);

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);
    });

    it('should get exception if user with pending announcement try to stake', async () => {
      const reason = 'Staking: User has a withdrawal announcement';

      await staking.connect(FIRST).stake(stakeAmount);
      await staking.connect(FIRST).announceWithdrawal();

      await expect(staking.connect(FIRST).stake(stakeAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if try to stake zero amount', async () => {
      const reason = 'Staking: Zero stake amount';

      await expect(staking.connect(FIRST).stake(0)).to.be.revertedWith(reason);
    });

    it('should get exception if try to stake amount that higher than balance', async () => {
      const reason = 'Staking: Not enough tokens to stake';

      await expect(staking.connect(FIRST).stake(tokensAmount.add(1))).to.be.revertedWith(reason);
    });

    it('should get exception if the validator tries to stake', async () => {
      const reason = 'Staking: User is already a validator';

      await expect(staking.connect(OWNER).stake(stakeAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if not whitelisted validator tries to stake', async () => {
      const reason = 'Staking: Not a whitelisted user';

      await expect(staking.connect(SECOND).stake(stakeAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if the slashed validator tries to stake', async () => {
      const reason = 'Staking: Validator was slashed';

      await dkg.setSlashedValidator(FIRST.address, true);

      await expect(staking.connect(FIRST).stake(stakeAmount)).to.be.revertedWith(reason);
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

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);

      expect(await staking.totalStake()).to.be.eq(stakeAmount.add(defMinimalStake));
      expect(await staking.getStake(FIRST.address)).to.be.eq(stakeAmount);
    });

    it('should get exception if not whitelisted validator tries to stake', async () => {
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

    it('should get exception if the slashed validator tries to stake', async () => {
      const reason = 'Staking: Validator was slashed';

      await dkg.setSlashedValidator(FIRST.address, true);

      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        stakeAmount.toString()
      );

      await expect(
        staking.connect(FIRST).stakeWithPermit(stakeAmount, deadline, sig.v, sig.r, sig.s)
      ).to.be.revertedWith(reason);
    });
  });

  describe('announceWithdrawal', () => {
    const stakeAmount = wei('150');
    const announceAmount = wei('30');
    let secondEpochId: BigNumber;
    let secondEpochStartTime: BigNumber;
    let activePeriodStartTime: BigNumber;
    let announceTime: BigNumber;

    beforeEach('setup', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await nerifToken.ownerMint(SECOND.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address, SECOND.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, stakeAmount);
      await staking.connect(FIRST).stake(stakeAmount);

      const firstSignature = FIRST.signMessage(msgToSign);
      secondEpochId = startEpochId.add('1');
      secondEpochStartTime = getDKGGenPeriodStartTime(startTime).add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());

      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);

      activePeriodStartTime = secondEpochStartTime.add(defGuaranteedWorkingEpochDuration).add('100');
      announceTime = activePeriodStartTime.add('100');

      await setTime(activePeriodStartTime.toNumber());

      expect(await dkg.getActiveEpochStatus()).to.be.eq(2);

      expect((await dkg.getDKGEpochInfo(secondEpochId)).epochStartTime).to.be.eq(secondEpochStartTime);
    });

    it('should correctly announce withdrawal for non validator', async () => {
      await nerifToken.connect(SECOND).approve(staking.address, announceAmount);
      await staking.connect(SECOND).stake(announceAmount);

      expect(await dkg.isValidator(SECOND.address)).to.be.eq(false);

      await setNextBlockTime(announceTime.toNumber());
      const tx = await staking.connect(SECOND).announceWithdrawal();

      await expect(tx).emit(staking, 'WithdrawalAnnounced').withArgs(SECOND.address, announceAmount, secondEpochId);

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(SECOND.address);

      expect(withdrawalInfo.withdrawalEpochId).to.be.eq(secondEpochId);
      expect(withdrawalInfo.tokensAmount).to.be.eq(announceAmount);

      expect(await dkg.getActiveEpochStatus()).to.be.eq(2);
    });

    it('should correctly announce withdrawal with DKG gen period creation', async () => {
      await setNextBlockTime(announceTime.toNumber());
      const tx = await staking.connect(FIRST).announceWithdrawal();

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);

      const expectedWithdrawalEpochId = secondEpochId.add('1');

      await expect(tx)
        .emit(staking, 'WithdrawalAnnounced')
        .withArgs(FIRST.address, stakeAmount, expectedWithdrawalEpochId);

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(FIRST.address);

      expect(withdrawalInfo.withdrawalEpochId).to.be.eq(expectedWithdrawalEpochId);
      expect(withdrawalInfo.tokensAmount).to.be.eq(stakeAmount);

      const expectedDKGGenPeriodStartTime = announceTime.add(defUpdateCollectionsEpochDuration);

      expect((await dkg.getDKGEpochInfo(secondEpochId)).dkgGenPeriodStartTime).to.be.eq(expectedDKGGenPeriodStartTime);

      expect(await dkg.getActiveEpochStatus()).to.be.eq(3);
    });

    it('should get exception if the user has nothing to withdraw', async () => {
      const reason = 'Staking: Nothing to withdraw';

      await expect(staking.connect(SECOND).announceWithdrawal()).to.be.revertedWith(reason);
    });

    it('should get exception if the validator already has withdrawal announcement', async () => {
      const reason = 'Staking: User already has withdrawal announcement';

      await staking.connect(FIRST).announceWithdrawal();

      await expect(staking.connect(FIRST).announceWithdrawal()).to.be.revertedWith(reason);
    });

    it('should get exception if the slashed validator tries to announce withdraw', async () => {
      const reason = 'Staking: Validator was slashed';

      await dkg.setSlashedValidator(FIRST.address, true);

      await expect(staking.connect(FIRST).announceWithdrawal()).to.be.revertedWith(reason);
    });
  });

  describe('withdraw', () => {
    const stakeAmount = wei('75');

    let secondEpochId: BigNumber;
    let secondEpochStartTime: BigNumber;
    let activePeriodStartTime: BigNumber;
    let announceTime: BigNumber;

    beforeEach('setup', async () => {
      await nerifToken.ownerMint(FIRST.address, tokensAmount);
      await nerifToken.ownerMint(SECOND.address, tokensAmount);
      await staking.updateWhitelistedUsers([FIRST.address, SECOND.address], true);

      await nerifToken.connect(FIRST).approve(staking.address, stakeAmount.mul(2));
      await staking.connect(FIRST).stake(stakeAmount.mul(2));

      const firstSignature = FIRST.signMessage(msgToSign);
      secondEpochId = startEpochId.add('1');
      secondEpochStartTime = getDKGGenPeriodStartTime(startTime).add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());

      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);

      activePeriodStartTime = secondEpochStartTime.add(defGuaranteedWorkingEpochDuration).add('100');
      announceTime = activePeriodStartTime.add('100');

      await setTime(activePeriodStartTime.toNumber());

      expect(await dkg.getActiveEpochStatus()).to.be.eq(2);

      await setNextBlockTime(announceTime.toNumber());
    });

    it('should correctly withdraw stake tokens for non validators', async () => {
      await nerifToken.connect(SECOND).approve(staking.address, stakeAmount);
      await staking.connect(SECOND).stake(stakeAmount);

      const withdrawalTime = announceTime.add('100');

      await setNextBlockTime(withdrawalTime.toNumber());
      await staking.connect(SECOND).announceWithdrawal();

      expect(await dkg.isValidator(SECOND.address)).to.be.eq(false);

      expect((await staking.getWithdrawalAnnouncement(SECOND.address)).withdrawalEpochId).to.be.eq(secondEpochId);

      const tx = await staking.connect(SECOND).withdraw();

      await expect(tx).to.emit(staking, 'TokensWithdrawn').withArgs(SECOND.address, stakeAmount);

      const withdrawalInfo = await staking.getWithdrawalAnnouncement(SECOND.address);

      expect(withdrawalInfo.withdrawalEpochId).to.be.eq(0);
      expect(withdrawalInfo.tokensAmount).to.be.eq(0);

      expect(await dkg.isActiveValidator(SECOND.address)).to.be.eq(false);
      expect(await dkg.isValidator(SECOND.address)).to.be.eq(false);

      const totalStake = stakeAmount.mul(2).add(defMinimalStake);

      expect(await staking.totalStake()).to.be.eq(totalStake);
      expect(await staking.getStake(SECOND.address)).to.be.eq(0);

      expect(await nerifToken.balanceOf(SECOND.address)).to.be.eq(tokensAmount);
      expect(await nerifToken.balanceOf(staking.address)).to.be.eq(totalStake);
    });

    it('should correctly withdraw tokens for validator', async () => {
      await staking.connect(FIRST).announceWithdrawal();

      const expectedWithdrawalEpochId = secondEpochId.add('1');
      const expectedDKGGenerationStartTime = announceTime.add(defUpdateCollectionsEpochDuration).add('10');

      expect((await staking.getWithdrawalAnnouncement(FIRST.address)).withdrawalEpochId).to.be.eq(
        expectedWithdrawalEpochId
      );

      await setNextBlockTime(expectedDKGGenerationStartTime.toNumber());

      const secondSignature = SECOND.signMessage(msgToSign);
      await dkg.voteSigner(secondEpochId, SECOND.address, secondSignature);
      await dkg.connect(FIRST).voteSigner(secondEpochId, SECOND.address, secondSignature);

      expect(await dkg.getActiveEpochId()).to.be.eq(expectedWithdrawalEpochId);
      const tx = await staking.connect(FIRST).withdraw();

      await expect(tx).to.emit(staking, 'TokensWithdrawn').withArgs(FIRST.address, stakeAmount.mul(2));

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      expect(await staking.totalStake()).to.be.eq(defMinimalStake);
      expect(await staking.getStake(SECOND.address)).to.be.eq(0);
    });

    it('should get exception if try to withdraw without announcement', async () => {
      const reason = 'Staking: User does not have withdrawal announcement';

      await expect(staking.connect(SECOND).withdraw()).to.be.revertedWith(reason);
    });

    it('should get exception if the withdrawal time has not come', async () => {
      await staking.connect(FIRST).announceWithdrawal();

      const reason = 'Staking: The time for withdrawal has not come';

      await expect(staking.connect(FIRST).withdraw()).to.be.revertedWith(reason);
    });

    it('should get exception if the slashed validator tries to withdraw', async () => {
      await staking.connect(FIRST).announceWithdrawal();

      const reason = 'Staking: Validator was slashed';

      await dkg.setSlashedValidator(FIRST.address, true);

      await expect(staking.connect(FIRST).withdraw()).to.be.revertedWith(reason);
    });
  });

  describe('getters', () => {
    const stakeAmount = wei('150');

    it('should correctly return validators info', async () => {
      await staking.updateWhitelistedUsers([FIRST.address, SECOND.address], true);

      const expectedValidatorsArr = [OWNER, FIRST, SECOND];
      const expectedValidatorsStakeArr = [defMinimalStake];

      for (let i = 1; i < expectedValidatorsArr.length; i++) {
        await nerifToken.ownerMint(expectedValidatorsArr[i].address, stakeAmount.mul(i + 1));
        await nerifToken.connect(expectedValidatorsArr[i]).approve(staking.address, stakeAmount.mul(i + 1));
        await staking.connect(expectedValidatorsArr[i]).stake(stakeAmount.mul(i + 1));

        expectedValidatorsStakeArr.push(stakeAmount.mul(i + 1));
      }
      expect(await dkg.getAllValidatorsCount()).to.be.eq(expectedValidatorsArr.length);

      const userStakesInfo = await staking.getUsersStakeInfo(0, 10);

      for (let i = 0; i < userStakesInfo.length; i++) {
        expect(userStakesInfo[i].userAddr).to.be.eq(expectedValidatorsArr[i].address);
        expect(userStakesInfo[i].userStakedAmount).to.be.eq(expectedValidatorsStakeArr[i]);
      }
    });
  });
});
