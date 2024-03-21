import { Deployer } from '@solarity/hardhat-migrate';

import {
  BillingManager__factory,
  ContractsRegistry__factory,
  GatewayFactory__factory,
  Gateway__factory,
  Registry__factory,
  SignerStorage__factory,
} from '../generated-types/ethers';
import { ConfigParser } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const configParser: ConfigParser = new ConfigParser();

  const contractsRegistry = await deployer.deployed(ContractsRegistry__factory);

  if (configParser.needToDeployOperationalContracts()) {
    const billingManagerImpl = await deployer.deploy(BillingManager__factory);
    const registryImpl = await deployer.deploy(Registry__factory);
    const gatewayFactoryImpl = await deployer.deploy(GatewayFactory__factory);

    await deployer.deploy(Gateway__factory);

    await contractsRegistry.addProxyContract(
      await contractsRegistry.BILLING_MANAGER_NAME(),
      billingManagerImpl.address
    );
    await contractsRegistry.addProxyContract(await contractsRegistry.REGISTRY_NAME(), registryImpl.address);
    await contractsRegistry.addProxyContract(
      await contractsRegistry.GATEWAY_FACTORY_NAME(),
      gatewayFactoryImpl.address
    );

    if (!(await contractsRegistry.isMainChain())) {
      const signerStorageImpl = await deployer.deploy(SignerStorage__factory);

      await contractsRegistry.addProxyContract(await contractsRegistry.SIGNER_GETTER_NAME(), signerStorageImpl.address);
    }
  }
};
