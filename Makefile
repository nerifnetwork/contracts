.PHONY: help
help: # Display this help
	@awk 'BEGIN{FS=":.*#";printf "Usage:\n  make <target>\n\nTargets:\n"}/^[a-zA-Z_-]+:.*?#/{printf"  %-10s %s\n",$$1,$$2}' $(MAKEFILE_LIST)

.PHONY: abigen
abigen: # Generate go files
	rm -rf artifacts/
	npm run compile
	npm run extract-abi
	docker build -f Dockerfile.abigen -t extract-abi .
	rm -rf pkg/*
	CONTAINER=`docker create extract-abi --name extract-abi`; \
	docker cp $$CONTAINER:/system pkg/system; \
	docker cp $$CONTAINER:/interfaces pkg/interfaces; \
	docker rm -v $$CONTAINER

.PHONY: deploy
deploy: deploy-system deploy-registries deploy-gateways

.PHONY: deploy-system
deploy-system:
	npx hardhat --network mumbai run scripts/deploy-mainchain.ts
	npx hardhat --network goerli run scripts/deploy-sidechain.ts
	npx hardhat --network bsc-testnet run scripts/deploy-sidechain.ts

.PHONY: deploy-registries
deploy-registries:
	npx hardhat --network mumbai run scripts/deploy-registry.ts
	npx hardhat --network goerli run scripts/deploy-registry.ts
	npx hardhat --network bsc-testnet run scripts/deploy-registry.ts

.PHONY: deploy-gateways
deploy-gateways:
	npx hardhat --network mumbai run scripts/deploy-gateway.ts
	npx hardhat --network goerli run scripts/deploy-gateway.ts
	npx hardhat --network bsc-testnet run scripts/deploy-gateway.ts