import { Deployer } from '@solarity/hardhat-migrate';

import {
  ContractsRegistry__factory,
  DKG__factory,
  SlashingVoting__factory,
  Staking__factory,
} from '../generated-types/ethers';
import { ConfigParser } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const configParser: ConfigParser = new ConfigParser();

  const contractsRegistry = await deployer.deployed(ContractsRegistry__factory);

  if (configParser.needToDeploySystemContracts() && (await contractsRegistry.isMainChain())) {
    const dkgImpl = await deployer.deploy(DKG__factory);
    const stakingImpl = await deployer.deploy(Staking__factory);
    const slashingVotingImpl = await deployer.deploy(SlashingVoting__factory);

    await contractsRegistry.addProxyContract(await contractsRegistry.DKG_NAME(), dkgImpl.address);
    await contractsRegistry.addProxyContract(await contractsRegistry.STAKING_NAME(), stakingImpl.address);
    await contractsRegistry.addProxyContract(
      await contractsRegistry.SLASHING_VOTING_NAME(),
      slashingVotingImpl.address
    );

    await contractsRegistry.addContract(
      await contractsRegistry.SIGNER_GETTER_NAME(),
      await contractsRegistry.getDKGContract()
    );
  }
};
