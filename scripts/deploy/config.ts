import { BaseContract } from 'ethers';
import { promises as fs } from 'fs';

const environment = process.env.ENVIRONMENT || 'local';

export interface ContractsConfig {
  [key: string]: string | undefined;
}

export async function readContractsConfig(): Promise<ContractsConfig> {
  try {
    const fileData = await fs.readFile(`deployed-contracts/${environment}/contracts.json`);
    return JSON.parse(fileData.toString());
  } catch (error) {
    return {};
  }
}

export async function writeContractsConfig(config: ContractsConfig): Promise<void> {
  try {
    await fs.writeFile(`deployed-contracts/${environment}/contracts.json`, JSON.stringify(config, null, 2));
  } catch (error) {
    console.error(error);
  }
}

export async function updateContractsConfig(config: ContractsConfig, data: Object) {
  for (const [key, value] of Object.entries(data)) {
    if (value !== undefined && value instanceof BaseContract) {
      config[key] = value.address;
    }
  }
}

export async function readChainContractsConfig(chainId: number): Promise<ContractsConfig> {
  try {
    const fileData = await fs.readFile(`deployed-contracts/${environment}/contracts-${chainId}.json`);
    return JSON.parse(fileData.toString());
  } catch (error) {
    return {};
  }
}

export async function writeChainContractsConfig(chainId: number, config: ContractsConfig): Promise<void> {
  try {
    await fs.writeFile(`deployed-contracts/${environment}/contracts-${chainId}.json`, JSON.stringify(config, null, 2));
  } catch (error) {
    console.error(error);
  }
}
