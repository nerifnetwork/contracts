import { ethers } from 'hardhat';
import { Registry } from '../../typechain';
import { Deployer } from './deployer';
import { BigNumber } from 'ethers';

// TODO: Define more elegant way
const internalWorkflows = [
  // Workflow registration workflow
  {
    id: '40505927788353901442144037336646356013',
    owner: '0x0000000000000000000000000000000000000000',
    hash: '0x00',
    status: 1,
    isInternal: true,
    totalSpent: 0,
  },
  // Workflow cancellation workflow
  {
    id: '219775546284901721155783592958414245131',
    owner: '0x0000000000000000000000000000000000000000',
    hash: '0x00',
    status: 1,
    isInternal: true,
    totalSpent: 0,
  },
];

const defaultRegistryDeploymentParameters: RegistryDeploymentParameters = {
  signerStorage: '',
  isMainchain: false,
  displayLogs: false,
  verify: false,
};

export async function deployRegistryContracts(options?: RegistryDeploymentOptions): Promise<RegistryDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  const res: RegistryDeployment = {
    registry: await deployer.deploy(ethers.getContractFactory('Registry'), 'Registry'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  // Get current config if exists
  let registryConfig = {
    performanceOverhead: BigNumber.from(0),
    performancePremiumThreshold: 0,
    registrationOverhead: BigNumber.from(0),
    cancellationOverhead: BigNumber.from(0),
    maxWorkflowsPerAccount: 0,
  };
  if (params.registryAddr && params.registryAddr.length > 0) {
    const existingRegistryFactory = await ethers.getContractFactory('Registry');
    const existingRegistry = existingRegistryFactory.attach(params.registryAddr);
    const existingConfig = await existingRegistry.config();

    registryConfig.performanceOverhead = existingConfig.performanceOverhead;
    registryConfig.performancePremiumThreshold = existingConfig.performancePremiumThreshold;
    registryConfig.registrationOverhead = existingConfig.registrationOverhead;
    registryConfig.cancellationOverhead = existingConfig.cancellationOverhead;
    registryConfig.maxWorkflowsPerAccount = existingConfig.maxWorkflowsPerAccount;
  }

  await deployer.sendTransaction(
    res.registry.initialize(params.isMainchain, params.signerStorage, internalWorkflows, registryConfig),
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

  if (options.registry !== undefined) {
    parameters.registryAddr = options.registry;
  }

  if (options.isMainchain !== undefined) {
    parameters.isMainchain = options.isMainchain;
  }

  if (options.signerStorage !== undefined) {
    parameters.signerStorage = options.signerStorage;
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
  registryAddr?: string;
  isMainchain: boolean;
  signerStorage: string;
  displayLogs: boolean;
  verify: boolean;
}

export interface RegistryDeploymentOptions {
  signerStorage: string;
  registry?: string;
  isMainchain?: boolean;
  displayLogs?: boolean;
  verify?: boolean;
}
