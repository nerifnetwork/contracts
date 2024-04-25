import { Deployer, Reporter } from '@solarity/hardhat-migrate';

import {
  BillingManager__factory,
  ContractsRegistry__factory,
  DKG__factory,
  GatewayFactory__factory,
  Gateway__factory,
  NerifToken__factory,
  Registry__factory,
  SlashingVoting__factory,
  Staking__factory,
  TokensVesting__factory,
} from '../generated-types/ethers';
import { ConfigParser, isZeroAddr } from './helpers/configParser';
import { ethers } from 'hardhat';

export = async (deployer: Deployer) => {
  const configParser: ConfigParser = new ConfigParser();

  const contractsRegistry = await deployer.deployed(ContractsRegistry__factory);

  const contractsInfoList: [string, string][] = [
    ['ContractsRegistry', contractsRegistry.address],
    ['----------', '----------'],
  ];

  const nerifTokenData = configParser.configData.nerifTokenData;

  if (nerifTokenData) {
    const nerifToken = await deployer.deployed(NerifToken__factory, await contractsRegistry.getNerifTokenContract());
    const tokensVesting = await deployer.deployed(
      TokensVesting__factory,
      await contractsRegistry.getTokensVestingContract()
    );

    await nerifToken.initialize(
      nerifTokenData.nerifTokenInitParams.tokenInitAmount,
      nerifTokenData.nerifTokenInitParams.tokenName,
      nerifTokenData.nerifTokenInitParams.tokenSymbol
    );
    await tokensVesting.initialize();

    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());

    contractsInfoList.push(
      ['NerifToken', nerifToken.address],
      ['TokensVesting', tokensVesting.address],
      ['----------', '----------']
    );
  }

  const systemContractsInitParams = configParser.configData.systemContractsInitParams;

  if (systemContractsInitParams) {
    const dkg = await deployer.deployed(DKG__factory, await contractsRegistry.getDKGContract());
    const staking = await deployer.deployed(Staking__factory, await contractsRegistry.getStakingContract());
    const slashingVoting = await deployer.deployed(
      SlashingVoting__factory,
      await contractsRegistry.getSlashingVotingContract()
    );

    await dkg.initialize(
      systemContractsInitParams.dkgInitParams.updatesCollectionEpochDuration,
      systemContractsInitParams.dkgInitParams.dkgGenerationEpochDuration,
      systemContractsInitParams.dkgInitParams.guaranteedWorkingEpochDuration
    );

    let stakeTokenAddr: string;

    if (isZeroAddr(systemContractsInitParams.stakingInitParams.stakingTokenAddr)) {
      stakeTokenAddr = await contractsRegistry.getNerifTokenContract();
    } else {
      stakeTokenAddr = systemContractsInitParams.stakingInitParams.stakingTokenAddr;
    }

    await staking.initialize(
      stakeTokenAddr,
      systemContractsInitParams.stakingInitParams.minimalStake,
      systemContractsInitParams.stakingInitParams.whitelistedUsers
    );
    await slashingVoting.initialize(systemContractsInitParams.slashingVotingInitParams.slashingThresoldPercentage);

    await contractsRegistry.injectDependencies(await contractsRegistry.DKG_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.STAKING_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.SLASHING_VOTING_NAME());

    contractsInfoList.push(
      ['DKG', dkg.address],
      ['Staking', staking.address],
      ['SlashingVoting', slashingVoting.address],
      ['----------', '----------']
    );
  }

  const operationalContractsInitParams = configParser.configData.operationalContractsInitParams;

  if (operationalContractsInitParams) {
    const billingManager = await deployer.deployed(
      BillingManager__factory,
      await contractsRegistry.getBillingManagerContract()
    );
    const registry = await deployer.deployed(Registry__factory, await contractsRegistry.getRegistryContract());
    const gatewayFactory = await deployer.deployed(
      GatewayFactory__factory,
      await contractsRegistry.getGatewayFactoryContract()
    );
    const gatewayImpl = await deployer.deployed(Gateway__factory);

    await billingManager.initialize({
      depositAssetKey: operationalContractsInitParams.nativeDepositAssetData.nativeDepositAssetKey,
      depositAssetData: {
        tokenAddr: ethers.constants.AddressZero,
        workflowExecutionDiscount: operationalContractsInitParams.nativeDepositAssetData.workflowExecutionDiscount,
        isEnabled: operationalContractsInitParams.nativeDepositAssetData.isEnabled,
        networkRewards: 0,
        isPermitable: false,
      },
    });
    await registry.initialize(operationalContractsInitParams.maxWorkflowsPerAccount);
    await gatewayFactory.initialize(gatewayImpl.address);

    await contractsRegistry.injectDependencies(await contractsRegistry.BILLING_MANAGER_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.REGISTRY_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.GATEWAY_FACTORY_NAME());

    contractsInfoList.push(
      ['BillingManager', billingManager.address],
      ['Registry', registry.address],
      ['GatewayFactory', gatewayFactory.address],
      ['GatewayImpl', gatewayImpl.address],
      ['----------', '----------']
    );
  }

  Reporter.reportContracts(...contractsInfoList);
};
