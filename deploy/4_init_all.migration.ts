import { Deployer, Reporter } from '@solarity/hardhat-migrate';

import {
  BillingManager__factory,
  ContractRegistry__factory,
  DKG__factory,
  GatewayFactory__factory,
  Gateway__factory,
  NerifToken,
  NerifToken__factory,
  Registry__factory,
  RewardDistributionPool__factory,
  SignerStorage__factory,
  SlashingVoting__factory,
  Staking__factory,
  TokensVesting__factory,
} from '../generated-types/ethers';
import { isZeroAddr, parseConfig } from './helpers/configParser';
import { ethers } from 'hardhat';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  let nerifToken: NerifToken;
  let contractsToLog: [string, string][] = [];

  if (config.isMainChain && config.systemContractsInitParams) {
    const contractsRegistry = await deployer.deployed(ContractRegistry__factory);
    const staking = await deployer.deployed(Staking__factory);
    const rewardsDistributionPool = await deployer.deployed(RewardDistributionPool__factory);
    const dkg = await deployer.deployed(DKG__factory);
    const slashingVoting = await deployer.deployed(SlashingVoting__factory);

    const tokensVesting = await deployer.deployed(TokensVesting__factory);

    if (isZeroAddr(config.systemContractsInitParams.stakingInitParams.stakingTokenAddr)) {
      nerifToken = await deployer.deployed(NerifToken__factory);
    } else {
      nerifToken = await deployer.deployed(
        NerifToken__factory,
        config.systemContractsInitParams.stakingInitParams.stakingTokenAddr
      );

      await deployer.save(NerifToken__factory, config.systemContractsInitParams.stakingInitParams.stakingTokenAddr);
    }

    await dkg.initialize(contractsRegistry.address, config.systemContractsInitParams.dkgDeadlinePeriod);
    await contractsRegistry.initialize(dkg.address);
    await staking.initialize(
      dkg.address,
      contractsRegistry.address,
      nerifToken.address,
      config.systemContractsInitParams.stakingInitParams.minimalStake,
      config.systemContractsInitParams.stakingInitParams.withdrawalPeriod
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

    if (config.systemTokenData) {
      await nerifToken.initialize(
        contractsRegistry.address,
        config.systemTokenData.nerifTokenInitParams.tokenInitAmount,
        config.systemTokenData.nerifTokenInitParams.tokenName,
        config.systemTokenData.nerifTokenInitParams.tokenSymbol
      );
      await tokensVesting.initialize();
    }

    await contractsRegistry.setContract(await slashingVoting.SLASHING_VOTING_KEY(), slashingVoting.address);
    await contractsRegistry.setContract(await dkg.DKG_KEY(), dkg.address);
    await contractsRegistry.setContract(await staking.STAKING_KEY(), staking.address);
    await contractsRegistry.setContract(
      await rewardsDistributionPool.REWARD_DISTRIBUTION_POOL_KEY(),
      rewardsDistributionPool.address
    );
    await contractsRegistry.setContract(await staking.TOKENS_VESTING_KEY(), tokensVesting.address);

    contractsToLog.push(
      ['ContractRegistry', contractsRegistry.address],
      ['DKG', dkg.address],
      ['Staking', staking.address],
      ['SlashingVoting', slashingVoting.address],
      ['RewardDistributionPool', rewardsDistributionPool.address],
      ['NerifToken', nerifToken.address],
      ['TokensVesting', tokensVesting.address]
    );
  }

  const billingManager = await deployer.deployed(BillingManager__factory);
  const registry = await deployer.deployed(Registry__factory);
  const gatewayFactory = await deployer.deployed(GatewayFactory__factory);
  const gatewayImpl = await deployer.deployed(Gateway__factory, 'GatewayImpl');
  const signerStorage = await deployer.deployed(SignerStorage__factory);

  await billingManager.initialize(registry.address, signerStorage.address, {
    depositAssetKey: config.operationalContractsInitParams.nativeDepositAssetData.nativeDepositAssetKey,
    depositAssetData: {
      tokenAddr: ethers.constants.AddressZero,
      workflowExecutionDiscount: config.operationalContractsInitParams.nativeDepositAssetData.workflowExecutionDiscount,
      isEnabled: config.operationalContractsInitParams.nativeDepositAssetData.isEnabled,
      networkRewards: 0,
      isPermitable: false,
    },
  });
  await gatewayFactory.initialize(registry.address, gatewayImpl.address);
  await registry.initialize(
    config.isMainChain,
    signerStorage.address,
    gatewayFactory.address,
    billingManager.address,
    config.operationalContractsInitParams.maxWorkflowsPerAccount
  );

  contractsToLog.push(
    ['BillingManager', billingManager.address],
    ['Gateway Impl', gatewayImpl.address],
    ['GatewayFactory', gatewayFactory.address],
    ['Registry', registry.address],
    ['SignerStorage', signerStorage.address]
  );

  Reporter.reportContracts(...contractsToLog);
};
