import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect, assert } from "chai";
import { ethers } from "hardhat";
import {Wallet, BigNumber, Signer} from "ethers";
import {SignerWithAddress} from "@nomiclabs/hardhat-ethers/signers";

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
            .withArgs(1, 43488, true);

            // Make sure the node got registered
            const expectedNode = await newNodeWallet.getAddress();
            expect(await registry.getActiveNodes()).to.include(expectedNode);
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
