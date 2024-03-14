import * as fs from 'fs';
import { BigNumber, BigNumberish } from 'ethers';

import * as dotenv from 'dotenv';
import { ethers } from 'hardhat';
dotenv.config();

export type Config = {
  isMainChain: boolean;

  systemContractsInitParams?: SystemContractsInitParams;
  operationalContractsInitParams: OperationalContractsInitParams;

  schedules: ScheduleData[];
  vestings: VestingData[];

  stakingKeys: string[];
};

export type SystemContractsInitParams = {
  dkgDeadlinePeriod: BigNumber;

  stakingInitParams: StakingInitParams;
  slashingVotingInitParams: SlashingVotingInitParams;
  nerifTokenInitParams: NerifTokenInitParams;
};

export type OperationalContractsInitParams = {
  maxWorkflowsPerAccount: BigNumber;
  nativeDepositAssetData: NativeDepositAssetData;
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

export type NativeDepositAssetData = {
  nativeDepositAssetKey: string;
  workflowExecutionDiscount: BigNumber;
  isEnabled: boolean;
};

export type ScheduleData = {
  secondsInPeriod: BigNumberish;
  durationInPeriods: BigNumberish;
  cliffInPeriods: BigNumberish;
};

export type VestingData = {
  vestingToken: string;
  beneficiary: string;
  vestingStartTime: BigNumberish;
  vestingAmount: BigNumberish;
  scheduleId: BigNumberish;
};

export type NerifTokenInitParams = {
  tokenName: string;
  tokenSymbol: string;
  tokenInitAmount: BigNumberish;
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
  nonUndefinedField(config.schedules, 'schedules');
  nonUndefinedField(config.vestings, 'vestings');

  config.schedules.forEach((el) => validateScheduleData(el));
  config.vestings.forEach((el) => validateVestingData(el));

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
  validateNerifTokenInitParams(systemContractsInitParams.nerifTokenInitParams);
}

function validateOperationalContractsInitParams(
  isMainChain: boolean,
  operationalContractsInitParams: OperationalContractsInitParams
) {
  nonEmptyField(operationalContractsInitParams.maxWorkflowsPerAccount, 'maxWorkflowsPerAccount');
  nonUndefinedField(operationalContractsInitParams.nativeDepositAssetData, 'nativeDepositAssetData');

  validateNativeDepositAssetData(operationalContractsInitParams.nativeDepositAssetData);

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

function validateNerifTokenInitParams(nerifTokenInitParams: NerifTokenInitParams) {
  nonEmptyField(nerifTokenInitParams.tokenName, 'tokenName');
  nonEmptyField(nerifTokenInitParams.tokenSymbol, 'tokenSymbol');
  nonEmptyField(nerifTokenInitParams.tokenInitAmount, 'tokenInitAmount');
}

function validateNativeDepositAssetData(nativeDepositAssetData: NativeDepositAssetData) {
  nonEmptyField(nativeDepositAssetData.nativeDepositAssetKey, 'nativeDepositAssetKey');
  nonEmptyField(nativeDepositAssetData.workflowExecutionDiscount, 'workflowExecutionDiscount');
  nonEmptyField(nativeDepositAssetData.isEnabled, 'isEnabled');
}

function validateScheduleData(scheduleData: ScheduleData) {
  nonEmptyField(scheduleData.secondsInPeriod, 'secondsInPeriod');
  nonEmptyField(scheduleData.durationInPeriods, 'durationInPeriods');
  nonEmptyField(scheduleData.cliffInPeriods, 'cliffInPeriods');
}

function validateVestingData(vestingData: VestingData) {
  nonZeroAddr(vestingData.beneficiary, 'beneficiary');
  nonUndefinedField(vestingData.vestingToken, 'vestingToken');
  nonEmptyField(vestingData.scheduleId, 'scheduleId');
  nonEmptyField(vestingData.vestingAmount, 'vestingAmount');
  nonEmptyField(vestingData.vestingStartTime, 'vestingStartTime');
}

function nonZeroAddr(fieldValue: any, fieldName: string) {
  if (isZeroAddr(fieldValue)) {
    throw new Error(`Invalid ${fieldName} address in config`);
  }
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

export function isZeroAddr(fieldValue: any) {
  return isEmptyField(fieldValue) || fieldValue == ethers.constants.AddressZero;
}

export function isEmptyField(fieldValue: any) {
  return isUndefined(fieldValue) || fieldValue == '';
}

export function isUndefined(fieldValue: any) {
  return fieldValue === undefined;
}
