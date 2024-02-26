import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { Gateway, GatewayFactory, TestGateway } from '../../generated-types/ethers';

describe('GatewayFactory', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let REGISTRY: SignerWithAddress;

  let gatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;

  before(async () => {
    [OWNER, FIRST, REGISTRY] = await ethers.getSigners();

    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');

    gatewayFactory = await GatewayFactoryFactory.deploy();
    gatewayImpl = await GatewayImplFactory.deploy();

    await gatewayFactory.initialize(REGISTRY.address, gatewayImpl.address);

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

      await expect(gatewayFactory.initialize(REGISTRY.address, gatewayImpl.address)).to.be.revertedWith(reason);
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
