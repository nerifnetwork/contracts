import { anyValue } from '@nomicfoundation/hardhat-chai-matchers/withArgs';
import { expect } from 'chai';
import { ethers } from 'hardhat';
import { Signer, Wallet } from 'ethers';
import { Registry } from '../../typechain';

const tccABI = require('../../artifacts/contracts/test/TestCustomerContract.sol/TestCustomerContract.json');
const registryContractABI = require('../../artifacts/contracts/common/Registry.sol/Registry.json');

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

    const MockNetworkAddress = await ethers.getContractFactory('MockNetworkAddress');
    const mockNetworkAddress = await MockNetworkAddress.deploy(await network.getAddress());

    const Registry = await ethers.getContractFactory('Registry');
    const registry = await Registry.deploy(mockNetworkAddress.address, mainchain);

    await expect(
      registry.setConfig({
        performanceOverhead: 0,
        performancePremiumThreshold: 0,
        registrationOverhead: 0,
        cancellationOverhead: 0,
      })
    );

    return registry;
  }

  async function deployTestCustomerContract() {
    const TestCustomerContract = await ethers.getContractFactory('TestCustomerContract');
    const testCustomerContract = await TestCustomerContract.deploy();

    return testCustomerContract;
  }

  async function fundBalance(registry: Registry, fundWallet: Wallet, amountRaw: string) {
    const nodeAddress = await fundWallet.getAddress();
    const amount = ethers.utils.parseEther(amountRaw);

    await expect(registry.connect(fundWallet.connect(ethers.provider)).fundBalance(nodeAddress, { value: amount }))
      .to.emit(registry, 'BalanceFunded')
      .withArgs(nodeAddress, amount);

    await expect(await registry.getBalance(nodeAddress)).to.equal(amount);
  }

  async function activateWorkflow(registry: Registry, signer: Signer, workflowID: number) {
    // Create a contract ABI instance
    const cc = new ethers.Contract(registry.address, registryContractABI.abi, signer);

    // Encode the function call with the selector
    const functionCallData = cc.interface.encodeFunctionData(cc.interface.getSighash('activateWorkflow'), [
      ethers.utils.solidityPack(['uint256'], [workflowID]),
    ]);

    // Activate the workflow
    await expect(
      registry
        .connect(signer)
        .perform('40505927788353901442144037336646356013', 300_000, functionCallData, registry.address)
    )
      .to.emit(registry, 'Performance')
      .withArgs('40505927788353901442144037336646356013', anyValue, true);
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

      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).withdrawBalance(newNodeWalletAddr))
        .to.emit(registry, 'BalanceWithdrawn')
        .withArgs(newNodeWalletAddr, ethers.utils.parseEther('1'));
    });

    it('SIDECHAIN: anyone can withdraw balance', async function () {
      const registry = await deployRegistry(false);

      const newNodeWallet = await createRandomWallet();
      const newNodeWalletAddr = await newNodeWallet.getAddress();
      await fundBalance(registry, newNodeWallet, '1');

      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).withdrawBalance(newNodeWalletAddr))
        .to.emit(registry, 'BalanceWithdrawn')
        .withArgs(newNodeWalletAddr, ethers.utils.parseEther('1'));
    });
  });

  describe('Register workflow', function () {
    it('MAINCHAIN: workflow owner can register workflow', async function () {
      const registry = await deployRegistry(true);

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();

      // Register workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(123, nodeAddress, []))
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

      // Create a contract ABI instance
      const cc = new ethers.Contract(registry.address, registryContractABI.abi);

      // Encode the function call with the selector
      const functionSignature = cc.interface.getSighash('registerWorkflow');
      const functionCallData = cc.interface.encodeFunctionData(functionSignature, [
        ethers.utils.solidityPack(['uint256'], [123]),
        ethers.utils.solidityPack(['address'], [nodeAddress]),
        ethers.utils.solidityPack(['bytes'], ['0x00']),
      ]);

      console.log(await registry.networkAddress());

      await expect(
        registry
          .connect(network)
          .perform('40505927788353901442144037336646356013', 300_000, functionCallData, registry.address)
      )
        .to.emit(registry, 'Performance')
        .withArgs('40505927788353901442144037336646356013', anyValue, true);

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

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(workflowID, nodeAddress, [])
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, workflowID, []);

      // Activate workflow as a network
      await activateWorkflow(registry, network, 123);

      // Pause workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).pauseWorkflow(workflowID))
        .to.emit(registry, 'WorkflowPaused')
        .withArgs(workflowID);
    });
  });

  describe('Resume workflow', function () {
    it('MAINCHAIN: workflow owner can resume workflow', async function () {
      const registry = await deployRegistry(true);
      const [network] = await ethers.getSigners();
      const workflowID = 123;

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(workflowID, nodeAddress, [])
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, workflowID, []);

      // Activate workflow
      await activateWorkflow(registry, network, workflowID);

      // Pause workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).pauseWorkflow(workflowID))
        .to.emit(registry, 'WorkflowPaused')
        .withArgs(workflowID);

      // Resume workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).resumeWorkflow(workflowID))
        .to.emit(registry, 'WorkflowResumed')
        .withArgs(workflowID);
    });
  });

  describe('Cancel workflow', function () {
    it('MAINCHAIN: workflow owner can cancel workflow', async function () {
      const registry = await deployRegistry(true);
      const workflowID = 123;

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();

      // Register workflow
      await expect(
        registry.connect(newNodeWallet.connect(ethers.provider)).registerWorkflow(workflowID, nodeAddress, [])
      )
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(nodeAddress, workflowID, []);

      // Cancel workflow
      await expect(registry.connect(newNodeWallet.connect(ethers.provider)).cancelWorkflow(workflowID))
        .to.emit(registry, 'WorkflowCancelled')
        .withArgs(workflowID);
    });

    it('SIDECHAIN: network can cancel workflow', async function () {
      const registry = await deployRegistry(false);
      const [network] = await ethers.getSigners();

      const newNodeWallet = await createRandomWallet();
      const nodeAddress = await newNodeWallet.getAddress();

      // Create a contract ABI instance
      const cc = new ethers.Contract(registry.address, registryContractABI.abi);

      // Encode the function call with the selector
      let functionCallData = cc.interface.encodeFunctionData(cc.interface.getSighash('registerWorkflow'), [
        ethers.utils.solidityPack(['uint256'], [123]),
        ethers.utils.solidityPack(['address'], [nodeAddress]),
        ethers.utils.solidityPack(['bytes'], ['0x00']),
      ]);

      // Register workflow
      await expect(
        registry
          .connect(network)
          .perform('40505927788353901442144037336646356013', 300_000, functionCallData, registry.address)
      )
        .to.emit(registry, 'Performance')
        .withArgs('40505927788353901442144037336646356013', anyValue, true);

      // Encode the function call with the selector
      functionCallData = cc.interface.encodeFunctionData(cc.interface.getSighash('cancelWorkflow'), [
        ethers.utils.solidityPack(['uint256'], [123]),
      ]);

      // Cancel the workflow
      await expect(
        registry
          .connect(network)
          .perform('219775546284901721155783592958414245131', 300_000, functionCallData, registry.address)
      )
        .to.emit(registry, 'Performance')
        .withArgs('219775546284901721155783592958414245131', anyValue, true);
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
      await expect(
        registry.connect(workflowOwner).fundBalance(workflowOwnerAddress, { value: ethers.utils.parseEther('1') })
      )
        .to.emit(registry, 'BalanceFunded')
        .withArgs(workflowOwnerAddress, anyValue);

      // Register workflow
      await expect(registry.connect(workflowOwner).registerWorkflow(workflowID, workflowOwnerAddress, []))
        .to.emit(registry, 'WorkflowRegistered')
        .withArgs(workflowOwnerAddress, workflowID, []);

      // Activate workflow
      const [network] = await ethers.getSigners();
      await activateWorkflow(registry, network, workflowID);

      // Create a contract ABI instance
      const cc = new ethers.Contract(testCustomerContract.address, tccABI.abi);

      const customerPerformData = ethers.utils.solidityPack(['uint256', 'uint', 'uint256'], [0, 0, 100000]);

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

      const tx = await registry
        .connect(network)
        .perform(workflowID, 300_000, functionCallData, testCustomerContract.address);
      // const receipt = await tx.wait();
      // console.log(BigNumber.from("12000000096").div(receipt.effectiveGasPrice))

      await expect(tx).to.emit(registry, 'Performance').withArgs(123, anyValue, true);
    });
  });
});
