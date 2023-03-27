import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";
import { Wallet } from "ethers";

describe("Registry", function () {
    // Get a provider to interact with the network
    const provider = ethers.getDefaultProvider();

    async function createWallet(balance: number) {
        // Generate a new random wallet
        const wallet = ethers.Wallet.createRandom().connect(provider);

        if (balance == 0 || true) {
            return wallet;
        }

        // TODO: Implement functionality to fund the balance

        // Contracts are deployed using the first signer/account by default
        const [owner, otherAccount] = await ethers.getSigners();

        // Send Ether to the new wallet
        const transaction = {
            to: wallet.address,
            value: ethers.utils.parseEther(balance.toString())
        };
        const tx = await owner.sendTransaction(transaction);
        const receipt = await tx.wait();
        console.log("owner balance",  await provider.getBalance(owner.address))
        console.log("transaction", transaction.value)

        // Log the wallet address and balance
        const walletBalance = await provider.getBalance(wallet.address);
        console.log(`Wallet Address: ${wallet.address}`);
        console.log(`Wallet Balance: ${ethers.utils.formatEther(walletBalance)} ETH`);

        return wallet;
    }

    async function deployRegistry(num: number) {
        let wallets: Wallet[] = [];
        let addresses: string[] = [];
        for (let i = 0; i < num; i++) {
            wallets[i] = await createWallet(1);
            addresses[i] = await wallets[i].getAddress();
        }

        const Registry = await ethers.getContractFactory("Registry");
        const registry = await Registry.deploy(addresses);

        return { registry, addresses };
    }

    describe("Deployment", function () {
        it("Successfully deployed with 4 nodes", async function () {
            const { registry, addresses } = await deployRegistry(4);

            expect(addresses.length).to.equal(4);
        });
    });

    describe("Fund balance", function () {
        it("Successfully funded the balance", async function () {
            const { registry, addresses } = await deployRegistry(4);

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
            const balance = await registry.getBalance(workflowOwnerAddress)
            console.log("balance", balance)
        });
    });
});
