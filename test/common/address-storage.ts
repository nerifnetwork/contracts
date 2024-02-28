import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('AddressStorage', function () {
  it('should  clear all addreses', async function () {
    const signer = (await ethers.getSigners())[0];

    const validatorsAddress = [
      '0x0000000000000000000000000000000000000000',
      '0x0000000000000000000000000000000000000001',
      '0x0000000000000000000000000000000000000002',
    ];

    const AddressStorage = await ethers.getContractFactory('AddressStorage');
    const addressStorage = await AddressStorage.deploy();
    await addressStorage.initialize(signer.address, validatorsAddress);

    expect(await addressStorage.size()).to.equal(3);
    await addressStorage.clear();
    expect(await addressStorage.size()).to.equal(0);
  });

  it('should not add an already added validator and not remove a non  xistent existing validator', async function () {
    const signer = (await ethers.getSigners())[0];

    const validatorsAddress = [
      '0x0000000000000000000000000000000000000000',
      '0x0000000000000000000000000000000000000001',
      '0x0000000000000000000000000000000000000002',
    ];
    const existingValidator = '0x0000000000000000000000000000000000000000';
    const nonExistingValidator = '0x0000000000000000000000000000000000000004';

    const AddressStorage = await ethers.getContractFactory('AddressStorage');
    const addressStorage = await AddressStorage.deploy();
    await addressStorage.initialize(signer.address, validatorsAddress);

    await expect(addressStorage.mustAdd(existingValidator)).to.be.revertedWith('AddressStorage: failed to add address');
    await expect(addressStorage.mustRemove(nonExistingValidator)).to.be.revertedWith(
      'AddressStorage: failed to remove address'
    );
  });
});
