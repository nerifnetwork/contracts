import { Deployer } from "@solarity/hardhat-migrate";
import { ethers } from "hardhat";
import { BigNumber, BigNumberish } from "ethers";

import {AddressStorage__factory, ContractRegistry__factory, DKG__factory, RewardDistributionPool__factory, SlashingVoting__factory, Staking__factory} from "@/generated-types/ethers";

const defaultSystemDeploymentParameters: MainchainDeploymentParameters = {
  minimalStake: ethers.utils.parseEther('0.01').toHexString(),
  stakeWithdrawalPeriod: 60,
  slashingEpochs: 3,
  slashingEpochPeriod: 1000,
  slashingBansThresold: 10,

  dkgDeadlinePeriod: 10000,

  displayLogs: false,
  verify: false,
  stakingKeys: [],
};

interface MainchainDeploymentParameters {
  minimalStake: BigNumberish;
  stakeWithdrawalPeriod: BigNumberish;

  slashingEpochs: BigNumberish;
  slashingEpochPeriod: BigNumberish;
  slashingBansThresold: BigNumberish;

  dkgDeadlinePeriod: BigNumberish;

  displayLogs: boolean;
  verify: boolean;
  stakingKeys: string[];
};

export = async (deployer: Deployer) => {
  const contractsRegistry = await deployer.deploy(ContractRegistry__factory);
  const addressStorage = await deployer.deploy(AddressStorage__factory);
  const staking = await deployer.deploy(Staking__factory);
  const rewardsDistributionPool = await deployer.deploy(RewardDistributionPool__factory);
  const dkg = await deployer.deploy(DKG__factory);
  const slashingVoting = await deployer.deploy(SlashingVoting__factory);

  await addressStorage.initialize(staking.address, []);
  await dkg.initialize(contractsRegistry.address, defaultSystemDeploymentParameters.dkgDeadlinePeriod);
  await contractsRegistry.initialize(dkg.address);
  await staking.initialize(
    dkg.address,
    defaultSystemDeploymentParameters.minimalStake,
    defaultSystemDeploymentParameters.stakeWithdrawalPeriod,
    contractsRegistry.address,
    addressStorage.address
  );
  await slashingVoting.initialize(
    dkg.address,
    staking.address,
    defaultSystemDeploymentParameters.slashingEpochPeriod,
    defaultSystemDeploymentParameters.slashingBansThresold,
    defaultSystemDeploymentParameters.slashingEpochs,
    contractsRegistry.address
  );
  await rewardsDistributionPool.initialize(
    contractsRegistry.address,
    dkg.address
  );

  await contractsRegistry.setContract(await slashingVoting.SLASHING_VOTING_KEY(), slashingVoting.address);
  await contractsRegistry.setContract(await dkg.DKG_KEY(), dkg.address);
  await contractsRegistry.setContract(await staking.STAKING_KEY(), staking.address);
  await contractsRegistry.setContract(await rewardsDistributionPool.REWARD_DISTRIBUTION_POOL_KEY(), rewardsDistributionPool.address);
};
