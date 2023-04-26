import hre from 'hardhat';
import { ethers } from 'hardhat';
import { BaseContract, Contract, ContractTransaction } from 'ethers';

export interface ContractsObject {
  [key: string]: Contract;
}

export interface ContractFactory<C extends Contract> {
  deploy(): Promise<C>;
}

export class Deployer {
  readonly displayLogs: boolean;

  constructor(displayLogs: boolean) {
    this.displayLogs = displayLogs;
  }

  public async deploy<C extends Contract>(factoryPromise: Promise<ContractFactory<C>>, name?: String): Promise<C> {
    if (name === undefined) {
      name = 'Contract';
    }

    if (this.displayLogs) {
      console.log(`Deploying ${name}`);
    }

    const factory = await factoryPromise;
    const contract = await factory.deploy();
    const chainId = ethers.provider.network.chainId;

    if (this.displayLogs) {
      console.log(`Deploying ${name} with address ${contract.address}`);
      console.log(`Deploying ${name} with chainId ${chainId}`);
    }

    await contract.deployed();

    if (this.displayLogs) {
      console.log(`Deployed\n`);
    }

    return contract;
  }

  public async sendTransaction(txPromise: Promise<ContractTransaction>, message?: String): Promise<void> {
    if (this.displayLogs) {
      console.log(message);
    }

    const tx = await txPromise;

    if (this.displayLogs) {
      console.log(`Sending transaction ${tx.hash}`);
    }

    await tx.wait();

    if (this.displayLogs) {
      console.log(`Confirmed\n`);
    }

    return;
  }

  public async verify(contract: BaseContract): Promise<void> {
    try {
      this.log(`Verifying ${contract.address}\n`);
      await hre.run('verify:verify', { address: contract.address });
    } catch (e) {
      if (this.displayLogs) {
        this.log(`Failed to verify contract ${contract.address}: ${e}`);
      }
    }

    this.log('');
  }

  public async verifyObjectValues(obj: Object): Promise<void> {
    this.log('Verifying contracts\n');

    for (const contract of Object.values(obj)) {
      if (contract instanceof BaseContract) {
        await this.verify(contract);
      }
    }

    this.log('Successfully verified contracts\n');
  }

  public log(message?: any, ...optionalParams: any[]): void {
    if (this.displayLogs) {
      console.log(message, ...optionalParams);
    }
  }
}
