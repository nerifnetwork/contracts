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
	docker cp $$CONTAINER:/bridge pkg/bridge; \
	docker cp $$CONTAINER:/system pkg/system; \
	docker cp $$CONTAINER:/interfaces pkg/interfaces; \
	docker rm -v $$CONTAINER
