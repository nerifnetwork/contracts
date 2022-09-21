import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.5.16",
      },
      {
        version: "0.8.17",
      }
    ]
  },
};

export default config;
