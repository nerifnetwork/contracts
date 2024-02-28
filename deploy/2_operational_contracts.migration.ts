import { Deployer, Reporter } from '@solarity/hardhat-migrate';

import {
  BillingManager__factory,
  ERC1967Proxy__factory,
  GatewayFactory__factory,
  Gateway__factory,
  Registry__factory,
  SignerStorage__factory,
} from '../generated-types/ethers';
import { parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  const billingManagerImpl = await deployer.deploy(BillingManager__factory, {name: "BillingManagerImpl"});
  const registryImpl = await deployer.deploy(Registry__factory, {name: "RegistryImpl"});
  const gatewayImpl = await deployer.deploy(Gateway__factory, {name: "GatewayImpl"});

  const billingManagerProxy = await deployer.deploy(ERC1967Proxy__factory, [billingManagerImpl.address, "0x"], {name: "BillingManagerProxy"});
  const registryProxy = await deployer.deploy(ERC1967Proxy__factory, [registryImpl.address, "0x"], {name: "RegistryProxy"});

  await deployer.save(BillingManager__factory, billingManagerProxy.address);
  await deployer.save(Registry__factory, registryProxy.address);

  const billingManager = await deployer.deployed(BillingManager__factory);
  const registry = await deployer.deployed(Registry__factory);

  const gatewayFactory = await deployer.deploy(GatewayFactory__factory);
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
