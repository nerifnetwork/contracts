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
  TestERC20,
  IBillingManager,
} from '../../generated-types/ethers';
import { setTime } from '../helpers/block-helper';

const { signPermit } = require('../helpers/signatures.js');

describe('BillingManager', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let signerStorage: SignerStorage;
  let registry: Registry;
  let billingManager: BillingManager;
  let gatewayFactory: GatewayFactory;
  let gatewayImpl: Gateway;

  let nerifToken: TestERC20;

  const OWNER_PK: string = 'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';

  const defaultWorkflowId: number = 13;
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
    const RegistryFactory = await ethers.getContractFactory('Registry');
    const BillingManagerFactory = await ethers.getContractFactory('BillingManager');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const GatewayFactoryFactory = await ethers.getContractFactory('GatewayFactory');
    const GatewayImplFactory = await ethers.getContractFactory('Gateway');
    const TestERC20Factory = await ethers.getContractFactory('TestERC20');

    signerStorage = await SignerStorageFactory.deploy();
    registry = await RegistryFactory.deploy();

    const billingManagerImpl = await BillingManagerFactory.deploy();
    const billingManagerProxy = await ERC1967ProxyFactory.deploy(billingManagerImpl.address, '0x');
    billingManager = BillingManagerFactory.attach(billingManagerProxy.address);

    gatewayFactory = await GatewayFactoryFactory.deploy();
    gatewayImpl = await GatewayImplFactory.deploy();

    nerifToken = await TestERC20Factory.deploy('Nerif', 'NFF', tokensAmount);

    await signerStorage.initialize(SIGNER.address);
    await registry.initialize(true, signerStorage.address, gatewayFactory.address, billingManager.address, 0);
    await billingManager.initialize(registry.address, signerStorage.address, {
      depositAssetKey: nativeDepositAssetKey,
      depositAssetData: nativeDepositAssetData,
    });
    await gatewayFactory.initialize(registry.address, gatewayImpl.address);

    await registry.registerWorkflows([
      {
        id: defaultWorkflowId,
        workflowOwner: OWNER.address,
        hash: '0x',
        requireGateway: true,
        deployGateway: true,
      },
    ]);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialize', async () => {
      expect(await billingManager.registry()).to.be.eq(registry.address);
      expect(await billingManager.signerGetter()).to.be.eq(signerStorage.address);

      const nativeAssetInfo = (await billingManager.getDepositAssetsInfo([nativeDepositAssetKey]))[0];

      expect(nativeAssetInfo.depositAssetData.tokenAddr).to.be.eq(ethers.constants.AddressZero);
      expect(nativeAssetInfo.depositAssetData.isEnabled).to.be.eq(true);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(
        billingManager.initialize(registry.address, signerStorage.address, {
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

      await testBillingManager.connect(SIGNER).upgradeTo(newBillingManagerImpl.address);

      expect(await testBillingManager.version()).to.be.eq('v2.0.0');
    });

    it('should get exception if not a signer try to call upgareTo function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(billingManager.upgradeTo(ethers.constants.AddressZero)).to.be.revertedWith(reason);
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
      const reason = 'SignerOwnable: only signer';

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
      const reason = 'SignerOwnable: only signer';

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
      const reason = 'SignerOwnable: only signer';

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

  describe('lockExecutionFunds', () => {
    const depositTokensAmount = wei('100');
    const tokensLockAmount = wei('20');
    const nativeLockAmount = wei('0.5');
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

    it('should correctly lock user funds', async () => {
      await billingManager.deposit(nerifAssetKey, OWNER.address, depositTokensAmount);

      const nextWorkflowExecutionId = await billingManager.nextWorkflowExecutionId();
      const tx = await billingManager
        .connect(SIGNER)
        .lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount);

      await expect(tx)
        .emit(billingManager, 'ExecutionFundsLocked')
        .withArgs(nerifAssetKey, defaultWorkflowId, OWNER.address, nextWorkflowExecutionId, tokensLockAmount);

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userAddr).to.be.eq(OWNER.address);
      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount);
      expect(userDepositInfo.userLockedAmount).to.be.eq(tokensLockAmount);
      expect(userDepositInfo.pendingWorkflowExecutionIds).to.be.deep.eq([nextWorkflowExecutionId]);

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(nextWorkflowExecutionId);

      expect(workflowExecutionInfo.depositAssetKey).to.be.eq(nerifAssetKey);
      expect(workflowExecutionInfo.workflowId).to.be.eq(defaultWorkflowId);
      expect(workflowExecutionInfo.workflowOwner).to.be.eq(OWNER.address);
      expect(workflowExecutionInfo.executionLockedAmount).to.be.eq(tokensLockAmount);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(0);
      expect(workflowExecutionInfo.status).to.be.eq(1);

      expect(await billingManager.getWorkflowExecutionStatus(nextWorkflowExecutionId)).to.be.eq(1);
      expect(await billingManager.getExecutionWorkflowId(nextWorkflowExecutionId)).to.be.eq(defaultWorkflowId);
      expect(await billingManager.getWorkflowExecutionOwner(nextWorkflowExecutionId)).to.be.eq(OWNER.address);
    });

    it('should correctly lock user funds several times', async () => {
      await billingManager.deposit(nerifAssetKey, OWNER.address, depositTokensAmount);

      const expectedWorkflowExecutionIds = [];

      expectedWorkflowExecutionIds.push(await billingManager.nextWorkflowExecutionId());
      await billingManager.connect(SIGNER).lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount);

      expectedWorkflowExecutionIds.push(await billingManager.nextWorkflowExecutionId());
      await billingManager
        .connect(SIGNER)
        .lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount.mul(2));

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userLockedAmount).to.be.eq(tokensLockAmount.mul(3));
      expect(userDepositInfo.pendingWorkflowExecutionIds).to.be.deep.eq(expectedWorkflowExecutionIds);
      expect(await billingManager.nextWorkflowExecutionId()).to.be.eq(2);

      let workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(expectedWorkflowExecutionIds[0]);

      expect(workflowExecutionInfo.workflowOwner).to.be.eq(OWNER.address);
      expect(workflowExecutionInfo.executionLockedAmount).to.be.eq(tokensLockAmount);

      workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(expectedWorkflowExecutionIds[1]);

      expect(workflowExecutionInfo.workflowOwner).to.be.eq(OWNER.address);
      expect(workflowExecutionInfo.executionLockedAmount).to.be.eq(tokensLockAmount.mul(2));
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(billingManager.lockExecutionFunds(nativeDepositAssetKey, 10, nativeLockAmount)).to.be.revertedWith(
        reason
      );
    });

    it('should get exception if try to lock execution funds with nonexisting deposit asset key', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      await expect(
        billingManager.connect(SIGNER).lockExecutionFunds('SOME_KEY', 10, tokensLockAmount)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if try to lock disabled deposit asset', async () => {
      const reason = 'BillingManager: Deposit asset is disabled';

      await billingManager.connect(SIGNER).updateDepositAssetEnabledStatus(nerifAssetKey, false);

      await expect(
        billingManager.connect(SIGNER).lockExecutionFunds(nerifAssetKey, 10, tokensLockAmount)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if workflow id does not exist', async () => {
      const reason = 'BillingManager: Workflow does not exist';

      await expect(
        billingManager.connect(SIGNER).lockExecutionFunds(nerifAssetKey, 10, tokensLockAmount)
      ).to.be.revertedWith(reason);
    });

    it('should get exception if user does not have available funds to lock', async () => {
      const reason = 'BillingManager: Not enough available funds to lock';

      await expect(
        billingManager.connect(SIGNER).lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount)
      ).to.be.revertedWith(reason);
    });
  });

  describe('completeExecution', () => {
    const depositTokensAmount = wei('200');
    const tokensLockAmount = wei('50');
    const tokensExecutionAmount = wei('30');
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

    beforeEach('setup', async () => {
      await billingManager.deposit(nerifAssetKey, OWNER.address, depositTokensAmount);

      await billingManager.connect(SIGNER).lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount);
      await billingManager
        .connect(SIGNER)
        .lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount.mul(2));
    });

    it('should correctly complete execution when lockedAmount > executionAmount', async () => {
      const tx = await billingManager.connect(SIGNER).completeExecution(0, tokensExecutionAmount);

      await expect(tx).to.emit(billingManager, 'ExecutionCompleted').withArgs(0, tokensExecutionAmount);
      expect((await tx.wait()).events?.length).to.be.eq(1);

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount.sub(tokensExecutionAmount));
      expect(userDepositInfo.userLockedAmount).to.be.eq(tokensLockAmount.mul(2));
      expect(userDepositInfo.pendingWorkflowExecutionIds).to.be.deep.eq([1]);

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(0);

      expect(workflowExecutionInfo.status).to.be.eq(2);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(tokensExecutionAmount);

      const workflowInfo = await registry.getWorkflowInfo(defaultWorkflowId);

      expect(workflowInfo.depositAssetKeys).to.be.deep.eq([nerifAssetKey]);
      expect(workflowInfo.depositAssetsInfo).to.be.deep.eq([[nerifAssetKey, tokensExecutionAmount]]);

      const depositAssetInfo = (await billingManager.getDepositAssetsInfo([nerifAssetKey]))[0];

      expect(depositAssetInfo.depositAssetData.networkRewards).to.be.eq(tokensExecutionAmount);
    });

    it('should correctly complete execution when lockedAmount < execution amount and enough funds', async () => {
      const tx = await billingManager.connect(SIGNER).completeExecution(0, tokensExecutionAmount.mul(2));

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(billingManager, 'ExecutionCompleted').withArgs(0, tokensExecutionAmount.mul(2));
      await expect(tx)
        .emit(billingManager, 'UnexpectedExecutionAmountFound')
        .withArgs(0, tokensLockAmount, tokensExecutionAmount.mul(2));

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount.sub(tokensExecutionAmount.mul(2)));
      expect(userDepositInfo.userLockedAmount).to.be.eq(tokensLockAmount.mul(2));

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(0);

      expect(workflowExecutionInfo.status).to.be.eq(2);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(tokensExecutionAmount.mul(2));
    });

    it('should correctly complete execution when lockedAmount < execution amount and not enough funds', async () => {
      const newExecutionAmount = wei('110');
      const expectedMaxExecutionAmount = wei('100');

      const tx = await billingManager.connect(SIGNER).completeExecution(0, newExecutionAmount);

      expect((await tx.wait()).events?.length).to.be.eq(2);
      await expect(tx).emit(billingManager, 'ExecutionCompleted').withArgs(0, expectedMaxExecutionAmount);
      await expect(tx)
        .emit(billingManager, 'UnexpectedExecutionAmountFound')
        .withArgs(0, tokensLockAmount, expectedMaxExecutionAmount);

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount.sub(expectedMaxExecutionAmount));
      expect(userDepositInfo.userLockedAmount).to.be.eq(tokensLockAmount.mul(2));

      const workflowExecutionInfo = await billingManager.getWorkflowExecutionInfo(0);

      expect(workflowExecutionInfo.status).to.be.eq(2);
      expect(workflowExecutionInfo.executionAmount).to.be.eq(expectedMaxExecutionAmount);

      const depositAssetInfo = (await billingManager.getDepositAssetsInfo([nerifAssetKey]))[0];

      expect(depositAssetInfo.depositAssetData.networkRewards).to.be.eq(expectedMaxExecutionAmount);
    });

    it('should get exception if not a signer try to call this function', async () => {
      const reason = 'SignerOwnable: only signer';

      await expect(billingManager.completeExecution(10, tokensLockAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if status is not equal to PENDING', async () => {
      const reason = 'BillingManager: Not a pending workflow execution';

      await expect(
        billingManager.connect(SIGNER).completeExecution(defaultWorkflowId + 1, tokensLockAmount)
      ).to.be.revertedWith(reason);
    });
  });

  describe('withdrawFunds', () => {
    const depositTokensAmount = wei('200');
    const depositNativeAmount = wei('1');
    const tokensLockAmount = wei('50');
    const withdrawAmount = wei('30');
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
      await billingManager.connect(SIGNER).lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount);
    });

    it('should correctly withdraw ERC20 tokens with specific amount', async () => {
      const tx = await billingManager.withdrawFunds(nerifAssetKey, withdrawAmount);

      await expect(tx).emit(billingManager, 'UserFundsWithdrawn').withArgs(OWNER.address, withdrawAmount);

      expect(await nerifToken.balanceOf(OWNER.address)).to.be.eq(
        tokensAmount.sub(depositTokensAmount).add(withdrawAmount)
      );
      expect(await nerifToken.balanceOf(billingManager.address)).to.be.eq(depositTokensAmount.sub(withdrawAmount));

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount.sub(withdrawAmount));
      expect(await billingManager.getUserAvailableFunds(OWNER.address, nerifAssetKey)).to.be.eq(
        depositTokensAmount.sub(tokensLockAmount).sub(withdrawAmount)
      );
    });

    it('should correctly withdraw all available funds', async () => {
      const availableFunds = await billingManager.getUserAvailableFunds(OWNER.address, nerifAssetKey);

      expect(availableFunds).to.be.eq(depositTokensAmount.sub(tokensLockAmount));

      const tx = await billingManager.withdrawFunds(nerifAssetKey, availableFunds);

      await expect(tx).emit(billingManager, 'UserFundsWithdrawn').withArgs(OWNER.address, availableFunds);

      expect(await nerifToken.balanceOf(OWNER.address)).to.be.eq(
        tokensAmount.sub(depositTokensAmount).add(availableFunds)
      );
      expect(await nerifToken.balanceOf(billingManager.address)).to.be.eq(depositTokensAmount.sub(availableFunds));

      const userDepositInfo = await billingManager.getUserDepositInfo(OWNER.address, nerifAssetKey);

      expect(userDepositInfo.userDepositedAmount).to.be.eq(depositTokensAmount.sub(availableFunds));
      expect(await billingManager.getUserAvailableFunds(OWNER.address, nerifAssetKey)).to.be.eq('0');
    });

    it('should correctly withdraw all user funds', async () => {
      await billingManager.deposit(nativeDepositAssetKey, FIRST.address, depositNativeAmount, {
        value: depositNativeAmount,
      });

      expect(await billingManager.getUserAvailableFunds(FIRST.address, nativeDepositAssetKey)).to.be.eq(
        depositNativeAmount
      );
      expect(await billingManager.getTotalUsersCount()).to.be.eq(2);

      const tx = await billingManager.connect(FIRST).withdrawFunds(nativeDepositAssetKey, depositNativeAmount);

      await expect(tx).to.changeEtherBalances(
        [FIRST, billingManager],
        [depositNativeAmount, depositNativeAmount.mul(-1)]
      );
      expect(await billingManager.getTotalUsersCount()).to.be.eq(1);
    });

    it('should get exception if try to withdraw more than the available', async () => {
      const reason = 'BillingManager: Not enough available funds to withdraw';

      await expect(billingManager.withdrawFunds(nerifAssetKey, withdrawAmount.mul(6))).to.be.revertedWith(reason);
    });

    it('should get exception if the deposit asset key does not exist', async () => {
      const reason = 'BillingManager: Deposit asset does not exist';

      await expect(billingManager.withdrawFunds('SOME_KEY', withdrawAmount)).to.be.revertedWith(reason);
    });

    it('should get exception if try to withdraw zero amount', async () => {
      const reason = 'BillingManager: Zero amount to withdraw';

      await expect(billingManager.withdrawFunds(nerifAssetKey, 0)).to.be.revertedWith(reason);
    });
  });

  describe('withdrawNetworkRewards', () => {
    const depositTokensAmount = wei('100');
    const tokensLockAmount = wei('25');
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
      await billingManager.connect(SIGNER).lockExecutionFunds(nerifAssetKey, defaultWorkflowId, tokensLockAmount);
      await billingManager.connect(SIGNER).completeExecution(0, tokensExecutionAmount);

      expect(await billingManager.getNetworkRewards(nerifAssetKey)).to.be.eq(tokensExecutionAmount);
    });

    it('should correctly withdraw network rewards', async () => {
      const tx = await billingManager.withdrawNetworkRewards(nerifAssetKey);

      await expect(tx).emit(billingManager, 'RewardsWithdrawn').withArgs(SIGNER.address, tokensExecutionAmount);

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

  describe('getters', () => {
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

      for (let i = 0; i < numberToWithdraw; i++) {
        await billingManager
          .connect(signers[signers.length - 1 - i])
          .withdrawFunds(nativeDepositAssetKey, nativeTokensAmount);

        expect(await billingManager.getTotalUsersCount()).to.be.eq(signers.length - 1 - i);
      }

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
