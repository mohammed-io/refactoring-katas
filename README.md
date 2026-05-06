# Refactoring Katas

A collection of **20 progressive refactoring katas** inspired by [Emily Bache's Tennis Refactoring Kata](https://github.com/emilybache/Tennis-Refactoring-Kata), covering all major code smells from [Martin Fowler's Refactoring Catalog](https://refactoring.com/catalog/).

## Kata List

| # | Kata | Core Smell | Key Refactorings |
|---|------|-----------|-----------------|
| 1 | **Magic Receipt** | Magic numbers/strings | Replace Magic Literal |
| 2 | **The Never-Ending Function** | Function doing 5+ things | Extract Function, Slide Statements |
| 3 | **Copy-Paste Payroll** | 3+ near-identical blocks | Extract Function, Parameterize Function |
| 4 | **Cryptic Calculator** | Single-letter names | Rename Variable/Function |
| 5 | **The Zombie Code** | Dead code, unreachable branches | Remove Dead Code |
| 6 | **Nested If Hell** | 4+ nested conditionals | Replace Nested Conditional with Guard Clauses |
| 7 | **The Switch Factory** | Giant switch on type | Replace Conditional with Polymorphism |
| 8 | **Jealous Function** | Obsessed with another class | Move Function, Hide Delegate |
| 9 | **Stringly Typed** | Primitives for domain | Replace Primitive with Object |
| 10 | **The Null Spiral** | Null checks everywhere | Introduce Special Case |
| 11 | **Parameter Soup** | 6+ parameters | Introduce Parameter Object |
| 12 | **God Object** | 15+ methods, 3+ concerns | Extract Class, Move Field |
| 13 | **The Chatterbox** | Exposed internals | Encapsulate Collection |
| 14 | **The Messenger** | Pure forwarding class | Remove Middle Man |
| 15 | **Inheritance Abuse** | Subclass breaks LSP | Push Down Field/Method, Remove Subclass |
| 16 | **Shotgun Surgery** | 1 change, 5+ files | Move Function/Field, Combine into Class |
| 17 | **The Taxonomy Trap** | Parallel hierarchies | Replace Inheritance with Delegation |
| 18 | **Big Bang Algorithm** | Opaque/complex algo | Substitute Algorithm, Split Loop |
| 19 | **Tangled Web** | Hardcoded dependencies | Extract Interface, Dependency Injection |
| 20 | **The Legacy Monolith** | **All smells combined** | Full pipeline |

## Languages

Each kata is available in:
- **Ruby** (4.0.0) — Minitest (stdlib)
- **Python** (3.12.0) — pytest
- **Go** (1.22.0) — stdlib `testing`
- **JavaScript** (Node 20.11.0) — built-in `node --test`
- **C#** (.NET 8.0) — xUnit

## Quick Start

### 1. Install languages

```bash
mise install   # Uses .tool-versions in each language directory
```

### 2. Pick a kata

```bash
cd ruby/01-magic-receipt
# or
cd python/07-switch-factory
# or
cd golang/12-god-object
```

### 3. Run tests (they should all pass on the bad code)

```bash
make test
```

### 4. Refactor the code in `src/`

- Read `README.md` for the scenario
- Read `HINTS.md` for smells to find and techniques to apply
- **Keep all tests green** — they test behavior, not implementation
- **Do not change tests or method signatures**

### 5. Validate your solution

```bash
make verify      # Prints the validation prompt to paste into an LLM
make doctor      # Check your tools are installed
```

Or use the built-in skill:
```bash
# Claude Code / OpenCode
/validate-kata

# Codex CLI
codex validate-kata
```

## Philosophy

- **No solutions included** — only starter code + tests + hints
- **Tests are perfect** — behavior-based, implementation-agnostic
- **Comments removed** from bad code — you must spot the smells yourself
- **Progressive difficulty** — start with simple naming issues, end with full legacy monoliths

## Project Structure

```
refactoring-katas/
├── README.md
├── ruby/
│   ├── .tool-versions
│   ├── 01-magic-receipt/
│   │   ├── Makefile
│   │   ├── README.md
│   │   ├── HINTS.md
│   │   ├── VALIDATION_PROMPT.md
│   │   ├── src/
│   │   │   └── receipt.rb
│   │   └── test/
│   │       └── receipt_test.rb
│   ├── 02-never-ending-function/
│   └── ...
├── python/
├── golang/
├── javascript/
├── csharp/
├── .opencode/skills/validate-kata/
├── .claude/skills/validate-kata/
└── .codex/skills/validate-kata/
```

## Kata Makefile Commands

Every kata directory has a `Makefile` with:

| Command | Purpose |
|---------|---------|
| `make install` | Install dependencies (if any) |
| `make test` | Run the test suite |
| `make doctor` | Check for missing tools (ruby, mise, etc.) |
| `make verify` | Print the LLM validation prompt |

## Validation

Each kata includes `VALIDATION_PROMPT.md` — a prompt you can paste into any LLM (Claude, ChatGPT, Gemini, Codex) to evaluate your refactored solution. It scores across:

- **Readability** (naming, formatting, clarity)
- **Single Responsibility** (function/class size and cohesion)
- **Design Patterns** (Fowler techniques applied correctly)
- **Test Preservation** (all original tests still pass)
- **Smells Eliminated** (checklist from HINTS.md)
- **Overall Score** (0-10)

## License

MIT — inspired by Emily Bache's katas and Martin Fowler's Refactoring.
