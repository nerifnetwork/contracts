import { ethers, network } from 'hardhat';
import { deployRegistryContracts } from './deploy/registry';
import {
    readChainContractsConfig,
    readContractsConfig,
    updateContractsConfig,
    writeChainContractsConfig,
    writeContractsConfig
} from './deploy/config';

async function main() {
    const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
    const chainId = network.config.chainId ?? 1;
    const blockNumber = await ethers.provider.getBlockNumber();

    const homeContractsConfig = await readContractsConfig();
    const contractsConfig = await readChainContractsConfig(chainId);
    const isMainchain = contractsConfig.networkName === undefined || homeContractsConfig.networkName === contractsConfig.networkName;
    const signerStorage = isMainchain ? homeContractsConfig.dkg : contractsConfig.signerStorage;

    const res = await deployRegistryContracts({
        gatewayStorage: contractsConfig.gatewayStorage,
        workflowStorage: contractsConfig.workflowStorage,
        signerStorage: signerStorage,
        isMainchain: isMainchain,
        displayLogs: true,
        verify: verify,
    });

    if (isMainchain) {
        updateContractsConfig(homeContractsConfig, res);
        await writeContractsConfig(homeContractsConfig);
    } else {
        contractsConfig.networkName = network.name;
        contractsConfig.startBlock = blockNumber.toString();

        updateContractsConfig(contractsConfig, res);
        await writeChainContractsConfig(chainId, contractsConfig);
    }
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
