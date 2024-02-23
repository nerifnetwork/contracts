.PHONY: help
help: # Display this help
	@awk 'BEGIN{FS=":.*#";printf "Usage:\n  make <target>\n\nTargets:\n"}/^[a-zA-Z_-]+:.*?#/{printf"  %-10s %s\n",$$1,$$2}' $(MAKEFILE_LIST)

.PHONY: init
init:
	npm install
	npm run compile
	cp .env.example .env

.PHONY: deploy
deploy: deploy-system deploy-registries

.PHONY: deploy-system
deploy-system:
	VERIFY=true npx hardhat --network mumbai run scripts/deploy-mainchain.ts
	VERIFY=true npx hardhat --network zkevm-testnet run scripts/deploy-sidechain.ts
	VERIFY=true npx hardhat --network goerli run scripts/deploy-sidechain.ts
	VERIFY=true npx hardhat --network bsc-testnet run scripts/deploy-sidechain.ts
	VERIFY=true npx hardhat --network gnosis-chiado run scripts/deploy-sidechain.ts
	VERIFY=true npx hardhat --network linea-testnet run scripts/deploy-sidechain.ts

.PHONY: deploy-registries
deploy-registries:
	VERIFY=true npx hardhat --network mumbai run scripts/deploy-registry.ts
	VERIFY=true npx hardhat --network zkevm-testnet run scripts/deploy-registry.ts
	VERIFY=true npx hardhat --network goerli run scripts/deploy-registry.ts
	VERIFY=true npx hardhat --network bsc-testnet run scripts/deploy-registry.ts
	VERIFY=true npx hardhat --network gnosis-chiado run scripts/deploy-registry.ts
	VERIFY=true npx hardhat --network linea-testnet run scripts/deploy-registry.ts

.PHONY: deploy-mainnet
deploy-prod: deploy-mainnet-system deploy-mainnet-registries

.PHONY: deploy-mainnet-system
deploy-mainnet-system:
	VERIFY=true npx hardhat --network polygon run scripts/deploy-mainchain.ts
	VERIFY=true npx hardhat --network zkevm run scripts/deploy-sidechain.ts
	VERIFY=true npx hardhat --network linea run scripts/deploy-sidechain.ts

.PHONY: deploy-mainnet-registries
deploy-mainnet-registries:
	VERIFY=true npx hardhat --network polygon run scripts/deploy-registry.ts
	VERIFY=true npx hardhat --network zkevm run scripts/deploy-registry.ts
	VERIFY=true npx hardhat --network linea run scripts/deploy-registry.ts
