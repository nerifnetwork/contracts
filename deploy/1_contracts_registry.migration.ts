import { Deployer } from '@solarity/hardhat-migrate';

import { ContractsRegistry__factory, ERC1967Proxy__factory } from '../generated-types/ethers';
import { ConfigParser } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const configParser: ConfigParser = new ConfigParser();

  const contractsRegistryImpl = await deployer.deploy(ContractsRegistry__factory, { name: 'ContractsRegistryImpl' });
  const contractsRegistryProxy = await deployer.deploy(ERC1967Proxy__factory, [contractsRegistryImpl.address, '0x'], {
    name: 'ContractsRegistryProxy',
  });

  await deployer.save(ContractsRegistry__factory, contractsRegistryProxy.address);

  const contractsRegistry = await deployer.deployed(ContractsRegistry__factory);

  await contractsRegistry.__OwnableContractsRegistry_init();

  await contractsRegistry.setIsMainChain(configParser.configData.isMainChain);
};
