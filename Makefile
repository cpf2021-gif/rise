.PHONY: lint
lint:
	@golangci-lint run -E gofumpt --fix

.PHONY: mod
mod:
	@go mod tidy