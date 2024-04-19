import * as fs from 'fs';
import { BigNumber, BigNumberish } from 'ethers';

import * as dotenv from 'dotenv';
import { ethers } from 'hardhat';
dotenv.config();

export type ConfigData = {
  isMainChain: boolean;

  nerifTokenData?: NerifTokenData;
  systemContractsInitParams?: SystemContractsInitParams;
  operationalContractsInitParams?: OperationalContractsInitParams;
};

export type NerifTokenData = {
  nerifTokenInitParams: NerifTokenInitParams;
  schedules: ScheduleData[];
  vestings: VestingData[];
};

export type SystemContractsInitParams = {
  dkgInitParams: DKGInitParams;
  stakingInitParams: StakingInitParams;
  slashingVotingInitParams: SlashingVotingInitParams;

  stakingKeys: string[];
};

export type OperationalContractsInitParams = {
  maxWorkflowsPerAccount: BigNumber;
  nativeDepositAssetData: NativeDepositAssetData;
  signer?: string;
};

export type NerifTokenInitParams = {
  tokenName: string;
  tokenSymbol: string;
  tokenInitAmount: BigNumberish;
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

export type DKGInitParams = {
  updatesCollectionEpochDuration: BigNumberish;
  dkgGenerationEpochDuration: BigNumberish;
  guaranteedWorkingEpochDuration: BigNumberish;
};

export type StakingInitParams = {
  stakingTokenAddr: string;
  minimalStake: BigNumber;
  whitelistedUsers: string[];
};

export type SlashingVotingInitParams = {
  slashingThresoldPercentage: BigNumber;
};

export type NativeDepositAssetData = {
  nativeDepositAssetKey: string;
  workflowExecutionDiscount: BigNumber;
  isEnabled: boolean;
};

export class ConfigParser {
  public configData: ConfigData;

  constructor(pathToConfig: string = '') {
    this.configData = this.parseConfig(pathToConfig);

    this.validateConfig();
  }

  public needToDeployNerifToken(): boolean {
    return !isUndefined(this.configData.nerifTokenData);
  }

  public needToDeploySystemContracts(): boolean {
    return !isUndefined(this.configData.systemContractsInitParams);
  }

  public needToDeployOperationalContracts(): boolean {
    return !isUndefined(this.configData.operationalContractsInitParams);
  }

  protected parseConfig(pathToConfig: string): ConfigData {
    pathToConfig = this.getConfigPath(pathToConfig);

    return JSON.parse(fs.readFileSync(pathToConfig, 'utf-8')) as ConfigData;
  }

  protected validateConfig() {
    this.nonUndefinedField(this.configData.isMainChain, 'isMainChain');

    if (this.configData.nerifTokenData) {
      this.validateNerifTokenData(this.configData.nerifTokenData);
    }

    if (this.configData.systemContractsInitParams) {
      this.validateSystemContractsInitParams(this.configData.systemContractsInitParams);
    }

    if (this.configData.operationalContractsInitParams) {
      this.validateOperationalContractsInitParams(this.configData.operationalContractsInitParams);
    }
  }

  protected validateNerifTokenData(nerifTokenData: NerifTokenData) {
    this.nonUndefinedField(nerifTokenData.schedules, 'schedules');
    this.nonUndefinedField(nerifTokenData.vestings, 'vestings');
    this.nonUndefinedField(nerifTokenData.nerifTokenInitParams, 'nerifTokenInitParams');

    this.validateNerifTokenInitParams(nerifTokenData.nerifTokenInitParams);

    nerifTokenData.schedules.forEach((el) => this.validateScheduleData(el));
    nerifTokenData.vestings.forEach((el) => this.validateVestingData(el));
  }

  protected validateSystemContractsInitParams(systemContractsInitParams: SystemContractsInitParams) {
    this.nonUndefinedField(systemContractsInitParams.stakingKeys, 'stakingKeys');

    this.validateDKGInitParams(systemContractsInitParams.dkgInitParams);
    this.validateStakingInitParams(systemContractsInitParams.stakingInitParams);
    this.validateSlashingVotingInitParams(systemContractsInitParams.slashingVotingInitParams);
  }

  protected validateOperationalContractsInitParams(operationalContractsInitParams: OperationalContractsInitParams) {
    this.nonEmptyField(operationalContractsInitParams.maxWorkflowsPerAccount, 'maxWorkflowsPerAccount');
    this.nonUndefinedField(operationalContractsInitParams.nativeDepositAssetData, 'nativeDepositAssetData');

    this.validateNativeDepositAssetData(operationalContractsInitParams.nativeDepositAssetData);

    if (!this.configData.isMainChain) {
      this.nonZeroAddr(operationalContractsInitParams.signer, 'signer');
    }
  }

  protected validateNerifTokenInitParams(nerifTokenInitParams: NerifTokenInitParams) {
    this.nonEmptyField(nerifTokenInitParams.tokenName, 'tokenName');
    this.nonEmptyField(nerifTokenInitParams.tokenSymbol, 'tokenSymbol');
    this.nonEmptyField(nerifTokenInitParams.tokenInitAmount, 'tokenInitAmount');
  }

  protected validateScheduleData(scheduleData: ScheduleData) {
    this.nonEmptyField(scheduleData.secondsInPeriod, 'secondsInPeriod');
    this.nonEmptyField(scheduleData.durationInPeriods, 'durationInPeriods');
    this.nonEmptyField(scheduleData.cliffInPeriods, 'cliffInPeriods');
  }

  protected validateVestingData(vestingData: VestingData) {
    this.nonZeroAddr(vestingData.beneficiary, 'beneficiary');
    this.nonUndefinedField(vestingData.vestingToken, 'vestingToken');
    this.nonEmptyField(vestingData.scheduleId, 'scheduleId');
    this.nonEmptyField(vestingData.vestingAmount, 'vestingAmount');
    this.nonEmptyField(vestingData.vestingStartTime, 'vestingStartTime');
  }

  protected validateDKGInitParams(dkgInitParams: DKGInitParams) {
    this.nonEmptyField(dkgInitParams.updatesCollectionEpochDuration, 'updatesCollectionEpochDuration');
    this.nonEmptyField(dkgInitParams.dkgGenerationEpochDuration, 'dkgGenerationEpochDuration');
    this.nonEmptyField(dkgInitParams.guaranteedWorkingEpochDuration, 'guaranteedWorkingEpochDuration');
  }

  protected validateStakingInitParams(stakingInitParams: StakingInitParams) {
    this.nonEmptyField(stakingInitParams.minimalStake, 'minimalStake');
    this.nonUndefinedField(stakingInitParams.whitelistedUsers, 'whitelistedUsers');

    if (!this.needToDeployNerifToken()) {
      this.nonZeroAddr(stakingInitParams.stakingTokenAddr, 'stakingTokenAddr');
    }
  }

  protected validateSlashingVotingInitParams(slashingVotingInitParams: SlashingVotingInitParams) {
    this.nonEmptyField(slashingVotingInitParams.slashingThresoldPercentage, 'slashingThresoldPercentage');
  }

  protected validateNativeDepositAssetData(nativeDepositAssetData: NativeDepositAssetData) {
    this.nonEmptyField(nativeDepositAssetData.nativeDepositAssetKey, 'nativeDepositAssetKey');
    this.nonEmptyField(nativeDepositAssetData.workflowExecutionDiscount, 'workflowExecutionDiscount');
    this.nonEmptyField(nativeDepositAssetData.isEnabled, 'isEnabled');
  }

  protected getConfigPath(configPath: string): string {
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

  protected nonZeroAddr(fieldValue: any, fieldName: string) {
    if (isZeroAddr(fieldValue)) {
      throw new Error(`Invalid ${fieldName} address in config`);
    }
  }

  protected nonEmptyField(fieldValue: any, fieldName: string) {
    if (isEmptyField(fieldValue)) {
      throw new Error(`Empty ${fieldName} config field`);
    }
  }

  protected nonUndefinedField(fieldValue: any, fieldName: string) {
    if (isUndefined(fieldValue)) {
      throw new Error(`Undefined ${fieldName} config field`);
    }
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
