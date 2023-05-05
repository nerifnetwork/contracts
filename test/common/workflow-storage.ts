import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('WorkflowStorage', function () {
  it('should  clear all workflows', async function () {
    const [network] = await ethers.getSigners();
    const workflows = [
      {
        id: '0x0000000000000000000000000000000000000001',
        owner: '0x0000000000000000000000000000000000000001',
        hash: '0x00',
        status: 1,
        isInternal: false,
        totalSpent: 0,
      },
      {
        id: '0x0000000000000000000000000000000000000002',
        owner: '0x0000000000000000000000000000000000000002',
        hash: '0x00',
        status: 1,
        isInternal: false,
        totalSpent: 0,
      },
    ];

    const WorkflowStorage = await ethers.getContractFactory('WorkflowStorage');
    const workflowStorage = await WorkflowStorage.deploy();
    await workflowStorage.initialize(network.address, workflows);

    expect(await workflowStorage.size()).to.equal(2);
    await workflowStorage.clear();
    expect(await workflowStorage.size()).to.equal(0);
  });

  it('should not add an already added workflow and not remove a non existing workflow', async function () {
    const [network] = await ethers.getSigners();
    const workflows = [
      {
        id: '0x0000000000000000000000000000000000000001',
        owner: '0x0000000000000000000000000000000000000001',
        hash: '0x00',
        status: 1,
        isInternal: false,
        totalSpent: 0,
      },
      {
        id: '0x0000000000000000000000000000000000000002',
        owner: '0x0000000000000000000000000000000000000002',
        hash: '0x00',
        status: 1,
        isInternal: false,
        totalSpent: 0,
      },
    ];

    const WorkflowStorage = await ethers.getContractFactory('WorkflowStorage');
    const workflowStorage = await WorkflowStorage.deploy();
    await workflowStorage.initialize(network.address, workflows);

    await expect(workflowStorage.mustAdd(workflows[0])).to.be.revertedWith('WorkflowStorage: failed to add workflow');
    await expect(workflowStorage.mustRemove('0x0000000000000000000000000000000000000003')).to.be.revertedWith(
      'WorkflowStorage: failed to remove workflow'
    );
  });
});
