import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";
import {Wallet, BigNumber, Signer} from "ethers";
import {SignerWithAddress} from "@nomiclabs/hardhat-ethers/signers";

const testCustomerContractABI = require("../artifacts/contracts/TestCustomerContract.sol/TestCustomerContract.json");

describe("Registry", function () {
    async function deployRegistry(num: number) {
        const [owner] = await ethers.getSigners();

        let wallets: Signer[] = [owner];
        let addresses: string[] = [owner.address];
        for (let i = 1; i < num; i++) {
            wallets[i] = await ethers.Wallet.createRandom();
            addresses[i] = await wallets[i].getAddress();
        }

        const Registry = await ethers.getContractFactory("Registry");
        const registry = await Registry.deploy(addresses, true);

        return { registry, wallets };
    }

    async function deployTestCustomerContract() {
        const TestCustomerContract = await ethers.getContractFactory("TestCustomerContract");
        const testCustomerContract = await TestCustomerContract.deploy();

        return { testCustomerContract };
    }

    describe("Deployment", function () {
        it("Successfully deployed with 4 nodes", async function () {
            const { registry, wallets } = await deployRegistry(4);

            expect(wallets.length).to.equal(4);
        });
    });

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
