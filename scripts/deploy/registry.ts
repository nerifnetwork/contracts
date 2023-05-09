import { ethers } from 'hardhat';
import { GatewayStorage, Registry, WorkflowStorage } from '../../typechain';
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
    gatewayStorage: await deployer.deploy(
      ethers.getContractFactory('GatewayStorage'),
      'GatewayStorage',
      params.gatewayStorageAddr
    ),
    workflowStorage: await deployer.deploy(
      ethers.getContractFactory('WorkflowStorage'),
      'WorkflowStorage',
      params.workflowStorageAddr
    ),
    registry: await deployer.deploy(ethers.getContractFactory('Registry'), 'Registry'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  if (!params.gatewayStorageAddr) {
    await deployer.sendTransaction(
      res.gatewayStorage.initialize(res.registry.address, []),
      'Initializing GatewayStorage'
    );
  } else {
    await deployer.sendTransaction(
      res.gatewayStorage.setRegistry(res.registry.address),
      'Updating registry in GatewayStorage'
    );
  }

  if (!params.workflowStorageAddr) {
    await deployer.sendTransaction(
      res.workflowStorage.initialize(res.registry.address, internalWorkflows),
      'Initializing WorkflowStorage'
    );
  } else {
    await deployer.sendTransaction(
      res.workflowStorage.setRegistry(res.registry.address),
      'Updating registry in WorkflowStorage'
    );
  }

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
    res.registry.initialize(
      params.isMainchain,
      res.workflowStorage.address,
      res.gatewayStorage.address,
      params.signerStorage,
      registryConfig
    ),
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

  if (options.gatewayStorage !== undefined) {
    parameters.gatewayStorageAddr = options.gatewayStorage;
  }

  if (options.workflowStorage !== undefined) {
    parameters.workflowStorageAddr = options.workflowStorage;
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
  gatewayStorage: GatewayStorage;
  workflowStorage: WorkflowStorage;
  registry: Registry;
}

export interface RegistryDeploymentParameters {
  gatewayStorageAddr?: string;
  workflowStorageAddr?: string;
  registryAddr?: string;
  isMainchain: boolean;
  signerStorage: string;
  displayLogs: boolean;
  verify: boolean;
}

export interface RegistryDeploymentOptions {
  signerStorage: string;
  gatewayStorage?: string;
  workflowStorage?: string;
  registry?: string;
  isMainchain?: boolean;
  displayLogs?: boolean;
  verify?: boolean;
}
