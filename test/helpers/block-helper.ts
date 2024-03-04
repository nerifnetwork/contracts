import hre from 'hardhat';

export async function setNextBlockTime(time: number) {
  await hre.network.provider.send('evm_setNextBlockTimestamp', [time]);
}

export async function setTime(time: number) {
  await setNextBlockTime(time);
  await mine();
}

export async function getCurrentBlockTime() {
  return (await hre.ethers.provider.getBlock(await hre.ethers.provider.getBlockNumber())).timestamp;
}

export async function mine(numberOfBlocks = 1) {
  for (let i = 0; i < numberOfBlocks; i++) {
    await hre.network.provider.send('evm_mine');
  }
}
