import { Deployer } from '@solarity/hardhat-migrate';

import { ContractsRegistry__factory, NerifToken__factory, TokensVesting__factory } from '../generated-types/ethers';
import { ConfigParser, isZeroAddr } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const configParser: ConfigParser = new ConfigParser();

  const contractsRegistry = await deployer.deployed(ContractsRegistry__factory);

  const nerifTokenData = configParser.configData.nerifTokenData;

  if (nerifTokenData) {
    const nerifToken = await deployer.deployed(NerifToken__factory, await contractsRegistry.getNerifTokenContract());
    const tokensVesting = await deployer.deployed(
      TokensVesting__factory,
      await contractsRegistry.getTokensVestingContract()
    );

    for (let i = 0; i < nerifTokenData.schedules.length; i++) {
      await tokensVesting.createSchedule({
        cliffInPeriods: nerifTokenData.schedules[i].cliffInPeriods,
        durationInPeriods: nerifTokenData.schedules[i].durationInPeriods,
        secondsInPeriod: nerifTokenData.schedules[i].secondsInPeriod,
      });

      console.log(`Created schedule with ID - ${await tokensVesting.scheduleId()}`);
    }

    for (let i = 0; i < nerifTokenData.vestings.length; i++) {
      await tokensVesting.createVesting({
        vestingToken: isZeroAddr(nerifTokenData.vestings[i].vestingToken)
          ? nerifToken.address
          : nerifTokenData.vestings[i].vestingToken,
        beneficiary: nerifTokenData.vestings[i].beneficiary,
        scheduleId: nerifTokenData.vestings[i].scheduleId,
        vestingAmount: nerifTokenData.vestings[i].vestingAmount,
        vestingStartTime: nerifTokenData.vestings[i].vestingStartTime,
        paidAmount: 0,
      });

      console.log(`Created vesting with ID - ${await tokensVesting.vestingId()}`);
    }
  }
};
