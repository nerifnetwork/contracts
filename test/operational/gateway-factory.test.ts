import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry, Gateway, GatewayFactory, TestGateway } from '../../generated-types/ethers';

describe('GatewayFactory', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let REGISTRY: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;
  let gatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;

  before(async () => {
    [OWNER, FIRST, REGISTRY] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const gatewayFactoryImpl = await GatewayFactoryFactory.deploy();
    gatewayImpl = await GatewayImplFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(
      await contractsRegistry.GATEWAY_FACTORY_NAME(),
      gatewayFactoryImpl.address
    );
    await contractsRegistry.addContract(await contractsRegistry.REGISTRY_NAME(), REGISTRY.address);

    gatewayFactory = GatewayFactoryFactory.attach(await contractsRegistry.getGatewayFactoryContract());

    await gatewayFactory.initialize(gatewayImpl.address);

    await contractsRegistry.injectDependencies(await contractsRegistry.GATEWAY_FACTORY_NAME());

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialization', async () => {
      expect(await gatewayFactory.owner()).to.be.eq(OWNER.address);
      expect(await gatewayFactory.registryAddr()).to.be.eq(REGISTRY.address);
      expect(await gatewayFactory.getGatewayImplementation()).to.be.eq(gatewayImpl.address);
      expect(await gatewayFactory.gatewayBeacon()).to.be.not.eq(ethers.constants.AddressZero);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(gatewayFactory.initialize(gatewayImpl.address)).to.be.revertedWith(reason);
    });
  });

  describe('setDependencies', () => {
    it('should correctly update dependencies', async () => {
      expect(await gatewayFactory.registryAddr()).to.be.eq(REGISTRY.address);

      await contractsRegistry.addContract(await contractsRegistry.REGISTRY_NAME(), FIRST.address);
      await contractsRegistry.injectDependencies(await contractsRegistry.GATEWAY_FACTORY_NAME());

      expect(await gatewayFactory.registryAddr()).to.be.eq(FIRST.address);
    });

    it('should get exception if not a contracts registry try to call this function', async () => {
      const reason = 'Dependant: not an injector';

      await expect(gatewayFactory.setDependencies(contractsRegistry.address, '0x')).to.be.revertedWith(reason);
    });
  });

  describe('setNewImplementation', () => {
    it('should correctly set new Gateway implementation', async () => {
      const NewGatewayImplFactory = await ethers.getContractFactory('TestGateway');

      const newGatewayImpl = await NewGatewayImplFactory.deploy();

      const expectedGatewayAddr = await gatewayFactory.connect(REGISTRY).callStatic.deployGateway(FIRST.address);

      await gatewayFactory.connect(REGISTRY).deployGateway(FIRST.address);

      const gateway: TestGateway = NewGatewayImplFactory.attach(expectedGatewayAddr);

      await expect(gateway.version()).to.be.revertedWithoutReason();

      await gatewayFactory.setNewImplementation(newGatewayImpl.address);

      expect(await gatewayFactory.getGatewayImplementation()).to.be.eq(newGatewayImpl.address);

      expect(await gateway.version()).to.be.eq('2.0.0');

      await gatewayFactory.setNewImplementation(newGatewayImpl.address);

      expect(await gatewayFactory.getGatewayImplementation()).to.be.eq(newGatewayImpl.address);
    });

    it('should get exception if nonowner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(gatewayFactory.connect(FIRST).setNewImplementation(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('deployGateway', () => {
    it('should correctly deploy new Gateway contract', async () => {
      const expectedGatewayAddr = await gatewayFactory.connect(REGISTRY).callStatic.deployGateway(FIRST.address);

      const tx = await gatewayFactory.connect(REGISTRY).deployGateway(FIRST.address);

      await expect(tx).to.be.emit(gatewayFactory, 'GatewayDeployed').withArgs(expectedGatewayAddr, FIRST.address);

      const GatewayImplFactory = await ethers.getContractFactory('Gateway');

      const gateway: Gateway = GatewayImplFactory.attach(expectedGatewayAddr);

      expect(await gateway.owner()).to.be.eq(FIRST.address);
      expect(await gateway.registry()).to.be.eq(REGISTRY.address);
    });

    it('should get exception if not a registry try to call this function', async () => {
      const reason = 'GatewayFactory: operation is not permitted';

      await expect(gatewayFactory.deployGateway(OWNER.address)).to.be.revertedWith(reason);
    });
  });
});
