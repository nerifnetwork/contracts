import { ethers } from 'hardhat';
import { Gateway } from '../../typechain';
import { Deployer } from './deployer';

const defaultGatewayDeploymentParameters: GatewayDeploymentParameters = {
  registry: '',
  displayLogs: false,
  verify: false,
};

export async function deployGatewayContracts(options?: GatewayDeploymentOptions): Promise<GatewayDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  const res: GatewayDeployment = {
    gateway: await deployer.deploy(ethers.getContractFactory('Gateway'), 'Gateway'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.gateway.initialize(params.registry), 'Initializing Gateway');

  deployer.log('Successfully initialized contracts\n');

  deployer.log('Adding gateway to registry\n');

  const [owner] = await ethers.getSigners();

  const registryFactory = await ethers.getContractFactory('Registry', owner);
  const registry = registryFactory.attach(params.registry);

  await deployer.sendTransaction(registry.setGateway(owner.address, res.gateway.address), 'Registering Gateway');

  deployer.log('Successfully added gateway to registry\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: GatewayDeploymentOptions): GatewayDeploymentParameters {
  let parameters = defaultGatewayDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.registry !== undefined) {
    parameters.registry = options.registry;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  return parameters;
}

export interface GatewayDeploymentResult extends GatewayDeployment, GatewayDeploymentParameters {}

export interface GatewayDeployment {
  gateway: Gateway;
}

export interface GatewayDeploymentParameters {
  registry: string;
  displayLogs: boolean;
  verify: boolean;
}

export interface GatewayDeploymentOptions {
  registry?: string;
  displayLogs?: boolean;
  verify?: boolean;
}
