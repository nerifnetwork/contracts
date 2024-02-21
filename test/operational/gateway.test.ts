import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { expect } from 'chai';
import { Reverter } from '../helpers/reverter';
import { Gateway, SignerStorage, TestRegistry, TestTarget } from '../../typechain';
import { BigNumberish } from 'ethers';

describe('Gateway', () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let SIGNER: SignerWithAddress;

  let signerStorage: SignerStorage;
  let gateway: Gateway;
  let testRegistry: TestRegistry;
  let testTarget: TestTarget;

  const defWorkflowId = 10;

  function getSetNumberCalldata(expectedNumber: BigNumberish): string {
    return testTarget.interface.encodeFunctionData('setNumber', [expectedNumber]);
  }

  before(async () => {
    [OWNER, FIRST, SECOND, SIGNER] = await ethers.getSigners();

    const GatewayFactory = await ethers.getContractFactory('Gateway');
    const SignerStorageFactory = await ethers.getContractFactory('SignerStorage');
    const TestRegistryFactory = await ethers.getContractFactory('TestRegistry');
    const TestTargetFactory = await ethers.getContractFactory('TestTarget');

    signerStorage = await SignerStorageFactory.deploy();
    gateway = await GatewayFactory.deploy();
    testRegistry = await TestRegistryFactory.deploy();
    testTarget = await TestTargetFactory.deploy();

    await signerStorage.initialize(SIGNER.address);
    await gateway.initialize(testRegistry.address, OWNER.address);
    await testRegistry.initialize(
      true,
      signerStorage.address,
      ethers.constants.AddressZero,
      ethers.constants.AddressZero,
      0
    );

    await testRegistry.setGateway(gateway.address);
    await testRegistry.registerWorkflow(defWorkflowId, OWNER.address, '0x', true);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe('creation', () => {
    it('should set correct data after initialization', async () => {
      expect(await gateway.owner()).to.be.eq(OWNER.address);
      expect(await gateway.registry()).to.be.eq(testRegistry.address);
    });

    it('should get exception if try to call init function twice', async () => {
      const reason = 'Initializable: contract is already initialized';

      await expect(gateway.initialize(testRegistry.address, OWNER.address)).to.be.revertedWith(reason);
    });
  });

  describe('perform', () => {
    const expectedNumber = 10;
    const workflowID = 20;

    beforeEach('setup', async () => {
      await testRegistry.connect(FIRST).registerWorkflow(workflowID, FIRST.address, '0x', false);
    });

    describe('allowed cases', () => {
      it('should correctly perform action from the workflow owner', async () => {
        await testRegistry.callPerform(
          gateway.address,
          defWorkflowId,
          testTarget.address,
          getSetNumberCalldata(expectedNumber)
        );

        expect(await testTarget.number()).to.be.eq(expectedNumber);
      });

      it('should correctly perform action with known workflow', async () => {
        await gateway.updateKnownWorkflows([
          {
            knownWorkflowToUpdate: workflowID,
            isAdding: true,
          },
        ]);

        await testRegistry.callPerform(
          gateway.address,
          workflowID,
          testTarget.address,
          getSetNumberCalldata(expectedNumber)
        );

        expect(await testTarget.number()).to.be.eq(expectedNumber);
      });

      it('should correctly perform action with known workflow owner', async () => {
        await gateway.updateKnownWorkflowOwners([
          {
            addressToUpdate: FIRST.address,
            isAdding: true,
          },
        ]);

        await testRegistry.callPerform(
          gateway.address,
          workflowID,
          testTarget.address,
          getSetNumberCalldata(expectedNumber)
        );

        expect(await testTarget.number()).to.be.eq(expectedNumber);
      });

      it('should correctly perform action with known customer contract address', async () => {
        await gateway.updateKnownCustomerContracts([
          {
            addressToUpdate: testTarget.address,
            isAdding: true,
          },
        ]);

        await testRegistry.callPerform(
          gateway.address,
          workflowID,
          testTarget.address,
          getSetNumberCalldata(expectedNumber)
        );

        expect(await testTarget.number()).to.be.eq(expectedNumber);
      });
    });

    it('should get exception if not a registry try to call this function', async () => {
      const reason = 'Gateway: operation is not permitted';

      await expect(gateway.perform(10, ethers.constants.AddressZero, [])).to.be.revertedWith(reason);
    });

    it('should get exception if not allowed', async () => {
      const reason = 'Gateway: is not allowed workflow';

      await expect(
        testRegistry.callPerform(gateway.address, workflowID, testTarget.address, getSetNumberCalldata(expectedNumber))
      ).to.be.revertedWith(reason);
    });

    it('should get exception if transaction on target contract was failed', async () => {
      const reason = 'Gateway: failed to execute customer contract';

      await expect(
        testRegistry.callPerform(
          gateway.address,
          defWorkflowId,
          testRegistry.address,
          getSetNumberCalldata(expectedNumber)
        )
      ).to.be.revertedWith(reason);
    });
  });

  describe('setRegistry', () => {
    it('should correctly set new Registry address', async () => {
      await gateway.setRegistry(FIRST.address);

      expect(await gateway.registry()).to.be.eq(FIRST.address);
    });

    it('should get exception if nonowner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(gateway.connect(FIRST).setRegistry(FIRST.address)).to.be.revertedWith(reason);
    });
  });

  describe('updateConfigData', () => {
    it('should correctly update config data', async () => {
      let updateConfigData = {
        updateKnownWorkflowsEntries: [
          {
            knownWorkflowToUpdate: 10,
            isAdding: true,
          },
          {
            knownWorkflowToUpdate: 20,
            isAdding: true,
          },
        ],
        updateKnownWorkflowOwnersEntries: [
          {
            addressToUpdate: OWNER.address,
            isAdding: true,
          },
          {
            addressToUpdate: FIRST.address,
            isAdding: true,
          },
        ],
        updateKnownCustomerContractsEntries: [],
      };

      const expectedKnownWorkflowsArr = [10, 20];
      const expectedKnownWorkflowOwnersArr = [OWNER.address, FIRST.address];

      await gateway.updateConfigData(updateConfigData);

      const currentConfig = await gateway.getConfigInfo();

      currentConfig.knownWorkflows.forEach((el, index) => {
        expect(el).to.be.eq(expectedKnownWorkflowsArr[index]);
      });
      expect(currentConfig.knownWorkflowOwners).to.be.deep.eq(expectedKnownWorkflowOwnersArr);
      expect(currentConfig.knownCustomerContracts).to.be.deep.eq([]);
    });

    it('should get exception if nonowner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(
        gateway.connect(FIRST).updateConfigData({
          updateKnownWorkflowsEntries: [],
          updateKnownWorkflowOwnersEntries: [
            {
              addressToUpdate: FIRST.address,
              isAdding: true,
            },
          ],
          updateKnownCustomerContractsEntries: [],
        })
      ).to.be.revertedWith(reason);
    });
  });

  describe('updateKnownWorkflows', () => {
    it('should correctly update known workflows', async () => {
      let updateKnownWorkflowsEntries = [
        {
          knownWorkflowToUpdate: 10,
          isAdding: true,
        },
        {
          knownWorkflowToUpdate: 20,
          isAdding: true,
        },
        {
          knownWorkflowToUpdate: 30,
          isAdding: true,
        },
      ];
      let expectedKnownWorkflowsArr = [10, 20, 30];

      await gateway.updateKnownWorkflows(updateKnownWorkflowsEntries);

      let currentConfig = await gateway.getConfigInfo();

      currentConfig.knownWorkflows.forEach((el, index) => {
        expect(el).to.be.eq(expectedKnownWorkflowsArr[index]);
      });

      updateKnownWorkflowsEntries = [
        {
          knownWorkflowToUpdate: 10,
          isAdding: false,
        },
      ];
      expectedKnownWorkflowsArr = [30, 20];

      await gateway.updateKnownWorkflows(updateKnownWorkflowsEntries);

      currentConfig = await gateway.getConfigInfo();

      currentConfig.knownWorkflows.forEach((el, index) => {
        expect(el).to.be.eq(expectedKnownWorkflowsArr[index]);
      });
    });

    it('should get exception if nonowner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(
        gateway.connect(FIRST).updateKnownWorkflows([
          {
            knownWorkflowToUpdate: 100,
            isAdding: true,
          },
        ])
      ).to.be.revertedWith(reason);
    });
  });

  describe('updateKnownWorkflowOwners', () => {
    it('should correctly update known workflows', async () => {
      let updateKnownWorkflowOwnersEntries = [
        {
          addressToUpdate: OWNER.address,
          isAdding: true,
        },
        {
          addressToUpdate: FIRST.address,
          isAdding: true,
        },
        {
          addressToUpdate: SECOND.address,
          isAdding: true,
        },
      ];
      let expectedKnownWorkflowOwnersArr = [OWNER.address, FIRST.address, SECOND.address];

      await gateway.updateKnownWorkflowOwners(updateKnownWorkflowOwnersEntries);

      let currentConfig = await gateway.getConfigInfo();

      expect(currentConfig.knownWorkflowOwners).to.be.deep.eq(expectedKnownWorkflowOwnersArr);

      updateKnownWorkflowOwnersEntries = [
        {
          addressToUpdate: OWNER.address,
          isAdding: false,
        },
        {
          addressToUpdate: SECOND.address,
          isAdding: false,
        },
      ];
      expectedKnownWorkflowOwnersArr = [FIRST.address];

      await gateway.updateKnownWorkflowOwners(updateKnownWorkflowOwnersEntries);

      currentConfig = await gateway.getConfigInfo();

      expect(currentConfig.knownWorkflowOwners).to.be.deep.eq(expectedKnownWorkflowOwnersArr);
    });

    it('should get exception if nonowner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(
        gateway.connect(FIRST).updateKnownWorkflowOwners([
          {
            addressToUpdate: FIRST.address,
            isAdding: true,
          },
        ])
      ).to.be.revertedWith(reason);
    });
  });

  describe('updateKnownCustomerContracts', () => {
    it('should correctly update known workflows', async () => {
      let updateKnownCustomerContractsEntries = [
        {
          addressToUpdate: OWNER.address,
          isAdding: true,
        },
        {
          addressToUpdate: FIRST.address,
          isAdding: true,
        },
        {
          addressToUpdate: SECOND.address,
          isAdding: true,
        },
      ];
      let expectedKnownCustomerContractsArr = [OWNER.address, FIRST.address, SECOND.address];

      await gateway.updateKnownCustomerContracts(updateKnownCustomerContractsEntries);

      let currentConfig = await gateway.getConfigInfo();

      expect(currentConfig.knownCustomerContracts).to.be.deep.eq(expectedKnownCustomerContractsArr);

      updateKnownCustomerContractsEntries = [
        {
          addressToUpdate: FIRST.address,
          isAdding: false,
        },
      ];
      expectedKnownCustomerContractsArr = [OWNER.address, SECOND.address];

      await gateway.updateKnownCustomerContracts(updateKnownCustomerContractsEntries);

      currentConfig = await gateway.getConfigInfo();

      expect(currentConfig.knownCustomerContracts).to.be.deep.eq(expectedKnownCustomerContractsArr);
    });

    it('should get exception if nonowner try to call this function', async () => {
      const reason = 'Ownable: caller is not the owner';

      await expect(
        gateway.connect(FIRST).updateKnownCustomerContracts([
          {
            addressToUpdate: FIRST.address,
            isAdding: true,
          },
        ])
      ).to.be.revertedWith(reason);
    });
  });
});
