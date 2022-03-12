BUILD_ID := $(shell git rev-parse --short HEAD 2>/dev/null || echo no-commit-id)

.DEFAULT_GOAL := help
help: ## List targets & descriptions
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Clean the build
	rm -rf ./build

get: ## Run go get
	go get .

build-service: ## Build the Go service
	mkdir -p ./build
	go build -v -o ./build/slack-slash-command

run: ## Build and run the project locally
	mkdir -p ./build
	go build -o ./build/slack-slash-command && ./build/slack-slash-command
