import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry } from '../../generated-types/ethers';

describe('ContractsRegistry', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;

  before(async () => {
    [OWNER, FIRST, SIGNER] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialization', async () => {
      expect(await contractsRegistry.owner()).to.be.eq(OWNER.address);
      expect(await contractsRegistry.getProxyUpgrader()).to.be.not.eq(ethers.constants.AddressZero);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(contractsRegistry.__OwnableContractsRegistry_init()).to.be.revertedWith(reason);
    });
  });

  describe('upgradability', () => {
    it('should correctly upgrade contracts registry', async () => {
      const TestContractsRegistryFactory = await ethers.getContractFactory('TestContractsRegistry');

      let testContractsRegistry = TestContractsRegistryFactory.attach(contractsRegistry.address);

      await expect(testContractsRegistry.version()).to.be.revertedWithoutReason();

      const newContractsRegistryImpl = await TestContractsRegistryFactory.deploy();

      await contractsRegistry.upgradeTo(newContractsRegistryImpl.address);

      expect(await testContractsRegistry.version()).to.be.eq('v2.0.0');
    });

    it('should get exception if not an owner try to call upgradeTo function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(contractsRegistry.connect(FIRST).upgradeTo(ethers.constants.AddressZero)).to.be.revertedWith(reason);
    });
  });

  describe('setIsMainChain', () => {
    it('should correctly set main chain flag', async () => {
      await contractsRegistry.setIsMainChain(true);

      expect(await contractsRegistry.isMainChain()).to.be.eq(true);
    });

    it('should get exception if not an owner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(contractsRegistry.connect(FIRST).setIsMainChain(false)).to.be.revertedWith(reason);
    });
  });

  describe('getSigner', () => {
    it('should return correct signer address', async () => {
      expect(await contractsRegistry.getSigner()).to.be.eq(ethers.constants.AddressZero);

      const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');

      const signerStorage = await SignerStorageFactory.deploy();

      await signerStorage.initialize(SIGNER.address);

      await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), signerStorage.address);

      expect(await contractsRegistry.getSigner()).to.be.eq(SIGNER.address);
    });
  });
});
