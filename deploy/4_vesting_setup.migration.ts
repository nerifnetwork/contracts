import { Deployer } from '@solarity/hardhat-migrate';

import { NerifToken__factory, TokensVesting__factory } from '../generated-types/ethers';
import { isZeroAddr, parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  if (config.isMainChain) {
    const nerifToken = await deployer.deployed(NerifToken__factory);
    const tokensVesting = await deployer.deployed(TokensVesting__factory);

    for (let i = 0; i < config.schedules.length; i++) {
      await tokensVesting.createSchedule({
        cliffInPeriods: config.schedules[i].cliffInPeriods,
        durationInPeriods: config.schedules[i].durationInPeriods,
        secondsInPeriod: config.schedules[i].secondsInPeriod,
      });
      
      console.log(`Created chedule with ID - ${await tokensVesting.scheduleId()}`);

      console.log(await tokensVesting.getSchedule(1));
    }

    for (let i = 0; i < config.vestings.length; i++) {
      await tokensVesting.createVesting({
        vestingToken: isZeroAddr(config.vestings[i].vestingToken) ? nerifToken.address : config.vestings[i].vestingToken,
        beneficiary: config.vestings[i].beneficiary,
        scheduleId: config.vestings[i].scheduleId,
        vestingAmount: config.vestings[i].vestingAmount,
        vestingStartTime: config.vestings[i].vestingStartTime,
        paidAmount: 0,
      });

      console.log(`Created vesting with ID - ${await tokensVesting.vestingId()}`);

      console.log(await tokensVesting.getVesting(1));
    }
  }
};
