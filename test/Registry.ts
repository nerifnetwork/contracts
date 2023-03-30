import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";
import {BigNumber, Signer, Wallet} from "ethers";
import {Registry} from "../typechain-types";

const testCustomerContractABI = require("../artifacts/contracts/TestCustomerContract.sol/TestCustomerContract.json");
const registryContractABI = require("../artifacts/contracts/Registry.sol/Registry.json");

describe("Registry", function () {
    async function createRandomWallet() {
        const [owner] = await ethers.getSigners();
        const wallet = await ethers.Wallet.createRandom();

        await expect(
            owner.sendTransaction({
                to: await wallet.getAddress(),
                value: ethers.utils.parseEther("10"),
                gasLimit: 500_000
            })
        )

        return wallet;
    }

    async function deployRegistry(num: number, mainchain: boolean = true) {
        const [owner] = await ethers.getSigners();

        let wallets: Signer[] = [owner];
        let addresses: string[] = [owner.address];
        for (let i = 1; i < num; i++) {
            wallets[i] = await createRandomWallet();
            addresses[i] = await wallets[i].getAddress();
        }

        const Registry = await ethers.getContractFactory("Registry");
        const registry = await Registry.deploy(addresses, mainchain);
        expect(await registry.isMainChain()).to.equal(mainchain);

        return { registry, wallets };
    }

    async function deployTestCustomerContract() {
        const TestCustomerContract = await ethers.getContractFactory("TestCustomerContract");
        const testCustomerContract = await TestCustomerContract.deploy();

        return { testCustomerContract };
    }

    async function registerNodeOnSidechain(registry: Registry, signer: Signer, newNode: Wallet) {
        // Create a contract ABI instance
        const cc = new ethers.Contract(registry.address, registryContractABI.abi, signer)

        const performData = ethers.utils.solidityPack(
            ['address'],
            [await newNode.getAddress()]
        )

        // Encode the function call with the selector
        const functionSignature = cc.interface.getSighash("registerNode");
        const functionCallData = cc.interface.encodeFunctionData(functionSignature, [performData]);

        // Network can add a new node
        await expect(
            registry
                .connect(signer)
                .perform(1, 300_000, functionCallData, registry.address, [])
        )
            .to.emit(registry, "Performance")
            .withArgs(1, anyValue, true);

        // Make sure the node got registered
        const expectedNode = await newNode.getAddress();
        expect(await registry.getActiveNodes()).to.include(expectedNode);
    }

    async function unregisterNodeOnSidechain(registry: Registry, signer: Signer, newNode: Wallet) {
        // Create a contract ABI instance
        const cc = new ethers.Contract(registry.address, registryContractABI.abi, signer)

        const performData = ethers.utils.solidityPack(
            ['address'],
            [await newNode.getAddress()]
        )

        // Encode the function call with the selector
        const functionSignature = cc.interface.getSighash("unregisterNode");
        const functionCallData = cc.interface.encodeFunctionData(functionSignature, [performData]);

        // Network can add a new node
        await expect(
            registry
                .connect(signer)
                .perform(1, 300_000, functionCallData, registry.address, [])
        )
            .to.emit(registry, "Performance")
            .withArgs(1, 22572, true);

        // Make sure the node got registered
        const expectedNode = await newNode.getAddress();
        expect(await registry.getActiveNodes()).to.not.include(expectedNode);
    }

    async function fundBalance(registry: Registry, signer: Signer, fundWallet: Wallet, amount: BigNumber) {
        const nodeAddress = await fundWallet.getAddress();

        await expect(
            registry
                .connect(fundWallet.connect(ethers.provider))
                .fundBalance(nodeAddress, { value: amount })
        )
            .to.emit(registry, "BalanceFunded")
            .withArgs(nodeAddress, amount);

        await expect(await registry.getBalance(nodeAddress))
            .to.equal(amount);
    }

    describe("Deployment", function () {
        it("Successfully deployed with 4 nodes", async function () {
            const { registry, wallets } = await deployRegistry(4, true);
            expect(wallets.length).to.equal(4);
        });
    });

    describe("Register node", function () {
        it("MAINCHAIN: existing node cannot add a new one", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();

            await expect(
                registry
                    .connect(wallets[0])
                    .registerNode(newNodeWallet.getAddress())
            )
            .to.be.revertedWith("Operation is not permitted");
        })

        it("MAINCHAIN: unknown node can add itself", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();

            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerNode(newNodeWallet.getAddress())
            )
            .to.emit(registry, "NodeRegistered")
            .withArgs(await newNodeWallet.getAddress());
        })

        it("MAINCHAIN: network cannot add a new node", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();

            // Create a contract ABI instance
            const cc = new ethers.Contract(registry.address, registryContractABI.abi, wallets[0])

            const registerNodePerformData = ethers.utils.solidityPack(
                ['address'],
                [await newNodeWallet.getAddress()]
            )

            // Encode the function call with the selector
            const functionSignature = cc.interface.getSighash("registerNode");
            const functionCallData = cc.interface.encodeFunctionData(functionSignature, [registerNodePerformData]);

            // Network can add a new node
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(1, 300_000, functionCallData, registry.address, [])
            )
            .to.emit(registry, "Performance")
            .withArgs(1, anyValue, false);

            // Make sure the node is not registered
            const expectedNode = await newNodeWallet.getAddress();
            expect(await registry.getActiveNodes()).to.not.include(expectedNode);
        })

        it("SIDECHAIN: existing node cannot add a new one", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();

            // Node operator cannot add a new node directly
            await expect(
                registry
                    .connect(wallets[0])
                    .registerNode(newNodeWallet.getAddress())
            )
            .to.be.revertedWith("Operation is not permitted");
        })

        it("SIDECHAIN: unknown node cannot add a new one", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const [_, other] = await ethers.getSigners();
            const newNodeWallet = await createRandomWallet();

            // Node operator cannot add a new node directly
            await expect(
                registry
                    .connect(other)
                    .registerNode(newNodeWallet.getAddress())
            )
            .to.be.revertedWith("Operation is not permitted");
        })

        it("SIDECHAIN: network add a new node", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();

            // Register node on the sidechain
            await registerNodeOnSidechain(registry, wallets[0], newNodeWallet)
        })
    })

    describe("Approve node registration", function () {
        it("MAINCHAIN: node operator can approve the node", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register node
            expect(
                registry
                    .connect(wallets[0])
                    .registerNode(nodeAddress)
            )
                .to.emit(registry, "NodeRegistered")
                .withArgs(nodeAddress);

            // Approve node
            expect(
                registry
                    .connect(wallets[0])
                    .approveRegistration(nodeAddress)
            )
            .to.emit(registry, "NodeApproved")
            .withArgs(nodeAddress);
        })

        it("MAINCHAIN: unknown node cannot approve the node", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const [_, other] = await ethers.getSigners();
            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register node
            expect(
                registry
                    .connect(wallets[0])
                    .registerNode(nodeAddress)
            )
                .to.emit(registry, "NodeRegistered")
                .withArgs(nodeAddress);

            // Approve node
            expect(
                registry
                    .connect(other)
                    .approveRegistration(nodeAddress)
            )
            .to.be.revertedWith("Operation is not permitted");
        })

        it("SIDECHAIN: operation is not permitted", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const [_, other] = await ethers.getSigners();
            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Approve node by other node
            expect(
                registry
                    .connect(other)
                    .approveRegistration(nodeAddress)
            )
            .to.be.revertedWith("Operation is not permitted");

            // Approve node by existing node
            expect(
                registry
                    .connect(wallets[0])
                    .approveRegistration(nodeAddress)
            )
            .to.be.revertedWith("Operation is not permitted");
        })
    })

    describe("Unregister node", function () {
        it("MAINCHAIN: active node can unregister itself", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register node
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerNode(nodeAddress)
            )
                .to.emit(registry, "NodeRegistered")
                .withArgs(nodeAddress);

            // Get 2/3 number of approvals
            const approvals = wallets.length / 3 * 2;

            // Approve node by network
            for (let i = 1; i < approvals; i++) {
                await expect(
                    registry
                        .connect(wallets[i].connect(ethers.provider))
                        .approveRegistration(nodeAddress)
                )
                .to.emit(registry, "NodeApproved")
                .withArgs(nodeAddress, await wallets[i].getAddress());
            }

            // Make sure the node is active
            await expect(await registry.getActiveNodes()).to.include(nodeAddress);

            // Unregister node
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .unregisterNode(nodeAddress)
            )
                .to.emit(registry, "NodeUnregistered")
                .withArgs(nodeAddress);
        })

        it("MAINCHAIN: active node cannot unregister other node", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            // Failed to unregister node
            await expect(
                registry
                    .connect(wallets[1].connect(ethers.provider))
                    .unregisterNode(await wallets[0].getAddress())
            )
                .to.be.revertedWith("Operation is not permitted")
        })

        it("MAINCHAIN: unknown node cannot unregister existing node", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Failed to unregister node
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .unregisterNode(await wallets[0].getAddress())
            )
                .to.be.revertedWith("Operation is not permitted")
        })

        it("SIDECHAIN: network can unregister the node", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register node on the sidechain
            await registerNodeOnSidechain(registry, wallets[0], newNodeWallet)

            // Unregister node on the sidechain
            await unregisterNodeOnSidechain(registry, wallets[0], newNodeWallet)
        })

        it("SIDECHAIN: existing node cannot unregister itself", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            // Failed to unregister node
            await expect(
                registry
                    .connect(wallets[0])
                    .unregisterNode(await wallets[0].getAddress())
            )
                .to.be.revertedWith("Operation is not permitted")
        })

        it("SIDECHAIN: unknown node cannot unregister active node", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Failed to unregister node
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .unregisterNode(await wallets[0].getAddress())
            )
                .to.be.revertedWith("Operation is not permitted")
        })
    })

    describe("Cancel node registration", function () {
        it("MAINCHAIN: pending node can cancel registration", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register node
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerNode(nodeAddress)
            )
                .to.emit(registry, "NodeRegistered")
                .withArgs(nodeAddress);

            // Cancel node registration
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .cancelNodeRegistration(nodeAddress)
            )
                .to.emit(registry, "NodeRegistrationCancelled")
                .withArgs(nodeAddress);
        })

        it("MAINCHAIN: other node cannot cancel pending node registration", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register node
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerNode(nodeAddress)
            )
                .to.emit(registry, "NodeRegistered")
                .withArgs(nodeAddress);

            // Cancel node registration
            await expect(
                registry
                    .connect(wallets[0])
                    .cancelNodeRegistration(nodeAddress)
            )
            .to.be.revertedWith("Operation is not permitted")
        })

        it("SIDECHAIN: operation is not permitted", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            await expect(
                registry
                    .connect(wallets[0])
                    .cancelNodeRegistration(nodeAddress)
            )
            .to.be.revertedWith("Operation is not permitted")

            await expect(
                registry
                    .connect(wallets[0])
                    .cancelNodeRegistration(await wallets[0].getAddress())
            )
            .to.be.revertedWith("Operation is not permitted")
        })
    });

    describe("Fund balance", function () {
        async function fundBalanceOnChain(mainchain: boolean) {
            const { registry, wallets } = await deployRegistry(4, mainchain);
            const newNodeWallet = await createRandomWallet();
            const amount = ethers.utils.parseEther("1")
            await fundBalance(registry, wallets[0], newNodeWallet, amount);
        }

        it("MAINCHAIN: anyone can fund balance", async function () {
            await fundBalanceOnChain(true)
        })

        it("SIDECHAIN: anyone can fund balance", async function () {
            await fundBalanceOnChain(false)
        })
    });

    describe("Withdraw balance", function () {
        it("MAINCHAIN: anyone can withdraw balance", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const newNodeWalletAddr = await newNodeWallet.getAddress();
            const amount = ethers.utils.parseEther("1")

            await fundBalance(registry, wallets[0], newNodeWallet, amount)

            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .withdrawBalance(newNodeWalletAddr)
            )
                .to.emit(registry, "BalanceWithdrawn")
                .withArgs(newNodeWalletAddr, amount);
        })

        it("SIDECHAIN: anyone can withdraw balance", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const newNodeWalletAddr = await newNodeWallet.getAddress();
            const amount = ethers.utils.parseEther("1")

            await fundBalance(registry, wallets[0], newNodeWallet, amount)

            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .withdrawBalance(newNodeWalletAddr)
            )
                .to.emit(registry, "BalanceWithdrawn")
                .withArgs(newNodeWalletAddr, amount);
        })
    });

    describe("Register workflow", function () {
        it("MAINCHAIN: workflow owner can register workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerWorkflow(123, nodeAddress, [], [])
            )
            .to.emit(registry, "WorkflowRegistered")
            .withArgs(nodeAddress, "123", [], []);
        })

        it("SIDECHAIN: network can register workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Create a contract ABI instance
            const cc = new ethers.Contract(registry.address, registryContractABI.abi, wallets[0])

            // Encode the function call with the selector
            const functionSignature = cc.interface.getSighash("registerWorkflow");
            const functionCallData = cc.interface.encodeFunctionData(functionSignature, [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
                ethers.utils.solidityPack(
                    ['address'],
                    [nodeAddress]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                )
            ]);

            await expect(
                registry
                    .connect(wallets[0])
                    .perform(3, 300_000, functionCallData, registry.address, [])
            )
            .to.emit(registry, "Performance")
            .withArgs(3, 100940, true);
        })
    })

    describe("Pause workflow", function () {
        it("MAINCHAIN: workflow owner can pause workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerWorkflow(123, nodeAddress, [], [])
            )
            .to.emit(registry, "WorkflowRegistered")
            .withArgs(nodeAddress, "123", [], []);

            // Pause workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .pauseWorkflow(123)
            )
            .to.emit(registry, "ChangeWorkflowStatus")
            .withArgs(123, 1);
        })

        it("SIDECHAIN: network can pause workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Create a contract ABI instance
            const cc = new ethers.Contract(registry.address, registryContractABI.abi, wallets[0])

            // Encode the function call with the selector
            let functionCallData = cc.interface.encodeFunctionData(
                cc.interface.getSighash("registerWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
                ethers.utils.solidityPack(
                    ['address'],
                    [nodeAddress]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                )
            ]);

            // Register workflow
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(3, 300_000, functionCallData, registry.address, [])
            )
            .to.emit(registry, "Performance")
            .withArgs(3, 100940, true);

            // Encode the function call with the selector
            functionCallData = cc.interface.encodeFunctionData(
                cc.interface.getSighash("pauseWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
            ]);

            // Pause the workflow
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(4, 300_000, functionCallData, registry.address, [])
            )
                .to.emit(registry, "Performance")
                .withArgs(4, 27245, true);
        })
    })

    describe("Resume workflow", function () {
        it("MAINCHAIN: workflow owner can resume workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerWorkflow(123, nodeAddress, [], [])
            )
            .to.emit(registry, "WorkflowRegistered")
            .withArgs(nodeAddress, "123", [], []);

            // Pause workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .pauseWorkflow(123)
            )
            .to.emit(registry, "ChangeWorkflowStatus")
            .withArgs(123, 1);

            // Resume workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .resumeWorkflow(123)
            )
            .to.emit(registry, "ChangeWorkflowStatus")
            .withArgs(123, 0);
        })

        it("SIDECHAIN: network can resume workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Create a contract ABI instance
            const cc = new ethers.Contract(registry.address, registryContractABI.abi, wallets[0])

            // Encode the function call with the selector
            let functionCallData = cc.interface.encodeFunctionData(
                cc.interface.getSighash("registerWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
                ethers.utils.solidityPack(
                    ['address'],
                    [nodeAddress]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                )
            ]);

            // Register workflow
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(3, 300_000, functionCallData, registry.address, [])
            )
            .to.emit(registry, "Performance")
            .withArgs(3, 100940, true);

            // Encode the function call with the selector
            functionCallData = cc.interface.encodeFunctionData(
                cc.interface.getSighash("pauseWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
            ]);

            // Pause the workflow
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(4, 300_000, functionCallData, registry.address, [])
            )
                .to.emit(registry, "Performance")
                .withArgs(4, 27245, true);

            // Encode the function call with the selector
            functionCallData = cc.interface.encodeFunctionData(
                cc.interface.getSighash("resumeWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
            ]);

            // Resume the workflow
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(4, 300_000, functionCallData, registry.address, [])
            )
                .to.emit(registry, "Performance")
                .withArgs(4, 10187, true);
        })
    })

    describe("Cancel workflow", function () {
        it("MAINCHAIN: workflow owner can cancel workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Register workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .registerWorkflow(123, nodeAddress, [], [])
            )
            .to.emit(registry, "WorkflowRegistered")
            .withArgs(nodeAddress, "123", [], []);

            // Cancel workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .cancelWorkflow(123)
            )
            .to.emit(registry, "ChangeWorkflowStatus")
            .withArgs(123, 2);
        })

        it("SIDECHAIN: network can cancel workflow", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await createRandomWallet();
            const nodeAddress = await newNodeWallet.getAddress();

            // Create a contract ABI instance
            const cc = new ethers.Contract(registry.address, registryContractABI.abi, wallets[0])

            // Encode the function call with the selector
            let functionCallData = cc.interface.encodeFunctionData(
                cc.interface.getSighash("registerWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
                ethers.utils.solidityPack(
                    ['address'],
                    [nodeAddress]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                ),
                ethers.utils.solidityPack(
                    ['bytes'],
                    ["0x00"]
                )
            ]);

            // Register workflow
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(3, 300_000, functionCallData, registry.address, [])
            )
            .to.emit(registry, "Performance")
            .withArgs(3, anyValue, true);

            // Encode the function call with the selector
            functionCallData = cc.interface.encodeFunctionData(
                cc.interface.getSighash("cancelWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [123]
                ),
            ]);

            // Cancel the workflow
            await expect(
                registry
                    .connect(wallets[0])
                    .perform(4, 300_000, functionCallData, registry.address, [])
            )
                .to.emit(registry, "Performance")
                .withArgs(4, anyValue, true);
        })
    })

    describe("Perform", function () {
        it("Execute customer contract", async () => {
            // Deploy contracts
            const { registry, wallets } = await deployRegistry(4);
            const { testCustomerContract } = await deployTestCustomerContract();

            // Fund balance
            const [workflowOwner] = await ethers.getSigners();
            const workflowOwnerAddress = await workflowOwner.getAddress();

            const amount = ethers.utils.parseEther("1");
            await expect(
                    registry
                        .connect(workflowOwner)
                        .fundBalance(workflowOwnerAddress, { value: amount })
                )
                .to.emit(registry, "BalanceFunded")
                .withArgs(workflowOwnerAddress, anyValue);

            // Register workflow
            await expect(
                    registry
                        .connect(workflowOwner)
                        .registerWorkflow(123, workflowOwnerAddress, [], [])
                )
                .to.emit(registry, "WorkflowRegistered")
                .withArgs(workflowOwnerAddress, "123", [], []);

            // Create a contract ABI instance
            const cc = new ethers.Contract(testCustomerContract.address, testCustomerContractABI.abi, workflowOwner)

            const customerPerformData = ethers.utils.solidityPack(
                ['uint256', 'uint', 'uint256'],
                [0, 0, 100000]
            )

            // Encode the function call with the selector
            const functionSignature = cc.interface.getSighash("perform");
            const functionCallData = cc.interface.encodeFunctionData(functionSignature, [customerPerformData]);

            await expect(
                    workflowOwner.sendTransaction({
                        to: testCustomerContract.address,
                        data: functionCallData,
                        gasLimit: 500_000
                    })
                )
                .to.emit(testCustomerContract, "Logger");

            await expect(
                    registry
                        .connect(wallets[0])
                        .perform(123, 300_000, functionCallData, testCustomerContract.address,[])
                )
                .to.emit(registry, "Performance")
                .withArgs(123, 81483, true);
        })
    })
});
