import { expect } from 'chai';
import { ethers } from 'hardhat';
import { deploySidechain } from '../utils/deploy';

describe.skip('SignerStorage', function () {
  it('should set, get and emit event', async function () {
    const [signer, newSigner, user] = await ethers.getSigners();

    const { signerStorage } = await deploySidechain();

    const balanceSignerBefore = await ethers.provider.getBalance(signer.address);
    const balanceNewSignerBefore = await ethers.provider.getBalance(newSigner.address);

    const estimatedTxFee = ethers.utils.parseEther('0.1');

    const value = balanceSignerBefore.sub(estimatedTxFee);

    const newUser = await ethers.getContractAt('SignerStorage', signerStorage.address, user);
    await expect(newUser.setAddress(newSigner.address, { value })).to.be.revertedWith('SignerOwnable: only signer');
    await expect(signerStorage.setAddress(newSigner.address, { value }))
      .to.emit(signerStorage, 'SignerUpdated')
      .withArgs(newSigner.address);
    expect(await signerStorage.getSignerAddress()).to.equal(newSigner.address);

    const balanceNewSignerAfter = await ethers.provider.getBalance(newSigner.address);
    expect(balanceNewSignerAfter).to.be.equal(balanceNewSignerBefore.add(value));

    await (
      await newSigner.sendTransaction({
        to: signer.address,
        value: value,
      })
    ).wait();
  });
});
