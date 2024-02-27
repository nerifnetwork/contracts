import { Deployer, Reporter } from '@solarity/hardhat-migrate';

import {
  BillingManager__factory,
  GatewayFactory__factory,
  Gateway__factory,
  Registry__factory,
  SignerStorage__factory,
} from '../generated-types/ethers';
import { parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  const billingManager = await deployer.deploy(BillingManager__factory);
  const gatewayImpl = await deployer.deploy(Gateway__factory);
  const gatewayFactory = await deployer.deploy(GatewayFactory__factory);
  const registry = await deployer.deploy(Registry__factory);
  const signerStorage = await deployer.deploy(SignerStorage__factory);

  if (config.operationalContractsInitParams.signer) {
    await signerStorage.initialize(config.operationalContractsInitParams.signer);
  }

  await billingManager.initialize(registry.address, signerStorage.address);
  await gatewayFactory.initialize(registry.address, gatewayImpl.address);
  await registry.initialize(
    config.isMainChain,
    signerStorage.address,
    gatewayFactory.address,
    billingManager.address,
    config.operationalContractsInitParams.maxWorkflowsPerAccount
  );

  Reporter.reportContracts(
    ["BillingManager", billingManager.address],
    ["Gateway Impl", gatewayImpl.address],
    ["GatewayFactory", gatewayFactory.address],
    ["Registry", registry.address],
    ["SignerStorage", signerStorage.address],
  );
};
