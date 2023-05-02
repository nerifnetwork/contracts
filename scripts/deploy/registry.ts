import { ethers } from 'hardhat';
import { Registry } from '../../typechain';
import { Deployer } from './deployer';

const defaultRegistryDeploymentParameters: RegistryDeploymentParameters = {
  signerStorage: '',
  isMainchain: false,
  displayLogs: false,
  verify: false,
};

export async function deployRegistryContracts(options?: RegistryDeploymentOptions): Promise<RegistryDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  const [owner] = await ethers.getSigners();

  const signerStorageAddress = params.signerStorage.length > 0 ? params.signerStorage : owner.address;

  deployer.log('Deploying contracts\n');

  const res: RegistryDeployment = {
    registry: await deployer.deploy(ethers.getContractFactory('Registry'), 'Registry'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(
    res.registry.initialize(signerStorageAddress, params.isMainchain),
    'Initializing Registry'
  );

  deployer.log('Successfully initialized contracts\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: RegistryDeploymentOptions): RegistryDeploymentParameters {
  let parameters = defaultRegistryDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.signerStorage !== undefined) {
    parameters.signerStorage = options.signerStorage;
  }

  if (options.isMainchain !== undefined) {
    parameters.isMainchain = options.isMainchain;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  return parameters;
}

export interface RegistryDeploymentResult extends RegistryDeployment, RegistryDeploymentParameters {}

export interface RegistryDeployment {
  registry: Registry;
}

export interface RegistryDeploymentParameters {
  signerStorage: string;
  isMainchain: boolean;
  displayLogs: boolean;
  verify: boolean;
}

export interface RegistryDeploymentOptions {
  signerStorage?: string;
  isMainchain?: boolean;
  displayLogs?: boolean;
  verify?: boolean;
}
