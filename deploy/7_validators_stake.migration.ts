import { Deployer } from '@solarity/hardhat-migrate';

import { ContractsRegistry__factory, NerifToken__factory, Staking__factory } from '../generated-types/ethers';
import { ConfigParser } from './helpers/configParser';
import { ethers } from 'hardhat';

export = async (deployer: Deployer) => {
  const configParser: ConfigParser = new ConfigParser();

  const contractsRegistry = await deployer.deployed(ContractsRegistry__factory);

  const systemContractsInitParams = configParser.configData.systemContractsInitParams;

  if (systemContractsInitParams) {
    const staking = await deployer.deployed(Staking__factory, await contractsRegistry.getStakingContract());
    const nerifToken = await deployer.deployed(NerifToken__factory, await contractsRegistry.getNerifTokenContract());

    const minimalStake = await staking.minimalStake();

    console.log(`Minimal stake is ${minimalStake.toString()}`);

    for (let i = 0; i < systemContractsInitParams.stakingKeys.length; ++i) {
      const signer = new ethers.Wallet(systemContractsInitParams.stakingKeys[i], ethers.provider);

      await nerifToken.ownerMint(signer.address, minimalStake);
      await nerifToken.connect(signer).approve(staking.address, minimalStake);
      await staking.connect(signer).stake(minimalStake);

      console.log(`Staked ${minimalStake} from the ${signer.address} address`);
    }
  }
};
