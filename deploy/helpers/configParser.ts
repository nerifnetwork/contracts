import * as fs from 'fs';
import { BigNumber } from 'ethers';

import * as dotenv from 'dotenv';
import { ethers } from 'hardhat';
dotenv.config();

export type Config = {
  isMainChain: boolean;

  systemContractsInitParams?: SystemContractsInitParams;
  operationalContractsInitParams: OperationalContractsInitParams;

  stakingKeys: string[];
};

export type SystemContractsInitParams = {
  dkgDeadlinePeriod: BigNumber;

  stakingInitParams: StakingInitParams;
  slashingVotingInitParams: SlashingVotingInitParams;
};

export type OperationalContractsInitParams = {
  maxWorkflowsPerAccount: BigNumber;
  signer?: string;
};

export type StakingInitParams = {
  minimalStake: BigNumber;
  withdrawalPeriod: BigNumber;
};

export type SlashingVotingInitParams = {
  epochPeriod: BigNumber;
  slashingThresold: BigNumber;
  slashingEpochs: BigNumber;
};

export function parseConfig(configPath: string = ''): Config {
  configPath = getConfigPath(configPath);

  const config: Config = JSON.parse(fs.readFileSync(configPath, 'utf-8')) as Config;

  nonUndefinedField(config.isMainChain, 'isMainChain');

  if (config.isMainChain) {
    if (config.systemContractsInitParams) {
      validateSystemContractsInitParams(config.systemContractsInitParams);
    } else {
      throw new Error(`Undefined systemContractsInitParams config field`);
    }
  }

  validateOperationalContractsInitParams(config.isMainChain, config.operationalContractsInitParams);

  nonUndefinedField(config.stakingKeys, 'stakingKeys');

  return config;
}

function getConfigPath(configPath: string): string {
  const envPath = process.env.DEPLOY_CONFIG_PATH;
  const defaultPath: string = 'data/config.json';

  if (!isEmptyField(configPath)) {
    return configPath;
  }

  if (envPath && !isEmptyField(envPath)) {
    return envPath;
  }

  return defaultPath;
}

function validateSystemContractsInitParams(systemContractsInitParams: SystemContractsInitParams) {
  nonEmptyField(systemContractsInitParams.dkgDeadlinePeriod, 'dkgDeadlinePeriod');

  validateStakingInitParams(systemContractsInitParams.stakingInitParams);
  validateSlashingVotingInitParams(systemContractsInitParams.slashingVotingInitParams);
}

function validateOperationalContractsInitParams(
  isMainChain: boolean,
  operationalContractsInitParams: OperationalContractsInitParams
) {
  nonEmptyField(operationalContractsInitParams.maxWorkflowsPerAccount, 'maxWorkflowsPerAccount');

  if (
    !isMainChain &&
    (isEmptyField(operationalContractsInitParams.signer) ||
      operationalContractsInitParams.signer == ethers.constants.AddressZero)
  ) {
    throw new Error(`Undefined signer config field`);
  }
}

function validateStakingInitParams(stakingInitParams: StakingInitParams) {
  nonEmptyField(stakingInitParams.minimalStake, 'minimalStake');
  nonEmptyField(stakingInitParams.withdrawalPeriod, 'withdrawalPeriod');
}

function validateSlashingVotingInitParams(slashingVotingInitParams: SlashingVotingInitParams) {
  nonEmptyField(slashingVotingInitParams.epochPeriod, 'epochPeriod');
  nonEmptyField(slashingVotingInitParams.slashingThresold, 'slashingThresold');
  nonEmptyField(slashingVotingInitParams.slashingEpochs, 'slashingEpochs');
}

function nonEmptyField(fieldValue: any, fieldName: string) {
  if (isEmptyField(fieldValue)) {
    throw new Error(`Empty ${fieldName} config field`);
  }
}

function nonUndefinedField(fieldValue: any, fieldName: string) {
  if (isUndefined(fieldValue)) {
    throw new Error(`Undefined ${fieldName} config field`);
  }
}

function isEmptyField(fieldValue: any) {
  return isUndefined(fieldValue) || fieldValue == '';
}

function isUndefined(fieldValue: any) {
  return fieldValue === undefined;
}
