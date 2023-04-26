/*
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
        const registry = await Registry.deploy(
            addresses,
            mainchain,
        );
        expect(await registry.isMainChain()).to.equal(mainchain);

        await expect(registry.setConfig({
            performanceOverhead: 71157,
            performancePremiumThreshold: 0,
            registrationOverhead: 0,
            cancellationOverhead: 0,
        }))

        return { registry, wallets };
    }

    async function deployTestCustomerContract() {
        const TestCustomerContract = await ethers.getContractFactory("TestCustomerContract");
        const testCustomerContract = await TestCustomerContract.deploy();

        return { testCustomerContract };
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

    async function activateWorkflow(registry: Registry, signer: Signer, workflowID: number) {
        // Create a contract ABI instance
        const cc = new ethers.Contract(registry.address, registryContractABI.abi, signer)

        // Encode the function call with the selector
        const functionCallData = cc.interface.encodeFunctionData(
            cc.interface.getSighash("activateWorkflow"), [
                ethers.utils.solidityPack(
                    ['uint256'],
                    [workflowID]
                ),
            ]);

        // Activate the workflow
        await expect(
            registry
                .connect(signer)
                .perform('40505927788353901442144037336646356013', 300_000, functionCallData, registry.address, [])
        )
            .to.emit(registry, "Performance")
            .withArgs('40505927788353901442144037336646356013', anyValue, true);
    }

    describe("Deployment", function () {
        it("Successfully deployed with 4 nodes", async function () {
            const { registry, wallets } = await deployRegistry(4, true);
            expect(wallets.length).to.equal(4);
        });
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

            // Status must be PENDING on MAINCHAIN
            const workflow = await registry.getWorkflow(123);
            expect(workflow.status).to.equal(0);
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
                    .perform('40505927788353901442144037336646356013', 300_000, functionCallData, registry.address, [])
            )
            .to.emit(registry, "Performance")
            .withArgs('40505927788353901442144037336646356013', anyValue, true);

            // Status must be ACTIVE on SIDECHAIN
            const workflow = await registry.getWorkflow(123);
            expect(workflow.status).to.equal(1);
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
            .withArgs(nodeAddress, 123, [], []);

            // Activate workflow
            await activateWorkflow(registry, wallets[0], 123);

            // Pause workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .pauseWorkflow(123)
            )
            .to.emit(registry, "WorkflowPaused")
            .withArgs(123);
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

            // Activate workflow
            await activateWorkflow(registry, wallets[0], 123);

            // Pause workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .pauseWorkflow(123)
            )
            .to.emit(registry, "WorkflowPaused")
            .withArgs(123);

            // Resume workflow
            await expect(
                registry
                    .connect(newNodeWallet.connect(ethers.provider))
                    .resumeWorkflow(123)
            )
            .to.emit(registry, "WorkflowResumed")
            .withArgs(123);
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
            .to.emit(registry, "WorkflowCancelled")
            .withArgs(123);
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
                    .perform('40505927788353901442144037336646356013', 300_000, functionCallData, registry.address, [])
            )
            .to.emit(registry, "Performance")
            .withArgs('40505927788353901442144037336646356013', anyValue, true);

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
                    .perform('219775546284901721155783592958414245131', 300_000, functionCallData, registry.address, [])
            )
                .to.emit(registry, "Performance")
                .withArgs('219775546284901721155783592958414245131', anyValue, true);
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

            // Activate workflow
            await activateWorkflow(registry, wallets[0], 123);

            // Create a contract ABI instance
            const cc = new ethers.Contract(testCustomerContract.address, testCustomerContractABI.abi, workflowOwner)

            const customerPerformData = ethers.utils.solidityPack(
                ['uint256', 'uint', 'uint256'],
                [0, 0, 100000]
            )

            // Encode the function call with the selector
            const functionSignature = cc.interface.getSighash("perform");
            const functionCallData = cc.interface.encodeFunctionData(functionSignature, [customerPerformData]);

            await expect(workflowOwner.sendTransaction({
                        to: testCustomerContract.address,
                        data: functionCallData,
                        gasLimit: 500_000
                    }))
                .to.emit(testCustomerContract, "Logger");

            const tx = await registry
                    .connect(wallets[0])
                    .perform(123, 300_000, functionCallData, testCustomerContract.address,[]);
            const receipt = await tx.wait();
            // console.log(BigNumber.from("12000000096").div(receipt.effectiveGasPrice))

            await expect(tx)
                .to.emit(registry, "Performance")
                .withArgs(123, receipt.cumulativeGasUsed.mul(receipt.effectiveGasPrice), true);
        })
    })
});
*/
