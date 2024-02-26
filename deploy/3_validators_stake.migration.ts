import { Deployer } from "@solarity/hardhat-migrate";
import { ethers } from "hardhat";
import { BigNumber } from "ethers";

import { Staking__factory } from "@/generated-types/ethers";

export = async (deployer: Deployer) => {
  const staking = await deployer.deployed(Staking__factory);

  const minimalStake = 1000;
  const stakingKeys: string[] = [];

  if (stakingKeys.length > 0) {
    for (const privateKey of stakingKeys) {
      const signer = new ethers.Wallet(privateKey, ethers.provider);
  
      await staking.connect(signer).stake({ value: minimalStake});
    }
  };
};
