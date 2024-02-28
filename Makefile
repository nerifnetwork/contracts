.PHONY: help
help: # Display this help
	@awk 'BEGIN{FS=":.*#";printf "Usage:\n  make <target>\n\nTargets:\n"}/^[a-zA-Z_-]+:.*?#/{printf"  %-10s %s\n",$$1,$$2}' $(MAKEFILE_LIST)

.PHONY: init
init:
	npm install
	npm run compile
	cp .env.example .env

.PHONY: deploy-testnets
deploy-testnets:
	DEPLOY_CONFIG_PATH='./data/testnet/config.mumbai.json' npm run deploy mumbai
	DEPLOY_CONFIG_PATH='./data/testnet/config.polygonZkEVMTestnet.json' npm run deploy polygonZkEVMTestnet
	DEPLOY_CONFIG_PATH='./data/testnet/config.goerli.json' npm run deploy goerli
	DEPLOY_CONFIG_PATH='./data/testnet/config.bscTestnet.json' npm run deploy bscTestnet
	DEPLOY_CONFIG_PATH='./data/testnet/config.chiado.json' npm run deploy chiado
	DEPLOY_CONFIG_PATH='./data/testnet/config.lineaTestnet.json' npm run deploy lineaTestnet

.PHONY: deploy-mainnets
deploy-mainnets:
	DEPLOY_CONFIG_PATH='./data/mainnet/config.polygon.json' npm run deploy polygon
	DEPLOY_CONFIG_PATH='./data/mainnet/config.polygonZkEVM.json' npm run deploy polygonZkEVM
	DEPLOY_CONFIG_PATH='./data/mainnet/config.linea.json' npm run deploy linea
