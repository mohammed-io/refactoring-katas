# add_kata_structure_cli

## Status: completed (20260506094009)

## Context
The repo needs a root-level CLI that uses Tree-sitter to summarize each kata's source and tests in a simple Markdown format. AI agents should compare these reports across languages without reading every source file directly.

## Value Proposition
The CLI gives a normalized, language-neutral view of public structure, control flow, literals, test names, and assertion shape. It helps detect kata drift while keeping parity judgment outside the tool.

## Alternatives considered (with trade-offs)
1. Full parity checker.
Trade-off: more automation, but too opinionated and likely brittle across languages.
2. Regex-only summarizer.
Trade-off: simpler install, but does not satisfy Tree-sitter requirement and misses structure.
3. Markdown-only Tree-sitter structure reporter.
Trade-off: lighter than strict parity, but best for AI-assisted comparison and simple maintenance.

Chosen approach: single Python CLI with Tree-sitter parsing and deterministic Markdown output.

## Todos
- [x] Inspect current root tooling and kata file layout
- [x] Add Tree-sitter tool dependency and root Makefile targets
- [x] Implement Markdown-only kata structure CLI
- [x] Verify output on kata 07 across all languages
- [x] Review edge cases and complete work

## Acceptance Criteria
- CLI accepts `--language` and `--kata`
- CLI prints Markdown only by default
- Source and test sections are separated and self-contained
- Test names and counts are visible
- Output includes enough structure for AI comparison across languages
- Root Makefile can install tool dependencies and run the CLI

## Notes
Keep logic report-oriented: tool states facts, AI decides parity.

Verification:
- `make install-tools` installs the Tree-sitter language pack and pre-downloads JavaScript, Python, Ruby, Go, and C# grammars.
- `python3 -m py_compile tools/kata_structure.py` passes.
- `make kata-structure LANG=<language> KATA=<01-20>` passes for all 100 language/kata pairs.
