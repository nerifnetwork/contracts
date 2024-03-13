import { Deployer, Reporter } from '@solarity/hardhat-migrate';

import {
  ContractRegistry__factory,
  DKG__factory,
  RewardDistributionPool__factory,
  SlashingVoting__factory,
  Staking__factory,
  TestERC20,
  TestERC20__factory,
} from '../generated-types/ethers';
import { isEmptyField, parseConfig } from './helpers/configParser';
import { wei } from '@/test/helpers/utils';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  if (config.isMainChain && config.systemContractsInitParams) {
    const contractsRegistry = await deployer.deploy(ContractRegistry__factory);
    const staking = await deployer.deploy(Staking__factory);
    const rewardsDistributionPool = await deployer.deploy(RewardDistributionPool__factory);
    const dkg = await deployer.deploy(DKG__factory);
    const slashingVoting = await deployer.deploy(SlashingVoting__factory);

    let stakeToken: TestERC20;

    if (isEmptyField(config.systemContractsInitParams.stakingInitParams.stakeTokenAddr)) {
      stakeToken = await deployer.deploy(TestERC20__factory, ["Test Nerif Token", "TNT", wei("100000").toString()]);
    } else {
      stakeToken = await deployer.deployed(TestERC20__factory, config.systemContractsInitParams.stakingInitParams.stakeTokenAddr);

      await deployer.save(TestERC20__factory, config.systemContractsInitParams.stakingInitParams.stakeTokenAddr);
    }

    await dkg.initialize(contractsRegistry.address, config.systemContractsInitParams.dkgDeadlinePeriod);
    await contractsRegistry.initialize(dkg.address);
    await staking.initialize(
      dkg.address,
      contractsRegistry.address,
      stakeToken.address,
      config.systemContractsInitParams.stakingInitParams.minimalStake,
      config.systemContractsInitParams.stakingInitParams.withdrawalPeriod,
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
      ["StakeToken", stakeToken.address],
    );
  }
};
