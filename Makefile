help:
	@echo "Available targets:"
	@echo "  help          - Show this help message"
	@echo "  install-tools - Install tree-sitter analysis tools"
	@echo "  kata-structure LANG=xxx KATA=yyy - Analyze a specific kata's structure"
	@echo ""
	@echo "Examples:"
	@echo "  make kata-structure LANG=python KATA=07"
	@echo "  make kata-structure LANG=ruby KATA=01"

.PHONY: help install-tools kata-structure

LANG ?=
KATA ?=

install-tools:
	mise exec -- python3 -m pip install -r requirements-tools.txt
	mise exec -- python3 -c "from tree_sitter_language_pack import download; download(['javascript', 'python', 'ruby', 'go', 'csharp'])"

kata-structure:
	@test -n "$(LANG)" || { echo "LANG is required, e.g. make kata-structure LANG=python KATA=07"; exit 1; }
	@test -n "$(KATA)" || { echo "KATA is required, e.g. make kata-structure LANG=python KATA=07"; exit 1; }
	mise exec -- python3 tools/kata_structure.py --language "$(LANG)" --kata "$(KATA)"
