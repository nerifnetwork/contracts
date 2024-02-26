import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { wei } from '../helpers/utils';
import { Gateway, GatewayFactory, Registry, BillingManager, SignerStorage } from '../../generated-types/ethers';

describe('BillingManager', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let signerStorage: SignerStorage;
  let registry: Registry;
  let billingManager: BillingManager;
  let gatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;

  const defaultWorkflowId: number = 13;

  before(async () => {
    [OWNER, FIRST, SIGNER] = await ethers.getSigners();

    const RegistryFactory = await ethers.getContractFactory('Registry');
    const BillingManagerFactory = await ethers.getContractFactory('BillingManager');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');

    signerStorage = await SignerStorageFactory.deploy();
    registry = await RegistryFactory.deploy();
    billingManager = await BillingManagerFactory.deploy();

    gatewayFactory = await GatewayFactoryFactory.deploy();
    gatewayImpl = await GatewayImplFactory.deploy();

    await signerStorage.initialize(SIGNER.address);
    await registry.initialize(true, signerStorage.address, gatewayFactory.address, billingManager.address, 0);
    await billingManager.initialize(registry.address, signerStorage.address);
    await gatewayFactory.initialize(registry.address, gatewayImpl.address);

    await registry.registerWorkflows([
      {
        id: defaultWorkflowId,
        workflowOwner: OWNER.address,
        hash: '0x',
        requireGateway: true,
        deployGateway: true,
      },
    ]);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialize', async () => {
      expect(await billingManager.registry()).to.be.eq(registry.address);
      expect(await billingManager.signerGetter()).to.be.eq(signerStorage.address);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(billingManager.initialize(registry.address, signerStorage.address)).to.be.revertedWith(reason);
    });
  });

  describe('receive', () => {
    it('should fund balance after sending native currency directly to the contract', async () => {
      const fundBalance = wei('1');

      const tx = await OWNER.sendTransaction({
        from: OWNER.address,
        to: billingManager.address,
        value: fundBalance,
      });

      await expect(tx).emit(billingManager, 'BalanceFunded').withArgs(OWNER.address, fundBalance, fundBalance);
      await expect(tx).to.changeEtherBalances([OWNER, billingManager], [fundBalance.mul(-1), fundBalance]);
    });
  });

  describe('fundBalance', () => {
    it('should correctly fund balance for msg.sender', async () => {
      const fundBalance = wei('1');

      const tx = await billingManager['fundBalance()']({ value: fundBalance });

      await expect(tx).emit(billingManager, 'BalanceFunded').withArgs(OWNER.address, fundBalance, fundBalance);
      await expect(tx).to.changeEtherBalances([OWNER, billingManager], [fundBalance.mul(-1), fundBalance]);

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userFundBalance).to.be.eq(fundBalance);
    });

    it('should correctly fund balance for specific recipient', async () => {
      const fundBalance = wei('1');

      let tx = await billingManager['fundBalance(address)'](FIRST.address, { value: fundBalance });

      await expect(tx).emit(billingManager, 'BalanceFunded').withArgs(FIRST.address, fundBalance, fundBalance);

      tx = await billingManager['fundBalance(address)'](FIRST.address, { value: fundBalance });

      await expect(tx).emit(billingManager, 'BalanceFunded').withArgs(FIRST.address, fundBalance.mul(2), fundBalance);

      const userFundsInfo = await billingManager.getUserFundsInfo(FIRST.address);

      expect(userFundsInfo.userFundBalance).to.be.eq(fundBalance.mul(2));
    });

    it('should get exception if pass zero msg.value', async () => {
      const reason = 'BillingManager: Zero funds to add';

      await expect(billingManager['fundBalance()']()).to.be.revertedWith(reason);
    });
  });

  describe('lockExecutionFunds', () => {
    const fundBalance = wei('1');
    const lockAmount = wei('0.25');

    it('should correctly lock user funds', async () => {
      await billingManager['fundBalance()']({ value: fundBalance });

      const workflowExecutionId = await billingManager.nextWorkflowExecutionId();

      expect(await billingManager.getWorkflowExecutionStatus(workflowExecutionId)).to.be.eq(0);

      const tx = await billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount);

      await expect(tx)
        .emit(billingManager, 'ExecutionFundsLocked')
        .withArgs(defaultWorkflowId, OWNER.address, workflowExecutionId, lockAmount);

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userLockedBalance).to.be.eq(lockAmount);
      expect(userFundsInfo.pendingWorkflowExecutionIds).to.be.deep.eq([workflowExecutionId]);

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(workflowExecutionId);

      expect(workflowExecutionInfo.workflowId).to.be.eq(defaultWorkflowId);
      expect(workflowExecutionInfo.workflowOwner).to.be.eq(OWNER.address);
      expect(workflowExecutionInfo.executionLockedAmount).to.be.eq(lockAmount);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(0);
      expect(workflowExecutionInfo.status).to.be.eq(1);

      expect(await billingManager.getWorkflowExecutionStatus(workflowExecutionId)).to.be.eq(1);
      expect(await billingManager.getExecutionWorkflowId(workflowExecutionId)).to.be.eq(defaultWorkflowId);
      expect(await billingManager.getWorkflowExecutionOwner(workflowExecutionId)).to.be.eq(OWNER.address);
    });

    it('should correctly lock user funds several times', async () => {
      await billingManager['fundBalance()']({ value: fundBalance });

      const expectedWorkflowExecutionIds = [];

      expectedWorkflowExecutionIds.push(await billingManager.nextWorkflowExecutionId());
      await billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount);

      expectedWorkflowExecutionIds.push(await billingManager.nextWorkflowExecutionId());
      await billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount.mul(2));

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userLockedBalance).to.be.eq(lockAmount.mul(3));
      expect(userFundsInfo.pendingWorkflowExecutionIds).to.be.deep.eq(expectedWorkflowExecutionIds);
      expect(await billingManager.nextWorkflowExecutionId()).to.be.eq(2);

      let workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(expectedWorkflowExecutionIds[0]);

      expect(workflowExecutionInfo.workflowOwner).to.be.eq(OWNER.address);
      expect(workflowExecutionInfo.executionLockedAmount).to.be.eq(lockAmount);

      workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(expectedWorkflowExecutionIds[1]);

      expect(workflowExecutionInfo.workflowOwner).to.be.eq(OWNER.address);
      expect(workflowExecutionInfo.executionLockedAmount).to.be.eq(lockAmount.mul(2));
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(billingManager.lockExecutionFunds(10, lockAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if workflow id does not exist', async () => {
      const reason = 'BillingManager: Workflow does not exist';

      await expect(billingManager.connect(SIGNER).lockExecutionFunds(10, lockAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if user does not have available funds to lock', async () => {
      const reason = 'BillingManager: Not enough funds to lock';

      await expect(billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount)).to.be.revertedWith(
        reason
      );
    });
  });

  describe('completeExecution', () => {
    const fundBalance = wei('1');
    const lockAmount = wei('0.25');
    const executionAmount = wei('0.2');

    beforeEach('setup', async () => {
      await billingManager['fundBalance()']({ value: fundBalance });
      await billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount);
      await billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount.mul(2));
    });

    it('should correctly complete execution when lockedAmount > executionAmount', async () => {
      const tx = await billingManager.connect(SIGNER).completeExecution(0, executionAmount);

      await expect(tx).to.emit(billingManager, 'ExecutionCompleted').withArgs(0, executionAmount);
      expect((await tx.wait()).events?.length).to.be.eq(1);

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userFundBalance).to.be.eq(fundBalance.sub(executionAmount));
      expect(userFundsInfo.userLockedBalance).to.be.eq(lockAmount.mul(2));
      expect(userFundsInfo.pendingWorkflowExecutionIds).to.be.deep.eq([1]);

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(0);

      expect(workflowExecutionInfo.status).to.be.eq(2);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(executionAmount);

      const workflowInfo = await registry.getWorkflow(defaultWorkflowId);

      expect(workflowInfo.totalSpent).to.be.eq(executionAmount);

      expect(await billingManager.networkRewards()).to.be.eq(executionAmount);
    });

    it('should correctly complete execution when lockedAmount < execution amount and enough funds', async () => {
      const tx = await billingManager.connect(SIGNER).completeExecution(0, executionAmount.mul(2));

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(billingManager, 'ExecutionCompleted').withArgs(0, executionAmount.mul(2));
      await expect(tx)
        .emit(billingManager, 'UnexpectedExecutionAmountFound')
        .withArgs(0, lockAmount, executionAmount.mul(2));

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userFundBalance).to.be.eq(fundBalance.sub(executionAmount.mul(2)));
      expect(userFundsInfo.userLockedBalance).to.be.eq(lockAmount.mul(2));

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(0);

      expect(workflowExecutionInfo.status).to.be.eq(2);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(executionAmount.mul(2));
    });

    it('should correctly complete execution when lockedAmount < execution amount and not enough funds', async () => {
      const newExecutionAmount = wei('0.6');
      const expectedMaxExecutionAmount = wei('0.5');

      const tx = await billingManager.connect(SIGNER).completeExecution(0, newExecutionAmount);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(billingManager, 'ExecutionCompleted').withArgs(0, expectedMaxExecutionAmount);
      await expect(tx)
        .emit(billingManager, 'UnexpectedExecutionAmountFound')
        .withArgs(0, lockAmount, expectedMaxExecutionAmount);

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userFundBalance).to.be.eq(fundBalance.sub(expectedMaxExecutionAmount));
      expect(userFundsInfo.userLockedBalance).to.be.eq(lockAmount.mul(2));

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(0);

      expect(workflowExecutionInfo.status).to.be.eq(2);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(expectedMaxExecutionAmount);

      expect(await billingManager.networkRewards()).to.be.eq(expectedMaxExecutionAmount);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(billingManager.completeExecution(10, lockAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if status is not equal to PENDING', async () => {
      const reason = 'BillingManager: Not a pending workflow execution';

      await expect(
        billingManager.connect(SIGNER).completeExecution(defaultWorkflowId + 1, lockAmount)
      ).to.be.revertedWith(reason);
    });
  });

  describe('withdrawFunds', () => {
    const fundBalance = wei('1');
    const lockAmount = wei('0.5');
    const withdrawAmount = wei('0.3');

    beforeEach('setup', async () => {
      await billingManager['fundBalance()']({ value: fundBalance });
      await billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount);
    });

    it('should correctly withdraw specific amount', async () => {
      const tx = await billingManager.withdrawFunds(withdrawAmount);

      await expect(tx).emit(billingManager, 'UserFundsWithdrawn').withArgs(OWNER.address, withdrawAmount);
      await expect(tx).to.changeEtherBalances([OWNER, billingManager], [withdrawAmount, withdrawAmount.mul(-1)]);

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userFundBalance).to.be.eq(fundBalance.sub(withdrawAmount));
      expect(await billingManager.getUserAvailableFunds(OWNER.address)).to.be.eq(
        fundBalance.sub(lockAmount).sub(withdrawAmount)
      );
    });

    it('should correctly withdraw all available funds', async () => {
      const availableFunds = await billingManager.getUserAvailableFunds(OWNER.address);

      expect(availableFunds).to.be.eq(fundBalance.sub(lockAmount));

      const tx = await billingManager.withdrawFunds(availableFunds);

      await expect(tx).emit(billingManager, 'UserFundsWithdrawn').withArgs(OWNER.address, availableFunds);
      await expect(tx).to.changeEtherBalances([OWNER, billingManager], [availableFunds, availableFunds.mul(-1)]);

      const userFundsInfo = await billingManager.getUserFundsInfo(OWNER.address);

      expect(userFundsInfo.userFundBalance).to.be.eq(fundBalance.sub(availableFunds));
      expect(await billingManager.getUserAvailableFunds(OWNER.address)).to.be.eq('0');
    });

    it('should get exception if try to withdraw more than the available', async () => {
      const reason = 'BillingManager: Not enough available funds to withdraw';

      await expect(billingManager.withdrawFunds(withdrawAmount.mul(2))).to.be.revertedWith(reason);
    });
  });

  describe('withdrawNetworkRewards', () => {
    const fundBalance = wei('1');
    const lockAmount = wei('0.25');
    const executionAmount = wei('0.2');

    beforeEach('setup', async () => {
      await billingManager['fundBalance()']({ value: fundBalance });
      await billingManager.connect(SIGNER).lockExecutionFunds(defaultWorkflowId, lockAmount);
      await billingManager.connect(SIGNER).completeExecution(0, executionAmount);

      expect(await billingManager.networkRewards()).to.be.eq(executionAmount);
    });

    it('should correctly withdraw network rewards', async () => {
      const tx = await billingManager.withdrawNetworkRewards();

      await expect(tx).emit(billingManager, 'RewardsWithdrawn').withArgs(SIGNER.address, executionAmount);
      await expect(tx).to.changeEtherBalances([SIGNER, billingManager], [executionAmount, executionAmount.mul(-1)]);

      expect(await billingManager.networkRewards()).to.be.eq('0');
    });

    it('should get execption if nothing to withdraw', async () => {
      await billingManager.withdrawNetworkRewards();

      const reason = 'BillingManager: No network rewards to withdraw';

      await expect(billingManager.withdrawNetworkRewards()).to.be.revertedWith(reason);
    });

    it('should get exception if signer is zero addr', async () => {
      const reason = 'BillingManager: Zero signer address';

      await signerStorage.connect(SIGNER).setAddress(ethers.constants.AddressZero);

      await expect(billingManager.withdrawNetworkRewards()).to.be.revertedWith(reason);
    });
  });
});
