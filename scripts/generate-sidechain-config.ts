import { readContractsConfig, createSidechainConfig, writeSidechainConfig } from './deploy/sidechain-config';
import { ContractsConfig } from './deploy/config';
import glob from 'glob';

async function main() {
  const files = await getContractsFiles();

  var configs: ContractsConfig[] = [];
  for (const path of files) {
    const config = await readContractsConfig(path);

    configs.push(config);
  }

  const bConfig = await createSidechainConfig(configs)

  await writeSidechainConfig(bConfig)
}

async function getContractsFiles(): Promise<string[]> {
  return new Promise((resolve, reject) => {
    glob('contracts-*.json', async function (err, files) {
      if (err) {
        reject(err);
      } else {
        resolve(files);
      }
    });
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
