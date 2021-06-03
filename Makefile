.DEFAULT_GOAL := build/cbp

BUILD_ID := $(shell git rev-parse --short HEAD 2>/dev/null || echo no-commit-id)
IMAGE_NAME := ipenguin/cbp

#######################################
# Helper targets:
#######################################

# Parse the Makefile and show any targets with comments
.PHONY: help
help: ## Show available targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Cleanup build files
	rm -rf build/*

#######################################
# Build process targets:
#######################################

.PHONY: fmt
fmt: ## Format Go code
	go fmt .

.PHONY: lint
lint: fmt ## Run linting on code
	golint .

.PHONY: vet
vet: lint ## vet code for non-obvious errors
	go vet .

.PHONY: test
test: vet ## Run the unit tests
	go test .

.PHONY: build/cbp
build/cbp: test ## Build the docker image
	go build -o build/cbp .

.PHONY: container
container: build/cbp ## Create the docker images
    docker build -t $(IMAGE_NAME):$(BUILD_ID) .
	docker tag $(IMAGE_NAME):$(BUILD_ID) $(IMAGE_NAME):latest

.PHONY: push
push: container ## Push the docker image to docker hub with tags 'latest' & 'BUILD_ID'
	docker push $(IMAGE_NAME):$(BUILD_ID)
	docker push $(IMAGE_NAME):latest
