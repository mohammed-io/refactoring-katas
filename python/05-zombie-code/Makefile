PYTHON := mise exec -- python3

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
	$(PYTHON) -m pip install pytest --quiet

test:
	$(PYTHON) -m pytest test -v

doctor:
	@command -v mise >/dev/null 2>&1 || { echo "mise is not installed"; exit 1; }
	@command -v $(PYTHON) >/dev/null 2>&1 || { echo "python3 is not installed"; exit 1; }
	@$(PYTHON) -m pytest --version >/dev/null

verify:
	@echo "Paste the contents of VALIDATION_PROMPT.md into your LLM"
