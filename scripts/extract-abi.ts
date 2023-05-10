import fs from 'fs';
import path from 'path';

async function main() {
  await extractAbi('artifacts/contracts', 'artifacts/abi', 'artifacts/bin');
}

async function extractAbi(from: string, toABI: string, toBIN: string): Promise<void> {
  if (!(await fileExists(toABI))) {
    await fs.promises.mkdir(toABI);
  }

  if (!(await fileExists(toBIN))) {
    await fs.promises.mkdir(toBIN);
  }

  const items = await fs.promises.readdir(from);
  for (const item of items) {
    const stats = await fs.promises.lstat(path.join(from, item));
    if (stats.isDirectory()) {
      await extractAbi(path.join(from, item), path.join(toABI, item), path.join(toBIN, item));
    } else {
      const data = JSON.parse((await fs.promises.readFile(path.join(from, item))).toString());
      if (data.abi != undefined) {
        await fs.promises.writeFile(path.join(toABI, item), JSON.stringify(data.abi));
      }

      if (data.bytecode != undefined) {
        await fs.promises.writeFile(path.join(toBIN, item.split('.')[0]+'.bin'), data.bytecode);
      }
    }
  }
}

async function fileExists(file: string): Promise<boolean> {
  return fs.promises
    .access(file, fs.constants.F_OK)
    .then(() => true)
    .catch(() => false);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
