import { ethers } from 'ethers';

export function wei(value: string | number | bigint, decimal: number = 18): ethers.BigNumber {
  if (typeof value == 'number' || typeof value == 'bigint') {
    value = value.toString();
  }

  return ethers.utils.parseUnits(value as string, decimal);
}

export function fromWei(value: string | number | bigint, decimal: number = 18): string {
  return (BigInt(value) / 10n ** BigInt(decimal)).toString();
}
