help:
	@echo "Available targets:"
	@echo "  help    - Show this help message"
	@echo "  install - Install dependencies"
	@echo "  test    - Run tests"
	@echo "  doctor  - Check environment"
	@echo "  verify  - Validate with LLM (paste VALIDATION_PROMPT.md)"

.PHONY: help install test doctor verify

install:
	mise install
	go mod tidy

test:
	GOCACHE=$(CURDIR)/.gocache go test ./...

doctor:
	@command -v mise >/dev/null 2>&1 || { echo "mise is not installed"; exit 1; }
	@command -v go >/dev/null 2>&1 || { echo "go is not installed"; exit 1; }
	@go version

verify:
	@echo "Paste the contents of VALIDATION_PROMPT.md into your LLM"
