##
##- Local Go:
##

run: ## run API with Go
	@go run cmd/book-api/main.go -level debug

build: ## build API with Go
	@go build -o book-api cmd/book-api/main.go

.PHONY: run build

##
##- Local Docker:
##

up: ## Up local docker stack
	@docker compose up --build --force-recreate -d

swagger: ## Start swagger ui locally
	@docker compose up -d --no-deps scalar

log: ## Logs local docker
	@docker compose logs -f


.PHONY: up log swagger

##
##- Tooling:
##

lint: ## Run golangci-lint
	@golangci-lint run ./...

pre-commit: ## Run pre-commit on all files
	@pre-commit run --all-files

setup-hooks: ## Install pre-commit hooks
	@pre-commit install

.PHONY: lint pre-commit setup-hooks

##
##- Test:
##

test: ## Run test
test:
	@go test -race ./...

test-cover: ## Run test with coverage
test-cover:
	@sh -c "go test -race ./... -coverprofile=/tmp/coverage.out && go tool cover -html=/tmp/coverage.out"

.PHONY: test test-cover

##

.DEFAULT_GOAL := help
help:
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

.PHONY: help
