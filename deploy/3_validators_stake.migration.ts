import { Deployer } from '@solarity/hardhat-migrate';
import { ethers } from 'hardhat';

import {Staking__factory, TestERC20__factory} from '../generated-types/ethers';
import { parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  if (config.isMainChain && config.stakingKeys.length > 0) {
    const token = await deployer.deployed(TestERC20__factory);
    const staking = await deployer.deployed(Staking__factory);
    const minimalStake = await staking.minimalStake();
    console.log(`Minimal stake is ${minimalStake.toString()}`);

    for (let i = 0; i < config.stakingKeys.length; i++) {
      const signer = new ethers.Wallet(config.stakingKeys[i], ethers.provider);

      await token.connect(signer).approve(signer.address, minimalStake);
      await token.connect(signer).mint(signer.address, minimalStake);
      await staking.connect(signer).stake(minimalStake);
      console.log(`Staked ${minimalStake} from the ${signer.address} address`);
    }
  }
};
