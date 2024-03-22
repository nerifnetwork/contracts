import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { wei } from '../helpers/utils';
import {
  Gateway,
  GatewayFactory,
  Registry,
  BillingManager,
  SignerStorage,
  TestTarget,
  IBillingManager,
  ContractsRegistry,
} from '../../generated-types/ethers';
import { BigNumberish } from 'ethers';

describe('Registry', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;

  let signerStorage: SignerStorage;
  let registry: Registry;
  let billingManager: BillingManager;
  let gatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;
  let testTarget: TestTarget;

  const tokensAmount = wei('10000');

  const nativeDepositAssetKey: string = 'NATIVE';
  const nativeDepositAssetData: IBillingManager.DepositAssetDataStruct = {
    tokenAddr: ethers.constants.AddressZero,
    workflowExecutionDiscount: 0,
    networkRewards: 0,
    isPermitable: false,
    isEnabled: true,
  };

  function getSetNumberCalldata(expectedNumber: BigNumberish): string {
    return testTarget.interface.encodeFunctionData('setNumber', [expectedNumber]);
  }

  before(async () => {
    [OWNER, FIRST, SIGNER] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const RegistryFactory = await ethers.getContractFactory('Registry');
    const BillingManagerFactory = await ethers.getContractFactory('BillingManager');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');
    const TokensVestingFactory = await ethers.getContractFactory('TokensVesting');
    const TestTargetFactory = await ethers.getContractFactory('TestTarget');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const registryImpl = await RegistryFactory.deploy();
    const billingManagerImpl = await BillingManagerFactory.deploy();
    const nerifTokenImpl = await NerifTokenFactory.deploy();

    const tokensVesting = await TokensVestingFactory.deploy();
    signerStorage = await SignerStorageFactory.deploy();
    gatewayFactory = await GatewayFactoryFactory.deploy();
    gatewayImpl = await GatewayImplFactory.deploy();

    testTarget = await TestTargetFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.REGISTRY_NAME(), registryImpl.address);
    await contractsRegistry.addProxyContract(
      await contractsRegistry.BILLING_MANAGER_NAME(),
      billingManagerImpl.address
    );
    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);

    await contractsRegistry.addContract(await contractsRegistry.GATEWAY_FACTORY_NAME(), gatewayFactory.address);
    await contractsRegistry.addContract(await contractsRegistry.TOKENS_VESTING_NAME(), tokensVesting.address);
    await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), signerStorage.address);

    registry = RegistryFactory.attach(await contractsRegistry.getRegistryContract());
    billingManager = BillingManagerFactory.attach(await contractsRegistry.getBillingManagerContract());
    const nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());

    await signerStorage.initialize(SIGNER.address);
    await registry.initialize(0);
    await billingManager.initialize({
      depositAssetKey: nativeDepositAssetKey,
      depositAssetData: nativeDepositAssetData,
    });
    await gatewayFactory.initialize(gatewayImpl.address);
    await nerifToken.initialize(tokensAmount, 'NERIF', 'NERIF');
    await tokensVesting.initialize();

    await contractsRegistry.injectDependencies(await contractsRegistry.REGISTRY_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.BILLING_MANAGER_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.GATEWAY_FACTORY_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());

    await contractsRegistry.setIsMainChain(true);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialize', async () => {
      expect(await registry.maxWorkflowsPerAccount()).to.be.eq(0);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(registry.initialize(0)).to.be.revertedWith(reason);
    });
  });

  describe('upgradability', () => {
    it('should correctly upgrade Registry contract', async () => {
      const TestRegistryFactory = await ethers.getContractFactory('TestRegistry');

      let testRegistry = TestRegistryFactory.attach(registry.address);

      await expect(testRegistry.version()).to.be.revertedWithoutReason();

      const newRegistryImpl = await TestRegistryFactory.deploy();

      await contractsRegistry.upgradeContract(await contractsRegistry.REGISTRY_NAME(), newRegistryImpl.address);

      expect(await testRegistry.version()).to.be.eq('v2.0.0');
    });
  });

  describe('setDependencies', () => {
    it('should correctly update dependencies', async () => {
      const TestRegistryFactory = await ethers.getContractFactory('TestRegistry');

      const newRegistryImpl = await TestRegistryFactory.deploy();
      await contractsRegistry.upgradeContract(await contractsRegistry.REGISTRY_NAME(), newRegistryImpl.address);

      const newRegistry = TestRegistryFactory.attach(registry.address);

      expect(await newRegistry.billingManagerAddress()).to.be.eq(billingManager.address);

      await contractsRegistry.addContract(await contractsRegistry.BILLING_MANAGER_NAME(), FIRST.address);
      await contractsRegistry.injectDependencies(await contractsRegistry.REGISTRY_NAME());

      expect(await newRegistry.billingManagerAddress()).to.be.eq(FIRST.address);
    });

    it('should get exception if not a contracts registry try to call this function', async () => {
      const reason = 'Dependant: not an injector';

      await expect(registry.setDependencies(contractsRegistry.address, '0x')).to.be.revertedWith(reason);
    });
  });

  describe('setMaxWorkflowsPerAccount', () => {
    it('should correctly set max workflows per account', async () => {
      const newMaxWorkflowsPerAccount = 5;

      await registry.connect(SIGNER).setMaxWorkflowsPerAccount(newMaxWorkflowsPerAccount);

      expect(await registry.maxWorkflowsPerAccount()).to.be.eq(newMaxWorkflowsPerAccount);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'Registry: Not a signer';

      await expect(registry.setMaxWorkflowsPerAccount(1)).to.be.revertedWith(reason);
    });
  });

  describe('setGateway', () => {
    it('should correctly set gateway contract', async () => {
      const tx = await registry.connect(FIRST).setGateway(OWNER.address);

      await expect(tx).to.emit(registry, 'GatewaySet').withArgs(FIRST.address, OWNER.address);

      expect(await registry.getGateway(FIRST.address)).to.be.eq(OWNER.address);
      expect(await registry.getTotalGatewaysCount()).to.be.eq(1);
    });
  });

  describe('deployAndSetGateway', () => {
    it('should correctly deploy and set gateway contract', async () => {
      const expectedGatewayAddress = await registry.connect(FIRST).callStatic.deployAndSetGateway();

      const tx = await registry.connect(FIRST).deployAndSetGateway();

      await expect(tx).to.emit(registry, 'GatewaySet').withArgs(FIRST.address, expectedGatewayAddress);

      expect(await registry.getGateway(FIRST.address)).to.be.eq(expectedGatewayAddress);
      expect(await registry.getTotalGatewaysCount()).to.be.eq(1);
    });
  });

  describe('updateWorkflowTotalSpent', () => {
    const workflowId = 10;
    const spentAmount = wei('0.1');

    let BILLING_MANAGER: SignerWithAddress;

    beforeEach('setup', async () => {
      BILLING_MANAGER = (await ethers.getSigners())[4];

      await contractsRegistry.addContract(await contractsRegistry.BILLING_MANAGER_NAME(), BILLING_MANAGER.address);

      await contractsRegistry.injectDependencies(await contractsRegistry.REGISTRY_NAME());

      await registry.registerWorkflows([
        {
          id: workflowId,
          workflowOwner: OWNER.address,
          hash: '0x',
          requireGateway: true,
          deployGateway: true,
        },
      ]);
    });

    it('should correctly update workflow total spent', async () => {
      await registry.connect(BILLING_MANAGER).updateWorkflowTotalSpent(nativeDepositAssetKey, workflowId, spentAmount);

      const workflowInfo = await registry.getWorkflowInfo(workflowId);

      expect(workflowInfo.depositAssetsInfo[0].depositAssetTotalSpent).to.be.eq(spentAmount);
    });

    it('should get exception if not a billing manager try to call this function', async () => {
      const reason = 'Registry: sender is not a billing manager';

      await expect(
        registry.updateWorkflowTotalSpent(nativeDepositAssetKey, workflowId, spentAmount)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if workflow id does not exist', async () => {
      const reason = 'Registry: workflow does not exist';

      await expect(
        registry.connect(BILLING_MANAGER).updateWorkflowTotalSpent(nativeDepositAssetKey, 0, spentAmount)
      ).to.be.revertedWith(reason);
    });
  });

  describe('perform', () => {
    const workflowId = 10;
    const workflowExecutionId = 0;
    const nativeDepositAmount = wei('1');
    const lockAmount = wei('0.5');
    const gasAmount = 1000000;

    beforeEach('setup', async () => {
      await registry.registerWorkflows([
        {
          id: workflowId,
          workflowOwner: OWNER.address,
          hash: '0x',
          requireGateway: true,
          deployGateway: true,
        },
      ]);

      await billingManager.deposit(nativeDepositAssetKey, OWNER.address, nativeDepositAmount, {
        value: nativeDepositAmount,
      });
      await billingManager.connect(SIGNER).lockExecutionFunds(nativeDepositAssetKey, workflowId, lockAmount);
    });

    it('should correctly perform function', async () => {
      const tx = await registry
        .connect(SIGNER)
        .perform(workflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address);

      await expect(tx).to.emit(registry, 'Performance').withArgs(workflowId, workflowExecutionId, true);
    });

    it('should get exception if not a signer call this function', async () => {
      const reason = 'Registry: Not a signer';

      await expect(
        registry.perform(workflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if workflow does not exist', async () => {
      const reason = 'Registry: workflow does not exist';

      await expect(
        registry
          .connect(SIGNER)
          .perform(workflowId + 1, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass invalid workflow execution ID', async () => {
      const reason = 'Registry: invalid workflow execution ID';

      const newWorkflowId = 20;
      const newWorkflowExecutionId = 1;

      await registry.connect(FIRST).registerWorkflows([
        {
          id: newWorkflowId,
          workflowOwner: FIRST.address,
          hash: '0x',
          requireGateway: true,
          deployGateway: true,
        },
      ]);

      await expect(
        registry
          .connect(SIGNER)
          .perform(newWorkflowId, newWorkflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);

      await billingManager
        .connect(FIRST)
        .deposit(nativeDepositAssetKey, FIRST.address, nativeDepositAmount, { value: nativeDepositAmount });
      await billingManager.connect(SIGNER).lockExecutionFunds(nativeDepositAssetKey, newWorkflowId, lockAmount);
      await billingManager.connect(SIGNER).completeExecution(newWorkflowExecutionId, lockAmount);

      await expect(
        registry
          .connect(SIGNER)
          .perform(newWorkflowId, newWorkflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);

      await expect(
        registry
          .connect(SIGNER)
          .perform(newWorkflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the target address equals to the registry address', async () => {
      const reason = 'Registry: operation is not permitted';

      await expect(
        registry
          .connect(SIGNER)
          .perform(workflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), registry.address)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the workflow owner does not have gateway contract', async () => {
      const reason = 'Registry: zero gateway address';

      await registry.setGateway(ethers.constants.AddressZero);

      await expect(
        registry
          .connect(SIGNER)
          .perform(workflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);
    });
  });

  describe('registerWorkflows', () => {
    const workflowId = 10;

    it('should correctly register new workflow on the main chain', async () => {
      const someHash = ethers.utils.solidityKeccak256(['string'], ['some string']);

      const tx = await registry.registerWorkflows([
        {
          id: workflowId,
          workflowOwner: OWNER.address,
          hash: someHash,
          requireGateway: true,
          deployGateway: true,
        },
      ]);

      await expect(tx).to.emit(registry, 'WorkflowRegistered').withArgs(OWNER.address, workflowId, someHash);

      const workflow = await registry.getWorkflowInfo(workflowId);

      expect(workflow.baseInfo.id).to.be.eq(workflowId);
      expect(workflow.baseInfo.owner).to.be.eq(OWNER.address);
      expect(await registry.workflowsPerAddress(OWNER.address)).to.be.eq(1);
    });

    it('should correctly register new workflow on the side chain', async () => {
      await contractsRegistry.setIsMainChain(false);

      const someHash = ethers.utils.solidityKeccak256(['string'], ['some string']);

      const tx = await registry.connect(SIGNER).registerWorkflows([
        {
          id: workflowId,
          workflowOwner: OWNER.address,
          hash: someHash,
          requireGateway: false,
          deployGateway: false,
        },
      ]);

      await expect(tx).to.emit(registry, 'WorkflowRegistered').withArgs(OWNER.address, workflowId, someHash);

      const workflow = await registry.getWorkflowInfo(workflowId);

      expect(workflow.baseInfo.id).to.be.eq(workflowId);
      expect(workflow.baseInfo.owner).to.be.eq(OWNER.address);
      expect(await registry.workflowsPerAddress(OWNER.address)).to.be.eq(1);

      await registry.connect(SIGNER).registerWorkflows([
        {
          id: workflowId + 1,
          workflowOwner: OWNER.address,
          hash: someHash,
          requireGateway: false,
          deployGateway: false,
        },
      ]);

      expect(await registry.workflowsPerAddress(OWNER.address)).to.be.eq(2);
    });

    it('should correctly register workflows and deploy new gateway contract', async () => {
      const expectedGatewayAddress = await registry.callStatic.deployAndSetGateway();

      expect(await registry.getGateway(OWNER.address)).to.be.eq(ethers.constants.AddressZero);

      const someHash = ethers.utils.solidityKeccak256(['string'], ['some string']);

      const tx = await registry.registerWorkflows([
        {
          id: workflowId,
          workflowOwner: OWNER.address,
          hash: someHash,
          requireGateway: true,
          deployGateway: true,
        },
        {
          id: workflowId + 1,
          workflowOwner: OWNER.address,
          hash: someHash,
          requireGateway: true,
          deployGateway: true,
        },
      ]);

      expect(await registry.getGateway(OWNER.address)).to.be.eq(expectedGatewayAddress);

      await expect(tx).to.emit(gatewayFactory, 'GatewayDeployed').withArgs(expectedGatewayAddress, OWNER.address);
      await expect(tx).to.emit(registry, 'WorkflowRegistered').withArgs(OWNER.address, workflowId, someHash);
      await expect(tx)
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(OWNER.address, workflowId + 1, someHash);
    });

    it('should get exception if the sender is not the specified addr or signer', async () => {
      await contractsRegistry.setIsMainChain(false);

      const reason = 'Registry: not a sender or signer';

      await expect(
        registry.registerWorkflows([
          {
            id: workflowId,
            workflowOwner: FIRST.address,
            hash: '0x',
            requireGateway: false,
            deployGateway: false,
          },
        ])
      ).to.be.revertedWith(reason);

      await expect(
        registry.registerWorkflows([
          {
            id: workflowId,
            workflowOwner: OWNER.address,
            hash: '0x',
            requireGateway: false,
            deployGateway: false,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the user has zero gateway address when required flag is true', async () => {
      const reason = 'Registry: gateway not found';

      await expect(
        registry.registerWorkflows([
          {
            id: workflowId,
            workflowOwner: OWNER.address,
            hash: '0x',
            requireGateway: true,
            deployGateway: false,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the max workflows number per address reached', async () => {
      const reason = 'Registry: reached max workflows capacity';

      await registry.connect(SIGNER).setMaxWorkflowsPerAccount(1);

      await registry.registerWorkflows([
        {
          id: workflowId,
          workflowOwner: OWNER.address,
          hash: '0x',
          requireGateway: true,
          deployGateway: true,
        },
      ]);

      await expect(
        registry.registerWorkflows([
          {
            id: workflowId + 1,
            workflowOwner: OWNER.address,
            hash: '0x',
            requireGateway: true,
            deployGateway: true,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass workflow id that already exists', async () => {
      const reason = 'Registry: workflow id is already exists';

      await expect(
        registry.registerWorkflows([
          {
            id: workflowId,
            workflowOwner: OWNER.address,
            hash: '0x',
            requireGateway: true,
            deployGateway: true,
          },
          {
            id: workflowId,
            workflowOwner: OWNER.address,
            hash: '0x',
            requireGateway: true,
            deployGateway: true,
          },
        ])
      ).to.be.revertedWith(reason);
    });
  });

  describe('getters', () => {
    const startWorkflowId = 10;

    let ownerGateway: string;
    let firstGateway: string;

    let ownerHash: string;
    let firstHash: string;

    beforeEach('setup', async () => {
      ownerGateway = await registry.callStatic.deployAndSetGateway();

      ownerHash = ethers.utils.solidityKeccak256(['string'], ['some string1']);
      firstHash = ethers.utils.solidityKeccak256(['string'], ['some string1']);

      await registry.registerWorkflows([
        {
          id: startWorkflowId,
          workflowOwner: OWNER.address,
          hash: ownerHash,
          requireGateway: true,
          deployGateway: true,
        },
        {
          id: startWorkflowId + 1,
          workflowOwner: OWNER.address,
          hash: ownerHash,
          requireGateway: true,
          deployGateway: true,
        },
      ]);

      firstGateway = await registry.callStatic.deployAndSetGateway();

      await registry.connect(FIRST).deployAndSetGateway();
      await registry.connect(FIRST).registerWorkflows([
        {
          id: startWorkflowId + 2,
          workflowOwner: FIRST.address,
          hash: firstHash,
          requireGateway: true,
          deployGateway: true,
        },
      ]);
    });

    it('should return correct info about gateways', async () => {
      await registry.connect(SIGNER).setGateway(SIGNER.address);

      expect(await registry.getTotalGatewaysCount()).to.be.eq(3);

      expect(await registry.getGateway(OWNER.address)).to.be.eq(ownerGateway);
      expect(await registry.getGateway(FIRST.address)).to.be.eq(firstGateway);
      expect(await registry.getGateway(SIGNER.address)).to.be.eq(SIGNER.address);

      const gatewaysInfo = await registry.getGatewaysInfo(0, 10);

      expect(gatewaysInfo.length).to.be.eq(3);

      const expectedGatewaysInfo = [
        {
          gatewayOwner: OWNER.address,
          gateway: ownerGateway,
        },
        {
          gatewayOwner: FIRST.address,
          gateway: firstGateway,
        },
        {
          gatewayOwner: SIGNER.address,
          gateway: SIGNER.address,
        },
      ];

      gatewaysInfo.forEach((el, index) => {
        expect(el.gatewayOwner).to.be.eq(expectedGatewaysInfo[index].gatewayOwner);
        expect(el.gateway).to.be.eq(expectedGatewaysInfo[index].gateway);
      });

      await registry.connect(SIGNER).setGateway(ethers.constants.AddressZero);

      expect(await registry.getTotalGatewaysCount()).to.be.eq(2);
      expect(await registry.getGateway(SIGNER.address)).to.be.eq(ethers.constants.AddressZero);
    });

    it('should return correct workflows info', async () => {
      const expectedWorkflowsInfo = [
        {
          workflowId: startWorkflowId,
          workflowOwner: OWNER.address,
          hash: ownerHash,
        },
        {
          workflowId: startWorkflowId + 1,
          workflowOwner: OWNER.address,
          hash: ownerHash,
        },
        {
          workflowId: startWorkflowId + 2,
          workflowOwner: FIRST.address,
          hash: firstHash,
        },
      ];

      expect(await registry.getTotalWorkflowsCount()).to.be.eq(expectedWorkflowsInfo.length);

      const workflowsInfo = await registry.getWorkflowsInfo(0, 10);

      workflowsInfo.forEach((el, index) => {
        expect(el.baseInfo.id).to.be.eq(expectedWorkflowsInfo[index].workflowId);
        expect(el.baseInfo.owner).to.be.eq(expectedWorkflowsInfo[index].workflowOwner);
        expect(el.baseInfo.hash).to.be.eq(expectedWorkflowsInfo[index].hash);
      });
    });
  });
});
