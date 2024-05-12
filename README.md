# Nerif Network Contracts

This repository contains Solidity contracts needed to bootstrap and run Nerif Network.

![structure.png](./docs/structure.png)

## Core contracts

### ContractsRegistry
The **ContractsRegistry** contract is the central contract of the protocol and is used to store, retrieve and update existing system contracts. The [EIP-6224](https://eips.ethereum.org/EIPS/eip-6224) standard was used to implement the **ContractsRegistry**.

The **ContractsRegistry** contract can also be used to conveniently retrieve *collective address* and information about whether the current network is a main network.

### SignerOwnable
**SignerOwnable** is an auxiliary contract that implements common logic for the **Nerif** protocol, namely storing the `signerGetter` address of the contract and the `onlySigner` modifier.

## System contracts
System contracts are the central part of the protocol and will only be deployed on the main network. The `isMainChain` function of the **ContractsRegistry** contract can be used to find out which network is the main network.

### DKG
The **DKG** contract is one of the central system contracts. This contract is used to keep track of epochs, perform DKG ([Distributed Key Generation](https://eprint.iacr.org/2021/1591.pdf)) procedures, store a list of active validators and the current *collective address* of the protocol.

#### Epochs logic
The central mechanism on which most of the system is based is the logic of epochs. The main essence of this approach is one set of validators within one epoch. When this validator set is changed, the epoch is changed as well. Each epoch consists of periods.

##### Epochs periods
1. The period of guaranteed work - the period from which any epoch begins
2. Period of active work - as the name makes clear - this is the period when the network is active and there are no changes
3. Updates collection period - the period when the start time of the DKG generation period has already been determined and updates to the validator list are being collected
4. DKG generation period is intended to perform DKG generation for a new validator set

##### Epochs flow
Since the epoch system is implemented quite flexibly, it is necessary to accurately understand all of the existing options that can occur within an epoch.

As stated earlier, any epoch starts with a period of guaranteed working. During this period, validators can perform any available actions. For example, announce a withdrawal from a **Staking** contract, staking an additional amount of tokens, etc. If the actions performed lead to a future change in the validator set, the start of the DKG generation period will be determined, if it was not determined earlier.

The start time of the DKG generation period will be as follows:

```
dkgGenPeriodStartTime = end time of the guaranteed working period + duration of the updates collection period
```

After the end of the guaranteed working period, there can be several options:
1. In case the start of the DKG generation period was defined, that is, there were some changes in the validator set, then the next period will be the updates collection period
2. If there were no changes, then the active period will simply begin

During the active working period, just like during the guaranteed working period, validators can perform any action related to their stake. The only difference will be in the definition of the start time of the DKG generation period.

In this case, it will be defined as follows:

```
dkgGenPeriodStartTime = current time + duration of the updates collection period
```

During the updates collection period, all actions are available to validators, but only until the DKG generation period starts.

During a DKG generation period that starts at a previously defined time and lasts for a period of time that is set by the `dkgGenerationEpochDuration` variable from the **DKG** contract, any action by validators is **FORBIDDEN**. This restriction is necessary to specifically define the status of the validators.

After a successful DKG procedure, a new epoch will be created and the list of active validators will be updated.

In case no new *collective address* was generated during the DKG generation period, the epoch will return to active status after its end.

#### DKG procedure
The DKG procedure itself can only take place during the DKG generation period. To accept a new *collective address*, the majority of active validators must vote in favour of its acceptance. This vote is done using the `voteSigner` function of the **DKG** contract.

#### Initial validators set
The epoch logic for the very first epoch is slightly different from all subsequent epochs. Once the **DKG** contract is deployed and initialized, the first epoch is created, but there are no validators at that point. Those addresses that were added to the initial whitelist in the **Staking** contract will be able to stake tokens and become active validators immediately from the current epoch. Then the first DKG generation takes place, the first *collective address* is accepted and the system starts stable operation from the second epoch.

#### Slashing logic
Once an active validator has been slashed, it immediately loses the ability to be an active validator anymore. Accordingly, he will no longer be able to participate in accepting a new *collective address* and will stop receiving system rewards. At the same time, he will have an opportunity to withdraw the rewards earned earlier.

### Staking
Staking contract is used to define protocol validators. In order to become a validator it is necessary to stake **NERIF** tokens for an amount that is greater than the minimum required. Also the address that wants to stake tokens must be in the **whitelist**, which is defined inside the **Staking** contract.

#### Whitelist logic
In this contract, whitelist logic is used to limit the number of users that can stake tokens. The initial list of whitelisted users is defined during contract initialization and can be changed further by the *collective address*, i.e. by the protocol itself.

#### Staking logic
To stake tokens, users need to have the right amount of **ERC20** tokens on their balance. The token address in the Staking contract is set at initiali`ation and is never changed. A **NERIF** token will be used for staking in the **Nerif** protocol.

In order to become an active validator, users need to stake an amount greater than the minimum amount and wait for the next epoch. The minimum amount is defined in the **Staking** contract by the variable `minimalStake`, which can be changed by the protocol *collective address*.

Staking process is available during all epoch periods **BEFORE** the DKG generation period.

The number of staked tokens will also affect the number of system rewards the validator will receive for its work.

#### Slashing logic
In case a validator is slashed, its entire stake will be sent to the **BillingManager** contract, from where it will be distributed among other validators as their rewards. Also, the address that has been slashed is permanently banned from staking and withdrawing tokens from the **Staking** contract.

### Slashing voting
The **SlashingVoting** contract was created to penalise validators who do not do the work required of them. Using this contract, any active validator can create a proposal to slash a particular validator for a specific reason. If the vote is successful, the validator will be slashed.

#### Proposal creation
Only active validators can create slashing proposals. The validator, on whose slashing proposal is created, must also be active.

Proposals can be created in any period of the epoch except the period of DKG generation. When creating a proposal, by analogy with staking, the start time of the DKG generation period will be determined, if it was not there before.

Voting always starts immediately after its creation and lasts until the beginning of the DKG generation period.

#### Voting process
Only active validators can also vote. After a simple majority of votes of the current set of active validators, the automatic execution of the proposal, i.e. slashing of the validator, takes place.

## Operational contracts
Operational contracts are the primary contracts that users will interact with. These contracts will be deployed on all networks that will be supported by the **Nerif** protocol.

### BillingManager
The **BillingManager** contract is designed to collect the supported assets with which users should top up their balances. These balances will be used by the protocol to ensure the execution of users workflow and other operations that will be performed by the system.

Also from this contract validators will be able to withdraw the rewards they have earned.

#### Deposit assets
Deposit assets, are any Assets with which users can top up their balances. When initializing a **BillingManager** contract, it is mandatory that `NativeDepositAsset` and `NerifTokenDepositAsset` are set. These are two assets that **MUST** always be in the protocol.

During the lifetime of the protocol, the *collective address* can add new deposit assets, but cannot delete existing ones. When adding a new asset, it is necessary to specify its name in string format and information about the asset itself.

Asset information includes the following fields:
1. `tokenAddr` - The **ERC20** address of the token that users will deposit. For `NativeDepositAsset` this parameter is `0x0000000000000000000000000000000000000000`
2. `workflowExecutionDiscount` - The discount percentage that will be factored into the workflow execution cost
3. `totalAssetAmount` - The total number of tokens that have been staked. When added, this parameter **MUST** be 0.
4. `isPermitable` - Flag that indicates whether this address implements **Permit** logic.
5. `isEnabled` - Flag that indicates whether the asset is enabled.

*Collective address* has the ability to change the `workflowExecutionDiscount` and `isEnabled` fields for all assets.

#### Deposit process
As stated earlier, a deposit is required for users to fund their account, from which funds will be deducted for various actions performed by the system.

Any user can deposit any amount of supported assets if they are enabled at the time of deposit. Users can also deposit a tokens using **Permit** logic if the asset supports this feature.

After a successful deposit, an `AssetDeposited` event will be created, with the help of which the backend will count the total balance of users.

The main feature of these deposits is that each user's balance is not stored on a contract. Instead, all users simply deposit tokens into a common pool, and the backend of the system calculates all balances. This is to prevent users from withdrawing tokens during a workflow execution, thus making the system at a loss.

#### Funds withdrawal process
To withdraw deposited funds, all users must first obtain a signature from the protocol backend. This step is necessary because only the backend knows about the user's current balance, taking into account all the costs that have been incurred by the system to perform workflow and other system actions.

The signature is generated according to the **EIP712** standard and includes a custom structure of the following form:

```
StructTypehash = FundsWithdraw(address userAddr,bytes32 depositAssetsHash,bytes32 amountsHash,uint256 nonce,uint256 deadline)
```

1. `userAddr` - address of the user who should send the transaction with the given signature
2. `depositAssetsHash` - `keccak256` hash of the array of keys of the deposit asset keys
3. `amountsHash` - `keccak256` hash of the array of token amounts to be sent to the user
4. `nonce` - random user's nonce, which avoids signature reusing
5. `deadline` - timestamp until which the signature will be valid

Thanks to this signature structure it is possible to perform at once batch withdrawal.

#### Rewards withdrawal process

Withdrawal of system rewards follows a similar flow with withdrawal of deposited funds. Before withdrawal it is also necessary to get a signature from the backend, but for this case a slightly different signature structure is used.

```
StructTypehash = RewardsWithdraw(address userAddr,bytes32 depositAssetsHash,bytes32 amountsHash,uint256 nonce,uint256 deadline)
```

The fields of the signature structure are similar to the funds withdrawal.

### Registry
The **Registry** contract is used by the system as an entry point to execute transactions from user's **Gateway** contracts. Users, in turn, can manage their **Gateway** contracts and register new gateways.

#### System flow
The system, namely the *collective address*, can modify the maximum number of workflows that can be created by users on the main network. Also, as mentioned above, the system must execute user transactions through a special `perform` function.

#### Gateway logic
Each user has the ability to set up a **Gateway** contract that will accept calls from the **Registry** contract. This contract is a layer between the user contracts and the **Nerif** protocol, which is necessary for additional security. All calls to user contracts will go through the **Gateway** contract set by the user himself, which will allow to precisely define the address that can call specific functions.

#### User flow
Users can manage their **Gateway** contract as they wish. There are two ways to set the **Gateway** contract:
1. Use the `setGateway` function and set any desired address. This approach is good for those who have their own implementation of the **Gateway** contract
2. Use the `deployAndSetGateway` function, which will deploy a new **Gateway** contract instance that offers the **Nerif** protocol, and set the created instance to the desired address.

Also, any user can register new workflows. When registering, the user can specify whether a **Gateway** is needed for a particular wokflow. If a **Gateway** is needed and the user does not have one, another flag can be passed, which will automatically deploy a new **Gateway** contract.

### GatewayFactory
The **GatewayFactory** contract is needed to deploy new **Gateway** contracts. The deployment uses the **ProxyBeacon** pattern, where **BeaconProxy** contracts are deployed instead of the usual proxies. This approach allows all deployed proxy contracts to be updated in a single transaction. This is very good, as the **Nerif** protocol will be able to update all **Gateway** contracts that have been deployed through this factory to fix various bugs or to add new functionality.

### Geteway
The **Geteway** contract, as discussed earlier, is required to enhance the security of the interaction between the **Nerif** protocol and user contracts.

In the standard implementation from the **Nerif** protocol, the user can control the following settings:
1. Modify the list of known wokflows
2. Modify the list of known wokflow owners
3. Modify the list of known user contracts

These settings are used by the **Gateway** contract for additional checks before sending the transaction to the user contract.

### SignerStorage
**SignerStorage** is an auxiliary contract that is used off the main network to store and retrieve *collective addresses*.

## Set up environment

The first step is to initialize environment by running the following command:

```bash
$ make init
```

The command installs node packages, compiles contracts, and creates `.env` config file with default mocked values.
Those values must be replaced with the real ones.

### Run Tests

The following command runs tests:

```bash
$ npm run test
```

## Deployment

The current hardhat configuration supports 3 ENV-compatible networks: Goerli, Amoy, BSC Testnet.
More networks could be added if needed.

In this example, `polygonAmoy` is going to be a mainchain.
That means this network is going to be used for system contracts deployment as well as operational ones.
Other sidechains include operational contracts only. 

In order to deploy to another chain, `--network` parameter should be changed.

The following command deploys all contracts in the proper order on all supported chains:

```bash
$ make deploy
```

*Node: take a look at the Makefile*

After system contracts are deployed on mainchain, the script asks to start DKG process in order to set up a collective signer address.
This address is going to be used within another contracts across all supported chains.

During the network bootstrap, each node requests a DKG contract address in order to start the process.
Each network participant (Nerif Node) **must stake** a specific number of tokens in order to join the network, and it should be done before starting the node.
If the current validator have not staked, the node will ask for approval to stake tokens.

### Verify Contracts

To verify contracts, you need to specify network, contract address and constructor parameters (if present).

```
$ npx hardhat verify --network <network-name> <CONTRACT_ADDRESS> <CONSTRUCTOR_PARAMETERS>
```

Or you can set `VERIFY` variable to `true` while deploying contracts to automatically verify them afterwards.

```
$ VERIFY=true make deploy
```