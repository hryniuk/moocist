.PHONY: default
default: build ut ## Default target - builds and runs UTs

.PHONY: generate
generate: ## Generate .go files with static directory, uses go-bindata
	go-bindata -pkg api -o pkg/api/data.go static/

.PHONY: build
build: generate ## Builds moocist binary
	go build

.PHONY: ut
ut: ## Runs all unit tests found in the project
	go test ./...

.PHONY: test
test: ## Runs tests on courses from test_slugs.txt file
	./test.sh

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
