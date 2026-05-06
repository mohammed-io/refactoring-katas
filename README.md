# Refactoring Katas

A collection of **20 progressive refactoring katas** inspired by [Emily Bache's Tennis Refactoring Kata](https://github.com/emilybache/Tennis-Refactoring-Kata).

Each kata is a small, working legacy-code exercise. The tests pass at the start. Your job is to improve the internal design while preserving behavior, using the tests as a safety net.

## Kata List

| # | Kata | Core Smell | Practice Focus |
|---|------|-----------|----------------|
| 1 | **Magic Receipt** | Magic numbers/strings | Naming hidden business rules |
| 2 | **The Never-Ending Function** | Function doing 5+ things | Finding boundaries inside one workflow |
| 3 | **Copy-Paste Payroll** | 3+ near-identical blocks | Reducing repetition safely |
| 4 | **Cryptic Calculator** | Single-letter names | Making intent visible |
| 5 | **The Zombie Code** | Dead code, unreachable branches | Removing noise without changing behavior |
| 6 | **Nested If Hell** | 4+ nested conditionals | Flattening decision flow |
| 7 | **The Switch Factory** | Giant switch on type | Organizing mode-specific rules |
| 8 | **Jealous Function** | Obsessed with another class | Moving behavior toward the data it uses |
| 9 | **Stringly Typed** | Primitives for domain | Making domain values explicit |
| 10 | **The Null Spiral** | Null checks everywhere | Clarifying missing-data behavior |
| 11 | **Parameter Soup** | 6+ parameters | Grouping related inputs |
| 12 | **God Object** | 15+ methods, 3+ concerns | Separating responsibilities |
| 13 | **The Chatterbox** | Exposed internals | Protecting object boundaries |
| 14 | **The Messenger** | Pure forwarding class | Removing low-value indirection |
| 15 | **Inheritance Abuse** | Subclass tree for small variations | Questioning hierarchy shape |
| 16 | **Shotgun Surgery** | 1 change, 5+ files | Finding scattered rules |
| 17 | **The Taxonomy Trap** | Parallel hierarchies | Avoiding class explosion |
| 18 | **Big Bang Algorithm** | Opaque/complex algorithm | Separating calculation phases |
| 19 | **Tangled Web** | Hardcoded dependencies | Untangling environment dependencies |
| 20 | **The Legacy Monolith** | **All smells combined** | Refactoring in small safe steps |

## Languages

Each kata is available in:
- **Ruby** (4.0.0) — Minitest (stdlib)
- **Python** (3.12.0) — pytest
- **Go** (1.22.0) — stdlib `testing`
- **JavaScript** (Node 20.11.0) — built-in `node --test`
- **C#** (.NET 8.0) — xUnit

Tool versions are managed with [`mise`](https://mise.jdx.dev/). Each language directory contains its own `.tool-versions` file.

## Requirements

- `mise`
- `make`
- A shell that can run the language toolchain commands

You do not need to install every language before starting. Install only the language for the kata you want to practice.

## Quick Start

### 1. Pick a kata

```bash
cd ruby/01-magic-receipt
```

Or choose the same kata in another language:

```bash
cd python/07-switch-factory
cd golang/12-god-object
```

### 2. Install tools and dependencies

```bash
make install
```

`make install` runs `mise install` from the kata directory and then installs any language-specific dependencies, such as `pytest` for Python or NuGet packages for C#.

### 3. Run tests

```bash
make test
```

The tests should pass before you change anything.

### 4. Refactor the code

- Read `README.md` for the scenario
- Read `HINTS.md` only if you want small nudges about what to look for
- **Keep all tests green** — they test behavior, not implementation
- **Do not change tests or public behavior**

### 5. Validate your solution

```bash
make verify
```

`make verify` points you to `VALIDATION_PROMPT.md`, which you can paste into an LLM for a refactoring review.

You can also check the local setup at any point:

```bash
make doctor
```

## Philosophy

- **No solutions included** — only starter code, tests, and small hints
- **Tests are the contract** — behavior-based and fast to run
- **Same kata across languages** — each kata number has the same domain and behavior in Ruby, Python, JavaScript, Go, and C#
- **Inspired by Tennis Kata** — small working programs with awkward design, meant for behavior-preserving refactoring
- **Progressive difficulty** — start with simple naming issues, end with a legacy monolith

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
| `make install` | Run `mise install` and install kata dependencies |
| `make test` | Run the test suite |
| `make doctor` | Check for missing tools |
| `make verify` | Point to the validation prompt |

## Kata Structure Reports

Use the root Tree-sitter structure tool when you want an AI agent to compare the same kata across languages.

Install the reporting tool once from the repo root:

```bash
make install-tools
```

Print a Markdown report for one language and kata:

```bash
make kata-structure LANG=python KATA=07
make kata-structure LANG=javascript KATA=07
make kata-structure LANG=ruby KATA=07
make kata-structure LANG=golang KATA=07
make kata-structure LANG=csharp KATA=07
```

The report separates source and tests, then lists declarations, control-flow shape, literals, identifiers, test names, test count, estimated runtime cases, assertions, parse errors, and a compact Tree-sitter node profile. The tool reports facts only; use the Markdown outputs to judge parity across languages.

## Validation

Each kata includes `VALIDATION_PROMPT.md` — a prompt you can paste into any LLM to evaluate your refactored solution. It scores across:

- **Readability** (naming, formatting, clarity)
- **Single Responsibility** (function/class size and cohesion)
- **Behavior Preservation** (same observable behavior)
- **Test Preservation** (all original tests still pass)
- **Smells Eliminated** (checklist from HINTS.md)
- **Overall Score** (0-10)

## License

MIT — inspired by Emily Bache's katas and Martin Fowler's Refactoring.
