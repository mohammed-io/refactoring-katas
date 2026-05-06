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
	@echo "Ruby stdlib includes minitest. No gems to install."

test:
	ruby -I src -I test test/*_test.rb

doctor:
	@command -v mise >/dev/null 2>&1 || { echo "mise is not installed"; exit 1; }
	@command -v ruby >/dev/null 2>&1 || { echo "ruby is not installed"; exit 1; }
	@ruby -e "require 'minitest/autorun'"

verify:
	@echo "Paste the contents of VALIDATION_PROMPT.md into your LLM"
