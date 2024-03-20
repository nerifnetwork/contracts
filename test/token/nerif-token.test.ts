import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { ContractsRegistry, NerifToken, SignerStorage } from '../../generated-types/ethers';
import { wei } from '../helpers/utils';
import { setTime } from '../helpers/block-helper';

const { signPermit } = require('../helpers/signatures.js');

describe('NerifToken', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let VESTING_CONTRACT: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;
  let nerifToken: NerifToken;
  let signerStorage: SignerStorage;

  const OWNER_PK: string = 'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';

  const nerifTokenName = 'Nerif Token';
  const nerifTokenSymbol = 'NRF';
  const tokensAmount = wei('1000');

  function createPermitSig(
    tokenName: string,
    tokenAddr: string,
    spenderAddr: string,
    deadline: string | number = 0,
    value: string | number = tokensAmount.toString(),
    ownerAddr: string = OWNER.address,
    nonce: string | number = 0
  ) {
    const domain = {
      name: tokenName,
      version: '1',
      verifyingContract: tokenAddr,
      chainId: '1',
    };

    const message = {
      owner: ownerAddr,
      spender: spenderAddr,
      value: value,
      nonce: nonce,
      deadline: deadline,
    };

    const buffer = Buffer.from(OWNER_PK, 'hex');

    return signPermit(domain, message, buffer);
  }

  before(async () => {
    [OWNER, FIRST, VESTING_CONTRACT, SIGNER] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const nerifTokenImpl = await NerifTokenFactory.deploy();

    signerStorage = await SignerStorageFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);

    await contractsRegistry.addContract(await contractsRegistry.TOKENS_VESTING_NAME(), VESTING_CONTRACT.address);

    nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());

    await nerifToken.initialize(tokensAmount, nerifTokenName, nerifTokenSymbol);
    await signerStorage.initialize(SIGNER.address);

    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after init', async () => {
      expect(await nerifToken.balanceOf(OWNER.address)).to.be.eq(tokensAmount);
      expect(await nerifToken.name()).to.be.eq(nerifTokenName);
      expect(await nerifToken.symbol()).to.be.eq(nerifTokenSymbol);
      expect(await nerifToken.owner()).to.be.eq(OWNER.address);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(nerifToken.initialize(tokensAmount, nerifTokenName, nerifTokenSymbol)).to.be.revertedWith(reason);
    });
  });

  describe('vestingMint', () => {
    it('should correctly mint tokens from the vesting', async () => {
      const tx = await nerifToken.connect(VESTING_CONTRACT).vestingMint(FIRST.address, tokensAmount);

      await expect(tx).to.changeTokenBalance(nerifToken, FIRST, tokensAmount);
      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(tokensAmount);
    });

    it('should get exception if not a vesting contract try to call this function', async () => {
      const reason = 'NerifToken: Not a tokens vesting';

      await expect(nerifToken.vestingMint(OWNER.address, tokensAmount)).to.be.revertedWith(reason);
    });
  });

  describe('ownerMint', () => {
    it('should correctly mint tokens from the owner address', async () => {
      const tx = await nerifToken.ownerMint(FIRST.address, tokensAmount);

      await expect(tx).to.changeTokenBalance(nerifToken, FIRST, tokensAmount);
      expect(await nerifToken.balanceOf(FIRST.address)).to.be.eq(tokensAmount);
    });

    it('should get exception if not a contract owner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(nerifToken.connect(FIRST).ownerMint(FIRST.address, tokensAmount)).to.be.revertedWith(reason);
    });
  });

  describe('permit', () => {
    const amountToApprove = wei('150');
    const currentTime = wei('100000', 1);
    const deadline = currentTime.add(1000);

    it('should correctly call permit function', async () => {
      await setTime(currentTime.toNumber());

      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        FIRST.address,
        deadline.toString(),
        amountToApprove.toString()
      );

      expect(await nerifToken.allowance(OWNER.address, FIRST.address)).to.be.eq(0);

      await nerifToken.permit(OWNER.address, FIRST.address, amountToApprove, deadline, sig.v, sig.r, sig.s);

      expect(await nerifToken.allowance(OWNER.address, FIRST.address)).to.be.eq(amountToApprove);
    });
  });
});
