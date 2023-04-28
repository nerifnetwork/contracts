import { promises as fs } from 'fs';
import { stringify } from 'yaml';
import { config } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import { ContractsConfig } from './config';

export interface SidechainConfig {
  [key: string]: Config | undefined;
}

export interface Config {
  chainId: Number;
  rpcURL: string;
  signerStorageContract: string;
  startBlock: Number;
}

export async function createSidechainConfig(contractConfigs: ContractsConfig[]): Promise<SidechainConfig> {
  let sidechainConfig: SidechainConfig = {};
  contractConfigs.forEach(async function (contractConfig) {
    if (
      contractConfig.signerStorage !== undefined &&
      contractConfig.startBlock !== undefined &&
      contractConfig.networkName !== undefined
    ) {
      const networkConfig = config.networks[contractConfig.networkName] as HttpNetworkConfig;
      if (networkConfig === undefined) {
        return;
      }

      const rpcURL = networkConfig.url;
      const chainId = networkConfig.chainId ?? 1;

      sidechainConfig[contractConfig.networkName] = {
        chainId: chainId,
        rpcURL: rpcURL,
        signerStorageContract: contractConfig.signerStorage,
        startBlock: Number(contractConfig.startBlock),
      };
    }
  });

  return sidechainConfig;
}

export async function writeSidechainConfig(bg: SidechainConfig): Promise<void> {
  try {
    await fs.writeFile('sidechainconfig.yaml', stringify(bg, null, 2));
  } catch (error) {
    console.error(error);
  }
}

export async function readContractsConfig(path: string): Promise<ContractsConfig> {
  try {
    const fileData = await fs.readFile(path);
    const config = JSON.parse(fileData.toString());

    return config;
  } catch (error) {
    return {};
  }
}
