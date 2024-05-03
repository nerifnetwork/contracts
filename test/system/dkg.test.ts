import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry, TestDKG } from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setNextBlockTime, setTime } from '../helpers/block-helper';
import { BigNumber } from 'ethers';

describe('DKG', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let THIRD: SignerWithAddress;
  let STAKING: SignerWithAddress;
  let SLASHING_VOTING: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;
  let dkg: TestDKG;

  const msgToSign = 'verify';
  const defUpdateCollectionsEpochDuration = wei('7200', 0);
  const defDKGGenerationEpochDuration = wei('600', 0);
  const defGuaranteedWorkingEpochDuration = wei('14400', 0);

  const startTime = wei('10000', 0);
  const startEpochId = wei('1', 0);

  function getDKGGenPeriodStartTime(startTime: BigNumber) {
    return startTime.add(defGuaranteedWorkingEpochDuration).add(defUpdateCollectionsEpochDuration);
  }

  before(async () => {
    [OWNER, FIRST, SECOND, THIRD, STAKING, SLASHING_VOTING] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const DKGFactory = await ethers.getContractFactory('TestDKG');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const dkgImpl = await DKGFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.DKG_NAME(), dkgImpl.address);
    await contractsRegistry.addContract(await contractsRegistry.STAKING_NAME(), STAKING.address);
    await contractsRegistry.addContract(await contractsRegistry.SLASHING_VOTING_NAME(), SLASHING_VOTING.address);

    dkg = DKGFactory.attach(await contractsRegistry.getDKGContract());

    await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), dkg.address);

    await setNextBlockTime(startTime.toNumber());

    await dkg.initialize(
      defUpdateCollectionsEpochDuration,
      defDKGGenerationEpochDuration,
      defGuaranteedWorkingEpochDuration
    );

    await contractsRegistry.injectDependencies(await contractsRegistry.DKG_NAME());

    await dkg.connect(STAKING).addValidator(OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await dkg.updatesCollectionEpochDuration()).to.be.eq(defUpdateCollectionsEpochDuration);
      expect(await dkg.dkgGenerationEpochDuration()).to.be.eq(defDKGGenerationEpochDuration);
      expect(await dkg.guaranteedWorkingEpochDuration()).to.be.eq(defGuaranteedWorkingEpochDuration);

      expect(await dkg.getActiveEpochId()).to.be.eq(startEpochId);
      expect(await dkg.getActiveEpochStatus()).to.be.eq(1);

      const epochInfo = await dkg.getDKGEpochInfo(startEpochId);
      const expectedDKGGenPeriodStartTime = getDKGGenPeriodStartTime(startTime);

      expect(epochInfo.epochStartTime).to.be.eq(startTime);
      expect(epochInfo.dkgGenPeriodStartTime).to.be.eq(expectedDKGGenPeriodStartTime);
      expect(epochInfo.epochSigner).to.be.eq(ethers.constants.AddressZero);
      expect(epochInfo.epochStatus).to.be.eq(1);

      expect(await dkg.isActiveValidator(OWNER.address)).to.be.eq(true);
      expect(await dkg.getActiveValidatorsCount()).to.be.eq(1);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(
        dkg.initialize(
          defUpdateCollectionsEpochDuration,
          defDKGGenerationEpochDuration,
          defGuaranteedWorkingEpochDuration
        )
      ).to.be.revertedWith(reason);
    });
  });

  describe('setDependencies', () => {
    it('should correctly update dependencies', async () => {
      expect(await dkg.getStaking()).to.be.eq(STAKING.address);

      await contractsRegistry.addContract(await contractsRegistry.STAKING_NAME(), FIRST.address);
      await contractsRegistry.injectDependencies(await contractsRegistry.DKG_NAME());

      expect(await dkg.getStaking()).to.be.eq(FIRST.address);
    });

    it('should get exception if not a registry contract try to call this function', async () => {
      const reason = 'Dependant: not an injector';

      await expect(dkg.setDependencies(contractsRegistry.address, '0x')).to.be.revertedWith(reason);
    });
  });

  describe('addValidator', () => {
    let firstSignature: string;

    beforeEach('setup', async () => {
      firstSignature = await FIRST.signMessage(msgToSign);
    });

    it('should corrctly add initial validators', async () => {
      const tx = await dkg.connect(STAKING).addValidator(FIRST.address);

      await expect(tx).emit(dkg, 'NewValidatorAdded').withArgs(FIRST.address, startEpochId);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);

      const validatorsData = await dkg.getValidatorInfo(FIRST.address);

      expect(validatorsData.validator).to.be.eq(FIRST.address);
      expect(validatorsData.startValidationEpochId).to.be.eq(startEpochId);
      expect(validatorsData.endValidationEpochId).to.be.eq(ethers.constants.MaxUint256);
    });

    it('should correctly add new validator with new DKG generation period creation', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      let expectedDKGGenerationStartTime = getDKGGenPeriodStartTime(startTime);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = expectedDKGGenerationStartTime.add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      expect(await dkg.getActiveEpochId()).to.be.eq(secondEpochId);
      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address]);

      let epochInfo = await dkg.getDKGEpochInfo(secondEpochId);

      expect(epochInfo.dkgGenPeriodStartTime).to.be.eq(0);

      const expectedStartValidationEpochId = secondEpochId.add('1');
      const tx = await dkg.connect(STAKING).addValidator(SECOND.address);

      await expect(tx).emit(dkg, 'NewValidatorAdded').withArgs(SECOND.address, expectedStartValidationEpochId);

      const validatorsData = await dkg.getValidatorInfo(SECOND.address);

      expect(validatorsData.validator).to.be.eq(SECOND.address);
      expect(validatorsData.startValidationEpochId).to.be.eq(expectedStartValidationEpochId);
      expect(validatorsData.endValidationEpochId).to.be.eq(ethers.constants.MaxUint256);

      const expectedDKGGenPeriodStartTime = secondEpochStartTime
        .add(defGuaranteedWorkingEpochDuration)
        .add(defUpdateCollectionsEpochDuration);
      epochInfo = await dkg.getDKGEpochInfo(secondEpochId);

      expect(epochInfo.dkgGenPeriodStartTime).to.be.eq(expectedDKGGenPeriodStartTime);
    });

    it('should correctly add validator during ACTIVE period', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      let expectedDKGGenerationStartTime = getDKGGenPeriodStartTime(startTime);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = expectedDKGGenerationStartTime.add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      const activePeriodStartTime = secondEpochStartTime.add(defGuaranteedWorkingEpochDuration).add('1000');
      const expectedStartValidationEpochId = secondEpochId.add('1');

      await setNextBlockTime(activePeriodStartTime.toNumber());
      await dkg.connect(STAKING).addValidator(SECOND.address);

      const validatorsData = await dkg.getValidatorInfo(SECOND.address);

      expect(validatorsData.startValidationEpochId).to.be.eq(expectedStartValidationEpochId);

      const expectedDKGGenPeriodStartTime = activePeriodStartTime.add(defUpdateCollectionsEpochDuration);
      const epochInfo = await dkg.getDKGEpochInfo(secondEpochId);

      expect(epochInfo.dkgGenPeriodStartTime).to.be.eq(expectedDKGGenPeriodStartTime);
    });

    it('should correctly add validator with second DKG generation period creation', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      let expectedDKGGenerationStartTime = getDKGGenPeriodStartTime(startTime);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = expectedDKGGenerationStartTime.add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      const activePeriodStartTime = secondEpochStartTime.add(defGuaranteedWorkingEpochDuration).add('1000');

      await setTime(activePeriodStartTime.sub('10').toNumber());
      expect(await dkg.getActiveEpochStatus()).to.be.eq(2);

      await setNextBlockTime(activePeriodStartTime.toNumber());
      await dkg.connect(STAKING).addValidator(SECOND.address);

      expect(await dkg.getActiveEpochStatus()).to.be.eq(3);

      let expectedDKGGenPeriodStartTime = activePeriodStartTime.add(defUpdateCollectionsEpochDuration);
      const newActivePeriodTime = expectedDKGGenPeriodStartTime.add(defDKGGenerationEpochDuration).add('100');

      await setTime(newActivePeriodTime.sub('10').toNumber());
      expect(await dkg.getActiveEpochStatus()).to.be.eq(2);

      await setNextBlockTime(newActivePeriodTime.toNumber());
      await dkg.connect(STAKING).addValidator(THIRD.address);

      expectedDKGGenPeriodStartTime = newActivePeriodTime.add(defUpdateCollectionsEpochDuration);

      const epochInfo = await dkg.getDKGEpochInfo(secondEpochId);

      expect(epochInfo.dkgGenPeriodStartTime).to.be.eq(expectedDKGGenPeriodStartTime);
    });

    it('should get exception if try to add validator during DKG generation period', async () => {
      const reason = 'DKG: Unable to collect updates during the DKG generation period';

      const expectedDKGGenPeriodStartTime = getDKGGenPeriodStartTime(startTime);

      await setNextBlockTime(expectedDKGGenPeriodStartTime.add('10').toNumber());

      await expect(dkg.connect(STAKING).addValidator(FIRST.address)).to.be.revertedWith(reason);
    });

    it('should get exception if not the Staking address try to call this function', async () => {
      const reason = 'DKG: Not a staking address';

      await expect(dkg.addValidator(OWNER.address)).to.be.revertedWith(reason);
    });

    it('should get exception if the validator already exists', async () => {
      const reason = 'DKG: Validator already exists';

      await dkg.connect(STAKING).addValidator(FIRST.address);

      await expect(dkg.connect(STAKING).addValidator(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('announceValidatorExit', () => {
    let firstSignature: string;
    let secondSignature: string;
    let secondEpochId: BigNumber;
    let secondEpochStartTime: BigNumber;

    beforeEach('setup', async () => {
      firstSignature = await FIRST.signMessage(msgToSign);
      secondSignature = await SECOND.signMessage(msgToSign);

      await dkg.connect(STAKING).addValidator(FIRST.address);
      await dkg.connect(STAKING).addValidator(SECOND.address);

      let expectedDKGGenerationStartTime = getDKGGenPeriodStartTime(startTime);

      secondEpochId = startEpochId.add('1');
      secondEpochStartTime = expectedDKGGenerationStartTime.add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);
    });

    it('should correctly announce validator exit', async () => {
      const tx = await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      const expectedEndValidationEpochId = secondEpochId.add('1');

      await expect(tx).emit(dkg, 'ValidatorExitAnnounced').withArgs(FIRST.address, expectedEndValidationEpochId);

      expect((await dkg.getValidatorInfo(FIRST.address)).endValidationEpochId).to.be.eq(expectedEndValidationEpochId);
    });

    it('should get exception if not the Staking address try to call this function', async () => {
      const reason = 'DKG: Not a staking address';

      await expect(dkg.announceValidatorExit(OWNER.address)).to.be.revertedWith(reason);
    });

    it('should get exception if validator is not active', async () => {
      const reason = 'DKG: Validator is not active';

      await expect(dkg.connect(STAKING).announceValidatorExit(THIRD.address)).to.be.revertedWith(reason);

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      const expectedDKGGenerationStartTime = secondEpochStartTime
        .add(defGuaranteedWorkingEpochDuration)
        .add(defUpdateCollectionsEpochDuration);

      await setTime(expectedDKGGenerationStartTime.add('10').toNumber());
      await dkg.voteSigner(secondEpochId, SECOND.address, secondSignature);
      await dkg.connect(SECOND).voteSigner(secondEpochId, SECOND.address, secondSignature);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);

      await expect(dkg.connect(STAKING).announceValidatorExit(FIRST.address)).to.be.revertedWith(reason);
    });

    it('should get exception if exit of the validator has already been announced', async () => {
      const reason = 'DKG: Exit of the validator has already been announced';

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      await expect(dkg.connect(STAKING).announceValidatorExit(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('createProposal', () => {
    it('should return correct info without new epoch creation', async () => {
      let dkgGenPeriodStartTime = await dkg.connect(SLASHING_VOTING).callStatic.createProposal();

      expect(dkgGenPeriodStartTime).to.be.eq(getDKGGenPeriodStartTime(startTime));

      const activePeriodTime = getDKGGenPeriodStartTime(startTime).add(defDKGGenerationEpochDuration).add('100');

      await setTime(activePeriodTime.toNumber());

      expect(await dkg.getActiveEpochStatus()).to.be.eq(2);

      dkgGenPeriodStartTime = await dkg.connect(SLASHING_VOTING).callStatic.createProposal();

      expect(dkgGenPeriodStartTime).to.be.eq(activePeriodTime.add(defUpdateCollectionsEpochDuration));
    });

    it('should get exception if not a slashing voting address tried to call this function', async () => {
      const reason = 'DKG: Not a slashing voting address';

      await expect(dkg.createProposal()).to.be.revertedWith(reason);
    });
  });

  describe('slashValidator', () => {
    it('should correctly slash validator', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidatorSlashed(FIRST.address)).to.be.eq(false);

      const tx = await dkg.connect(SLASHING_VOTING).slashValidator(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidatorSlashed(FIRST.address)).to.be.eq(true);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(dkg, 'ValidatorSlashed').withArgs(FIRST.address);
      await expect(tx).emit(dkg, 'ValidatorRemoved').withArgs(FIRST.address);
    });

    it('should get exception if not a slashing voting address tried to call this function', async () => {
      const reason = 'DKG: Not a slashing voting address';

      await expect(dkg.slashValidator(FIRST.address)).to.be.revertedWith(reason);
    });

    it('should get exception if validator has already slashed', async () => {
      const reason = 'DKG: Validator has already slashed';

      await dkg.connect(SLASHING_VOTING).slashValidator(FIRST.address);

      await expect(dkg.connect(SLASHING_VOTING).slashValidator(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('voteSigner', () => {
    let firstSignature: string;
    let dkgGenPeriodStartTime: BigNumber;

    beforeEach('setup', async () => {
      firstSignature = await FIRST.signMessage(msgToSign);

      await dkg.connect(STAKING).addValidator(FIRST.address);
      await dkg.connect(STAKING).addValidator(SECOND.address);

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);

      dkgGenPeriodStartTime = getDKGGenPeriodStartTime(startTime);

      await setTime(dkgGenPeriodStartTime.add('10').toNumber());

      expect(await dkg.getActiveEpochStatus()).to.be.eq(4);
    });

    it('should correctly vote for new signer', async () => {
      let tx = await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'SignerVoted').withArgs(startEpochId, OWNER.address, FIRST.address);

      expect(await dkg.getSignerVotesCount(startEpochId, FIRST.address)).to.be.eq(1);
      expect(await dkg.getValidatorVote(startEpochId, OWNER.address)).to.be.eq(FIRST.address);
      expect(await dkg.getSignerAddress()).to.be.eq(ethers.constants.AddressZero);
      expect((await dkg.getDKGEpochInfo(startEpochId)).epochSigner).to.be.eq(ethers.constants.AddressZero);

      const newEpochStartTime = dkgGenPeriodStartTime.add('20');

      await setNextBlockTime(newEpochStartTime.toNumber());
      tx = await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);

      expect((await tx.wait()).events?.length).to.be.eq(3);
      await expect(tx).emit(dkg, 'SignerVoted').withArgs(startEpochId, FIRST.address, FIRST.address);
      await expect(tx).emit(dkg, 'SignerAddressUpdated').withArgs(startEpochId, FIRST.address);
      await expect(tx).emit(dkg, 'NewEpochCreated').withArgs(startEpochId.add(1), newEpochStartTime);

      expect(await dkg.getSignerVotesCount(startEpochId, FIRST.address)).to.be.eq(2);
      expect(await dkg.getValidatorVote(startEpochId, FIRST.address)).to.be.eq(FIRST.address);
      expect(await dkg.getSignerAddress()).to.be.eq(FIRST.address);
      expect((await dkg.getDKGEpochInfo(startEpochId)).epochSigner).to.be.eq(FIRST.address);
    });

    it('should get exception if not an active validator try to call this function', async () => {
      const reason = 'DKG: Not an active validator';

      await expect(dkg.connect(THIRD).voteSigner(startEpochId, FIRST.address, firstSignature)).to.be.revertedWith(
        reason
      );
    });

    it('should get exception if try to call this function not during the DKG gen period', async () => {
      await setTime(dkgGenPeriodStartTime.add(defDKGGenerationEpochDuration).add('10').toNumber());

      expect(await dkg.getActiveEpochStatus()).to.be.eq(2);

      const reason = 'DKG: Not a DKG generation period';

      await expect(dkg.voteSigner(startEpochId, FIRST.address, firstSignature)).to.be.revertedWith(reason);
    });

    it('should get exception if pass invalid signature', async () => {
      const reason = 'DKG: Signature is invalid';

      await expect(dkg.voteSigner(startEpochId, OWNER.address, firstSignature)).to.be.revertedWith(reason);
    });

    it('should get exception if the validator has already voted', async () => {
      const reason = 'DKG: Already voted';

      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      await expect(dkg.voteSigner(startEpochId, FIRST.address, firstSignature)).to.be.revertedWith(reason);
    });
  });

  describe('getEpochStatus', () => {
    it('should return correct epoch statuses', async () => {
      expect(await dkg.getEpochStatus(123)).to.be.eq(0); // NONE
      expect(await dkg.getEpochStatus(startEpochId)).to.be.eq(1); // GUARANTEED_WORKING

      const expectedDKGGenPeriodStartTime = getDKGGenPeriodStartTime(startTime);

      await setTime(expectedDKGGenPeriodStartTime.sub('10').toNumber());

      expect(await dkg.getEpochStatus(startEpochId)).to.be.eq(3); // UPDATES_COLLECTION

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = expectedDKGGenPeriodStartTime.add('10');

      await setTime(secondEpochStartTime.toNumber());

      expect(await dkg.getEpochStatus(startEpochId)).to.be.eq(4); // DKG_GENERATION

      const firstSignature = await FIRST.signMessage(msgToSign);
      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      expect(await dkg.getEpochStatus(startEpochId)).to.be.eq(5); // FINISHED
      expect(await dkg.getEpochStatus(secondEpochId)).to.be.eq(1); // GUARANTEED_WORKING

      await setTime(secondEpochStartTime.add(defGuaranteedWorkingEpochDuration).add('100').toNumber());

      expect(await dkg.getEpochStatus(secondEpochId)).to.be.eq(2); // ACTIVE
    });
  });

  describe('isActiveValidator', () => {
    it('should return correct result for the first epoch', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect((await dkg.getValidatorInfo(FIRST.address)).startValidationEpochId).to.be.eq(startEpochId);
    });

    it('should return correct result if the validation start time has not come yet', async () => {
      const expectedDKGGenPeriodStartTime = getDKGGenPeriodStartTime(startTime);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = expectedDKGGenPeriodStartTime.add('10');

      await setTime(secondEpochStartTime.toNumber());

      const firstSignature = await FIRST.signMessage(msgToSign);
      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect((await dkg.getValidatorInfo(FIRST.address)).startValidationEpochId).to.be.eq(secondEpochId.add('1'));
    });

    it('should return correct result for slashed validator', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);

      await dkg.connect(SLASHING_VOTING).slashValidator(FIRST.address);

      expect(await dkg.isValidatorSlashed(FIRST.address)).to.be.eq(true);
      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
    });
  });

  describe('getters', () => {
    it('should return correct values', async () => {
      const firstSignature = await FIRST.signMessage(msgToSign);

      await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address]);

      const secondEpochStartTime = getDKGGenPeriodStartTime(startTime).add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());

      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);

      await dkg.connect(STAKING).addValidator(SECOND.address);

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address]);
      expect(await dkg.getAllValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);

      expect(await dkg.getActiveValidatorsCount()).to.be.eq(2);
      expect(await dkg.getAllValidatorsCount()).to.be.eq(3);
    });

    it('getDKGPeriodEndTime should revert if pass invalid epochId', async () => {
      const firstSignature = await FIRST.signMessage(msgToSign);

      await dkg.connect(STAKING).addValidator(FIRST.address);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = getDKGGenPeriodStartTime(startTime).add('10');

      await setNextBlockTime(secondEpochStartTime.sub('1').toNumber());

      await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);
      await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);

      const reason = 'DKG: DKG generation period does not exist';

      await expect(dkg.getDKGPeriodEndTime(secondEpochId)).to.be.revertedWith(reason);
    });
  });
});
