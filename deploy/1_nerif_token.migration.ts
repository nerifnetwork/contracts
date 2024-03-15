import { Deployer } from '@solarity/hardhat-migrate';

import { NerifToken__factory, TokensVesting__factory } from '../generated-types/ethers';
import { parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  if (config.systemTokenData) {
    await deployer.deploy(NerifToken__factory);
    await deployer.deploy(TokensVesting__factory);
  }
};
