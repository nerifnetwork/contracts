import { expect } from 'chai';
import { utils } from 'ethers';
import { ethers } from 'hardhat';
import { deployBridge } from '../utils/deploy';

describe('SignerStorage', function () {
  it('should set, get and emit event', async function () {
    const [signer, newSigner, user] = await ethers.getSigners();

    const { signerStorage } = await deployBridge();

    const balanceSignerBefore = await ethers.provider.getBalance(signer.address);
    const balanceNewSignerBefore = await ethers.provider.getBalance(newSigner.address);

    const estimatedTxFee = utils.parseEther('0.04');

    const value = balanceSignerBefore.sub(estimatedTxFee);

    const newUser = await ethers.getContractAt('SignerStorage', signerStorage.address, user);
    await expect(newUser.setAddress(newSigner.address, { value })).to.be.revertedWith('SignerStorage: only signer');
    await expect(signerStorage.setAddress(newSigner.address, { value }))
      .to.emit(signerStorage, 'SignerUpdated')
      .withArgs(newSigner.address);
    expect(await signerStorage.getAddress()).to.equal(newSigner.address);

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
