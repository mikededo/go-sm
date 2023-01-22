.PHONY: all
all: help
help: ## Display help screen
	@echo "Usage:"
	@echo "	make [COMMAND]"
	@echo "	make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint: ## Run the linter using golangci-lint
	golangci-lint run ./... --out-format tab

.PHONY: test
test: ## Run all tests and print the output to the console
	gotestsum --format pkgname --hide-summary output -- ./...

.PHONY: test-raw
test-raw: ## Same as test but without the usage of gotestsum
	go test -v ./...

.PHONY: test-domain
test-domain: ## Runs all the tests in the domain folder
	gotestsum --format testname --hide-summary output -- -v ./domain/...

.PHONY: test-domain-raw
test-domain-raw: ## Same as test-domain but without the usage of gotestsum
	go test -v ./domain/...

.PHONY: test-app
test-app: ## Runs all the tests in the application folder
	gotestsum --format testname --hide-summary output -- -v ./application/...

.PHONY: test-app-raw
test-app-raw: ## Same as test-app but without the usage of gotestsum
	go test -v ./application/...

.PHONY: test-infra
test-infra: ## Runs all the tests in the infrastructure folder
	gotestsum --format testname --hide-summary output -- -v ./infrastructure/...

.PHONY: test-infra-raw
test-infra-raw: ## Same as test-infra but without the usage of gotestsum 
	go test -v ./infrastructure/...

.PHONY: test-all
test-all: ## Runs all the tests
	make test-domain
	make test-app
	make test-infra

.PHONY: coverage
coverage: ## Run all tests and generate a coverage report in HTML format
	go test -covermode=count -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: serve
serve: ## Runs the application in a local env
	clear
	go run main.go
