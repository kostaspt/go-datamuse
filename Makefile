PKG_LIST = $(shell go list ./...  | grep -v /vendor/)

vet: ## Vet the files
	@go vet ${PKG_LIST}

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: dep ## Run tests
	@go test -short ${PKG_LIST}

test-race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

test-coverage: dep ## Run tests with coverage
	@go test -v -cover -race -coverprofile=coverage.out ${PKG_LIST}

dep: ## Get the dependencies
	@dep ensure
	@export GO111MODULE=on
	@go mod tidy

help: ## Display help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'