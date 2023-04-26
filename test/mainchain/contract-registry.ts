import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('ContractRegistry', function () {
  it('should set and get contract address', async function () {
    const [, other] = await ethers.getSigners();

    const key = 'test1';
    const otherKey = 'test2';
    const address = '0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF';

    const { contractRegistry } = await deploySystem();

    const contractRegistryOther = await ethers.getContractAt('ContractRegistry', contractRegistry.address, other);

    await contractRegistry.setContract(key, address);
    await expect(contractRegistryOther.setContract(key, address)).to.be.revertedWith('SignerOwnable: only signer');

    expect(await contractRegistry.getContract(key)).to.equal(address);
    await expect(contractRegistry.getContract(otherKey)).to.be.revertedWith(
      'ContractRegistry: no address was found for the specified key'
    );
  });
});
