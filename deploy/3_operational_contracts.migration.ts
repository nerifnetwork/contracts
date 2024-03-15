import { Deployer } from '@solarity/hardhat-migrate';

import {
  BillingManager__factory,
  DKG__factory,
  ERC1967Proxy__factory,
  GatewayFactory__factory,
  Gateway__factory,
  Registry__factory,
  SignerStorage,
  SignerStorage__factory,
} from '../generated-types/ethers';
import { parseConfig } from './helpers/configParser';

export = async (deployer: Deployer) => {
  const config = parseConfig();

  const billingManagerImpl = await deployer.deploy(BillingManager__factory, {name: "BillingManagerImpl"});
  const registryImpl = await deployer.deploy(Registry__factory, {name: "RegistryImpl"});
  await deployer.deploy(Gateway__factory, {name: "GatewayImpl"});

  const billingManagerProxy = await deployer.deploy(ERC1967Proxy__factory, [billingManagerImpl.address, "0x"], {name: "BillingManagerProxy"});
  const registryProxy = await deployer.deploy(ERC1967Proxy__factory, [registryImpl.address, "0x"], {name: "RegistryProxy"});

  await deployer.save(BillingManager__factory, billingManagerProxy.address);
  await deployer.save(Registry__factory, registryProxy.address);

  await deployer.deploy(GatewayFactory__factory);

  let signerStorage: SignerStorage;

  if (config.isMainChain) {
    const dkgAddress = (await deployer.deployed(DKG__factory)).address;

    signerStorage = await deployer.deployed(SignerStorage__factory, dkgAddress);

    await deployer.save(SignerStorage__factory, dkgAddress);
  } else if (config.operationalContractsInitParams.signer) {
    signerStorage = await deployer.deploy(SignerStorage__factory);

    await signerStorage.initialize(config.operationalContractsInitParams.signer);
  } else {
    throw new Error("Undefined signer address");
  }
};
