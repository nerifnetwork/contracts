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

export async function deployMainchain(options?: MainchainDeploymentOptions): Promise<MainchainDeploymentResult> {
  return await deployMainchainContracts(options);
}

export async function deploySidechain(options?: SidechainDeploymentOptions): Promise<SidechainDeploymentResult> {
  return await deploySidechainContracts(options);
}
