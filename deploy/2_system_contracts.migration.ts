import { Deployer } from '@solarity/hardhat-migrate';

import {
  ContractRegistry__factory,
  DKG__factory,
  RewardDistributionPool__factory,
  SlashingVoting__factory,
  Staking__factory,
} from '../generated-types/ethers';
import { parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  if (config.isMainChain && config.systemContractsInitParams) {
    await deployer.deploy(ContractRegistry__factory);
    await deployer.deploy(Staking__factory);
    await deployer.deploy(RewardDistributionPool__factory);
    await deployer.deploy(DKG__factory);
    await deployer.deploy(SlashingVoting__factory);
  }
};
