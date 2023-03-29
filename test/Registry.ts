import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";
import {Signer, Wallet} from "ethers";
import {Registry} from "../typechain-types";

const testCustomerContractABI = require("../artifacts/contracts/TestCustomerContract.sol/TestCustomerContract.json");
const registryContractABI = require("../artifacts/contracts/Registry.sol/Registry.json");

describe("Registry", function () {
    async function deployRegistry(num: number, mainchain: boolean = true) {
        const [owner] = await ethers.getSigners();

        let wallets: Signer[] = [owner];
        let addresses: string[] = [owner.address];
        for (let i = 1; i < num; i++) {
            wallets[i] = await ethers.Wallet.createRandom();
            addresses[i] = await wallets[i].getAddress();

            await expect(
                owner.sendTransaction({
                    to: addresses[i],
                    value: ethers.utils.parseEther("10"),
                    gasLimit: 500_000
                })
            )
            .to.not.be.reverted;
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
            .withArgs(1, 43488, true);

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

    describe("Deployment", function () {
        it("Successfully deployed with 4 nodes", async function () {
            const { registry, wallets } = await deployRegistry(4, true);
            expect(wallets.length).to.equal(4);
        });
    });

    describe("Register node", function () {
        it("MAINCHAIN: existing node can add a new one", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await ethers.Wallet.createRandom();

            await expect(
                registry
                    .connect(wallets[0])
                    .registerNode(newNodeWallet.getAddress())
            )
                .to.emit(registry, "NodeRegistered")
                .withArgs(await newNodeWallet.getAddress());
        })

        it("MAINCHAIN: unknown node cannot add a new one", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const [_, other] = await ethers.getSigners();
            const newNodeWallet = await ethers.Wallet.createRandom();

            await expect(
                registry
                    .connect(other)
                    .registerNode(newNodeWallet.getAddress())
            )
            .to.be.revertedWith("Operation is not permitted");
        })

        it("MAINCHAIN: network cannot add a new node", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await ethers.Wallet.createRandom();

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
            .withArgs(1, 12100, false);

            // Make sure the node is not registered
            const expectedNode = await newNodeWallet.getAddress();
            expect(await registry.getActiveNodes()).to.not.include(expectedNode);
        })

        it("SIDECHAIN: existing node cannot add a new one", async function () {
            const { registry, wallets } = await deployRegistry(4, false);

            const newNodeWallet = await ethers.Wallet.createRandom();

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
            const newNodeWallet = await ethers.Wallet.createRandom();

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

            const newNodeWallet = await ethers.Wallet.createRandom();

            // Register node on the sidechain
            await registerNodeOnSidechain(registry, wallets[0], newNodeWallet)
        })
    })

    describe("Approve node registration", function () {
        it("MAINCHAIN: node operator can approve the node", async function () {
            const { registry, wallets } = await deployRegistry(4, true);

            const newNodeWallet = await ethers.Wallet.createRandom();
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
            const newNodeWallet = await ethers.Wallet.createRandom();
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
            const newNodeWallet = await ethers.Wallet.createRandom();
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

            const newNodeWallet = await ethers.Wallet.createRandom();
            const nodeAddress = await newNodeWallet.getAddress();

            await expect(
                wallets[0].sendTransaction({
                    to: nodeAddress,
                    value: ethers.utils.parseEther("10"),
                    gasLimit: 500_000
                })
            )
                .to.not.be.reverted;

            // Register node
            expect(
                registry
                    .connect(wallets[0])
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
            expect(await registry.getActiveNodes()).to.include(nodeAddress);

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

            const newNodeWallet = await ethers.Wallet.createRandom();
            const nodeAddress = await newNodeWallet.getAddress();

            await expect(
                wallets[0].sendTransaction({
                    to: nodeAddress,
                    value: ethers.utils.parseEther("10"),
                    gasLimit: 500_000
                })
            )
                .to.not.be.reverted;

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

            const newNodeWallet = await ethers.Wallet.createRandom();
            const nodeAddress = await newNodeWallet.getAddress();

            await expect(
                wallets[0].sendTransaction({
                    to: nodeAddress,
                    value: ethers.utils.parseEther("10"),
                    gasLimit: 500_000
                })
            )
                .to.not.be.reverted;

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

            const newNodeWallet = await ethers.Wallet.createRandom();
            const nodeAddress = await newNodeWallet.getAddress();

            await expect(
                wallets[0].sendTransaction({
                    to: nodeAddress,
                    value: ethers.utils.parseEther("10"),
                    gasLimit: 500_000
                })
            )
                .to.not.be.reverted;

            // Failed to unregister node
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .unregisterNode(await wallets[0].getAddress())
            )
                .to.be.revertedWith("Operation is not permitted")
        })
    })

    describe("Fund balance", function () {
        it("Successfully funded the balance", async function () {
            const { registry, wallets } = await deployRegistry(4);

            // Contracts are deployed using the first signer/account by default
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

            // Check the balance
            // const balance = await registry.getBalance(workflowOwnerAddress)
            // console.log("balance", balance)
        });
    });

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
                .withArgs(123, 81483);
        })
    })
});
