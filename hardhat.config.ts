import * as dotenv from "dotenv";

import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@nomiclabs/hardhat-solhint";
import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "solidity-coverage";

dotenv.config();

task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

const accounts = process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : [];

const config: HardhatUserConfig = {
  solidity: "0.8.18",
  networks: {
    goerli: {
      chainId: 5,
      url: `https://goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts,
    },
    'bsc-testnet': {
      chainId: 97,
      url: 'https://data-seed-prebsc-1-s1.binance.org:8545',
      accounts,
    },
    mumbai: {
      chainId: 80001,
      url: 'https://rpc-mumbai.maticvigil.com',
      accounts,
    },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
    customChains: [
      {
        network: "goerli",
        chainId: 5,
        urls: {
          apiURL: "https://goerli.etherscan.io/api",
          browserURL: "https://goerli.etherscan.io"
        }
      },
      {
        network: "bsc-testnet",
        chainId: 97,
        urls: {
          apiURL: "https://bscscan.com/api",
          browserURL: "https://testnet.bscscan.com"
        }
      },
      {
        network: "mumbai",
        chainId: 80001,
        urls: {
          apiURL: "https://mumbai.polygonscan.com/api",
          browserURL: "https://mumbai.polygonscan.com"
        }
      },
    ]
  },
};

export default config;
