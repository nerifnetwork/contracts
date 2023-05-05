import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('GatewayStorage', function () {
  it('should  clear all gateways', async function () {
    const [network] = await ethers.getSigners();
    const gateways = [
      {
        gateway: '0x0000000000000000000000000000000000000000',
        owner: '0x0000000000000000000000000000000000000000',
      },
      {
        gateway: '0x0000000000000000000000000000000000000001',
        owner: '0x0000000000000000000000000000000000000001',
      },
    ];

    const GatewayStorage = await ethers.getContractFactory('GatewayStorage');
    const gatewayStorage = await GatewayStorage.deploy();
    await gatewayStorage.initialize(network.address, gateways);

    expect(await gatewayStorage.size()).to.equal(2);
    await gatewayStorage.clear();
    expect(await gatewayStorage.size()).to.equal(0);
  });

  it('should not remove a non existing gateway', async function () {
    const [network] = await ethers.getSigners();
    const gateways = [
      {
        gateway: '0x0000000000000000000000000000000000000000',
        owner: '0x0000000000000000000000000000000000000000',
      },
      {
        gateway: '0x0000000000000000000000000000000000000001',
        owner: '0x0000000000000000000000000000000000000001',
      },
    ];

    const nonExistingGateway = '0x0000000000000000000000000000000000000002';

    const GatewayStorage = await ethers.getContractFactory('GatewayStorage');
    const gatewayStorage = await GatewayStorage.deploy();
    await gatewayStorage.initialize(network.address, gateways);

    await expect(gatewayStorage.mustRemove(nonExistingGateway)).to.be.revertedWith(
      'GatewayStorage: failed to remove gateway'
    );
  });
});
