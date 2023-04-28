import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import {
  deployMainchainContracts,
  MainchainDeploymentOptions,
  MainchainDeploymentResult,
} from '../../scripts/deploy/mainchain';
import {
  SidechainDeploymentOptions,
  SidechainDeploymentResult,
  deploySidechainContracts,
} from '../../scripts/deploy/sidechain';
import { MockToken, MockDEXRouter } from '../../typechain';

export async function deployMainchain(options?: MainchainDeploymentOptions): Promise<MainchainDeploymentResult> {
  return await deployMainchainContracts(options);
}

export async function deploySidechain(options?: SidechainDeploymentOptions): Promise<SidechainDeploymentResult> {
  return await deploySidechainContracts(options);
}

export async function deployMainchainWithMocks(
  options?: MainchainDeploymentOptions
): Promise<SystemWithMocksDeploymentResult> {
  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Mock Token', 'MOCK', BigNumber.from('10000000000000000000000000'));
  await mockToken.deployed();

  const MockDEXRouter = await ethers.getContractFactory('MockDEXRouter');
  const mockDEXRouter = await MockDEXRouter.deploy();
  await mockDEXRouter.deployed();

  await mockToken.approve(mockDEXRouter.address, BigNumber.from('10000000000000000000000'));

  options = { router: mockDEXRouter.address };
  const deployment = await deployMainchainContracts(options);

  return {
    mockToken: mockToken,
    mockDEXRouter: mockDEXRouter,
    ...deployment,
  };
}

export interface SystemWithMocksDeploymentResult extends MainchainDeploymentResult {
  mockToken: MockToken;
  mockDEXRouter: MockDEXRouter;
}
