import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';
import { ethers } from 'hardhat';

describe('BridgeAppFactory', function () {
  it('should create bridge app', async function () {
    const [owner] = await ethers.getSigners();

    const { bridgeAppFactory } = await deploySystem();

    const appAddress = ethers.utils.getContractAddress({
      from: bridgeAppFactory.address,
      nonce: await ethers.provider.getTransactionCount(bridgeAppFactory.address),
    });

    await expect(bridgeAppFactory.createApp())
      .to.emit(bridgeAppFactory, 'BridgeAppCreated')
      .withArgs(appAddress, owner.address);
    expect(await bridgeAppFactory.apps(0)).to.equal(appAddress);
  });

  it('should set contract address', async function () {
    const [owner, user] = await ethers.getSigners();
    const chainId = 123;
    const contractAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { bridgeAppFactory } = await deploySystem();

    const appAddress = ethers.utils.getContractAddress({
      from: bridgeAppFactory.address,
      nonce: await ethers.provider.getTransactionCount(bridgeAppFactory.address),
    });

    await bridgeAppFactory.createApp();

    const BridgeAppOwner = await ethers.getContractAt('BridgeApp', appAddress, owner);
    await expect(BridgeAppOwner.setContractAddress(chainId, contractAddress))
      .to.emit(BridgeAppOwner, 'ContractAddressUpdated')
      .withArgs(chainId, contractAddress);

    expect(await BridgeAppOwner.contractAddresses(chainId)).to.equal(contractAddress);

    const BridgeAppUser = await ethers.getContractAt('BridgeApp', appAddress, user);
    await expect(BridgeAppUser.setContractAddress(chainId, contractAddress)).to.be.revertedWith(
      'Ownable: caller is not the owner'
    );
  });

  it('should set mediator', async function () {
    const [owner] = await ethers.getSigners();
    const mediatorAddress = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { bridgeAppFactory } = await deploySystem();

    const appAddress = ethers.utils.getContractAddress({
      from: bridgeAppFactory.address,
      nonce: await ethers.provider.getTransactionCount(bridgeAppFactory.address),
    });

    await bridgeAppFactory.createApp();

    const BridgeAppOwner = await ethers.getContractAt('BridgeApp', appAddress, owner);
    await expect(BridgeAppOwner.setMediator(mediatorAddress))
      .to.emit(BridgeAppOwner, 'UpdatedMediatorAddress')
      .withArgs(mediatorAddress);
  });
});
