import { promises as fs } from 'fs';
import { stringify } from 'yaml';
import { config } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import { ContractsConfig } from './config';

export interface BridgeConfig {
  [key: string]: Config | undefined;
}

export interface Config {
  chainId: Number;
  rpcURL: string;
  bridgeContract: string;
  signerStorageContract: string;
  startBlock: Number;
}

export async function createBridgeConfig(contractConfigs: ContractsConfig[]): Promise<BridgeConfig> {
  var bridgeConfig: BridgeConfig = {};
  contractConfigs.forEach(async function (contractConfig) {
    if (
      contractConfig.relayBridge !== undefined &&
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

      var con: Config = {
        chainId: chainId,
        rpcURL: rpcURL,
        bridgeContract: contractConfig.relayBridge,
        signerStorageContract: contractConfig.signerStorage,
        startBlock: Number(contractConfig.startBlock),
      };

      bridgeConfig[contractConfig.networkName] = con;
    }
  });

  return bridgeConfig;
}

export async function writeBridgeConfig(bg: BridgeConfig): Promise<void> {
  try {
    await fs.writeFile('bridgeconfig.yaml', stringify(bg, null, 2));
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
