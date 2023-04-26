import { ethers, network } from 'hardhat';
import { deployBridgeContracts } from './deploy/sidechain';
import { readChainContractsConfig, readContractsConfig, updateContractsConfig, writeChainContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const chainId = network.config.chainId ?? 1;
  const blockNumber = await ethers.provider.getBlockNumber();

  const homeContractsConfig = await readContractsConfig();
  const homeNetwork = homeContractsConfig.networkName;
  const dkgAddress = homeContractsConfig.dkg;

  const contractsConfig = await readChainContractsConfig(chainId);

  const res = await deployBridgeContracts({
    homeNetwork: homeNetwork,
    homeDKGAddress: dkgAddress,
    displayLogs: true,
    verify: verify,
  });

  contractsConfig.networkName = network.name;
  contractsConfig.startBlock = blockNumber.toString();

  updateContractsConfig(contractsConfig, res);
  await writeChainContractsConfig(chainId, contractsConfig);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
