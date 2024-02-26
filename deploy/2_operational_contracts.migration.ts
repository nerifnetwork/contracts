import { Deployer } from "@solarity/hardhat-migrate";
import { ethers } from "hardhat";
import { BigNumber } from "ethers";

import { BillingManager__factory, GatewayFactory__factory, Gateway__factory, Registry__factory, SignerStorage__factory } from "@/generated-types/ethers";

const defaultSystemDeploymentParameters: MainchainDeploymentParameters = {
  minimalStake: ethers.utils.parseEther('0.01'),
  stakeWithdrawalPeriod: BigNumber.from(60),
  slashingEpochs: BigNumber.from(3),
  slashingEpochPeriod: BigNumber.from(1000),
  slashingBansThresold: BigNumber.from(10),

  dkgDeadlinePeriod: BigNumber.from(10000),

  displayLogs: false,
  verify: false,
  stakingKeys: [],
  isMainChain: true,
};

interface MainchainDeploymentParameters {
  minimalStake: BigNumber;
  stakeWithdrawalPeriod: BigNumber;

  slashingEpochs: BigNumber;
  slashingEpochPeriod: BigNumber;
  slashingBansThresold: BigNumber;

  dkgDeadlinePeriod: BigNumber;

  displayLogs: boolean;
  verify: boolean;
  stakingKeys: string[];
  isMainChain: boolean;
};

export = async (deployer: Deployer) => {
  const billingManager = await deployer.deploy(BillingManager__factory);
  const gatewayImpl = await deployer.deploy(Gateway__factory);
  const gatewayFactory = await deployer.deploy(GatewayFactory__factory);
  const registry = await deployer.deploy(Registry__factory);
  const signerStorage = await deployer.deploy(SignerStorage__factory);

  await billingManager.initialize(registry.address, signerStorage.address);
  await gatewayFactory.initialize(registry.address, gatewayImpl.address);
  await registry.initialize(
    defaultSystemDeploymentParameters.isMainChain,
    signerStorage.address,
    gatewayFactory.address,
    billingManager.address,
    0
  );
};
