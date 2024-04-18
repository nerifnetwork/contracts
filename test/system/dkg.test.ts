import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry, IDKG, TestDKG } from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setNextBlockTime, setTime } from '../helpers/block-helper';
import { BigNumber, BigNumberish } from 'ethers';

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

  function getEndDKGPeriodTime(startEpochTime: BigNumber) {
    return startEpochTime.add(defUpdateCollectionsEpochDuration).add(defDKGGenerationEpochDuration);
  }

  function getEpochEndTime(startEpochTime: BigNumber) {
    return startEpochTime
      .add(defUpdateCollectionsEpochDuration)
      .add(defDKGGenerationEpochDuration)
      .add(defGuaranteedWorkingEpochDuration);
  }

  function checkValidationData(
    actualValidationData: IDKG.ValidationDataStruct,
    expectedValidationData: BigNumberish[]
  ) {
    expect(actualValidationData.validationTime).to.be.eq(expectedValidationData[0]);
    expect(actualValidationData.validationEpoch).to.be.eq(expectedValidationData[1]);
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

    await dkg.initialize(
      defUpdateCollectionsEpochDuration,
      defDKGGenerationEpochDuration,
      defGuaranteedWorkingEpochDuration
    );

    await contractsRegistry.injectDependencies(await contractsRegistry.DKG_NAME());

    await setNextBlockTime(startTime.toNumber());

    await dkg.connect(STAKING).addValidator(OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await dkg.updatesCollectionEpochDuration()).to.be.eq(defUpdateCollectionsEpochDuration);
      expect(await dkg.dkgGenerationEpochDuration()).to.be.eq(defDKGGenerationEpochDuration);
      expect(await dkg.guaranteedWorkingEpochDuration()).to.be.eq(defGuaranteedWorkingEpochDuration);

      expect(await dkg.getCurrentEpochId()).to.be.eq(startEpochId);
      expect(await dkg.getCurrentEpochStatus()).to.be.eq(2);

      const epochInfo = await dkg.getDKGEpochInfo(startEpochId);

      expect(epochInfo.mainEpochInfo.epochStartTime).to.be.eq(startTime);
      expect(epochInfo.epochSigner).to.be.eq(ethers.constants.AddressZero);
      expect(epochInfo.epochStatus).to.be.eq(2);

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
    it('should correctly add initial validators', async () => {
      const DKGFactory = await ethers.getContractFactory('DKG');
      const newDKG = await DKGFactory.deploy();

      await newDKG.initialize(
        defUpdateCollectionsEpochDuration,
        defDKGGenerationEpochDuration,
        defGuaranteedWorkingEpochDuration
      );
      await newDKG.setDependencies(contractsRegistry.address, '0x');

      await setNextBlockTime(startTime.mul(2).toNumber());

      const tx = await newDKG.connect(STAKING).addValidator(FIRST.address);

      const valInfo = await newDKG.getValidatorInfo(FIRST.address);

      expect(valInfo.validator).to.be.eq(FIRST.address);
      checkValidationData(valInfo.startValidationData, [startTime.mul(2), startEpochId]);
      checkValidationData(valInfo.endValidationData, [ethers.constants.MaxUint256, 0]);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(newDKG, 'NewValidatorAdded').withArgs(FIRST.address, startTime.mul(2), startEpochId);
      await expect(tx).emit(newDKG, 'NewEpochCreated').withArgs(startEpochId, startTime.mul(2));

      expect(await newDKG.isLastEpoch(startEpochId)).to.be.eq(true);
      expect(await newDKG.isCurrentEpoch(startEpochId)).to.be.eq(true);
      expect(await newDKG.getCurrentEpochStatus()).to.be.eq(2);

      expect(await newDKG.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await newDKG.getActiveValidators()).to.be.deep.eq([FIRST.address]);
    });

    it('should correctly add validator during first update collection epoch', async () => {
      const expectedStartValidationTime = startTime.add('10');

      await setNextBlockTime(expectedStartValidationTime.toNumber());
      const tx = await dkg.connect(STAKING).addValidator(FIRST.address);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx)
        .emit(dkg, 'NewValidatorAdded')
        .withArgs(FIRST.address, expectedStartValidationTime, startEpochId);

      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.getAllValidators()).to.be.deep.eq([OWNER.address, FIRST.address]);

      const valInfo = await dkg.getValidatorInfo(FIRST.address);

      expect(valInfo.validator).to.be.eq(FIRST.address);
      checkValidationData(valInfo.startValidationData, [expectedStartValidationTime, startEpochId]);
      checkValidationData(valInfo.endValidationData, [ethers.constants.MaxUint256, 0]);

      expect(await dkg.isLastEpoch(startEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(startEpochId)).to.be.eq(true);
    });

    it('should correctly add validator during next update collection epochs', async () => {
      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = getEpochEndTime(startTime).add('100');

      await setNextBlockTime(newEpochStartTime.toNumber());
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const expectedStartValidationTime = getEndDKGPeriodTime(newEpochStartTime);
      const tx = await dkg.connect(STAKING).addValidator(SECOND.address);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'NewValidatorAdded').withArgs(SECOND.address, expectedStartValidationTime, newEpochId);

      const valInfo = await dkg.getValidatorInfo(SECOND.address);

      expect(valInfo.validator).to.be.eq(SECOND.address);
      checkValidationData(valInfo.startValidationData, [expectedStartValidationTime, newEpochId]);
      checkValidationData(valInfo.endValidationData, [ethers.constants.MaxUint256, 0]);

      expect(await dkg.isValidator(SECOND.address)).to.be.eq(true);
      expect(await dkg.isActiveValidator(SECOND.address)).to.be.eq(false);
    });

    it('should correctly add validator when the epoch status is ACTIVE with new epoch creation', async () => {
      const expetcedActivePeriodStartTime = getEpochEndTime(startTime).add('100');

      await setTime(expetcedActivePeriodStartTime.toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = expetcedActivePeriodStartTime.add('100');

      await setNextBlockTime(newEpochStartTime.toNumber());

      const tx = await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.isLastEpoch(newEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(newEpochId)).to.be.eq(true);

      const dkgEpochInfo = await dkg.getDKGEpochInfo(newEpochId);

      expect(dkgEpochInfo.mainEpochInfo.epochId).to.be.eq(newEpochId);
      expect(dkgEpochInfo.mainEpochInfo.epochStartTime).to.be.eq(newEpochStartTime);
      expect(dkgEpochInfo.mainEpochInfo.dkgGenPeriodEndTime).to.be.eq(getEndDKGPeriodTime(newEpochStartTime));
      expect(dkgEpochInfo.epochStatus).to.be.eq(2);

      const expectedStartValidationTime = newEpochStartTime
        .add(defUpdateCollectionsEpochDuration)
        .add(defDKGGenerationEpochDuration);

      expect(await dkg.getDKGPeriodEndTime(newEpochId)).to.be.eq(expectedStartValidationTime);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(dkg, 'NewValidatorAdded').withArgs(FIRST.address, expectedStartValidationTime, newEpochId);
      await expect(tx).emit(dkg, 'NewEpochCreated').withArgs(newEpochId, newEpochStartTime);

      const valInfo = await dkg.getValidatorInfo(FIRST.address);

      expect(valInfo.validator).to.be.eq(FIRST.address);
      checkValidationData(valInfo.startValidationData, [expectedStartValidationTime, newEpochId]);
      checkValidationData(valInfo.endValidationData, [ethers.constants.MaxUint256, 0]);
    });

    it('should correctly add validator during DKG generation epoch with next epoch creation', async () => {
      const expectedDKGGenerationStartTime = startTime.add(defUpdateCollectionsEpochDuration);

      await setTime(expectedDKGGenerationStartTime.add('1').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(3);

      const tx = await dkg.connect(STAKING).addValidator(FIRST.address);

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = expectedDKGGenerationStartTime
        .add(defDKGGenerationEpochDuration)
        .add(defGuaranteedWorkingEpochDuration);

      expect(await dkg.getEpochEndTime(startEpochId)).to.be.eq(newEpochStartTime);

      expect(await dkg.isLastEpoch(newEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(newEpochId)).to.be.eq(false);

      const dkgEpochInfo = await dkg.getDKGEpochInfo(newEpochId);

      expect(dkgEpochInfo.mainEpochInfo.epochId).to.be.eq(newEpochId);
      expect(dkgEpochInfo.mainEpochInfo.epochStartTime).to.be.eq(newEpochStartTime);
      expect(dkgEpochInfo.epochStatus).to.be.eq(1);

      const expectedStartValidationTime = newEpochStartTime
        .add(defUpdateCollectionsEpochDuration)
        .add(defDKGGenerationEpochDuration);

      expect(await dkg.getDKGPeriodEndTime(newEpochId)).to.be.eq(expectedStartValidationTime);

      const valInfo = await dkg.getValidatorInfo(FIRST.address);

      expect(valInfo.validator).to.be.eq(FIRST.address);
      checkValidationData(valInfo.startValidationData, [expectedStartValidationTime, newEpochId]);
      checkValidationData(valInfo.endValidationData, [ethers.constants.MaxUint256, 0]);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(dkg, 'NewValidatorAdded').withArgs(FIRST.address, expectedStartValidationTime, newEpochId);
      await expect(tx).emit(dkg, 'NewEpochCreated').withArgs(newEpochId, newEpochStartTime);

      await setTime(newEpochStartTime.toNumber());
      expect(await dkg.getCurrentEpochId()).to.be.eq(newEpochId);
    });

    it('should correctly add validator during guaranteed working epoch with next epoch creation', async () => {
      const expectedGuaranteedWorkingStartTime = startTime
        .add(defUpdateCollectionsEpochDuration)
        .add(defDKGGenerationEpochDuration);

      await setTime(expectedGuaranteedWorkingStartTime.add('1').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(4);

      const tx = await dkg.connect(STAKING).addValidator(FIRST.address);

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = expectedGuaranteedWorkingStartTime.add(defGuaranteedWorkingEpochDuration);

      expect(await dkg.getEpochEndTime(startEpochId)).to.be.eq(newEpochStartTime);

      expect(await dkg.isLastEpoch(newEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(newEpochId)).to.be.eq(false);

      const dkgEpochInfo = await dkg.getDKGEpochInfo(newEpochId);

      expect(dkgEpochInfo.mainEpochInfo.epochId).to.be.eq(newEpochId);
      expect(dkgEpochInfo.mainEpochInfo.epochStartTime).to.be.eq(newEpochStartTime);
      expect(dkgEpochInfo.epochStatus).to.be.eq(1);

      const expectedStartValidationTime = newEpochStartTime
        .add(defUpdateCollectionsEpochDuration)
        .add(defDKGGenerationEpochDuration);

      expect(await dkg.getDKGPeriodEndTime(newEpochId)).to.be.eq(expectedStartValidationTime);

      const valInfo = await dkg.getValidatorInfo(FIRST.address);

      expect(valInfo.validator).to.be.eq(FIRST.address);
      checkValidationData(valInfo.startValidationData, [expectedStartValidationTime, newEpochId]);
      checkValidationData(valInfo.endValidationData, [ethers.constants.MaxUint256, 0]);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(dkg, 'NewValidatorAdded').withArgs(FIRST.address, expectedStartValidationTime, newEpochId);
      await expect(tx).emit(dkg, 'NewEpochCreated').withArgs(newEpochId, newEpochStartTime);

      await setTime(newEpochStartTime.toNumber());
      expect(await dkg.getCurrentEpochId()).to.be.eq(newEpochId);
    });

    it('should correctly add validator during DKG generation epoch without next epoch creation', async () => {
      const expectedDKGGenerationStartTime = startTime.add(defUpdateCollectionsEpochDuration);

      await setTime(expectedDKGGenerationStartTime.add('1').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(3);

      await dkg.connect(STAKING).addValidator(FIRST.address);

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = expectedDKGGenerationStartTime
        .add(defDKGGenerationEpochDuration)
        .add(defGuaranteedWorkingEpochDuration);

      expect(await dkg.getEpochEndTime(startEpochId)).to.be.eq(newEpochStartTime);

      expect(await dkg.isLastEpoch(newEpochId)).to.be.eq(true);
      expect(await dkg.isCurrentEpoch(newEpochId)).to.be.eq(false);

      const dkgEpochInfo = await dkg.getDKGEpochInfo(newEpochId);

      expect(dkgEpochInfo.mainEpochInfo.epochId).to.be.eq(newEpochId);
      expect(dkgEpochInfo.mainEpochInfo.epochStartTime).to.be.eq(newEpochStartTime);
      expect(dkgEpochInfo.epochStatus).to.be.eq(1);

      const tx = await dkg.connect(STAKING).addValidator(SECOND.address);

      expect(await dkg.isLastEpoch(newEpochId)).to.be.eq(true);

      const expectedStartValidationTime = newEpochStartTime
        .add(defUpdateCollectionsEpochDuration)
        .add(defDKGGenerationEpochDuration);

      expect(await dkg.getDKGPeriodEndTime(newEpochId)).to.be.eq(expectedStartValidationTime);

      const valInfo = await dkg.getValidatorInfo(SECOND.address);

      expect(valInfo.validator).to.be.eq(SECOND.address);
      checkValidationData(valInfo.startValidationData, [expectedStartValidationTime, newEpochId]);
      checkValidationData(valInfo.endValidationData, [ethers.constants.MaxUint256, 0]);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'NewValidatorAdded').withArgs(SECOND.address, expectedStartValidationTime, newEpochId);
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
    it('should correctly announce exit when the epoch status is ACTIVE with new epoch creation', async () => {
      const startValidationTime = startTime.add('10');

      await setNextBlockTime(startValidationTime.toNumber());
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const currentEpochEndTime = await dkg.getEpochEndTime(startEpochId);

      await setTime(currentEpochEndTime.add('10').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = currentEpochEndTime.add('100');

      const endValidationTime = await dkg.connect(STAKING).callStatic.announceValidatorExit(FIRST.address);

      expect(getEndDKGPeriodTime(currentEpochEndTime.add('10'))).to.be.eq(endValidationTime);

      await setNextBlockTime(newEpochStartTime.toNumber());

      const tx = await dkg.connect(STAKING).announceValidatorExit(FIRST.address);
      const expectedEndValidationTime = getEndDKGPeriodTime(newEpochStartTime);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx)
        .emit(dkg, 'ValidatorExitAnnounced')
        .withArgs(FIRST.address, expectedEndValidationTime, newEpochId);
      await expect(tx).emit(dkg, 'NewEpochCreated').withArgs(newEpochId, newEpochStartTime);

      const valInfo = await dkg.getValidatorInfo(FIRST.address);

      expect(valInfo.validator).to.be.eq(FIRST.address);
      checkValidationData(valInfo.startValidationData, [startValidationTime, startEpochId]);
      checkValidationData(valInfo.endValidationData, [expectedEndValidationTime, newEpochId]);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);

      await setTime(expectedEndValidationTime.toNumber());
      await dkg.setSigner(newEpochId, OWNER.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);
    });

    it('should correctly announce exit during DKG generation period with new epoch creation', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const startDKGGenerationTime = startTime.add(defUpdateCollectionsEpochDuration);

      await setTime(startDKGGenerationTime.add('1').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(3);

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = startDKGGenerationTime
        .add(defDKGGenerationEpochDuration)
        .add(defGuaranteedWorkingEpochDuration);

      const expectedEndValidationTime = getEndDKGPeriodTime(newEpochStartTime);
      const endValidationTime = await dkg.connect(STAKING).callStatic.announceValidatorExit(OWNER.address);

      expect(expectedEndValidationTime).to.be.eq(endValidationTime);

      const tx = await dkg.connect(STAKING).announceValidatorExit(OWNER.address);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx)
        .emit(dkg, 'ValidatorExitAnnounced')
        .withArgs(OWNER.address, expectedEndValidationTime, newEpochId);
      await expect(tx).emit(dkg, 'NewEpochCreated').withArgs(newEpochId, newEpochStartTime);

      const valInfo = await dkg.getValidatorInfo(OWNER.address);

      expect(valInfo.validator).to.be.eq(OWNER.address);
      checkValidationData(valInfo.startValidationData, [startTime, startEpochId]);
      checkValidationData(valInfo.endValidationData, [expectedEndValidationTime, newEpochId]);

      expect(await dkg.isActiveValidator(OWNER.address)).to.be.eq(true);
      expect(await dkg.isValidator(OWNER.address)).to.be.eq(true);

      await setTime(newEpochStartTime.add(defUpdateCollectionsEpochDuration).add('1').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(3);

      expect(await dkg.isActiveValidator(OWNER.address)).to.be.eq(true);
      expect(await dkg.isValidator(OWNER.address)).to.be.eq(true);

      await setTime(expectedEndValidationTime.add('1').toNumber());
      await dkg.setSigner(newEpochId, OWNER.address);

      expect(await dkg.isActiveValidator(OWNER.address)).to.be.eq(false);
      expect(await dkg.isValidator(OWNER.address)).to.be.eq(true);

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(4);
    });

    it('should correctly announce exit withot creation of new epoch', async () => {
      expect(await dkg.getCurrentEpochStatus()).to.be.eq(2);

      const expetcedStartValidationTime = startTime.add('10');

      await setNextBlockTime(expetcedStartValidationTime.toNumber());
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const firstEpochEndTime = getEpochEndTime(startTime);

      await setTime(firstEpochEndTime.add('10').toNumber());

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address]);
      expect(await dkg.getCurrentEpochStatus()).to.be.eq(5);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = firstEpochEndTime.add('100');

      await setNextBlockTime(secondEpochStartTime.toNumber());

      await dkg.connect(STAKING).addValidator(SECOND.address);

      expect(await dkg.isCurrentEpoch(secondEpochId)).to.be.eq(true);
      expect(await dkg.isLastEpoch(secondEpochId)).to.be.eq(true);

      const secondEpochStartDKGGenTime = secondEpochStartTime.add(defUpdateCollectionsEpochDuration);

      await setTime(secondEpochStartDKGGenTime.add('1').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(3);

      const thirdEpochId = secondEpochId.add('1');
      const thirdEpochStartTime = getEpochEndTime(secondEpochStartTime);

      await dkg.connect(STAKING).addValidator(THIRD.address);

      expect(await dkg.isCurrentEpoch(thirdEpochId)).to.be.eq(false);
      expect(await dkg.isLastEpoch(thirdEpochId)).to.be.eq(true);

      const tx = await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      const expectedEndValidationTime = getEndDKGPeriodTime(thirdEpochStartTime);

      const valInfo = await dkg.getValidatorInfo(FIRST.address);

      expect(valInfo.validator).to.be.eq(FIRST.address);
      checkValidationData(valInfo.startValidationData, [expetcedStartValidationTime, startEpochId]);
      checkValidationData(valInfo.endValidationData, [expectedEndValidationTime, thirdEpochId]);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx)
        .emit(dkg, 'ValidatorExitAnnounced')
        .withArgs(FIRST.address, expectedEndValidationTime, thirdEpochId);

      expect(await dkg.isLastEpoch(thirdEpochId)).to.be.eq(true);
    });

    it('should get exception if not the Staking address try to call this function', async () => {
      const reason = 'DKG: Not a staking address';

      await expect(dkg.announceValidatorExit(OWNER.address)).to.be.revertedWith(reason);
    });

    it('should get exception if validator is not active', async () => {
      const reason = 'DKG: Validator is not active';

      await expect(dkg.connect(STAKING).announceValidatorExit(FIRST.address)).to.be.revertedWith(reason);

      await setTime(getEpochEndTime(startTime).add('100').toNumber());

      await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);

      await expect(dkg.connect(STAKING).announceValidatorExit(FIRST.address)).to.be.revertedWith(reason);
    });

    it('should get exception if exit of the validator has already been announced', async () => {
      const reason = 'DKG: Exit of the validator has already been announced';

      await dkg.connect(STAKING).addValidator(FIRST.address);

      await setTime((await dkg.getDKGPeriodEndTime(startEpochId)).toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      await expect(dkg.connect(STAKING).announceValidatorExit(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('removeValidator', () => {
    it('should correctly remove validator', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const firstEpochEndTime = getEpochEndTime(startTime);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = firstEpochEndTime.add('10');

      await setNextBlockTime(secondEpochStartTime.toNumber());

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);

      const expectedEndValidationTime = getEndDKGPeriodTime(secondEpochStartTime);

      checkValidationData((await dkg.getValidatorInfo(FIRST.address)).endValidationData, [
        expectedEndValidationTime,
        secondEpochId,
      ]);

      await setTime(expectedEndValidationTime.add('10').toNumber());
      await dkg.setSigner(secondEpochId, OWNER.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(true);

      const tx = await dkg.connect(STAKING).removeValidator(FIRST.address);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'ValidatorRemoved').withArgs(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidator(FIRST.address)).to.be.eq(false);
    });

    it('should get exception if not the Staking address try to call this function', async () => {
      const reason = 'DKG: Not a staking address';

      await expect(dkg.removeValidator(OWNER.address)).to.be.revertedWith(reason);
    });

    it('should get exception if the validator does not exist', async () => {
      const reason = 'DKG: Validator does not exist';

      await expect(dkg.connect(STAKING).removeValidator(FIRST.address)).to.be.revertedWith(reason);
    });

    it('should get exception if try to remove validator before end validation time', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const firstEpochEndTime = getEpochEndTime(startTime);

      await setTime(firstEpochEndTime.add('10').toNumber());

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      const reason = "DKG: Validator can't be removed yet";

      await expect(dkg.connect(STAKING).removeValidator(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('createProposal', () => {
    it('should return correct info without new epoch creation', async () => {
      const mainEpochInfo = await dkg.connect(SLASHING_VOTING).callStatic.createProposal();

      expect(mainEpochInfo.epochId).to.be.eq(startEpochId);
      expect(mainEpochInfo.epochStartTime).to.be.eq(startTime);
      expect(mainEpochInfo.dkgGenPeriodEndTime).to.be.eq(getEndDKGPeriodTime(startTime));
    });

    it('should return correct info with new epoch creation', async () => {
      await setTime(getEndDKGPeriodTime(startTime).add('10').toNumber());

      const mainEpochInfo = await dkg.connect(SLASHING_VOTING).callStatic.createProposal();
      const tx = await dkg.connect(SLASHING_VOTING).createProposal();

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = getEpochEndTime(startTime);

      expect(mainEpochInfo.epochId).to.be.eq(newEpochId);
      expect(mainEpochInfo.epochStartTime).to.be.eq(newEpochStartTime);
      expect(mainEpochInfo.dkgGenPeriodEndTime).to.be.eq(getEndDKGPeriodTime(newEpochStartTime));

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'NewEpochCreated').withArgs(newEpochId, newEpochStartTime);
    });

    it('should get exception if not a slashing voting address tried to call this function', async () => {
      const reason = 'DKG: Not a slashing voting address';

      await expect(dkg.createProposal()).to.be.revertedWith(reason);
    });
  });

  describe('slashValidator', () => {
    it('should correctly slash validator', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      await dkg.setSigner(startEpochId, OWNER.address);
      await setTime(getEpochEndTime(startTime).toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isValidatorSlashed(FIRST.address)).to.be.eq(false);

      const tx = await dkg.connect(SLASHING_VOTING).slashValidator(FIRST.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
      expect(await dkg.isValidatorSlashed(FIRST.address)).to.be.eq(true);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'ValidatorSlashed').withArgs(FIRST.address);
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

  describe('updateAllValidators', async () => {
    it('should correctly update active validators', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);
      await dkg.connect(STAKING).addValidator(SECOND.address);

      expect(await dkg.getAllValidatorsCount()).to.be.eq('3');
      expect(await dkg.getAllValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);
      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);

      const secondEpochStartTime = getEpochEndTime(startTime).add('10');

      await setNextBlockTime(secondEpochStartTime.toNumber());

      await dkg.connect(STAKING).addValidator(THIRD.address);
      await dkg.connect(STAKING).announceValidatorExit(SECOND.address);

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);

      await dkg.setSigner(startEpochId, OWNER.address);

      const secondEpochId = startEpochId.add('1');
      const secondEpochDKGGenEndTime = getEndDKGPeriodTime(secondEpochStartTime);

      checkValidationData((await dkg.getValidatorInfo(THIRD.address)).startValidationData, [
        secondEpochDKGGenEndTime,
        secondEpochId,
      ]);
      checkValidationData((await dkg.getValidatorInfo(SECOND.address)).endValidationData, [
        secondEpochDKGGenEndTime,
        secondEpochId,
      ]);

      expect(await dkg.getAllValidatorsCount()).to.be.eq('4');
      expect(await dkg.getAllValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address, THIRD.address]);
      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);

      await dkg.updateAllValidators();

      expect(await dkg.getAllValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address, THIRD.address]);

      await setNextBlockTime(secondEpochDKGGenEndTime.add('10').toNumber());
      await dkg.setSigner(secondEpochId, OWNER.address);

      expect(await dkg.getActiveValidatorsCount()).to.be.eq('3');

      const tx = await dkg.updateAllValidators();

      expect(await dkg.getAllValidatorsCount()).to.be.eq('3');
      expect(await dkg.getActiveValidators()).to.be.deep.eq(await dkg.getAllValidators());
      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, THIRD.address]);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'ValidatorRemoved').withArgs(SECOND.address);
    });

    it('should correctly update start validation data for users', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);
      await dkg.connect(STAKING).addValidator(SECOND.address);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = getEpochEndTime(startTime).add('10');

      await setNextBlockTime(secondEpochStartTime.toNumber());

      await dkg.connect(STAKING).addValidator(THIRD.address);

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);

      let expectedStartValidationTime = getEndDKGPeriodTime(secondEpochStartTime);
      checkValidationData((await dkg.getValidatorInfo(THIRD.address)).startValidationData, [
        expectedStartValidationTime,
        secondEpochId,
      ]);

      await setTime(expectedStartValidationTime.add('10').toNumber());

      expect(await dkg.isActiveValidator(THIRD.address)).to.be.eq(false);

      const thirdEpochId = secondEpochId.add('1');
      const thirdEpochStartTime = getEpochEndTime(secondEpochStartTime);

      await setNextBlockTime(thirdEpochStartTime.toNumber());

      await dkg.updateAllValidators();

      expectedStartValidationTime = getEndDKGPeriodTime(thirdEpochStartTime);

      checkValidationData((await dkg.getValidatorInfo(THIRD.address)).startValidationData, [
        expectedStartValidationTime,
        thirdEpochId,
      ]);
    });

    it('should correctly update end validation data for users', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);
      await dkg.connect(STAKING).addValidator(SECOND.address);

      const secondEpochId = startEpochId.add('1');
      const secondEpochStartTime = getEpochEndTime(startTime).add('10');

      await setNextBlockTime(secondEpochStartTime.toNumber());

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      expect(await dkg.getActiveValidators()).to.be.deep.eq([OWNER.address, FIRST.address, SECOND.address]);

      let expectedEndValidationTime = getEndDKGPeriodTime(secondEpochStartTime);
      checkValidationData((await dkg.getValidatorInfo(FIRST.address)).endValidationData, [
        expectedEndValidationTime,
        secondEpochId,
      ]);

      await setTime(expectedEndValidationTime.add('10').toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);

      const thirdEpochId = secondEpochId.add('1');
      const thirdEpochStartTime = getEpochEndTime(secondEpochStartTime);

      await setNextBlockTime(thirdEpochStartTime.toNumber());

      await dkg.updateAllValidators();

      expectedEndValidationTime = getEndDKGPeriodTime(thirdEpochStartTime);

      checkValidationData((await dkg.getValidatorInfo(FIRST.address)).endValidationData, [
        expectedEndValidationTime,
        thirdEpochId,
      ]);
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

      dkgGenPeriodStartTime = startTime.add(defUpdateCollectionsEpochDuration);

      await setTime(dkgGenPeriodStartTime.add('10').toNumber());

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(3);
    });

    it('should correctly vote for new signer', async () => {
      let tx = await dkg.voteSigner(startEpochId, FIRST.address, firstSignature);

      expect((await tx.wait()).events?.length).to.be.eq(1);
      await expect(tx).emit(dkg, 'SignerVoted').withArgs(startEpochId, OWNER.address, FIRST.address);

      expect(await dkg.getSignerVotesCount(startEpochId, FIRST.address)).to.be.eq(1);
      expect(await dkg.getValidatorVote(startEpochId, OWNER.address)).to.be.eq(FIRST.address);
      expect(await dkg.getSignerAddress()).to.be.eq(ethers.constants.AddressZero);
      expect((await dkg.getDKGEpochInfo(startEpochId)).epochSigner).to.be.eq(ethers.constants.AddressZero);

      tx = await dkg.connect(FIRST).voteSigner(startEpochId, FIRST.address, firstSignature);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(dkg, 'SignerVoted').withArgs(startEpochId, FIRST.address, FIRST.address);
      await expect(tx).emit(dkg, 'SignerAddressUpdated').withArgs(startEpochId, FIRST.address);

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

      expect(await dkg.getCurrentEpochStatus()).to.be.eq(4);

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

      const firstEpochId = startEpochId.add('1');
      const firstEpochStartTime = getEpochEndTime(startTime).add('10');

      await setTime(firstEpochStartTime.toNumber());

      await dkg.connect(STAKING).addValidator(FIRST.address);

      expect(await dkg.getEpochStatus(startEpochId)).to.be.eq(6); // FINISHED
      expect(await dkg.getEpochStatus(firstEpochId)).to.be.eq(2); // UPDATES_COLLECTION

      const firstEpochDKGGenPeriodStartTime = firstEpochStartTime.add(defUpdateCollectionsEpochDuration);

      await setTime(firstEpochDKGGenPeriodStartTime.add('10').toNumber());

      expect(await dkg.getEpochStatus(firstEpochId)).to.be.eq(3); // DKG_GENERATION

      await dkg.connect(STAKING).addValidator(SECOND.address);

      const secondEpochId = firstEpochId.add('1');

      expect(await dkg.getLastEpochId()).to.be.eq(secondEpochId);
      expect(await dkg.getCurrentEpochId()).to.be.eq(firstEpochId);

      expect(await dkg.getEpochStatus(secondEpochId)).to.be.eq(1); // NOT_STARTED

      await setTime(firstEpochDKGGenPeriodStartTime.add(defGuaranteedWorkingEpochDuration).toNumber());

      expect(await dkg.getEpochStatus(firstEpochId)).to.be.eq(4); // GUARANTEED_WORKING

      await setTime((await dkg.getEpochEndTime(secondEpochId)).add('100').toNumber());

      expect(await dkg.getEpochStatus(secondEpochId)).to.be.eq(5); // GUARANTEED_WORKING
    });
  });

  describe('isActiveValidator', () => {
    it('should return correct result for the first epoch', async () => {
      const expectedValidationStartTime = startTime.add('10');

      await setNextBlockTime(expectedValidationStartTime.toNumber());
      await dkg.connect(STAKING).addValidator(FIRST.address);

      checkValidationData((await dkg.getValidatorInfo(FIRST.address)).startValidationData, [
        expectedValidationStartTime,
        startEpochId,
      ]);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
    });

    it('should return correct result if the validation start time has not come yet', async () => {
      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = getEpochEndTime(startTime).add('100');

      await setNextBlockTime(newEpochStartTime.toNumber());
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const expectedValidationStartTime = getEndDKGPeriodTime(newEpochStartTime);

      checkValidationData((await dkg.getValidatorInfo(FIRST.address)).startValidationData, [
        expectedValidationStartTime,
        newEpochId,
      ]);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
    });

    it('should return correct result if the dkg generetion period was not successful', async () => {
      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = getEpochEndTime(startTime).add('100');

      await setNextBlockTime(newEpochStartTime.toNumber());
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const expectedValidationStartTime = getEndDKGPeriodTime(newEpochStartTime);

      checkValidationData((await dkg.getValidatorInfo(FIRST.address)).startValidationData, [
        expectedValidationStartTime,
        newEpochId,
      ]);

      await setTime(expectedValidationStartTime.add('100').toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
    });

    it('should return correct result when user can unstake his tokens', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const expectedValidationStartTime = getEndDKGPeriodTime(startTime);

      await setTime(expectedValidationStartTime.add('100').toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      const newEpochId = startEpochId.add('1');
      const expectedEndValidationTime = getEndDKGPeriodTime(getEpochEndTime(startTime));

      await setTime(expectedEndValidationTime.add('100').toNumber());
      await dkg.setSigner(newEpochId, OWNER.address);

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(false);
    });

    it('should return correct result when user wants to withdraw but dkg generetion was not successful', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);

      const expectedValidationStartTime = getEndDKGPeriodTime(startTime);

      await setTime(expectedValidationStartTime.add('100').toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      const expectedEndValidationTime = getEndDKGPeriodTime(getEpochEndTime(startTime));

      await setTime(expectedEndValidationTime.add('100').toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
    });

    it('should return correct result for active validators', async () => {
      await dkg.connect(STAKING).addValidator(FIRST.address);
      await dkg.connect(STAKING).addValidator(SECOND.address);

      const expectedValidationStartTime = getEndDKGPeriodTime(startTime);

      await setTime(expectedValidationStartTime.add('100').toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
      expect(await dkg.isActiveValidator(SECOND.address)).to.be.eq(true);

      const newEpochId = startEpochId.add('1');
      const newEpochStartTime = getEpochEndTime(startTime);

      await dkg.connect(STAKING).announceValidatorExit(FIRST.address);

      const expectedEndValidationTime = getEndDKGPeriodTime(newEpochStartTime);

      checkValidationData((await dkg.getValidatorInfo(FIRST.address)).endValidationData, [
        expectedEndValidationTime,
        newEpochId,
      ]);

      await setTime(expectedEndValidationTime.add('100').toNumber());

      expect(await dkg.isActiveValidator(FIRST.address)).to.be.eq(true);
    });
  });
});
