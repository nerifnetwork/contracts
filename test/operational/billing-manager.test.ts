import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { wei } from '../helpers/utils';
import {
  Gateway,
  GatewayFactory,
  Registry,
  BillingManager,
  SignerStorage,
  IBillingManager,
  ContractsRegistry,
  NerifToken,
} from '../../generated-types/ethers';
import { setNextBlockTime, setTime } from '../helpers/block-helper';

const { signPermit, signFundsWithdraw, signRewardsWithdraw } = require('../helpers/signatures.js');

describe('BillingManager', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;
  let STAKING: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;

  let signerStorage: SignerStorage;
  let registry: Registry;
  let billingManager: BillingManager;
  let gatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;

  let nerifToken: NerifToken;

  const OWNER_PK: string = 'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';
  const SIGNER_PK: string = '5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a';

  const defaultWorkflowId: number = 13;
  const tokensAmount = wei('1000');

  const nativeDepositAssetKey: string = 'NATIVE';
  const nerifTokenDepositAssetKey: string = 'NERIF';

  let nativeDepositAssetData: IBillingManager.DepositAssetDataStruct = {
    tokenAddr: ethers.constants.AddressZero,
    workflowExecutionDiscount: 0,
    totalAssetAmount: 0,
    isPermitable: false,
    isEnabled: true,
  };
  let nerifTokenDepositAssetData: IBillingManager.DepositAssetDataStruct;

  function createPermitSig(
    tokenName: string,
    tokenAddr: string,
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
      spender: billingManager.address,
      value: value,
      nonce: nonce,
      deadline: deadline,
    };

    const buffer = Buffer.from(OWNER_PK, 'hex');

    return signPermit(domain, message, buffer);
  }

  function createFundsWithdrawSig(
    userAddr: string,
    depositAssetKeys: string[],
    deadline: string | number = 0,
    withdrawAmounts: string[] | number[] = [tokensAmount.toString()],
    nonce: string | number = 0,
    signerPK: string = SIGNER_PK
  ) {
    const domain = {
      name: 'BillingManager',
      version: '1',
      verifyingContract: billingManager.address,
      chainId: '1',
    };

    const depositAssetsHash = ethers.utils.keccak256(
      ethers.utils.defaultAbiCoder.encode(['string[]'], [depositAssetKeys])
    );
    const amountsHash = ethers.utils.keccak256(ethers.utils.defaultAbiCoder.encode(['uint256[]'], [withdrawAmounts]));

    const message = {
      userAddr: userAddr,
      depositAssetsHash: depositAssetsHash,
      amountsHash: amountsHash,
      nonce: nonce,
      deadline: deadline,
    };

    const buffer = Buffer.from(signerPK, 'hex');

    return signFundsWithdraw(domain, message, buffer);
  }

  function createRewardsWithdrawSig(
    userAddr: string,
    depositAssetKeys: string[],
    deadline: string | number = 0,
    withdrawAmounts: string[] | number[] = [tokensAmount.toString()],
    nonce: string | number = 0,
    signerPK: string = SIGNER_PK
  ) {
    const domain = {
      name: 'BillingManager',
      version: '1',
      verifyingContract: billingManager.address,
      chainId: '1',
    };

    const depositAssetsHash = ethers.utils.keccak256(
      ethers.utils.defaultAbiCoder.encode(['string[]'], [depositAssetKeys])
    );
    const amountsHash = ethers.utils.keccak256(ethers.utils.defaultAbiCoder.encode(['uint256[]'], [withdrawAmounts]));

    const message = {
      userAddr: userAddr,
      depositAssetsHash: depositAssetsHash,
      amountsHash: amountsHash,
      nonce: nonce,
      deadline: deadline,
    };

    const buffer = Buffer.from(signerPK, 'hex');

    return signRewardsWithdraw(domain, message, buffer);
  }

  before(async () => {
    [OWNER, FIRST, SIGNER, STAKING] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory('ERC1967Proxy');
    const ContractsRegistryFactory = await ethers.getContractFactory('ContractsRegistry');
    const RegistryFactory = await ethers.getContractFactory('Registry');
    const BillingManagerFactory = await ethers.getContractFactory('BillingManager');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');
    const NerifTokenFactory = await ethers.getContractFactory('NerifToken');
    const TokensVestingFactory = await ethers.getContractFactory('TokensVesting');

    const contractsRegistryImpl = await ContractsRegistryFactory.deploy();
    const contractsRegistryProxy = await ERC1967ProxyFactory.deploy(contractsRegistryImpl.address, '0x');

    const registryImpl = await RegistryFactory.deploy();
    const billingManagerImpl = await BillingManagerFactory.deploy();
    const nerifTokenImpl = await NerifTokenFactory.deploy();

    const tokensVesting = await TokensVestingFactory.deploy();
    signerStorage = await SignerStorageFactory.deploy();
    gatewayFactory = await GatewayFactoryFactory.deploy();
    gatewayImpl = await GatewayImplFactory.deploy();

    contractsRegistry = ContractsRegistryFactory.attach(contractsRegistryProxy.address);

    await contractsRegistry.__OwnableContractsRegistry_init();

    await contractsRegistry.addProxyContract(await contractsRegistry.REGISTRY_NAME(), registryImpl.address);
    await contractsRegistry.addProxyContract(
      await contractsRegistry.BILLING_MANAGER_NAME(),
      billingManagerImpl.address
    );
    await contractsRegistry.addProxyContract(await contractsRegistry.NERIF_TOKEN_NAME(), nerifTokenImpl.address);

    await contractsRegistry.addContract(await contractsRegistry.GATEWAY_FACTORY_NAME(), gatewayFactory.address);
    await contractsRegistry.addContract(await contractsRegistry.TOKENS_VESTING_NAME(), tokensVesting.address);
    await contractsRegistry.addContract(await contractsRegistry.STAKING_NAME(), STAKING.address);
    await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), signerStorage.address);

    registry = RegistryFactory.attach(await contractsRegistry.getRegistryContract());
    billingManager = BillingManagerFactory.attach(await contractsRegistry.getBillingManagerContract());
    nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());

    nativeDepositAssetData = {
      tokenAddr: ethers.constants.AddressZero,
      workflowExecutionDiscount: 0,
      totalAssetAmount: 0,
      isPermitable: false,
      isEnabled: true,
    };
    nerifTokenDepositAssetData = {
      tokenAddr: nerifToken.address,
      workflowExecutionDiscount: 10,
      totalAssetAmount: 0,
      isPermitable: true,
      isEnabled: true,
    };

    await signerStorage.initialize(SIGNER.address);
    await registry.initialize(0);
    await billingManager.initialize(
      { depositAssetKey: nativeDepositAssetKey, depositAssetData: nativeDepositAssetData },
      { depositAssetKey: nerifTokenDepositAssetKey, depositAssetData: nerifTokenDepositAssetData }
    );
    await gatewayFactory.initialize(gatewayImpl.address);
    await nerifToken.initialize(tokensAmount, 'NERIF', 'NERIF');
    await tokensVesting.initialize();

    await contractsRegistry.injectDependencies(await contractsRegistry.REGISTRY_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.BILLING_MANAGER_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.GATEWAY_FACTORY_NAME());
    await contractsRegistry.injectDependencies(await contractsRegistry.NERIF_TOKEN_NAME());

    await contractsRegistry.setIsMainChain(true);

    await registry.registerWorkflows([
      {
        id: defaultWorkflowId,
        workflowOwner: OWNER.address,
        requireGateway: true,
        deployGateway: true,
      },
    ]);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialize', async () => {
      const nativeAssetInfo = (await billingManager.getDepositAssetsInfo([nativeDepositAssetKey]))[0];

      expect(nativeAssetInfo.depositAssetData.tokenAddr).to.be.eq(ethers.constants.AddressZero);
      expect(nativeAssetInfo.depositAssetData.isEnabled).to.be.eq(true);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(
        billingManager.initialize(
          { depositAssetKey: nativeDepositAssetKey, depositAssetData: nativeDepositAssetData },
          { depositAssetKey: nerifTokenDepositAssetKey, depositAssetData: nerifTokenDepositAssetData }
        )
      ).to.be.revertedWith(reason);
    });
  });

  describe('receive', () => {
    const depositAmount = wei('1');

    it('should correctly receive native currency', async () => {
      const tx = await OWNER.sendTransaction({
        from: OWNER.address,
        to: billingManager.address,
        value: depositAmount,
      });

      await expect(tx)
        .emit(billingManager, 'AssetDeposited')
        .withArgs(nativeDepositAssetKey, OWNER.address, OWNER.address, depositAmount);
      await expect(tx).to.changeEtherBalances([OWNER, billingManager], [depositAmount.mul(-1), depositAmount]);

      expect(await billingManager.getTotalDepositAssetAmount(nativeDepositAssetKey)).to.be.eq(depositAmount);
    });

    it('should get exception if the native deposit asset is disabled', async () => {
      const reason = 'BillingManager: Deposit asset is disabled';

      await billingManager.connect(SIGNER).updateDepositAssetEnabledStatus(nativeDepositAssetKey, false);

      await expect(
        OWNER.sendTransaction({
          from: OWNER.address,
          to: billingManager.address,
          value: depositAmount,
        })
      ).to.be.revertedWith(reason);
    });
  });

  describe('upgradability', () => {
    it('should correctly upgrade BillingManager contract', async () => {
      const TestBillingManagerFactory = await ethers.getContractFactory('TestBillingManager');

      let testBillingManager = TestBillingManagerFactory.attach(billingManager.address);

      await expect(testBillingManager.version()).to.be.revertedWithoutReason();

      const newBillingManagerImpl = await TestBillingManagerFactory.deploy();

      await contractsRegistry.upgradeContract(
        await contractsRegistry.BILLING_MANAGER_NAME(),
        newBillingManagerImpl.address
      );

      expect(await testBillingManager.version()).to.be.eq('v2.0.0');
    });
  });

  describe('setDependencies', () => {
    it('should correctly update dependencies', async () => {
      const TestBillingManagerFactory = await ethers.getContractFactory('TestBillingManager');

      const newBillingManagerImpl = await TestBillingManagerFactory.deploy();
      await contractsRegistry.upgradeContract(
        await contractsRegistry.BILLING_MANAGER_NAME(),
        newBillingManagerImpl.address
      );

      const newBillingManager = TestBillingManagerFactory.attach(billingManager.address);

      expect(await newBillingManager.registry()).to.be.eq(registry.address);

      await contractsRegistry.addContract(await contractsRegistry.REGISTRY_NAME(), FIRST.address);
      await contractsRegistry.injectDependencies(await contractsRegistry.BILLING_MANAGER_NAME());

      expect(await newBillingManager.registry()).to.be.eq(FIRST.address);
    });

    it('should get exception if not a contracts registry try to call this function', async () => {
      const reason = 'Dependant: not an injector';

      await expect(billingManager.setDependencies(contractsRegistry.address, '0x')).to.be.revertedWith(reason);
    });
  });

  describe('addDepositAsset', () => {
    const newNerifAssetKey: string = 'NERIF_V2';
    let newNerifAssetData: IBillingManager.DepositAssetDataStruct;

    beforeEach('setup', async () => {
      newNerifAssetData = {
        tokenAddr: nerifToken.address,
        workflowExecutionDiscount: 20,
        totalAssetAmount: 0,
        isPermitable: true,
        isEnabled: true,
      };
    });

    it('should correctly add new deposit asset key', async () => {
      const tx = await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: newNerifAssetKey,
          depositAssetData: newNerifAssetData,
        },
      ]);

      await expect(tx)
        .emit(billingManager, 'DepositAssetAdded')
        .withArgs(newNerifAssetKey, Object.values(newNerifAssetData));

      expect(await billingManager.isDepositAssetEnabled(newNerifAssetKey)).to.be.eq(true);
      expect(await billingManager.isDepositAssetSupported(newNerifAssetKey)).to.be.eq(true);

      const newNerifAssetInfo = (await billingManager.getDepositAssetsInfo([newNerifAssetKey]))[0];

      expect(newNerifAssetInfo.depositAssetData.tokenAddr).to.be.eq(nerifToken.address);
      expect(newNerifAssetInfo.depositAssetData.workflowExecutionDiscount).to.be.eq(
        newNerifAssetData.workflowExecutionDiscount
      );
      expect(newNerifAssetInfo.depositAssetData.totalAssetAmount).to.be.eq(0);
      expect(newNerifAssetInfo.depositAssetData.isEnabled).to.be.eq(true);

      expect(await billingManager.getSupportedDepositAssetKeys()).to.be.deep.eq([
        nativeDepositAssetKey,
        nerifTokenDepositAssetKey,
        newNerifAssetKey,
      ]);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'BillingManager: Not a signer';

      await expect(
        billingManager.addDepositAssets([
          {
            depositAssetKey: newNerifAssetKey,
            depositAssetData: newNerifAssetData,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset key already exists', async () => {
      const reason = 'BillingManager: Deposit asset already exists';

      await expect(
        billingManager.connect(SIGNER).addDepositAssets([
          {
            depositAssetKey: nativeDepositAssetKey,
            depositAssetData: nativeDepositAssetData,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass invalid token address', async () => {
      const reason = 'BillingManager: Invalid deposit asset token address';

      newNerifAssetData.tokenAddr = ethers.constants.AddressZero;

      await expect(
        billingManager.connect(SIGNER).addDepositAssets([
          {
            depositAssetKey: newNerifAssetKey,
            depositAssetData: newNerifAssetData,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass empty deposit asset key', async () => {
      const reason = 'BillingManager: Invalid deposit asset key';

      await expect(
        billingManager.connect(SIGNER).addDepositAssets([
          {
            depositAssetKey: '',
            depositAssetData: newNerifAssetData,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass not empty total asset amount', async () => {
      const reason = 'BillingManager: Invalid init total asset amount value';

      newNerifAssetData.totalAssetAmount = wei('1');

      await expect(
        billingManager.connect(SIGNER).addDepositAssets([
          {
            depositAssetKey: newNerifAssetKey,
            depositAssetData: newNerifAssetData,
          },
        ])
      ).to.be.revertedWith(reason);
    });
  });

  describe('updateWorkflowExecutionDiscount', () => {
    const newDiscount = 40;

    it('should correctly update workflow execution discount', async () => {
      const tx = await billingManager
        .connect(SIGNER)
        .updateWorkflowExecutionDiscount(nativeDepositAssetKey, newDiscount);

      await expect(tx)
        .emit(billingManager, 'WorkflowExecutionDiscountUpdated')
        .withArgs(nativeDepositAssetKey, newDiscount, 0);

      expect(
        (await billingManager.getDepositAssetsInfo([nativeDepositAssetKey]))[0].depositAssetData
          .workflowExecutionDiscount
      ).to.be.eq(newDiscount);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'BillingManager: Not a signer';

      await expect(
        billingManager.updateWorkflowExecutionDiscount(nativeDepositAssetKey, newDiscount)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset does not exist', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      await expect(
        billingManager.connect(SIGNER).updateWorkflowExecutionDiscount('SOME_ASSET_KEY', newDiscount)
      ).to.be.revertedWith(reason);
    });
  });

  describe('updateDepositAssetEnabledStatus', () => {
    const newEnabledStatus: boolean = false;

    it('should correctly update deposit asset enabled status', async () => {
      const tx = await billingManager
        .connect(SIGNER)
        .updateDepositAssetEnabledStatus(nativeDepositAssetKey, newEnabledStatus);

      await expect(tx)
        .emit(billingManager, 'DepositAssetEnabledStatusUpdated')
        .withArgs(nativeDepositAssetKey, newEnabledStatus);

      expect(await billingManager.isDepositAssetEnabled(nativeDepositAssetKey)).to.be.eq(false);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'BillingManager: Not a signer';

      await expect(
        billingManager.updateDepositAssetEnabledStatus(nativeDepositAssetKey, newEnabledStatus)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset does not exist', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      await expect(
        billingManager.connect(SIGNER).updateDepositAssetEnabledStatus('SOME_ASSET_KEY', newEnabledStatus)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the new enabled status is equal to the current', async () => {
      const reason = 'BillingManager: Invalid new enabled status';

      await expect(
        billingManager.connect(SIGNER).updateDepositAssetEnabledStatus(nativeDepositAssetKey, true)
      ).to.be.revertedWith(reason);
    });
  });

  describe('deposit', () => {
    const depositTokensAmount = wei('100');
    const depositNativeAmount = wei('1');

    it('should correctly deposit native currency', async () => {
      const tx = await billingManager.deposit(nativeDepositAssetKey, OWNER.address, depositNativeAmount, {
        value: depositNativeAmount,
      });

      await expect(tx)
        .emit(billingManager, 'AssetDeposited')
        .withArgs(nativeDepositAssetKey, OWNER.address, OWNER.address, depositNativeAmount);
      await expect(tx).to.changeEtherBalances(
        [OWNER, billingManager],
        [depositNativeAmount.mul(-1), depositNativeAmount]
      );

      expect(await billingManager.getTotalDepositAssetAmount(nativeDepositAssetKey)).to.be.eq(depositNativeAmount);
    });

    it('should correctly deposit ERC20 tokens', async () => {
      await nerifToken.approve(billingManager.address, depositTokensAmount);

      const tx = await billingManager.deposit(nerifTokenDepositAssetKey, FIRST.address, depositTokensAmount);

      await expect(tx)
        .emit(billingManager, 'AssetDeposited')
        .withArgs(nerifTokenDepositAssetKey, OWNER.address, FIRST.address, depositTokensAmount);

      expect(await nerifToken.balanceOf(OWNER.address)).to.be.eq(tokensAmount.sub(depositTokensAmount));
      expect(await nerifToken.balanceOf(billingManager.address)).to.be.eq(depositTokensAmount);

      expect(await billingManager.getTotalDepositAssetAmount(nerifTokenDepositAssetKey)).to.be.eq(depositTokensAmount);
    });

    it('should get exception if the deposit asset key does not exist', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      await expect(billingManager.deposit('SOME_KEY', OWNER.address, depositTokensAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset is disabled', async () => {
      const reason = 'BillingManager: Deposit asset is disabled';

      await billingManager.connect(SIGNER).updateDepositAssetEnabledStatus(nativeDepositAssetKey, false);

      await expect(
        billingManager.deposit(nativeDepositAssetKey, OWNER.address, depositNativeAmount)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass zero deposit amount', async () => {
      const reason = 'BillingManager: Zero deposit amount';

      await expect(billingManager.deposit(nativeDepositAssetKey, OWNER.address, 0)).to.be.revertedWith(reason);
    });

    it('should get exception if the pass invalid msg.value', async () => {
      const reason = 'BillingManager: Invalid msg.value';

      await expect(
        billingManager.deposit(nativeDepositAssetKey, OWNER.address, depositNativeAmount)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the pass msg.value to the non native deposit asset', async () => {
      const reason = 'BillingManager: Not a native deposit asset';

      await expect(
        billingManager.deposit(nerifTokenDepositAssetKey, OWNER.address, depositTokensAmount, {
          value: depositNativeAmount,
        })
      ).to.be.revertedWith(reason);
    });
  });

  describe('depositWithPermit', () => {
    const depositTokensAmount = wei('100');
    const currentTime = wei('100000', 1);

    beforeEach('setup', async () => {
      await setTime(currentTime.toNumber());
    });

    it('should correctly deposit with permit', async () => {
      const deadline = currentTime.add(1000);

      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        depositTokensAmount.toString()
      );

      const tx = await billingManager.depositWithPermit(nerifTokenDepositAssetKey, OWNER.address, depositTokensAmount, {
        sigExpirationTime: deadline,
        v: sig.v,
        r: sig.r,
        s: sig.s,
      });

      await expect(tx)
        .emit(billingManager, 'AssetDeposited')
        .withArgs(nerifTokenDepositAssetKey, OWNER.address, OWNER.address, depositTokensAmount);

      expect(await billingManager.getTotalDepositAssetAmount(nerifTokenDepositAssetKey)).to.be.eq(depositTokensAmount);
    });

    it('should get exception if the deposit asset key does not exist', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      const deadline = currentTime.add(1000);
      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        depositTokensAmount.toString()
      );

      await expect(
        billingManager.depositWithPermit('SOME_KEY', OWNER.address, depositTokensAmount, {
          sigExpirationTime: deadline,
          v: sig.v,
          r: sig.r,
          s: sig.s,
        })
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset is disabled', async () => {
      const reason = 'BillingManager: Deposit asset is disabled';

      const deadline = currentTime.add(1000);
      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        depositTokensAmount.toString()
      );

      await billingManager.connect(SIGNER).updateDepositAssetEnabledStatus(nerifTokenDepositAssetKey, false);

      await expect(
        billingManager.depositWithPermit(nerifTokenDepositAssetKey, OWNER.address, depositTokensAmount, {
          sigExpirationTime: deadline,
          v: sig.v,
          r: sig.r,
          s: sig.s,
        })
      ).to.be.revertedWith(reason);
    });

    it('should get exception if try to use deposit with permit function with native deposit asset key', async () => {
      const reason = 'BillingManager: Unable to deposit native currency with permit';

      const deadline = currentTime.add(1000);
      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        depositTokensAmount.toString()
      );

      await expect(
        billingManager.depositWithPermit(nativeDepositAssetKey, OWNER.address, depositTokensAmount, {
          sigExpirationTime: deadline,
          v: sig.v,
          r: sig.r,
          s: sig.s,
        })
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the asset is not permitable', async () => {
      const reason = 'BillingManager: Deposit asset is not permitable';

      const newNerifAssetKey: string = 'NERIF_V2';
      const newNerifAssetData: IBillingManager.DepositAssetDataStruct = {
        ...nerifTokenDepositAssetData,
        isPermitable: false,
      };

      await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: newNerifAssetKey,
          depositAssetData: newNerifAssetData,
        },
      ]);

      const deadline = currentTime.add(1000);
      const sig = createPermitSig(
        await nerifToken.name(),
        nerifToken.address,
        deadline.toString(),
        depositTokensAmount.toString()
      );

      await expect(
        billingManager.depositWithPermit(newNerifAssetKey, OWNER.address, depositTokensAmount, {
          sigExpirationTime: deadline,
          v: sig.v,
          r: sig.r,
          s: sig.s,
        })
      ).to.be.revertedWith(reason);
    });
  });

  describe('withdrawFunds', () => {
    const currentTime = wei('100000', 1);
    const depositTokensAmount = wei('100');
    const depositNativeAmount = wei('1');
    const tokensAmountToWithdraw = wei('30');
    const nativeAmountToWithdraw = wei('0.5');

    beforeEach('setup', async () => {
      await nerifToken.approve(billingManager.address, tokensAmount);
      await billingManager.deposit(nerifTokenDepositAssetKey, OWNER.address, depositTokensAmount);
      await billingManager.deposit(nativeDepositAssetKey, OWNER.address, depositNativeAmount, {
        value: depositNativeAmount,
      });

      await setTime(currentTime.toNumber());
    });

    it('should correctly withdraw funds', async () => {
      const deadline = currentTime.add(1000);
      const userNonce = 100;

      const sig = createFundsWithdrawSig(
        OWNER.address,
        [nerifTokenDepositAssetKey, nativeDepositAssetKey],
        deadline.toString(),
        [tokensAmountToWithdraw.toString(), nativeAmountToWithdraw.toString()],
        userNonce.toString()
      );

      const tx = await billingManager.withdrawFunds({
        depositAssetKeys: [nerifTokenDepositAssetKey, nativeDepositAssetKey],
        withdrawAmounts: [tokensAmountToWithdraw.toString(), nativeAmountToWithdraw.toString()],
        nonce: userNonce.toString(),
        sigData: {
          sigExpirationTime: deadline,
          v: sig.v,
          r: sig.r,
          s: sig.s,
        },
      });

      expect(await nerifToken.balanceOf(OWNER.address)).to.be.eq(
        tokensAmount.sub(depositTokensAmount).add(tokensAmountToWithdraw)
      );
      expect(await billingManager.isUserNonceUsed(OWNER.address, userNonce)).to.be.eq(true);
      expect(await billingManager.getTotalDepositAssetAmount(nerifTokenDepositAssetKey)).to.be.eq(
        depositTokensAmount.sub(tokensAmountToWithdraw)
      );
      expect(await billingManager.getTotalDepositAssetAmount(nativeDepositAssetKey)).to.be.eq(
        depositNativeAmount.sub(nativeAmountToWithdraw)
      );

      await expect(tx)
        .to.emit(billingManager, 'FundsWithdrawn')
        .withArgs(
          OWNER.address,
          [nerifTokenDepositAssetKey, nativeDepositAssetKey],
          [tokensAmountToWithdraw, nativeAmountToWithdraw],
          userNonce
        );
    });

    it('should get exception if pass invalid signature', async () => {
      const reason = 'BillingManager: Not a signer';
      const deadline = currentTime.add(1000);
      const userNonce = 100;

      const sig = createFundsWithdrawSig(
        OWNER.address,
        [nerifTokenDepositAssetKey],
        deadline.toString(),
        [tokensAmountToWithdraw.toString()],
        userNonce.toString(),
        OWNER_PK
      );

      await expect(
        billingManager.withdrawFunds({
          depositAssetKeys: [nerifTokenDepositAssetKey],
          withdrawAmounts: [tokensAmountToWithdraw.toString()],
          nonce: userNonce.toString(),
          sigData: {
            sigExpirationTime: deadline,
            v: sig.v,
            r: sig.r,
            s: sig.s,
          },
        })
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the signature has expired', async () => {
      const reason = 'BillingManager: Expired deadline';
      const deadline = currentTime.add(1000);
      const userNonce = 100;

      const sig = createFundsWithdrawSig(
        OWNER.address,
        [nerifTokenDepositAssetKey],
        deadline.toString(),
        [tokensAmountToWithdraw.toString()],
        userNonce.toString()
      );

      await setNextBlockTime(deadline.add(100).toNumber());

      await expect(
        billingManager.withdrawFunds({
          depositAssetKeys: [nerifTokenDepositAssetKey],
          withdrawAmounts: [tokensAmountToWithdraw.toString()],
          nonce: userNonce.toString(),
          sigData: {
            sigExpirationTime: deadline,
            v: sig.v,
            r: sig.r,
            s: sig.s,
          },
        })
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the user tries execute several signatures with the same nonce', async () => {
      const reason = 'BillingManager: Nonce has already been used';
      const deadline = currentTime.add(1000);
      const userNonce = 100;

      const sig1 = createFundsWithdrawSig(
        OWNER.address,
        [nerifTokenDepositAssetKey],
        deadline.toString(),
        [tokensAmountToWithdraw.toString()],
        userNonce.toString()
      );
      const sig2 = createFundsWithdrawSig(
        OWNER.address,
        [nerifTokenDepositAssetKey],
        deadline.toString(),
        [tokensAmountToWithdraw.mul('2').toString()],
        userNonce.toString()
      );

      await billingManager.withdrawFunds({
        depositAssetKeys: [nerifTokenDepositAssetKey],
        withdrawAmounts: [tokensAmountToWithdraw.toString()],
        nonce: userNonce.toString(),
        sigData: {
          sigExpirationTime: deadline,
          v: sig1.v,
          r: sig1.r,
          s: sig1.s,
        },
      });

      await expect(
        billingManager.withdrawFunds({
          depositAssetKeys: [nerifTokenDepositAssetKey],
          withdrawAmounts: [tokensAmountToWithdraw.toString()],
          nonce: userNonce.toString(),
          sigData: {
            sigExpirationTime: deadline,
            v: sig2.v,
            r: sig2.r,
            s: sig2.s,
          },
        })
      ).to.be.revertedWith(reason);

      expect(await billingManager.getTotalDepositAssetAmount(nerifTokenDepositAssetKey)).to.be.eq(
        depositTokensAmount.sub(tokensAmountToWithdraw)
      );
      expect(await billingManager.isUserNonceUsed(OWNER.address, userNonce)).to.be.eq(true);
    });

    it('should get exception if try to withdraw more than total amount', async () => {
      const reason = 'BillingManager: Not enough deposit asset';
      const deadline = currentTime.add(1000);
      const userNonce = 100;

      const sig = createFundsWithdrawSig(
        OWNER.address,
        [nerifTokenDepositAssetKey],
        deadline.toString(),
        [depositTokensAmount.mul('2').toString()],
        userNonce.toString()
      );

      await expect(
        billingManager.withdrawFunds({
          depositAssetKeys: [nerifTokenDepositAssetKey],
          withdrawAmounts: [depositTokensAmount.mul('2').toString()],
          nonce: userNonce.toString(),
          sigData: {
            sigExpirationTime: deadline,
            v: sig.v,
            r: sig.r,
            s: sig.s,
          },
        })
      ).to.be.revertedWith(reason);
    });
  });

  describe('rewardsWithdraw', () => {
    const currentTime = wei('100000', 1);
    const depositTokensAmount = wei('100');
    const depositNativeAmount = wei('1');
    const tokensAmountToWithdraw = wei('30');
    const nativeAmountToWithdraw = wei('0.5');

    beforeEach('setup', async () => {
      await nerifToken.approve(billingManager.address, tokensAmount);
      await billingManager.deposit(nerifTokenDepositAssetKey, OWNER.address, depositTokensAmount);
      await billingManager.deposit(nativeDepositAssetKey, OWNER.address, depositNativeAmount, {
        value: depositNativeAmount,
      });

      await setTime(currentTime.toNumber());
    });

    it('should correctly withdraw rewards', async () => {
      const deadline = currentTime.add(1000);
      const userNonce = 100;

      const sig = createRewardsWithdrawSig(
        OWNER.address,
        [nerifTokenDepositAssetKey, nativeDepositAssetKey],
        deadline.toString(),
        [tokensAmountToWithdraw.toString(), nativeAmountToWithdraw.toString()],
        userNonce.toString()
      );

      const tx = await billingManager.withdrawRewards({
        depositAssetKeys: [nerifTokenDepositAssetKey, nativeDepositAssetKey],
        withdrawAmounts: [tokensAmountToWithdraw.toString(), nativeAmountToWithdraw.toString()],
        nonce: userNonce.toString(),
        sigData: {
          sigExpirationTime: deadline,
          v: sig.v,
          r: sig.r,
          s: sig.s,
        },
      });

      expect(await nerifToken.balanceOf(OWNER.address)).to.be.eq(
        tokensAmount.sub(depositTokensAmount).add(tokensAmountToWithdraw)
      );
      expect(await billingManager.isUserNonceUsed(OWNER.address, userNonce)).to.be.eq(true);
      expect(await billingManager.getTotalDepositAssetAmount(nerifTokenDepositAssetKey)).to.be.eq(
        depositTokensAmount.sub(tokensAmountToWithdraw)
      );
      expect(await billingManager.getTotalDepositAssetAmount(nativeDepositAssetKey)).to.be.eq(
        depositNativeAmount.sub(nativeAmountToWithdraw)
      );

      await expect(tx)
        .to.emit(billingManager, 'RewardsWithdrawn')
        .withArgs(
          OWNER.address,
          [nerifTokenDepositAssetKey, nativeDepositAssetKey],
          [tokensAmountToWithdraw, nativeAmountToWithdraw],
          userNonce
        );
    });
  });
});
