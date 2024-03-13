import "@nomicfoundation/hardhat-chai-matchers";
import "@nomiclabs/hardhat-waffle";
import "@nomiclabs/hardhat-solhint";
import "@solarity/hardhat-gobind";
import "@solarity/hardhat-migrate";
import "@typechain/hardhat";
import "hardhat-contract-sizer";
import "hardhat-gas-reporter";
import "solidity-coverage";
import "tsconfig-paths/register";

import { HardhatUserConfig, task } from "hardhat/config";

import * as dotenv from "dotenv";
dotenv.config();

task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

const accounts = process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : [];

const mainnets = {
  ethereum: {
    url: `https://mainnet.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
    gasMultiplier: 1.2,
  },
  polygon: {
    url: `https://polygon-mainnet.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
    gasMultiplier: 1.2,
  },
  polygonZkEVM: {
    url: `https://zkevm-rpc.com`,
    accounts,
    gasMultiplier: 1.2,
  },
  bsc: {
    url: 'https://rpc.ankr.com/bsc',
    accounts,
    gasMultiplier: 1.2,
  },
  gnosis: {
    url: 'https://rpc.gnosischain.com',
    accounts,
    gasMultiplier: 1.2,
  },
  linea: {
    url: `https://linea-mainnet.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
    gasMultiplier: 1.2,
  },
}

const testnets = {
  mumbai: {
    url: `https://polygon-mumbai.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
    gasMultiplier: 1.2,
  },
  polygonZkEVMTestnet: {
    url: `https://rpc.public.zkevm-test.net`,
    accounts,
    gasMultiplier: 1.2,
  },
  goerli: {
    url: `https://goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
    gasMultiplier: 1.2,
  },
  bscTestnet: {
    url: 'https://data-seed-prebsc-1-s1.binance.org:8545',
    accounts,
    gasMultiplier: 1.2,
  },
  chiado: {
    url: 'https://rpc.chiadochain.net',
    accounts,
    gasMultiplier: 1.2,
  },
  lineaTestnet: {
    url: `https://linea-goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
    gasMultiplier: 1.2,
  },
}

const config: HardhatUserConfig = {
  networks: {
    ...mainnets,
    ...testnets,
    localhost: {
      url: "http://127.0.0.1:8545",
      gasMultiplier: 1.2,
    },
    hardhat: {
      initialDate: "1970-01-01T00:00:00Z",
      chainId: 1,
    },
  },
  solidity: {
    compilers: [
      {
        version: "0.8.18",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
          evmVersion: "paris",
        },
      }
    ]
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: {
      // Testnets
      polygonMumbai: process.env.POLYGONMUMBAISCAN_API_KEY || '',
      polygonZkEVMTestnet: process.env.ZKEVMTESTNETSCAN_API_KEY || '',
      goerli: process.env.ETHERGOERLISCAN_API_KEY || '',
      bscTestnet: process.env.BSCTESTNETSCAN_API_KEY || '',
      chiado: process.env.GNOSISCHIADOSCAN_API_KEY || '',
      lineaTestnet: process.env.LINEATESTNETSCAN_API_KEY || '',

      // Mainnets
      polygon: process.env.POLYGONSCAN_API_KEY || '',
      polygonZkEVM: process.env.ZKEVMSCAN_API_KEY || '',
      ethereum: process.env.ETHERSCAN_API_KEY || '',
      bsc: process.env.BSCSCAN_API_KEY || '',
      gnosis: process.env.GNOSISSCAN_API_KEY || '',
      linea: process.env.LINEASCAN_API_KEY || '',
    },
    customChains: [
      {
        network: "lineaTestnet",
        chainId: 59140,
        urls: {
          apiURL: "https://api-goerli.lineascan.build/api",
          browserURL: "https://goerli.lineascan.build"
        }
      },

      // Mainnets
      {
        network: "linea",
        chainId: 59144,
        urls: {
          apiURL: "https://explorer.linea.build/api",
          browserURL: "https://explorer.linea.build"
        }
      }
    ]
  },
  migrate: {
    pathToMigrations: "./deploy/",
  },
  gobind: {
    outdir: "./pkg",
    deployable: true,
    runOnCompile: false,
    verbose: false,
    onlyFiles: ["./contracts"],
    skipFiles: ["./contracts/test", "./contracts/interfaces"],
  },
  typechain: {
    outDir: "generated-types/ethers",
    target: "ethers-v5",
    alwaysGenerateOverloads: true,
    discriminateTypes: true,
  },
};

export default config;
