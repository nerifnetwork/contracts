FROM ethereum/client-go:alltools-v1.10.17

WORKDIR /

COPY artifacts/ artifacts/

RUN mkdir common
RUN mkdir operational
RUN mkdir system
RUN mkdir interfaces

# Common contracts
RUN abigen --abi artifacts/abi/common/AddressStorage.sol/AddressStorage.json \
           --bin artifacts/bin/common/AddressStorage.sol/AddressStorage.bin \
           --pkg common \
           --type AddressStorage \
           --out common/addressstorage.go
RUN abigen --abi artifacts/abi/common/ContractRegistry.sol/ContractRegistry.json \
           --bin artifacts/bin/common/ContractRegistry.sol/ContractRegistry.bin \
           --pkg common \
           --type ContractRegistry \
           --out common/contractregistry.go

# Operational contracts
RUN abigen --abi artifacts/abi/operational/Gateway.sol/Gateway.json \
           --bin artifacts/bin/operational/Gateway.sol/Gateway.bin \
           --pkg operational \
           --type Gateway \
           --out operational/gateway.go
RUN abigen --abi artifacts/abi/operational/GatewayStorage.sol/GatewayStorage.json \
           --bin artifacts/bin/operational/GatewayStorage.sol/GatewayStorage.bin \
           --pkg operational \
           --type GatewayStorage \
           --out operational/gatewaystorage.go
RUN abigen --abi artifacts/abi/operational/Registry.sol/Registry.json \
           --bin artifacts/bin/operational/Registry.sol/Registry.bin \
           --pkg operational \
           --type Registry \
           --out operational/registry.go
RUN abigen --abi artifacts/abi/operational/SignerStorage.sol/SignerStorage.json \
           --bin artifacts/bin/operational/SignerStorage.sol/SignerStorage.bin \
           --pkg operational \
           --type SignerStorage \
           --out operational/signerstorage.go
RUN abigen --abi artifacts/abi/operational/WorkflowStorage.sol/WorkflowStorage.json \
           --bin artifacts/bin/operational/WorkflowStorage.sol/WorkflowStorage.bin \
           --pkg operational \
           --type WorkflowStorage \
           --out operational/workflowstorage.go

# System contracts
RUN abigen --abi artifacts/abi/system/DKG.sol/DKG.json \
           --bin artifacts/bin/system/DKG.sol/DKG.bin \
           --pkg system \
           --type DKG \
           --out system/dkg.go
RUN abigen --abi artifacts/abi/system/SlashingVoting.sol/SlashingVoting.json \
           --bin artifacts/bin/system/SlashingVoting.sol/SlashingVoting.bin \
           --pkg system \
           --type SlashingVoting \
           --out system/slashingvoting.go
RUN abigen --abi artifacts/abi/system/Staking.sol/Staking.json \
           --bin artifacts/bin/system/Staking.sol/Staking.bin \
           --pkg system \
           --type Staking \
           --out system/staking.go
RUN abigen --abi artifacts/abi/system/RewardDistributionPool.sol/RewardDistributionPool.json \
           --bin artifacts/bin/system/RewardDistributionPool.sol/RewardDistributionPool.bin \
           --pkg system \
           --type RewardDistributionPool \
           --out system/rewarddistributionpool.go

# Interfaces
RUN abigen --abi artifacts/abi/interfaces/ISignerAddress.sol/ISignerAddress.json \
           --bin artifacts/bin/interfaces/ISignerAddress.sol/ISignerAddress.bin \
           --pkg interfaces \
           --type ISignerAddress \
           --out interfaces/isigneraddress.go
RUN abigen --abi artifacts/abi/interfaces/IGateway.sol/IGateway.json \
           --bin artifacts/bin/interfaces/IGateway.sol/IGateway.bin \
           --pkg interfaces \
           --type IGateway \
           --out interfaces/igateway.go
RUN abigen --abi artifacts/abi/interfaces/SignerOwnable.sol/SignerOwnable.json \
           --bin artifacts/bin/interfaces/SignerOwnable.sol/SignerOwnable.bin \
           --pkg interfaces \
           --type SignerOwnable \
           --out interfaces/signerownable.go
