import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers, network } from 'hardhat';
import { deployBridge, deployBridgeWithMocks } from '../utils/deploy';

describe('RelayBridge', function () {
  it('should send data', async function () {
    const [appContract] = await ethers.getSigners();

    const sourceChain = network.config.chainId ?? 31337;
    const destinationChain = 1;
    const gasLimit = 1;
    const nonce = 0;
    const sendAmount = utils.parseEther('1');

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforsend']);

    const { relayBridge, bridgeValidatorFeePool } = await deployBridge();

    const hash = await relayBridge.dataHash(appContract.address, sourceChain, destinationChain, gasLimit, data, nonce);
    await expect(relayBridge.send(destinationChain, gasLimit, data, { value: sendAmount }))
      .to.emit(relayBridge, 'Sent')
      .withArgs(hash, appContract.address, sourceChain, destinationChain, data, gasLimit, nonce, sendAmount);

    expect(await relayBridge.sentData(hash)).to.equals(data);
    expect(await ethers.provider.getBalance(bridgeValidatorFeePool)).to.equals(sendAmount);
  });

  it('should emit event Reverted in appContract', async function () {
    const destinationChain = 1;
    const gasLimit = 1;
    const nonce = 0;
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforsend']);

    const { relayBridge, mockBridgeApp } = await deployBridgeWithMocks();

    await expect(
      relayBridge.revertSend(mockBridgeApp.address, destinationChain, gasLimit, data, nonce, leader)
    ).to.be.revertedWith('RelayBridge: data never sent');
    await mockBridgeApp.send(destinationChain, gasLimit, data);

    const leaderBeforeRevert = await relayBridge.leaderHistoryLength();

    const hash = ethers.utils.keccak256(data);
    await expect(relayBridge.revertSend(mockBridgeApp.address, destinationChain, gasLimit, data, nonce, leader))
      .to.emit(mockBridgeApp, 'Reverted')
      .withArgs(hash);

    await expect(
      relayBridge.revertSend(mockBridgeApp.address, destinationChain, gasLimit, data, nonce, leader)
    ).to.be.revertedWith('RelayBridge: data already reverted');
    expect(await relayBridge.leaderHistory(0)).to.equal(leader);

    const leaderAfterRevert = await relayBridge.leaderHistoryLength();
    expect(leaderAfterRevert.sub(leaderBeforeRevert)).to.equal(1);
  });

  it('should execute data', async function () {
    const sourceChain = 1;
    const destinationChain = network.config.chainId ?? 31337;
    const gasLimit = 1;
    const nonce = 1;
    const leader = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforsend']);

    const { relayBridge, mockBridgeApp } = await deployBridgeWithMocks();

    const hash = await relayBridge.dataHash(
      mockBridgeApp.address,
      sourceChain,
      destinationChain,
      gasLimit,
      data,
      nonce
    );

    const leaderBeforeExecute = await relayBridge.leaderHistoryLength();

    await expect(relayBridge.execute(mockBridgeApp.address, sourceChain, gasLimit, data, nonce, leader))
      .to.emit(relayBridge, 'Executed')
      .withArgs(hash, sourceChain, destinationChain);

    expect(await relayBridge.leaderHistory(0)).to.equal(leader);

    const leaderAfterExecute = await relayBridge.leaderHistoryLength();
    expect(leaderAfterExecute.sub(leaderBeforeExecute)).to.equal(1);

    await expect(
      relayBridge.execute(mockBridgeApp.address, sourceChain, gasLimit, data, nonce, leader)
    ).to.be.revertedWith('RelayBridge: data already executed');

    expect(await relayBridge.executed(hash)).to.equals(true);
    expect(await mockBridgeApp.value()).to.equals('dataforsend');

    expect(await mockBridgeApp.bridgeAppAddress()).to.equals('0x0000000000000000000000000000000000000000');
  });

  it('should register failed events', async function () {
    const [appContract, user] = await ethers.getSigners();

    const sourceChain = network.config.chainId ?? 31337;
    const destinationChain = 1;
    const gasLimit = 1;
    const nonce = 0;
    const validatorFee = utils.parseEther('1');

    const abiCoder = ethers.utils.defaultAbiCoder;
    const data = abiCoder.encode(['string'], ['dataforsend']);

    const { relayBridge } = await deployBridge();

    const newRelayBridge = await ethers.getContractAt('RelayBridge', relayBridge.address, user);

    const hash = await relayBridge.dataHash(appContract.address, sourceChain, destinationChain, gasLimit, data, nonce);

    await expect(
      newRelayBridge.failedSend(appContract.address, sourceChain, destinationChain, data, gasLimit, nonce, validatorFee)
    ).to.be.revertedWith('SignerOwnable: only signer');
    await expect(
      relayBridge.failedSend(appContract.address, sourceChain, destinationChain, data, gasLimit, nonce, validatorFee)
    )
      .to.emit(relayBridge, 'FailedSend')
      .withArgs(hash, appContract.address, sourceChain, destinationChain, data, gasLimit, nonce, validatorFee);

    expect(await relayBridge.failed(hash)).to.equals(true);

    await expect(
      relayBridge.failedSend(appContract.address, sourceChain, destinationChain, data, gasLimit, nonce, validatorFee)
    ).to.be.revertedWith('RelayBridge: data already failed');
  });
});
