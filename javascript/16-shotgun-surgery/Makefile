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
	@echo "Node built-in test runner available. No packages to install."

test:
	node --test test/*.js

doctor:
	@command -v mise >/dev/null 2>&1 || { echo "mise is not installed"; exit 1; }
	@command -v node >/dev/null 2>&1 || { echo "node is not installed"; exit 1; }

verify:
	@echo "Paste the contents of VALIDATION_PROMPT.md into your LLM"
