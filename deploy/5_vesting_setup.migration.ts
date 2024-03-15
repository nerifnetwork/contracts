import { Deployer } from '@solarity/hardhat-migrate';

import { NerifToken__factory, TokensVesting__factory } from '../generated-types/ethers';
import { isZeroAddr, parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  if (config.systemTokenData) {
    const nerifToken = await deployer.deployed(NerifToken__factory);
    const tokensVesting = await deployer.deployed(TokensVesting__factory);

    for (let i = 0; i < config.systemTokenData.schedules.length; i++) {
      await tokensVesting.createSchedule({
        cliffInPeriods: config.systemTokenData.schedules[i].cliffInPeriods,
        durationInPeriods: config.systemTokenData.schedules[i].durationInPeriods,
        secondsInPeriod: config.systemTokenData.schedules[i].secondsInPeriod,
      });

      console.log(`Created chedule with ID - ${await tokensVesting.scheduleId()}`);
    }

    for (let i = 0; i < config.systemTokenData.vestings.length; i++) {
      await tokensVesting.createVesting({
        vestingToken: isZeroAddr(config.systemTokenData.vestings[i].vestingToken)
          ? nerifToken.address
          : config.systemTokenData.vestings[i].vestingToken,
        beneficiary: config.systemTokenData.vestings[i].beneficiary,
        scheduleId: config.systemTokenData.vestings[i].scheduleId,
        vestingAmount: config.systemTokenData.vestings[i].vestingAmount,
        vestingStartTime: config.systemTokenData.vestings[i].vestingStartTime,
        paidAmount: 0,
      });

      console.log(`Created vesting with ID - ${await tokensVesting.vestingId()}`);
    }
  }
};
