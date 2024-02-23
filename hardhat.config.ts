import * as dotenv from "dotenv";

import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@nomiclabs/hardhat-solhint";
import "@nomicfoundation/hardhat-chai-matchers";
import "@typechain/hardhat";
import "@solarity/hardhat-gobind";
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

const mainnets = {
  ethereum: {
    chainId: 1,
    url: `https://mainnet.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
  },
  polygon: {
    chainId: 137,
    url: `https://polygon-mainnet.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts
  },
  zkevm: {
    chainId: 1101,
    url: `https://zkevm-rpc.com`,
    accounts
  },
  bsc: {
    chainId: 56,
    url: 'https://rpc.ankr.com/bsc',
    accounts
  },
  gnosis: {
    chainId: 100,
    url: 'https://rpc.gnosischain.com',
    accounts
  },
  linea: {
    chainId: 59144,
    url: `https://linea-mainnet.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
  },
}

const testnets = {
  mumbai: {
    chainId: 80001,
    url: `https://polygon-mumbai.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
  },
  'zkevm-testnet': {
    chainId: 1442,
    url: `https://rpc.public.zkevm-test.net`,
    accounts,
  },
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
  'gnosis-chiado': {
    chainId: 10200,
    url: 'https://rpc.chiadochain.net',
    gasPrice: 1000000000,
    accounts,
  },
  'linea-testnet': {
    chainId: 59140,
    url: `https://linea-goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
    accounts,
  },
}

const config: HardhatUserConfig = {
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
  networks: {
    ...mainnets,
    ...testnets,
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: {
      // Testnets
      mumbai: process.env.POLYGONMUMBAISCAN_API_KEY || '',
      "zkevm-testnet": process.env.ZKEVMTESTNETSCAN_API_KEY || '',
      goerli: process.env.ETHERGOERLISCAN_API_KEY || '',
      "bsc-testnet": process.env.BSCTESTNETSCAN_API_KEY || '',
      "gnosis-chiado": process.env.GNOSISCHIADOSCAN_API_KEY || '',
      "linea-testnet": process.env.LINEATESTNETSCAN_API_KEY || '',

      // Mainnets
      polygon: process.env.POLYGONSCAN_API_KEY || '',
      zkevm: process.env.ZKEVMSCAN_API_KEY || '',
      ethereum: process.env.ETHERSCAN_API_KEY || '',
      bsc: process.env.BSCSCAN_API_KEY || '',
      gnosis: process.env.GNOSISSCAN_API_KEY || '',
      linea: process.env.LINEASCAN_API_KEY || '',
    },
    customChains: [
      // Testnets
      {
        network: "mumbai",
        chainId: 80001,
        urls: {
          apiURL: "https://api-testnet.polygonscan.com/api",
          browserURL: "https://mumbai.polygonscan.com"
        }
      },
      {
        network: "zkevm-testnet",
        chainId: 1442,
        urls: {
          apiURL: "https://api-testnet-zkevm.polygonscan.com/api",
          browserURL: "https://testnet-zkevm.polygonscan.com"
        }
      },
      {
        network: "goerli",
        chainId: 5,
        urls: {
          apiURL: "https://api-goerli.etherscan.io/api",
          browserURL: "https://goerli.etherscan.io"
        }
      },
      {
        network: "bsc-testnet",
        chainId: 97,
        urls: {
          apiURL: "https://api-testnet.bscscan.com/api",
          browserURL: "https://testnet.bscscan.com"
        }
      },
      {
        network: "gnosis-chiado",
        chainId: 10200,
        urls: {
          apiURL: "https://blockscout.com/gnosis/chiado/api",
          browserURL: "https://blockscout.com/gnosis/chiado"
        }
      },
      {
        network: "linea-testnet",
        chainId: 59140,
        urls: {
          apiURL: "https://api-goerli.lineascan.build/api",
          browserURL: "https://goerli.lineascan.build"
        }
      },

      // Mainnets
      {
        network: "polygon",
        chainId: 137,
        urls: {
          apiURL: "https://api.polygonscan.com/api",
          browserURL: "https://polygonscan.com"
        }
      },
      {
        network: "zkevm",
        chainId: 1101,
        urls: {
          apiURL: "https://api-zkevm.polygonscan.com/api",
          browserURL: "https://zkevm.polygonscan.com"
        }
      },
      {
        network: "ethereum",
        chainId: 1,
        urls: {
          apiURL: "https://api.etherscan.io/api",
          browserURL: "https://etherscan.io"
        }
      },
      {
        network: "bsc",
        chainId: 56,
        urls: {
          apiURL: "https://api.bscscan.com/api",
          browserURL: "https://bscscan.com"
        }
      },
      {
        network: "gnosis",
        chainId: 100,
        urls: {
          apiURL: "https://api.gnosisscan.io/api",
          browserURL: "https://gnosisscan.io"
        }
      },
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
  gobind: {
    outdir: "./pkg",
    deployable: true,
    runOnCompile: false,
    verbose: false,
    onlyFiles: ["./contracts"],
    skipFiles: ["./contracts/test", "./contracts/interfaces"],
  },
};

export default config;
