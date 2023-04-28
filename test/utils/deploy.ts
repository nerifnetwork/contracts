import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import { deploySystemContracts, SystemDeploymentOptions, SystemDeploymentResult } from '../../scripts/deploy/mainchain';
import { BridgeDeploymentOptions, BridgeDeploymentResult, deployBridgeContracts } from '../../scripts/deploy/sidechain';
import { MockToken, MockDEXRouter } from '../../typechain';

export async function deploySystem(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  return await deploySystemContracts(options);
}

export async function deployBridge(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  return await deployBridgeContracts(options);
}

export async function deploySystemWithMocks(
  options?: SystemDeploymentOptions
): Promise<SystemWithMocksDeploymentResult> {
  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Mock Token', 'MOCK', BigNumber.from('10000000000000000000000000'));
  await mockToken.deployed();

  const MockDEXRouter = await ethers.getContractFactory('MockDEXRouter');
  const mockDEXRouter = await MockDEXRouter.deploy();
  await mockDEXRouter.deployed();

  await mockToken.approve(mockDEXRouter.address, BigNumber.from('10000000000000000000000'));

  options = { router: mockDEXRouter.address };
  const deployment = await deploySystemContracts(options);

  return {
    mockToken: mockToken,
    mockDEXRouter: mockDEXRouter,
    ...deployment,
  };
}

export interface SystemWithMocksDeploymentResult extends SystemDeploymentResult {
  mockToken: MockToken;
  mockDEXRouter: MockDEXRouter;
}
