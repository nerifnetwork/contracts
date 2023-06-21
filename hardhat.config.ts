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

const mainnets = {
  ethereum: {
    chainId: 1,
    url: 'https://ethereum-mainnet.core.chainstack.com/fb4dbc0292ed03c29997c2013ce7e37f',
    accounts,
  },
  polygon: {
    chainId: 137,
    url: 'https://nd-599-387-994.p2pify.com/023ce1a632fbefebcbe0547357917c1c',
    accounts
  },
  bsc: {
    chainId: 56,
    url: 'https://nd-570-006-784.p2pify.com/07b6b4ba1e2bc91374788a14fb46cdfa',
    accounts
  },
  gnosis: {
    chainId: 100,
    url: 'https://nd-044-823-524.p2pify.com/3d355b18c32dfadcb5670b019615379e',
    accounts
  }
}

const testnets = {
  mumbai: {
    chainId: 80001,
    url: 'https://rpc-mumbai.maticvigil.com',
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
    url: 'https://rpc.goerli.linea.build',
    accounts,
  },
}

const config: HardhatUserConfig = {
  solidity: "0.8.18",
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
      goerli: process.env.ETHERGOERLISCAN_API_KEY || '',
      "bsc-testnet": process.env.BSCTESTNETSCAN_API_KEY || '',
      "gnosis-chiado": process.env.GNOSISCHIADOSCAN_API_KEY || '',
      "linea-testnet": process.env.LINEATESTNETSCAN_API_KEY || '',

      // Mainnets
      polygon: process.env.POLYGONSCAN_API_KEY || '',
      ethereum: process.env.ETHERSCAN_API_KEY || '',
      bsc: process.env.BSCSCAN_API_KEY || '',
      gnosis: process.env.GNOSISSCAN_API_KEY || '',
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
      }
    ]
  },
};

export default config;
