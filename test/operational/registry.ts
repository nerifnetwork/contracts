import { anyValue } from '@nomicfoundation/hardhat-chai-matchers/withArgs';
import { expect } from 'chai';
import { ethers } from 'hardhat';
import { Wallet } from 'ethers';
import { Registry } from '../../typechain';

const tccABI = require('../../artifacts/contracts/test/TestCustomerContract.sol/TestCustomerContract.json');

describe('Registry', function () {
  async function createRandomWallet(fund: string = '10') {
    const [owner] = await ethers.getSigners();
    const wallet = await ethers.Wallet.createRandom();

    await expect(
      owner.sendTransaction({
        to: await wallet.getAddress(),
        value: ethers.utils.parseEther(fund),
        gasLimit: 500_000,
      })
    );

    return wallet;
  }

  async function deployRegistry(mainchain: boolean = true) {
    const [network] = await ethers.getSigners();

    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const signerStorage = await SignerStorageFactory.deploy();
    await expect(await signerStorage.initialize(await network.getAddress()));

    const Registry = await ethers.getContractFactory('Registry');
    const registry = await Registry.deploy();

    await expect(
      await registry.initialize(mainchain, signerStorage.address, ethers.constants.AddressZero, {
        performanceOverhead: 0,
        performancePremiumThreshold: 0,
        maxWorkflowsPerAccount: 0,
      })
    );

    return registry;
  }

  async function deployTestCustomerContract() {
    const TestCustomerContract = await ethers.getContractFactory('TestCustomerContract');
    return await TestCustomerContract.deploy();
  }

  async function fundBalance(registry: Registry, fundWallet: Wallet, amountRaw: string) {
    const nodeAddress = await fundWallet.getAddress();
    const amount = ethers.utils.parseEther(amountRaw);

    await expect(registry.connect(fundWallet.connect(ethers.provider)).fundBalance({ value: amount }))
      .to.emit(registry, 'BalanceFunded')
      .withArgs(nodeAddress, amount);

    await expect(await registry.getBalance(nodeAddress)).to.equal(amount);
  }

  describe('Deployment', function () {
    it('Successfully deployed on mainchain', async function () {
      const registry = await deployRegistry(true);
      expect(await registry.isMainChain()).to.equal(true);
    });

    it('Successfully deployed on sidechain', async function () {
      const registry = await deployRegistry(false);
      expect(await registry.isMainChain()).to.equal(false);
    });
  });

  describe('Fund balance', function () {
    async function fundBalanceOnChain(mainchain: boolean) {
      const registry = await deployRegistry(mainchain);
      const newNodeWallet = await createRandomWallet();
      await fundBalance(registry, newNodeWallet, '1');
    }

    it('MAINCHAIN: anyone can fund balance', async function () {
      await fundBalanceOnChain(true);
    });

    it('SIDECHAIN: anyone can fund balance', async function () {
      await fundBalanceOnChain(false);
    });
  });

  describe('Withdraw balance', function () {
    it('MAINCHAIN: anyone can withdraw balance', async function () {
      const registry = await deployRegistry(true);

      const newNodeWallet = await createRandomWallet();
      const newNodeWalletAddr = await newNodeWallet.getAddress();
      await fundBalance(registry, newNodeWallet, '1');

      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).withdrawBalance())
        .to.emit(registry, 'BalanceWithdrawn')
        .withArgs(newNodeWalletAddr, ethers.utils.parseEther('1'));
    });

    it('SIDECHAIN: anyone can withdraw balance', async function () {
      const registry = await deployRegistry(false);

      const newNodeWallet = await createRandomWallet();
      const newNodeWalletAddr = await newNodeWallet.getAddress();
      await fundBalance(registry, newNodeWallet, '1');

      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).withdrawBalance())
        .to.emit(registry, 'BalanceWithdrawn')
        .withArgs(newNodeWalletAddr, ethers.utils.parseEther('1'));
    });
  });

  describe('Register workflow', function () {
    it('MAINCHAIN: workflow owner can register workflow with gateway', async function () {
      const registry = await deployRegistry(true);

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();
      const newNodeProvider = newNodeWallet.connect(ethers.provider);

      // Register gateway
      const Gateway = await ethers.getContractFactory('Gateway');
      const gateway = await Gateway.connect(newNodeProvider).deploy();
      await expect(registry.connect(newNodeProvider).setGateway(gateway.address))
        .to.emit(registry, 'GatewaySet')
        .withArgs(nodeAddress, gateway.address);

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(123, nodeAddress, [], true)
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, '123', []);

      // Status must be PENDING on MAINCHAIN
      const workflow = await registry.getWorkflow(123);
      expect(workflow.status).to.equal(0);
    });

    it('MAINCHAIN: workflow owner can register workflow without gateway', async function () {
      const registry = await deployRegistry(true);

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(123, nodeAddress, [], false)
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, '123', []);

      // Status must be PENDING on MAINCHAIN
      const workflow = await registry.getWorkflow(123);
      expect(workflow.status).to.equal(0);
    });

    it('SIDECHAIN: network can register workflow', async function () {
      const registry = await deployRegistry(false);
      const [network] = await ethers.getSigners();

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();
      const newNodeProvider = newNodeWallet.connect(ethers.provider);

      // Register gateway
      const Gateway = await ethers.getContractFactory('Gateway');
      const gateway = await Gateway.connect(newNodeProvider).deploy();
      await expect(registry.connect(newNodeProvider).setGateway(gateway.address))
        .to.emit(registry, 'GatewaySet')
        .withArgs(nodeAddress, gateway.address);

      // Register workflow
      await expect(registry.connect(network).registerWorkflow(123, nodeAddress, [], true))
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(await network.getAddress(), '123', []);

      // Status must be ACTIVE on SIDECHAIN
      const workflow = await registry.getWorkflow(123);
      expect(workflow.status).to.equal(1);
    });
  });

  describe('Pause workflow', function () {
    it('MAINCHAIN: workflow owner can pause workflow', async function () {
      const registry = await deployRegistry(true);
      const [network] = await ethers.getSigners();
      const workflowID = 123;

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();
      const newNodeProvider = newNodeWallet.connect(ethers.provider);

      // Register gateway
      const Gateway = await ethers.getContractFactory('Gateway');
      const gateway = await Gateway.connect(newNodeProvider).deploy();
      await expect(registry.connect(newNodeProvider).setGateway(gateway.address))
        .to.emit(registry, 'GatewaySet')
        .withArgs(nodeAddress, gateway.address);

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(workflowID, nodeAddress, [], true)
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, workflowID, []);

      // Activate workflow
      await expect(registry.connect(network).activateWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 1);

      // Pause workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).pauseWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 2);
    });
  });

  describe('Resume workflow', function () {
    it('MAINCHAIN: workflow owner can resume workflow', async function () {
      const registry = await deployRegistry(true);
      const [network] = await ethers.getSigners();
      const workflowID = 123;

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();
      const newNodeProvider = newNodeWallet.connect(ethers.provider);

      // Register gateway
      const Gateway = await ethers.getContractFactory('Gateway');
      const gateway = await Gateway.connect(newNodeProvider).deploy();
      await expect(registry.connect(newNodeProvider).setGateway(gateway.address))
        .to.emit(registry, 'GatewaySet')
        .withArgs(nodeAddress, gateway.address);

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(workflowID, nodeAddress, [], true)
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, workflowID, []);

      // Activate workflow
      await expect(registry.connect(network).activateWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 1);

      // Pause workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).pauseWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 2);

      // Resume workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).resumeWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 1);
    });
  });

  describe('Cancel workflow', function () {
    it('MAINCHAIN: workflow owner can cancel workflow', async function () {
      const registry = await deployRegistry(true);
      const workflowID = 123;

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();
      const newNodeProvider = newNodeWallet.connect(ethers.provider);

      // Register gateway
      const Gateway = await ethers.getContractFactory('Gateway');
      const gateway = await Gateway.connect(newNodeProvider).deploy();
      await expect(registry.connect(newNodeProvider).setGateway(gateway.address))
        .to.emit(registry, 'GatewaySet')
        .withArgs(nodeAddress, gateway.address);

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(workflowID, nodeAddress, [], true)
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, workflowID, []);

      // Cancel workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).cancelWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 3);
    });

    it('SIDECHAIN: network can cancel workflow', async function () {
      const registry = await deployRegistry(false);
      const workflowID = 123;
      const [network] = await ethers.getSigners();

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();
      const newNodeProvider = newNodeWallet.connect(ethers.provider);

      // Register gateway
      const Gateway = await ethers.getContractFactory('Gateway');
      const gateway = await Gateway.connect(newNodeProvider).deploy();
      await expect(registry.connect(newNodeProvider).setGateway(gateway.address))
        .to.emit(registry, 'GatewaySet')
        .withArgs(nodeAddress, gateway.address);

      // Register workflow
      await expect(registry.connect(network).registerWorkflow(workflowID, nodeAddress, [], true))
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(await network.getAddress(), workflowID, []);

      // Cancel workflow
      await expect(registry.connect(network).cancelWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 3);
    });
  });

  describe('Perform', function () {
    it('Execute customer contract', async () => {
      // Deploy contracts
      const registry = await deployRegistry(true);
      const testCustomerContract = await deployTestCustomerContract();
      const workflowID = 123;

      // Create a workflow owner
      const [, workflowOwner] = await ethers.getSigners();
      const workflowOwnerAddress = await workflowOwner.getAddress();

      // Fund balance
      await expect(registry.connect(workflowOwner).fundBalance({ value: ethers.utils.parseEther('1') }))
        .to.emit(registry, 'BalanceFunded')
        .withArgs(workflowOwnerAddress, anyValue);

      // Register gateway
      const Gateway = await ethers.getContractFactory('Gateway');
      const gateway = await Gateway.connect(workflowOwner).deploy();
      await gateway.initialize(registry.address, workflowOwnerAddress);
      await expect(registry.connect(workflowOwner).setGateway(gateway.address))
        .to.emit(registry, 'GatewaySet')
        .withArgs(workflowOwnerAddress, gateway.address);

      // Register workflow
      await expect(registry.connect(workflowOwner).registerWorkflow(workflowID, workflowOwnerAddress, [], true))
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(workflowOwnerAddress, workflowID, []);

      // Activate workflow
      const [network] = await ethers.getSigners();
      await expect(registry.connect(network).activateWorkflow(workflowID))
        .to.emit(registry, 'WorkflowStatusChanged')
        .withArgs(workflowID, 1);

      // Create a contract ABI instance
      const cc = new ethers.Contract(testCustomerContract.address, tccABI.abi);

      const customerPerformData = ethers.utils.solidityPack(['uint256', 'uint', 'uint256'], [7, 0, 100000]);

      // Encode the function call with the selector
      const functionSignature = cc.interface.getSighash('perform');
      const functionCallData = cc.interface.encodeFunctionData(functionSignature, [customerPerformData]);

      await expect(
        workflowOwner.sendTransaction({
          to: testCustomerContract.address,
          data: functionCallData,
          gasLimit: 500_000,
        })
      ).to.emit(testCustomerContract, 'Logger');
      await expect(await testCustomerContract.counter(7)).to.equal(1);

      const tx = await registry
        .connect(network)
        .perform(workflowID, 300_000, functionCallData, testCustomerContract.address);
      // const receipt = await tx.wait();
      // console.log(BigNumber.from("12000000096").div(receipt.effectiveGasPrice))

      await expect(tx).to.emit(registry, 'Performance').withArgs(123, anyValue, true);
      await expect(await testCustomerContract.counter(7)).to.equal(2);
    });
  });
});
