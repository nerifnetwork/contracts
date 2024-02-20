import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { wei } from '../helpers/utils';
import { Gateway, GatewayFactory, Registry, BillingManager, SignerStorage, TestTarget } from '../../typechain';
import { BigNumberish } from 'ethers';

describe('Registry', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let signerStorage: SignerStorage;
  let registry: Registry;
  let billingManager: BillingManager;
  let gatewayFactory: GatewayFactory;
  let sideChainRegistry: Registry;
  let sideChainBillingManager: BillingManager;
  let sideChainGatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;
  let testTarget: TestTarget;

  function getSetNumberCalldata(expectedNumber: BigNumberish): string {
    return testTarget.interface.encodeFunctionData('setNumber', [expectedNumber]);
  }

  before(async () => {
    [OWNER, FIRST, SIGNER] = await ethers.getSigners();

    const RegistryFactory = await ethers.getContractFactory('Registry');
    const BillingManagerFactory = await ethers.getContractFactory('BillingManager');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');
    const TestTargetFactory = await ethers.getContractFactory('TestTarget');

    signerStorage = await SignerStorageFactory.deploy();
    registry = await RegistryFactory.deploy();
    billingManager = await BillingManagerFactory.deploy();

    gatewayFactory = await GatewayFactoryFactory.deploy();
    gatewayImpl = await GatewayImplFactory.deploy();

    sideChainRegistry = await RegistryFactory.deploy();
    sideChainGatewayFactory = await GatewayFactoryFactory.deploy();
    sideChainBillingManager = await BillingManagerFactory.deploy();

    testTarget = await TestTargetFactory.deploy();

    await signerStorage.initialize(SIGNER.address);
    await registry.initialize(true, signerStorage.address, gatewayFactory.address, billingManager.address, {
      performanceOverhead: 0,
      maxWorkflowsPerAccount: 0,
    });
    await billingManager.initialize(registry.address, signerStorage.address);
    await gatewayFactory.initialize(registry.address, gatewayImpl.address);

    await sideChainRegistry.initialize(
      false,
      signerStorage.address,
      sideChainGatewayFactory.address,
      sideChainBillingManager.address,
      {
        performanceOverhead: 0,
        maxWorkflowsPerAccount: 0,
      }
    );
    await sideChainBillingManager.initialize(sideChainRegistry.address, signerStorage.address);
    await sideChainGatewayFactory.initialize(sideChainRegistry.address, gatewayImpl.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialize', async () => {
      expect(await registry.isMainChain()).to.be.eq(true);

      const config = await registry.config();

      expect(config.performanceOverhead).to.be.eq(0);
      expect(config.maxWorkflowsPerAccount).to.be.eq(0);

      expect(await registry.gatewayFactory()).to.be.eq(gatewayFactory.address);
      expect(await registry.billingManager()).to.be.eq(billingManager.address);
      expect(await registry.signerGetter()).to.be.eq(signerStorage.address);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(
        registry.initialize(true, signerStorage.address, gatewayFactory.address, billingManager.address, {
          performanceOverhead: 0,
          maxWorkflowsPerAccount: 0,
        })
      ).to.be.revertedWith(reason);
    });
  });

  describe('setConfig', () => {
    it('should correctly update config', async () => {
      const newConfig = {
        performanceOverhead: 5,
        maxWorkflowsPerAccount: 5,
      };

      await registry.connect(SIGNER).setConfig(newConfig);

      const config = await registry.config();

      expect(config.performanceOverhead).to.be.eq(newConfig.performanceOverhead);
      expect(config.maxWorkflowsPerAccount).to.be.eq(newConfig.maxWorkflowsPerAccount);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(
        registry.setConfig({
          performanceOverhead: 10,
          maxWorkflowsPerAccount: 1,
        })
      ).to.be.revertedWith(reason);
    });
  });

  describe('setGatewayFactory', () => {
    it('should correctly update gateway factory address', async () => {
      await registry.connect(SIGNER).setGatewayFactory(FIRST.address);

      expect(await registry.gatewayFactory()).to.be.eq(FIRST.address);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(registry.setGatewayFactory(OWNER.address)).to.be.revertedWith(reason);
    });
  });

  describe('setGateway', () => {
    it('should correctly set gateway contract', async () => {
      const tx = await registry.connect(FIRST).setGateway(OWNER.address);

      await expect(tx).to.emit(registry, 'GatewaySet').withArgs(FIRST.address, OWNER.address);

      expect(await registry.getGateway(FIRST.address)).to.be.eq(OWNER.address);
    });
  });

  describe('deployAndSetGateway', () => {
    it('should correctly deploy and set gateway contract', async () => {
      const expectedGatewayAddress = await registry.connect(FIRST).callStatic.deployAndSetGateway();

      const tx = await registry.connect(FIRST).deployAndSetGateway();

      await expect(tx).to.emit(registry, 'GatewaySet').withArgs(FIRST.address, expectedGatewayAddress);

      expect(await registry.getGateway(FIRST.address)).to.be.eq(expectedGatewayAddress);
    });
  });

  describe('updateWorkflowTotalSpent', () => {
    const workflowId = 10;
    const spentAmount = wei('0.1');

    let newRegistry: Registry;
    let BILLING_MANAGER: SignerWithAddress;

    beforeEach('setup', async () => {
      BILLING_MANAGER = (await ethers.getSigners())[4];

      const RegistryFactory = await ethers.getContractFactory('Registry');
      const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');

      newRegistry = await RegistryFactory.deploy();
      const newGatewayFactory = await GatewayFactoryFactory.deploy();

      await newRegistry.initialize(true, signerStorage.address, newGatewayFactory.address, BILLING_MANAGER.address, {
        performanceOverhead: 0,
        maxWorkflowsPerAccount: 0,
      });
      await newGatewayFactory.initialize(newRegistry.address, gatewayImpl.address);

      await newRegistry.deployAndSetGateway();
      await newRegistry.registerWorkflow(workflowId, OWNER.address, '0x', true);
    });

    it('should correctly update workflow total spent', async () => {
      await newRegistry.connect(BILLING_MANAGER).updateWorkflowTotalSpent(workflowId, spentAmount);

      const workflowInfo = await newRegistry.getWorkflow(workflowId);

      expect(workflowInfo.totalSpent).to.be.eq(spentAmount);
    });

    it('should get exception if not a billing manager try to call this function', async () => {
      const reason = 'Registry: sender is not a billing manager';

      await expect(newRegistry.updateWorkflowTotalSpent(workflowId, spentAmount)).to.be.revertedWith(reason);
    });
  });

  describe('pauseWorkflow', () => {
    const workflowId = 10;

    beforeEach('setup', async () => {
      await registry.deployAndSetGateway();
      await registry.registerWorkflow(workflowId, OWNER.address, '0x', true);
      await registry.connect(SIGNER).activateWorkflow(workflowId);
    });

    it('should correctly pause workflow', async () => {
      const tx = await registry.pauseWorkflow(workflowId);

      await expect(tx).to.emit(registry, 'WorkflowStatusChanged').withArgs(workflowId, 2);
      expect((await registry.getWorkflow(workflowId)).status).to.be.eq(2);
    });

    it('should get exception if try to call this function not on a main chain', async () => {
      const reason = 'Registry: operation is not permitted';

      await expect(sideChainRegistry.pauseWorkflow(workflowId)).to.be.revertedWith(reason);
    });

    it('should get exception if workflow id does not exist', async () => {
      const reason = 'Registry: workflow does not exist';

      await expect(registry.pauseWorkflow(0)).to.be.revertedWith(reason);
    });

    it('should get exception if the sender is not the workflow owner', async () => {
      const reason = 'Registry: operation not permitted';

      await expect(registry.connect(FIRST).pauseWorkflow(workflowId)).to.be.revertedWith(reason);
    });

    it('should get exception if the workflow status is not ACTIVE', async () => {
      const reason = 'Registry: only active workflows could be paused';

      await registry.pauseWorkflow(workflowId);

      await expect(registry.pauseWorkflow(workflowId)).to.be.revertedWith(reason);
    });
  });

  describe('resumeWorkflow', () => {
    const workflowId = 10;

    beforeEach('setup', async () => {
      await registry.deployAndSetGateway();
      await registry.registerWorkflow(workflowId, OWNER.address, '0x', true);
      await registry.connect(SIGNER).activateWorkflow(workflowId);
    });

    it('should correctly resume workflow', async () => {
      await registry.pauseWorkflow(workflowId);

      const tx = await registry.resumeWorkflow(workflowId);

      await expect(tx).to.emit(registry, 'WorkflowStatusChanged').withArgs(workflowId, 1);
      expect((await registry.getWorkflow(workflowId)).status).to.be.eq(1);
    });

    it('should get exception if try to call this function not on a main chain', async () => {
      const reason = 'Registry: operation is not permitted';

      await expect(sideChainRegistry.resumeWorkflow(workflowId)).to.be.revertedWith(reason);
    });

    it('should get exception if workflow id does not exist', async () => {
      const reason = 'Registry: workflow does not exist';

      await expect(registry.resumeWorkflow(0)).to.be.revertedWith(reason);
    });

    it('should get exception if the sender is not the workflow owner', async () => {
      const reason = 'Registry: operation not permitted';

      await expect(registry.connect(FIRST).resumeWorkflow(workflowId)).to.be.revertedWith(reason);
    });

    it('should get exception if the workflow status is not PAUSED', async () => {
      const reason = 'Registry: only paused workflows could be resumed';

      await expect(registry.resumeWorkflow(workflowId)).to.be.revertedWith(reason);
    });
  });

  describe('perform', () => {
    const workflowId = 10;
    const workflowExecutionId = 0;
    const fundBalance = wei('1');
    const lockAmount = wei('0.5');
    const gasAmount = 1000000;

    beforeEach('setup', async () => {
      await registry.deployAndSetGateway();
      await registry.registerWorkflow(workflowId, OWNER.address, '0x', true);
      await registry.connect(SIGNER).activateWorkflow(workflowId);

      await billingManager['fundBalance()']({ value: fundBalance });
      await billingManager.connect(SIGNER).lockExecutionFunds(workflowId, lockAmount);
    });

    it('should correctly perform function', async () => {
      const tx = await registry
        .connect(SIGNER)
        .perform(workflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address);

      await expect(tx).to.emit(registry, 'Performance').withArgs(workflowId, workflowExecutionId, true);
    });

    it('should get exception if not a signer call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(
        registry.perform(workflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if workflow status is not ACTIVE', async () => {
      const reason = 'Registry: workflow must be active';

      await registry.pauseWorkflow(workflowId);

      await expect(
        registry
          .connect(SIGNER)
          .perform(workflowId, workflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass invalid workflow execution ID', async () => {
      const reason = 'Registry: invalid workflow execution ID';

      const newWorkflowId = 20;
      const newWorkflowExecutionId = 1;

      await registry.connect(FIRST).deployAndSetGateway();
      await registry.connect(FIRST).registerWorkflow(newWorkflowId, FIRST.address, '0x', true);
      await registry.connect(SIGNER).activateWorkflow(newWorkflowId);

      await expect(
        registry
          .connect(SIGNER)
          .perform(newWorkflowId, newWorkflowExecutionId, gasAmount, getSetNumberCalldata(10), testTarget.address)
      ).to.be.revertedWith(reason);

      await billingManager.connect(FIRST)['fundBalance()']({ value: fundBalance });
      await billingManager.connect(SIGNER).lockExecutionFunds(newWorkflowId, lockAmount);
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

  describe('registerWorkflow', () => {
    const workflowId = 10;

    it('should correctly register new workflow on the main chain', async () => {
      await registry.deployAndSetGateway();

      const someHash = ethers.utils.solidityKeccak256(['string'], ['some string']);

      const tx = await registry.registerWorkflow(workflowId, OWNER.address, someHash, true);

      await expect(tx).to.emit(registry, 'WorkflowRegistered').withArgs(OWNER.address, workflowId, someHash);

      const workflow = await registry.getWorkflow(workflowId);

      expect(workflow.status).to.be.eq(0);
      expect(workflow.id).to.be.eq(workflowId);
      expect(workflow.owner).to.be.eq(OWNER.address);
    });

    it('should correctly register new workflow on the side chain', async () => {
      const someHash = ethers.utils.solidityKeccak256(['string'], ['some string']);

      const tx = await sideChainRegistry.connect(SIGNER).registerWorkflow(workflowId, OWNER.address, someHash, false);

      await expect(tx).to.emit(sideChainRegistry, 'WorkflowRegistered').withArgs(OWNER.address, workflowId, someHash);

      const workflow = await sideChainRegistry.getWorkflow(workflowId);

      expect(workflow.status).to.be.eq(1);
      expect(workflow.id).to.be.eq(workflowId);
      expect(workflow.owner).to.be.eq(OWNER.address);
    });

    it('should get exception if the sender is not the specified addr or signer', async () => {
      const reason = 'Registry: operation is not permitted';

      await expect(registry.registerWorkflow(workflowId, FIRST.address, '0x', false)).to.be.revertedWith(reason);

      await expect(sideChainRegistry.registerWorkflow(workflowId, OWNER.address, '0x', false)).to.be.revertedWith(
        reason
      );
    });

    it('should get exception if the user has zero gateway address when required flag is true', async () => {
      const reason = 'Registry: gateway not found';

      await expect(registry.registerWorkflow(workflowId, OWNER.address, '0x', true)).to.be.revertedWith(reason);
    });

    it('should get exception if the max workflows number per address reached', async () => {
      const reason = 'Registry: reached max workflows capacity';

      await registry.connect(SIGNER).setConfig({
        performanceOverhead: 0,
        maxWorkflowsPerAccount: 1,
      });

      await registry.deployAndSetGateway();
      await registry.registerWorkflow(workflowId, OWNER.address, '0x', true);

      await expect(registry.registerWorkflow(workflowId + 1, OWNER.address, '0x', true)).to.be.revertedWith(reason);
    });

    it('should get exception if pass workflow id that already exists', async () => {
      const reason = 'Registry: failed to add workflow';

      await registry.deployAndSetGateway();
      await registry.registerWorkflow(workflowId, OWNER.address, '0x', true);

      await expect(registry.registerWorkflow(workflowId, OWNER.address, '0x', true)).to.be.revertedWith(reason);
    });
  });

  describe('activateWorkflow', () => {
    const workflowId = 10;

    beforeEach('setup', async () => {
      await registry.deployAndSetGateway();
      await registry.registerWorkflow(workflowId, OWNER.address, '0x', true);

      await sideChainRegistry.connect(FIRST).deployAndSetGateway();
      await sideChainRegistry.connect(SIGNER).registerWorkflow(workflowId + 1, FIRST.address, '0x', true);
    });

    it('should correctly activate workflow', async () => {
      const tx = await registry.connect(SIGNER).activateWorkflow(workflowId);

      await expect(tx).to.emit(registry, 'WorkflowStatusChanged').withArgs(workflowId, 1);

      const workflow = await registry.getWorkflow(workflowId);

      expect(workflow.status).to.be.eq(1);
      expect(workflow.id).to.be.eq(workflowId);
    });

    it('should get exception if try to call this function not on the mainchain', async () => {
      const reason = 'Registry: operation is not permitted';

      await expect(sideChainRegistry.connect(SIGNER).activateWorkflow(workflowId)).to.be.revertedWith(reason);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(registry.activateWorkflow(workflowId)).to.be.revertedWith(reason);
    });

    it('should get exception if workflow if the does not exist', async () => {
      const reason = 'Registry: workflow does not exist';

      await expect(registry.connect(SIGNER).activateWorkflow(workflowId + 1)).to.be.revertedWith(reason);
    });

    it('should get exception if the workflow status is not PENDING', async () => {
      const reason = 'Registry: workflow must be pending';

      await registry.connect(SIGNER).activateWorkflow(workflowId);

      await expect(registry.connect(SIGNER).activateWorkflow(workflowId)).to.be.revertedWith(reason);
    });
  });

  describe('cancelWorkflow', () => {
    const workflowId = 10;

    beforeEach('setup', async () => {
      await registry.deployAndSetGateway();
      await registry.registerWorkflow(workflowId, OWNER.address, '0x', true);

      await sideChainRegistry.connect(FIRST).deployAndSetGateway();
      await sideChainRegistry.connect(SIGNER).registerWorkflow(workflowId + 1, FIRST.address, '0x', true);
    });

    it('should correctly cancel workflow', async () => {
      let tx = await registry.cancelWorkflow(workflowId);

      await expect(tx).to.emit(registry, 'WorkflowStatusChanged').withArgs(workflowId, 3);

      let workflow = await registry.getWorkflow(workflowId);

      expect(workflow.status).to.be.eq(3);
      expect(workflow.id).to.be.eq(workflowId);

      tx = await sideChainRegistry.connect(SIGNER).cancelWorkflow(workflowId + 1);

      await expect(tx)
        .to.emit(sideChainRegistry, 'WorkflowStatusChanged')
        .withArgs(workflowId + 1, 3);

      workflow = await sideChainRegistry.getWorkflow(workflowId + 1);

      expect(workflow.status).to.be.eq(3);
      expect(workflow.id).to.be.eq(workflowId + 1);
    });

    it('should get exception if workflow id does not exist', async () => {
      const reason = 'Registry: workflow does not exist';

      await expect(registry.cancelWorkflow(workflowId + 1)).to.be.revertedWith(reason);
    });

    it('should get exception if the sender is not the workflow owner or signer', async () => {
      const reason = 'Registry: operation is not permitted';

      await expect(registry.connect(FIRST).cancelWorkflow(workflowId)).to.be.revertedWith(reason);

      await expect(sideChainRegistry.connect(FIRST).cancelWorkflow(workflowId + 1)).to.be.revertedWith(reason);
    });

    it('should get exception if the workflow status is CANCELED', async () => {
      const reason = 'Registry: workflow is already cancelled';

      await registry.cancelWorkflow(workflowId);

      await expect(registry.cancelWorkflow(workflowId)).to.be.revertedWith(reason);
    });
  });
});
