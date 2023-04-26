import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import { deploySystemContracts, SystemDeploymentOptions, SystemDeploymentResult } from '../../scripts/deploy/mainchain';
import { BridgeDeploymentOptions, BridgeDeploymentResult, deployBridgeContracts } from '../../scripts/deploy/sidechain';
import { MockMintableBurnableToken, MockBridgeApp, MockToken, MockDEXRouter } from '../../typechain';

export async function deploySystem(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  return await deploySystemContracts(options);
}

export async function deployBridge(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  return await deployBridgeContracts(options);
}

export async function deployBridgeWithMocks(
  options?: BridgeDeploymentOptions
): Promise<BridgeWithMocksDeploymentResult> {
  const chainId = BigNumber.from(853);

  const deployment = await deployBridgeContracts(options);

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Mock Token', 'MOCK', BigNumber.from('10000000000000000000000000'));
  await mockToken.deployed();

  const MockMintableBurnableToken = await ethers.getContractFactory('MockMintableBurnableToken');
  const mockMintableBurnableToken = await MockMintableBurnableToken.deploy('Mintable Mock Token', 'MINT');
  await mockMintableBurnableToken.deployed();

  const MockBridgeApp = await ethers.getContractFactory('MockBridgeApp');
  const mockBridgeApp = await MockBridgeApp.deploy();
  await mockBridgeApp.deployed();

  await (await mockBridgeApp.initialize(deployment.relayBridge.address)).wait();

  return {
    mockChainId: chainId,
    mockToken: mockToken,
    mockMintableBurnableToken: mockMintableBurnableToken,
    mockBridgeApp: mockBridgeApp,
    ...deployment,
  };
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

export interface BridgeWithMocksDeploymentResult extends BridgeDeploymentResult {
  mockChainId: BigNumber;
  mockToken: MockToken;
  mockMintableBurnableToken: MockMintableBurnableToken;
  mockBridgeApp: MockBridgeApp;
}

export interface SystemWithMocksDeploymentResult extends SystemDeploymentResult {
  mockToken: MockToken;
  mockDEXRouter: MockDEXRouter;
}
