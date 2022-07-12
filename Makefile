MODULE = $(shell go list -m)

.PHONY: build
build: ## build the API server binary
	go build -o server $(MODULE)/cmd/server
