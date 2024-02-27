import { Deployer } from '@solarity/hardhat-migrate';
import { ethers } from 'hardhat';

import { Staking__factory } from '../generated-types/ethers';
import { parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  if (config.isMainChain && config.stakingKeys.length > 0) {
    const staking = await deployer.deployed(Staking__factory);
    const minimalStake = await staking.minimalStake();

    for (let i = 0; i < config.stakingKeys.length; i++) {
      const signer = new ethers.Wallet(config.stakingKeys[i], ethers.provider);

      await staking.connect(signer).stake({ value: minimalStake });
      console.log(`Staked ${minimalStake} from the ${signer.address} address`);
    }
  }
};
