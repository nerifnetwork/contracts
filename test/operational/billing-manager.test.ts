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
import { setTime } from '../helpers/block-helper';

const { signPermit } = require('../helpers/signatures.js');

describe('BillingManager', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let contractsRegistry: ContractsRegistry;

  let signerStorage: SignerStorage;
  let registry: Registry;
  let billingManager: BillingManager;
  let gatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;

  let nerifToken: NerifToken;

  const OWNER_PK: string = 'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';

  const defaultWorkflowId: number = 13;
  const defaultWithdrawReason = 0;
  const tokensAmount = wei('1000');

  const nativeDepositAssetKey: string = 'NATIVE';
  const nativeDepositAssetData: IBillingManager.DepositAssetDataStruct = {
    tokenAddr: ethers.constants.AddressZero,
    workflowExecutionDiscount: 0,
    networkRewards: 0,
    isPermitable: false,
    isEnabled: true,
  };

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

  before(async () => {
    [OWNER, FIRST, SIGNER] = await ethers.getSigners();

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
    await contractsRegistry.addContract(await contractsRegistry.SIGNER_GETTER_NAME(), signerStorage.address);

    registry = RegistryFactory.attach(await contractsRegistry.getRegistryContract());
    billingManager = BillingManagerFactory.attach(await contractsRegistry.getBillingManagerContract());
    nerifToken = NerifTokenFactory.attach(await contractsRegistry.getNerifTokenContract());

    await signerStorage.initialize(SIGNER.address);
    await registry.initialize(0);
    await billingManager.initialize({
      depositAssetKey: nativeDepositAssetKey,
      depositAssetData: nativeDepositAssetData,
    });
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
        billingManager.initialize({
          depositAssetKey: nativeDepositAssetKey,
          depositAssetData: nativeDepositAssetData,
        })
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

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nativeDepositAssetKey);

      expect(userDepositInfo.userAddr).to.be.eq(OWNER.address);
      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositAmount);
      expect(await billingManager.getTotalUsersCount()).to.be.eq(1);
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
    const nerifAssetKey: string = 'NERIF';
    let nerifAssetData: IBillingManager.DepositAssetDataStruct;

    beforeEach('setup', async () => {
      nerifAssetData = {
        tokenAddr: nerifToken.address,
        workflowExecutionDiscount: 10,
        networkRewards: 0,
        isPermitable: true,
        isEnabled: true,
      };
    });

    it('should correctly add new deposit asset key', async () => {
      const tx = await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: nerifAssetKey,
          depositAssetData: nerifAssetData,
        },
      ]);

      await expect(tx).emit(billingManager, 'DepositAssetAdded').withArgs(nerifAssetKey, Object.values(nerifAssetData));

      expect(await billingManager.isDepositAssetEnabled(nerifAssetKey)).to.be.eq(true);
      expect(await billingManager.isDepositAssetSupported(nerifAssetKey)).to.be.eq(true);

      const nativeAssetInfo = (await billingManager.getDepositAssetsInfo([nerifAssetKey]))[0];

      expect(nativeAssetInfo.depositAssetData.tokenAddr).to.be.eq(nerifToken.address);
      expect(nativeAssetInfo.depositAssetData.workflowExecutionDiscount).to.be.eq(
        nerifAssetData.workflowExecutionDiscount
      );
      expect(nativeAssetInfo.depositAssetData.networkRewards).to.be.eq(0);
      expect(nativeAssetInfo.depositAssetData.isEnabled).to.be.eq(true);

      expect(await billingManager.getSupportedDepositAssetKeys()).to.be.deep.eq([nativeDepositAssetKey, nerifAssetKey]);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'BillingManager: Not a signer';

      await expect(
        billingManager.addDepositAssets([
          {
            depositAssetKey: nerifAssetKey,
            depositAssetData: nerifAssetData,
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

      nerifAssetData.tokenAddr = ethers.constants.AddressZero;

      await expect(
        billingManager.connect(SIGNER).addDepositAssets([
          {
            depositAssetKey: nerifAssetKey,
            depositAssetData: nerifAssetData,
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
            depositAssetData: nerifAssetData,
          },
        ])
      ).to.be.revertedWith(reason);
    });

    it('should get exception if pass not empty network rewards', async () => {
      const reason = 'BillingManager: Invalid init network rewards value';

      nerifAssetData.networkRewards = wei('1');

      await expect(
        billingManager.connect(SIGNER).addDepositAssets([
          {
            depositAssetKey: nerifAssetKey,
            depositAssetData: nerifAssetData,
          },
        ])
      ).to.be.revertedWith(reason);
    });
  });

  describe('updateWorkflowExecutionDiscount', () => {
    const newDiscount = 40;
    const nerifAssetKey: string = 'NERIF';

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
        billingManager.connect(SIGNER).updateWorkflowExecutionDiscount(nerifAssetKey, newDiscount)
      ).to.be.revertedWith(reason);
    });
  });

  describe('updateDepositAssetEnabledStatus', () => {
    const newEnabledStatus: boolean = false;
    const nerifAssetKey: string = 'NERIF';

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
        billingManager.connect(SIGNER).updateDepositAssetEnabledStatus(nerifAssetKey, newEnabledStatus)
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
    const nerifAssetKey: string = 'NERIF';
    let nerifAssetData: IBillingManager.DepositAssetDataStruct;

    beforeEach('setup', async () => {
      nerifAssetData = {
        tokenAddr: nerifToken.address,
        workflowExecutionDiscount: 10,
        networkRewards: 0,
        isPermitable: true,
        isEnabled: true,
      };

      await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: nerifAssetKey,
          depositAssetData: nerifAssetData,
        },
      ]);
    });

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

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nativeDepositAssetKey);

      expect(userDepositInfo.userAddr).to.be.eq(OWNER.address);
      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositNativeAmount);

      expect(await billingManager.getExistingUsers(0, 10)).to.be.deep.eq([OWNER.address]);
      expect(await billingManager.getUserDepositAssetKeys(OWNER.address)).to.be.deep.eq([nativeDepositAssetKey]);
    });

    it('should correctly deposit ERC20 tokens', async () => {
      await nerifToken.approve(billingManager.address, depositTokensAmount);

      const tx = await billingManager.deposit(nerifAssetKey, FIRST.address, depositTokensAmount);

      await expect(tx)
        .emit(billingManager, 'AssetDeposited')
        .withArgs(nerifAssetKey, OWNER.address, FIRST.address, depositTokensAmount);

      expect(await nerifToken.balanceOf(OWNER.address)).to.be.eq(tokensAmount.sub(depositTokensAmount));
      expect(await nerifToken.balanceOf(billingManager.address)).to.be.eq(depositTokensAmount);

      const userDepositInfo = await billingManager.getUserDepositInfo(FIRST.address, nerifAssetKey);

      expect(userDepositInfo.userAddr).to.be.eq(FIRST.address);
      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount);
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
        billingManager.deposit(nerifAssetKey, OWNER.address, depositTokensAmount, { value: depositNativeAmount })
      ).to.be.revertedWith(reason);
    });
  });

  describe('depositWithPermit', () => {
    const depositTokensAmount = wei('100');
    const currentTime = wei('100000', 1);
    const nerifAssetKey: string = 'NERIF';
    let nerifAssetData: IBillingManager.DepositAssetDataStruct;

    beforeEach('setup', async () => {
      nerifAssetData = {
        tokenAddr: nerifToken.address,
        workflowExecutionDiscount: 10,
        networkRewards: 0,
        isPermitable: true,
        isEnabled: true,
      };

      await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: nerifAssetKey,
          depositAssetData: nerifAssetData,
        },
      ]);

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

      const tx = await billingManager.depositWithPermit(
        nerifAssetKey,
        OWNER.address,
        depositTokensAmount,
        deadline,
        sig.v,
        sig.r,
        sig.s
      );

      await expect(tx)
        .emit(billingManager, 'AssetDeposited')
        .withArgs(nerifAssetKey, OWNER.address, OWNER.address, depositTokensAmount);

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userAddr).to.be.eq(OWNER.address);
      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount);
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
        billingManager.depositWithPermit('SOME_KEY', OWNER.address, depositTokensAmount, deadline, sig.v, sig.r, sig.s)
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

      await billingManager.connect(SIGNER).updateDepositAssetEnabledStatus(nerifAssetKey, false);

      await expect(
        billingManager.depositWithPermit(
          nerifAssetKey,
          OWNER.address,
          depositTokensAmount,
          deadline,
          sig.v,
          sig.r,
          sig.s
        )
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
        billingManager.depositWithPermit(
          nativeDepositAssetKey,
          OWNER.address,
          depositTokensAmount,
          deadline,
          sig.v,
          sig.r,
          sig.s
        )
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the asset is not permitable', async () => {
      const reason = 'BillingManager: Deposit asset is not permitable';

      const newNerifAssetKey = 'NERIF_V2';
      nerifAssetData.isPermitable = false;

      await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: newNerifAssetKey,
          depositAssetData: nerifAssetData,
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
        billingManager.depositWithPermit(
          newNerifAssetKey,
          OWNER.address,
          depositTokensAmount,
          deadline,
          sig.v,
          sig.r,
          sig.s
        )
      ).to.be.revertedWith(reason);
    });
  });

  describe.skip('withdrawNetworkRewards', () => {
    const depositTokensAmount = wei('100');
    const tokensExecutionAmount = wei('20');
    const nerifAssetKey: string = 'NERIF';
    let nerifAssetData: IBillingManager.DepositAssetDataStruct;

    beforeEach('setup', async () => {
      nerifAssetData = {
        tokenAddr: nerifToken.address,
        workflowExecutionDiscount: 10,
        networkRewards: 0,
        isPermitable: true,
        isEnabled: true,
      };

      await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: nerifAssetKey,
          depositAssetData: nerifAssetData,
        },
      ]);

      await nerifToken.approve(billingManager.address, tokensAmount);
      await billingManager.deposit(nerifAssetKey, OWNER.address, depositTokensAmount);

      expect(await billingManager.getNetworkRewards(nerifAssetKey)).to.be.eq(tokensExecutionAmount);
    });

    it('should correctly withdraw network rewards', async () => {
      const tx = await billingManager.withdrawNetworkRewards(nerifAssetKey);

      await expect(tx)
        .emit(billingManager, 'RewardsWithdrawn')
        .withArgs(nerifAssetKey, SIGNER.address, tokensExecutionAmount);

      expect(await nerifToken.balanceOf(billingManager.address)).to.be.eq(
        depositTokensAmount.sub(tokensExecutionAmount)
      );

      expect(await billingManager.getNetworkRewards(nerifAssetKey)).to.be.eq('0');
    });

    it('should get execption if nothing to withdraw', async () => {
      const reason = 'BillingManager: No network rewards to withdraw';

      await expect(billingManager.withdrawNetworkRewards(nativeDepositAssetKey)).to.be.revertedWith(reason);
    });

    it('should get exception if signer is zero addr', async () => {
      const reason = 'BillingManager: Zero signer address';

      await signerStorage.connect(SIGNER).setAddress(ethers.constants.AddressZero);

      await expect(billingManager.withdrawNetworkRewards(nerifAssetKey)).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset key does not exist', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      await expect(billingManager.withdrawNetworkRewards('SOME_KEY')).to.be.revertedWith(reason);
    });
  });

  describe.skip('networkWithdraw', () => {
    const depositTokensAmount = wei('200');
    const withdrawAmount = wei('75');
    const nerifAssetKey: string = 'NERIF';

    let nerifAssetData: IBillingManager.DepositAssetDataStruct;

    beforeEach('setup', async () => {
      nerifAssetData = {
        tokenAddr: nerifToken.address,
        workflowExecutionDiscount: 10,
        networkRewards: 0,
        isPermitable: true,
        isEnabled: true,
      };

      await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: nerifAssetKey,
          depositAssetData: nerifAssetData,
        },
      ]);

      await nerifToken.approve(billingManager.address, tokensAmount);
      await billingManager.deposit(nerifAssetKey, OWNER.address, depositTokensAmount);
    });

    it('should correctly withdraw user available tokens', async () => {
      const tx = await billingManager
        .connect(SIGNER)
        .networkWithdraw(nerifAssetKey, OWNER.address, withdrawAmount, defaultWithdrawReason);

      await expect(tx)
        .emit(billingManager, 'NetworkWithdrawCompleted')
        .withArgs(nerifAssetKey, OWNER.address, defaultWithdrawReason, withdrawAmount);

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount.sub(withdrawAmount));

      expect(await billingManager.getNetworkRewards(nerifAssetKey)).to.be.eq(withdrawAmount);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'BillingManager: Not a signer';

      await expect(
        billingManager.networkWithdraw(nerifAssetKey, OWNER.address, 10, defaultWithdrawReason)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset key does not exist', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      await expect(
        billingManager.connect(SIGNER).networkWithdraw('SOME_KEY', OWNER.address, 10, defaultWithdrawReason)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if try to withdraw zero amount', async () => {
      const reason = 'BillingManager: Zero amount to update';

      await expect(
        billingManager.connect(SIGNER).networkWithdraw(nerifAssetKey, OWNER.address, 0, defaultWithdrawReason)
      ).to.be.revertedWith(reason);
    });
  });

  describe.skip('getters', () => {
    const nativeTokensAmount = wei('1');
    const nerifAssetKey: string = 'NERIF';
    let nerifAssetData: IBillingManager.DepositAssetDataStruct;

    beforeEach('setup', async () => {
      nerifAssetData = {
        tokenAddr: nerifToken.address,
        workflowExecutionDiscount: 10,
        networkRewards: 0,
        isPermitable: true,
        isEnabled: true,
      };

      await billingManager.connect(SIGNER).addDepositAssets([
        {
          depositAssetKey: nerifAssetKey,
          depositAssetData: nerifAssetData,
        },
      ]);

      await nerifToken.approve(billingManager.address, tokensAmount);
    });

    it('should return correct array with existing users addresses', async () => {
      const signers = await ethers.getSigners();
      const expectedArr = [];

      for (let i = 0; i < signers.length; i++) {
        await billingManager
          .connect(signers[i])
          .deposit(nativeDepositAssetKey, signers[i].address, nativeTokensAmount, { value: nativeTokensAmount });

        expect(await billingManager.getTotalUsersCount()).to.be.eq(i + 1);
        expectedArr.push(signers[i].address);
      }

      expect(await billingManager.getExistingUsers(0, 10)).to.be.deep.eq(expectedArr.slice(0, 10));
      expect(await billingManager.getExistingUsers(5, signers.length)).to.be.deep.eq(
        expectedArr.slice(5, signers.length)
      );

      const numberToWithdraw = 3;

      // for (let i = 0; i < numberToWithdraw; i++) {
      //   await billingManager
      //     .connect(signers[signers.length - 1 - i])
      //     .withdrawFunds(nativeDepositAssetKey, nativeTokensAmount);

      //   expect(await billingManager.getTotalUsersCount()).to.be.eq(signers.length - 1 - i);
      // }

      expect(await billingManager.getTotalUsersCount()).to.be.eq(signers.length - numberToWithdraw);
      expect(await billingManager.getExistingUsers(0, signers.length)).to.be.deep.eq(
        expectedArr.slice(0, signers.length - numberToWithdraw)
      );
    });

    it('should return correct users funds info', async () => {
      await billingManager.deposit(nativeDepositAssetKey, OWNER.address, nativeTokensAmount, {
        value: nativeTokensAmount,
      });
      await billingManager
        .connect(FIRST)
        .deposit(nativeDepositAssetKey, FIRST.address, nativeTokensAmount.mul(2), { value: nativeTokensAmount.mul(2) });

      const usersFundsInfo = await billingManager.getUsersDepositInfo(nativeDepositAssetKey, 0, 10);

      expect(usersFundsInfo.length).to.be.eq(2);

      expect(usersFundsInfo[0].userAddr).to.be.eq(OWNER.address);
      expect(usersFundsInfo[0].userDepositedAmount).to.be.eq(nativeTokensAmount);

      expect(usersFundsInfo[1].userAddr).to.be.eq(FIRST.address);
      expect(usersFundsInfo[1].userDepositedAmount).to.be.eq(nativeTokensAmount.mul(2));
    });
  });
});
