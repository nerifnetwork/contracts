import hre from 'hardhat';
import { ethers } from 'hardhat';
import { expect } from 'chai';
import { deploySystem } from '../utils/deploy';

describe('DKG', function () {
  it('should broadcast all rounds', async function () {
    const initialGeneration = 0;
    const newGeneration = 1;
    const badGaneration = 4;

    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const [, v1, v2, signer, other] = await ethers.getSigners();

    const signerAddress = signer.address;
    const message = 'verify';

    const signature = await signer.signMessage(message);
    const signatureOther = await other.signMessage(message);

    const { dkg, staking, minimalStake } = await deploySystem();

    expect(await dkg.getGenerationsCount()).to.equal(1);

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await expect(dkg.roundBroadcast(initialGeneration, 1, data1)).to.be.revertedWith('DKG: not a validator');
    expect(await dkg.isValidator(initialGeneration, v1.address)).to.equal(false);
    expect(await dkg.isValidator(0, v2.address)).to.equal(false);
    expect(await dkg.getValidators(0)).to.deep.equal([]);
    expect(await dkg.getValidatorsCount(0)).to.equal(0);

    await staking1.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });

    expect(await dkg.isCurrentValidator(v1.address)).to.equal(true);
    expect(await dkg.isCurrentValidator(v2.address)).to.equal(true);

    expect(await dkg.getGenerationsCount()).to.equal(2);
    expect(await dkg.isValidator(badGaneration, v1.address)).to.equal(false);
    expect(await dkg.isValidator(newGeneration, v1.address)).to.equal(true);
    expect(await dkg.isValidator(newGeneration, v2.address)).to.equal(true);
    expect(await dkg.getValidators(newGeneration)).to.deep.equal([v1.address, v2.address]);
    expect(await dkg.getValidatorsCount(newGeneration)).to.equal(2);

    const dkgV1 = await ethers.getContractAt('DKG', dkg.address, v1);
    const dkgV2 = await ethers.getContractAt('DKG', dkg.address, v2);

    await expect(dkgV1.roundBroadcast(newGeneration, 2, data2)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV2.roundBroadcast(newGeneration, 2, data2)).to.be.revertedWith('DKG: round was not filled');

    // round1 - v1

    await expect(dkgV1.roundBroadcast(newGeneration, 1, data1))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(newGeneration, 1, v1.address);

    await expect(dkgV1.roundBroadcast(newGeneration, 1, data1)).to.be.revertedWith('DKG: round data already provided');

    expect(await dkg.getRoundBroadcastData(newGeneration, 1, v1.address)).to.equal(data1);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 1)).to.equal(1);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 2)).to.equal(0);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(newGeneration, 1)).to.equal(false);

    // round1 - v2

    await expect(dkgV2.roundBroadcast(newGeneration, 1, data1))
      .to.emit(dkgV2, 'RoundDataFilled')
      .withArgs(newGeneration, 1);

    expect(await dkg.getRoundBroadcastData(newGeneration, 1, v2.address)).to.equal(data1);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 2)).to.equal(0);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(newGeneration, 1)).to.equal(true);

    // round2 - v1

    await expect(dkgV1.roundBroadcast(newGeneration, 3, data2)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV1.roundBroadcast(newGeneration, 2, data2))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(newGeneration, 2, v1.address);

    expect(await dkg.getRoundBroadcastData(newGeneration, 2, v1.address)).to.equal(data2);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 2)).to.equal(1);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(newGeneration, 2)).to.equal(false);

    // round2 - v2

    await expect(dkgV2.roundBroadcast(newGeneration, 3, data2)).to.be.revertedWith('DKG: round was not filled');
    await expect(dkgV2.roundBroadcast(newGeneration, 2, data2))
      .to.emit(dkgV2, 'RoundDataFilled')
      .withArgs(newGeneration, 2);

    expect(await dkg.getRoundBroadcastData(newGeneration, 2, v2.address)).to.equal(data2);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 1)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 2)).to.equal(2);
    expect(await dkg.getRoundBroadcastCount(newGeneration, 3)).to.equal(0);
    expect(await dkg.isRoundFilled(newGeneration, 2)).to.equal(true);

    // round3 - v1

    await expect(dkgV1.voteSigner(newGeneration, signerAddress, signature)).to.be.revertedWith(
      'DKG: round was not filled'
    );
    await expect(dkgV1.roundBroadcast(newGeneration, 3, data3))
      .to.emit(dkgV1, 'RoundDataProvided')
      .withArgs(newGeneration, 3, v1.address);

    expect(await dkgV1.getRoundBroadcastData(newGeneration, 3, v1.address)).to.equal(data3);
    expect(await dkgV1.getRoundBroadcastCount(newGeneration, 1)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(newGeneration, 2)).to.equal(2);
    expect(await dkgV1.getRoundBroadcastCount(newGeneration, 3)).to.equal(1);
    expect(await dkg.isRoundFilled(newGeneration, 3)).to.equal(false);

    // round3 - v2

    await expect(dkgV2.voteSigner(newGeneration, signerAddress, signature)).to.be.revertedWith(
      'DKG: round was not filled'
    );
    await expect(dkgV2.roundBroadcast(newGeneration, 3, data3))
      .to.emit(dkgV2, 'RoundDataFilled')
      .withArgs(newGeneration, 3);

    expect(await dkgV2.getRoundBroadcastData(newGeneration, 3, v1.address)).to.equal(data3);
    expect(await dkgV2.getRoundBroadcastCount(newGeneration, 1)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(newGeneration, 2)).to.equal(2);
    expect(await dkgV2.getRoundBroadcastCount(newGeneration, 3)).to.equal(2);
    expect(await dkg.isRoundFilled(newGeneration, 3)).to.equal(true);

    await expect(dkgV1.voteSigner(newGeneration, signerAddress, signature))
      .to.emit(dkgV1, 'SignerVoted')
      .withArgs(newGeneration, v1.address, signerAddress);

    await expect(dkgV1.voteSigner(newGeneration, signerAddress, signatureOther)).to.be.revertedWith(
      'DKG: signature is invalid'
    );
    await expect(dkgV2.voteSigner(newGeneration, signerAddress, signatureOther)).to.be.revertedWith(
      'DKG: signature is invalid'
    );

    await expect(dkgV1.voteSigner(newGeneration, signerAddress, signature)).to.be.revertedWith('DKG: already voted');
    await expect(dkgV2.voteSigner(newGeneration, signerAddress, signature))
      .to.emit(dkgV2, 'SignerAddressUpdated')
      .withArgs(newGeneration, signerAddress);

    expect(await dkg.signerToGeneration(signerAddress)).to.equal(1);
  });

  it('should get active and pending status', async function () {
    const [, signer] = await ethers.getSigners();
    const { dkg, staking, minimalStake } = await deploySystem();

    const PENDING: number = 0;
    const ACTIVE: number = 2;

    const generation = 1;
    const message = 'verify';
    const signature = await signer.signMessage(message);
    const data1 = ethers.utils.keccak256([1]);
    const data2 = ethers.utils.keccak256([2]);
    const data3 = ethers.utils.keccak256([3]);

    const dkgSigner = await ethers.getContractAt('DKG', dkg.address, signer);
    const stakingSigner = await ethers.getContractAt('Staking', staking.address, signer);

    await staking.stake({ value: minimalStake });
    await stakingSigner.stake({ value: minimalStake });

    expect(await dkgSigner.getStatus(generation)).to.equal(PENDING);

    await dkg.roundBroadcast(generation, 1, data1);
    await dkgSigner.roundBroadcast(generation, 1, data1);

    await dkg.roundBroadcast(generation, 2, data2);
    await dkgSigner.roundBroadcast(generation, 2, data2);

    await dkg.roundBroadcast(generation, 3, data3);
    await dkgSigner.roundBroadcast(generation, 3, data3);

    expect(await dkgSigner.getStatus(generation)).to.equal(0);

    await dkg.voteSigner(generation, signer.address, signature);
    await dkgSigner.voteSigner(generation, signer.address, signature);

    expect(await dkgSigner.getStatus(generation)).to.equal(ACTIVE);
  });

  it('should get expired status', async function () {
    const [, v1, v2] = await ethers.getSigners();
    const { dkg, staking, minimalStake, dkgDeadlinePeriod } = await deploySystem();
    const data = ethers.utils.keccak256([1]);

    const PENDING: number = 0;
    const EXPIRED: number = 1;
    const newGeneration = 1;

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking1.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });

    expect(await dkg.getStatus(newGeneration)).to.equal(PENDING);
    await hre.network.provider.send('hardhat_mine', [dkgDeadlinePeriod.add(1).toHexString()]);
    expect(await dkg.getStatus(newGeneration)).to.equal(EXPIRED);

    const dkg1 = await ethers.getContractAt('DKG', dkg.address, v1);
    await expect(dkg1.roundBroadcast(newGeneration, 1, data)).to.be.revertedWith('DKG: not a pending generation');
  });

  it('should check if we can set deadline period', async function () {
    const [, other] = await ethers.getSigners();
    const { dkg } = await deploySystem();
    const dkgOther = await ethers.getContractAt('DKG', dkg.address, other);

    await dkg.setDeadlinePeriod(150);

    expect(await dkg.deadlinePeriod()).to.equals(150);

    await expect(dkgOther.setDeadlinePeriod(170)).to.be.revertedWith('DKG: not a active signer');
  });

  it('should do not create new generartion if it not diferent from previous', async function () {
    const [, v1, v2] = await ethers.getSigners();
    const { dkg, staking, minimalStake } = await deploySystem();

    const staking1 = await ethers.getContractAt('Staking', staking.address, v1);
    const staking2 = await ethers.getContractAt('Staking', staking.address, v2);

    await staking1.stake({ value: minimalStake });
    await staking2.stake({ value: minimalStake });

    expect(await dkg.getGenerationsCount()).to.equal(2);

    await dkg.updateGeneration();
    expect(await dkg.getGenerationsCount()).to.equal(2);
  });
});
