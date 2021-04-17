.DEFAULT_GOAL:help
MAKEFLAGS=--no-print-directory

ENV := development
LINTER_PATH := $(shell go env GOPATH)/bin/golint
PWD := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
BUILD_DIR = $(PWD)bin
SOURCE_DIR := $(PWD)src
CONFIG_DIR := $(PWD)config
TARGET_NAME = subtitle-downloader

help: ## List all Makefile targets
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

init: ## Initialize the project
ifeq ("$(wildcard $(CONFIG_DIR)/config.$(ENV).json)","") # Config file doesn't exist - We have to create it
	cp $(CONFIG_DIR)/config.json.dist $(CONFIG_DIR)/config.$(ENV).json
else
	@echo "$(CONFIG_DIR)/config.$(ENV).json already created"
endif

build: init ## Build the project. Use BUILD_TARGET=`path to build target` (Default '../bin/subtitle-downloader')
	cd $(SOURCE_DIR) && \
	go build -o $(BUILD_DIR)/$(TARGET_NAME) . && \
	cd -

tests: ## Run all tests
	cd $(SOURCE_DIR) && \
	go test input --tags="all" -failfast -v -cover && \
	go test downloader --tags="all" -failfast -v -cover && \
	go test utils --tags="all" -failfast -v -cover && \
	go test command --tags="all" -failfast -v  -cover && \
	cd -

execute: ## Run executable (use OPT='...' to propagate options)
ifeq ("$(wildcard $(BUILD_DIR)/$(TARGET_NAME))","") # Executable file doesn't exist - We have to create it
	@$(MAKE) -s build # Create the executable file
endif
	$(BUILD_DIR)/$(TARGET_NAME) $(OPT)

vet : ## Run Vet
	cd $(SOURCE_DIR) && \
	go vet ./... && \
	cd -

lint: ## Run linter
	$(LINTER_PATH) $(SOURCE_DIR)/...