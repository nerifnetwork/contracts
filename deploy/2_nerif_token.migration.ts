import { Deployer } from '@solarity/hardhat-migrate';

import { ContractsRegistry__factory, NerifToken__factory, TokensVesting__factory } from '../generated-types/ethers';
import { ConfigParser } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const configParser: ConfigParser = new ConfigParser();

  const contractsRegistry = await deployer.deployed(ContractsRegistry__factory);

  if (configParser.needToDeployNerifToken()) {
    const nerifTokenImpl = await deployer.deploy(NerifToken__factory);
    const tokensVestingImpl = await deployer.deploy(TokensVesting__factory);

    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);
    await contractsRegistry.addProxyContract(await contractsRegistry.TOKENS_VESTING_NAME(), tokensVestingImpl.address);
  }
};
