import { Deployer, Reporter } from '@solarity/hardhat-migrate';

import {
  AddressStorage__factory,
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
    const contractsRegistry = await deployer.deploy(ContractRegistry__factory);
    const addressStorage = await deployer.deploy(AddressStorage__factory);
    const staking = await deployer.deploy(Staking__factory);
    const rewardsDistributionPool = await deployer.deploy(RewardDistributionPool__factory);
    const dkg = await deployer.deploy(DKG__factory);
    const slashingVoting = await deployer.deploy(SlashingVoting__factory);

    await addressStorage.initialize(staking.address, []);
    await dkg.initialize(contractsRegistry.address, config.systemContractsInitParams.dkgDeadlinePeriod);
    await contractsRegistry.initialize(dkg.address);
    await staking.initialize(
      dkg.address,
      config.systemContractsInitParams.stakingInitParams.minimalStake,
      config.systemContractsInitParams.stakingInitParams.withdrawalPeriod,
      contractsRegistry.address,
      addressStorage.address
    );
    await slashingVoting.initialize(
      dkg.address,
      staking.address,
      config.systemContractsInitParams.slashingVotingInitParams.epochPeriod,
      config.systemContractsInitParams.slashingVotingInitParams.slashingThresold,
      config.systemContractsInitParams.slashingVotingInitParams.slashingEpochs,
      contractsRegistry.address
    );
    await rewardsDistributionPool.initialize(contractsRegistry.address, dkg.address);

    await contractsRegistry.setContract(await slashingVoting.SLASHING_VOTING_KEY(), slashingVoting.address);
    await contractsRegistry.setContract(await dkg.DKG_KEY(), dkg.address);
    await contractsRegistry.setContract(await staking.STAKING_KEY(), staking.address);
    await contractsRegistry.setContract(
      await rewardsDistributionPool.REWARD_DISTRIBUTION_POOL_KEY(),
      rewardsDistributionPool.address
    );

    Reporter.reportContracts(
      ["ContractRegistry", contractsRegistry.address],
      ["DKG", dkg.address],
      ["Staking", staking.address],
      ["SlashingVoting", slashingVoting.address],
      ["RewardDistributionPool", rewardsDistributionPool.address],
      ["AddressStorage", addressStorage.address],
    );
  }
};
