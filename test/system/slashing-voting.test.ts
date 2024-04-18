import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry, NerifToken, TestSlashingVoting, Staking, TestDKG } from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setNextBlockTime, setTime } from '../helpers/block-helper';
import { BigNumber, BigNumberish } from 'ethers';

describe('SlashingVoting', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let THIRD: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;
  let dkg: TestDKG;
  let slashingVoting: TestSlashingVoting;
  let staking: Staking;
  let nerifToken: NerifToken;

  const tokensAmount = wei('1000');
  const defMinimalStake = wei('100');
  const defVotingThresholdPercentage = wei('50', 25);

  const defUpdateCollectionsEpochDuration = wei('7200', 0);
  const defDKGGenerationEpochDuration = wei('600', 0);
  const defGuaranteedWorkingEpochDuration = wei('14400', 0);

  const startTime = wei('20000', 0);
  const startEpochId = wei('1', 0);

  function getEndDKGPeriodTime(startEpochTime: BigNumber) {
    return startEpochTime.add(defUpdateCollectionsEpochDuration).add(defDKGGenerationEpochDuration);
  }

  function getEpochEndTime(startEpochTime: BigNumber) {
    return startEpochTime
      .add(defUpdateCollectionsEpochDuration)
      .add(defDKGGenerationEpochDuration)
      .add(defGuaranteedWorkingEpochDuration);
  }

  async function stake(
    signer: SignerWithAddress,
    stakeTime: BigNumber = BigNumber.from('0'),
    stakeAmount: BigNumberish = defMinimalStake
  ) {
    await nerifToken.ownerMint(signer.address, tokensAmount);
    await staking.connect(SIGNER).updateWhitelistedUsers([signer.address], true);
    await nerifToken.connect(signer).approve(staking.address, tokensAmount);

    if (!stakeTime.eq('0')) {
      await setNextBlockTime(stakeTime.toNumber());
    }

    await staking.connect(signer).stake(stakeAmount);
  }

  before(async () => {
    [OWNER, FIRST, SECOND, THIRD, SIGNER] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const DKGFactory = await ethers.getContractFactory('TestDKG');
    const SlashingVotingFactory = await ethers.getContractFactory('TestSlashingVoting');
    const StakingFactory = await ethers.getContractFactory('Staking');
    const RewardDistributionPoolFactory = await ethers.getContractFactory('RewardDistributionPool');
    const TokensVestingFactory = await ethers.getContractFactory('TokensVesting');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const stakingImpl = await StakingFactory.deploy();
    const dkgImpl = await DKGFactory.deploy();
    const slashingVotingImpl = await SlashingVotingFactory.deploy();
    const rewardsDistributionPoolImpl = await RewardDistributionPoolFactory.deploy();
    const nerifTokenImpl = await NerifTokenFactory.deploy();

    const tokensVesting = await TokensVestingFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.DKG_NAME(), dkgImpl.address);
    await contractsRegistry.addProxyContract(await contractsRegistry.STAKING_NAME(), stakingImpl.address);
    await contractsRegistry.addProxyContract(
      await contractsRegistry.SLASHING_VOTING_NAME(),
      slashingVotingImpl.address
    );
    await contractsRegistry.addProxyContract(
      await contractsRegistry.REWARDS_DISTRIBUTION_POOL_NAME(),
      rewardsDistributionPoolImpl.address
    );
    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);

    staking = StakingFactory.attach(await contractsRegistry.getStakingContract());
    dkg = DKGFactory.attach(await contractsRegistry.getDKGContract());
    slashingVoting = SlashingVotingFactory.attach(await contractsRegistry.getSlashingVotingContract());
    nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());

    await contractsRegistry.addContract(await contractsRegistry.TOKENS_VESTING_NAME(), tokensVesting.address);
    await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), dkg.address);

    await nerifToken.initialize(tokensAmount, 'NERIF', 'NERIF');
    await staking.initialize(nerifToken.address, defMinimalStake, [OWNER.address]);
    await slashingVoting.initialize(defVotingThresholdPercentage);

    await dkg.initialize(
      defUpdateCollectionsEpochDuration,
      defDKGGenerationEpochDuration,
      defGuaranteedWorkingEpochDuration
    );

    await contractsRegistry.injectDependencies(await contractsRegistry.DKG_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.STAKING_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.SLASHING_VOTING_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.REWARDS_DISTRIBUTION_POOL_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());

    await setNextBlockTime(startTime.toNumber());

    await nerifToken.approve(staking.address, tokensAmount);
    await staking.stake(defMinimalStake);

    await dkg.setSigner(startEpochId, OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await slashingVoting.votingThresholdPercentage()).to.be.eq(defVotingThresholdPercentage);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(slashingVoting.initialize(defVotingThresholdPercentage)).to.be.revertedWith(reason);
    });
  });

  describe('setDependencies', () => {
    it('should correctly update dependencies', async () => {
      expect(await slashingVoting.getStaking()).to.be.eq(staking.address);

      await contractsRegistry.addContract(await contractsRegistry.STAKING_NAME(), SIGNER.address);
      await contractsRegistry.injectDependencies(await contractsRegistry.SLASHING_VOTING_NAME());

      expect(await slashingVoting.getStaking()).to.be.eq(SIGNER.address);
    });

    it('should get exception if not an injector try to call this function', async () => {
      const reason = 'Dependant: not an injector';

      await expect(slashingVoting.setDependencies(contractsRegistry.address, '0x')).to.be.revertedWith(reason);
    });
  });

  describe('setVotingThresholdPercentage', () => {
    it('should correctly set voting threshold percentage', async () => {
      const newValue = wei(60, 25);

      await slashingVoting.setVotingThresholdPercentage(newValue);

      expect(await slashingVoting.votingThresholdPercentage()).to.be.eq(newValue);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SlashingVoting: Not a signer';

      await expect(slashingVoting.connect(FIRST).setVotingThresholdPercentage(0)).to.be.revertedWith(reason);
    });
  });

  describe('createProposal', () => {
    const someReason = 'Some reason';
    let expectedValidators: SignerWithAddress[];
    let expectedValidatorAddresses: string[];
    let currentTime: BigNumber;

    beforeEach('setup', async () => {
      expectedValidators = [OWNER, FIRST, SECOND];
      expectedValidatorAddresses = expectedValidators.map((signer) => {
        return signer.address;
      });

      await dkg.setSigner(startEpochId, SIGNER.address);

      for (let i = 0; i < expectedValidators.length; i++) {
        await stake(expectedValidators[i]);
      }

      currentTime = getEpochEndTime(startTime).add('10');

      await setTime(currentTime.toNumber());

      expect(await dkg.getActiveValidators()).to.be.deep.eq(expectedValidatorAddresses);
      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);
    });

    it('should correctly create proposal and vote', async () => {
      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = currentTime.add('100');

      await setNextBlockTime(newEpochStartTime.toNumber());
      const newProposalId = await slashingVoting.callStatic.createProposal(FIRST.address, someReason);
      const tx = await slashingVoting.createProposal(FIRST.address, someReason);

      const expectedProposalId = 1;

      expect(await slashingVoting.lastProposalId()).to.be.eq(expectedProposalId);
      expect(newProposalId).to.be.eq(expectedProposalId);

      expect((await tx.wait()).events?.length).to.be.eq(3);
      await expect(tx).to.emit(dkg, 'NewEpochCreated').withArgs(newEpochId, newEpochStartTime);
      await expect(tx).to.emit(slashingVoting, 'ProposalCreated').withArgs(expectedProposalId, FIRST.address);
      await expect(tx).to.emit(slashingVoting, 'ProposalVoted').withArgs(expectedProposalId, OWNER.address);

      const propInfo = await slashingVoting.getDetailedProposalInfo(expectedProposalId);

      expect(propInfo.votedValidators.length).to.be.eq(1);
      expect(propInfo.votedValidators).to.be.deep.eq([OWNER.address]);

      expect(propInfo.baseProposalInfo.epochId).to.be.eq(newEpochId);
      expect(propInfo.baseProposalInfo.votingStartTime).to.be.eq(newEpochStartTime);
      expect(propInfo.baseProposalInfo.votingEndTime).to.be.eq(getEndDKGPeriodTime(newEpochStartTime));

      expect(await slashingVoting.hasPendingSlashingProposal(FIRST.address)).to.be.eq(true);
    });

    it('should correctly create voting without vote', async () => {
      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = currentTime.add('100');

      await stake(THIRD, secondEpochStartTime);

      await dkg.setSigner(secondEpochId, SIGNER.address);
      await setTime(getEndDKGPeriodTime(secondEpochStartTime).toNumber());

      const thirdEpochId = secondEpochId.add('1');
      const thirdEpochStartTime = getEpochEndTime(secondEpochStartTime);

      const tx = await slashingVoting.createProposal(FIRST.address, someReason);

      const expectedProposalId = 1;

      expect(await slashingVoting.lastProposalId()).to.be.eq(expectedProposalId);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).to.emit(dkg, 'NewEpochCreated').withArgs(thirdEpochId, thirdEpochStartTime);
      await expect(tx).to.emit(slashingVoting, 'ProposalCreated').withArgs(expectedProposalId, FIRST.address);

      const propInfo = await slashingVoting.getDetailedProposalInfo(expectedProposalId);

      expect(propInfo.votedValidators.length).to.be.eq(0);
      expect(propInfo.votedValidators).to.be.deep.eq([]);

      expect(propInfo.baseProposalInfo.epochId).to.be.eq(thirdEpochId);
      expect(propInfo.baseProposalInfo.votingStartTime).to.be.eq(thirdEpochStartTime);
      expect(propInfo.baseProposalInfo.votingEndTime).to.be.eq(getEndDKGPeriodTime(thirdEpochStartTime));

      expect(await slashingVoting.hasPendingSlashingProposal(FIRST.address)).to.be.eq(true);
    });

    it('should get exception if not an active validator tried to call this function', async () => {
      const reason = 'SlashingVoting: Not an active system validator';

      await expect(slashingVoting.connect(SIGNER).createProposal(FIRST.address, someReason)).to.be.revertedWith(reason);
    });

    it('should get exception if try to create proposal for non active validator', async () => {
      const reason = 'SlashingVoting: Target is not an active validator';

      await expect(slashingVoting.createProposal(THIRD.address, someReason)).to.be.revertedWith(reason);
    });

    it('should get exception if the validator already has pending proposal', async () => {
      const reason = 'SlashingVoting: Validator already has pending proposal';

      slashingVoting.createProposal(FIRST.address, someReason);

      await expect(slashingVoting.createProposal(FIRST.address, someReason)).to.be.revertedWith(reason);
    });
  });

  describe('vote', () => {
    const proposalId = 1;
    const activeValidatorsCount = 5;
    const someReason = 'Some reason';

    let expectedValidators: SignerWithAddress[];
    let expectedValidatorAddresses: string[];
    let currentTime: BigNumber;
    let newEpochStartTime: BigNumber;

    beforeEach('setup', async () => {
      expectedValidators = (await ethers.getSigners()).slice(0, activeValidatorsCount);
      expectedValidatorAddresses = expectedValidators.map((signer) => {
        return signer.address;
      });

      await dkg.setSigner(startEpochId, SIGNER.address);

      for (let i = 0; i < expectedValidators.length; i++) {
        await stake(expectedValidators[i]);
      }

      currentTime = getEpochEndTime(startTime).add('10');

      await setTime(currentTime.toNumber());

      expect(await dkg.getActiveValidators()).to.be.deep.eq(expectedValidatorAddresses);
      expect(await dkg.getActiveValidatorsCount()).to.be.eq(activeValidatorsCount);
      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);

      newEpochStartTime = currentTime.add('100');

      await setNextBlockTime(newEpochStartTime.toNumber());
      await slashingVoting.createProposal(FIRST.address, someReason);
    });

    it('should correctly vote without execution', async () => {
      const tx = await slashingVoting.connect(SECOND).vote(proposalId);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).to.emit(slashingVoting, 'ProposalVoted').withArgs(proposalId, SECOND.address);

      const propInfo = await slashingVoting.getDetailedProposalInfo(proposalId);

      expect(propInfo.votedValidators.length).to.be.eq(2);
      expect(propInfo.votedValidators).to.be.deep.eq([OWNER.address, SECOND.address]);
    });

    it('should correctly vote with execution', async () => {
      await slashingVoting.connect(SECOND).vote(proposalId);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidatorSlashed(FIRST.address)).to.be.eq(false);

      const tx = await slashingVoting.connect(THIRD).vote(proposalId);

      expect((await tx.wait()).events?.length).to.be.eq(3);
      await expect(tx).to.emit(dkg, 'ValidatorSlashed').withArgs(FIRST.address);
      await expect(tx).to.emit(slashingVoting, 'ProposalVoted').withArgs(proposalId, THIRD.address);
      await expect(tx).to.emit(slashingVoting, 'ProposalExecuted').withArgs(proposalId, FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidatorSlashed(FIRST.address)).to.be.eq(true);

      expect(await slashingVoting.isProposalExecuted(proposalId)).to.be.eq(true);

      const propInfo = await slashingVoting.getDetailedProposalInfo(proposalId);

      expect(propInfo.votedValidators.length).to.be.eq(3);
      expect(propInfo.votedValidators).to.be.deep.eq([OWNER.address, SECOND.address, THIRD.address]);
      expect(propInfo.isExecuted).to.be.eq(true);
    });

    it('should get exception if not an active validator tried to call this function', async () => {
      const reason = 'SlashingVoting: Not an active system validator';

      await expect(slashingVoting.connect((await ethers.getSigners())[10]).vote(proposalId)).to.be.revertedWith(reason);
    });

    it('should get exception if proposal is already executed', async () => {
      const reason = 'SlashingVoting: Proposal has already executed';

      await slashingVoting.connect(SECOND).vote(proposalId);
      await slashingVoting.connect(THIRD).vote(proposalId);

      await expect(slashingVoting.connect(SECOND).vote(proposalId)).to.be.revertedWith(reason);
    });

    it('should get exception if the voting has not started yet or finished', async () => {
      const reason = "SlashingVoting: Voting hasn't started yet or finished";

      await setTime(getEndDKGPeriodTime(newEpochStartTime).add('100').toNumber());

      await expect(slashingVoting.connect(SECOND).vote(proposalId)).to.be.revertedWith(reason);

      const newProposalId = proposalId + 1;

      await slashingVoting.createProposal(FIRST.address, someReason);

      await expect(slashingVoting.connect(SECOND).vote(newProposalId)).to.be.revertedWith(reason);
    });

    it('should get exception if the validator has already voted', async () => {
      const reason = 'SlashingVoting: Validator has already voted';

      await expect(slashingVoting.vote(proposalId)).to.be.revertedWith(reason);
    });
  });

  describe('getValidatorsPercentage', () => {
    let expectedValidators: SignerWithAddress[];
    let expectedValidatorAddresses: string[];
    let currentTime: BigNumber;

    beforeEach('setup', async () => {
      expectedValidators = (await ethers.getSigners()).slice(0, 8);
      expectedValidatorAddresses = expectedValidators.map((signer) => {
        return signer.address;
      });

      await dkg.setSigner(startEpochId, SIGNER.address);

      for (let i = 0; i < expectedValidators.length; i++) {
        await stake(expectedValidators[i]);
      }

      currentTime = getEpochEndTime(startTime).add('10');

      await setTime(currentTime.toNumber());

      expect(await dkg.getActiveValidators()).to.be.deep.eq(expectedValidatorAddresses);
      expect(await dkg.getActiveValidatorsCount()).to.be.eq(8);
      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);
    });

    it('should return correct percentages', async () => {
      expect(await slashingVoting.getValidatorsPercentage(4)).to.be.eq(wei(50, 25));
      expect(await slashingVoting.getValidatorsPercentage(1)).to.be.eq(wei(125, 24));
      expect(await slashingVoting.getValidatorsPercentage(5)).to.be.eq(wei(625, 24));
    });
  });
});
