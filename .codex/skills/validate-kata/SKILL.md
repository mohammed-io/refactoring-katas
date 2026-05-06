---
name: validate-kata
description: Validate a refactoring kata by collecting the kata prompt, implementation, and tests into one evaluation-ready markdown prompt.
---

# Validate Refactoring Kata

## When To Use

Use this skill when user wants to evaluate a completed refactoring kata solution, inspect whether smells were removed, or generate a review prompt for another LLM.

## Invocation

- OpenCode or Claude-compatible hosts: `/validate-kata [language-or-path]`
- Codex CLI environments may expose equivalent command wiring such as `codex validate-kata [path]`

If current working directory already contains a kata with `VALIDATION_PROMPT.md`, auto-detect it. Otherwise, use provided path or language directory to locate target kata.

## Required Inputs

- Kata directory containing `VALIDATION_PROMPT.md`
- Refactored source code in `src/` for Ruby, Python, JavaScript, and C#, or package-root `*.go` files for Go
- Tests in `test/` for Ruby, Python, JavaScript, and C#, or `*_test.go` files for Go

## Steps

1. Resolve kata directory.
   - Prefer current directory if it contains `VALIDATION_PROMPT.md`.
   - Otherwise resolve supplied path or infer from language argument.
2. Read `VALIDATION_PROMPT.md`.
3. Read all user source files.
   - Non-Go: include files under `src/`.
   - Go: include package source files while excluding `*_test.go`.
4. Read all test files.
   - Non-Go: include files under `test/`.
   - Go: include `*_test.go`.
5. Assemble one markdown document with these sections:
   - `# Kata Validation: <kata-name> (<language>)`
   - `## Validation Criteria`
   - `## User's Refactored Code`
   - `## Test Suite`
   - `## Evaluation Request`
6. In evaluation request, ask reviewer to score:
   - Readability
   - Single Responsibility
   - Design Patterns
   - Test Preservation
   - Smells Eliminated
   - Overall
7. Output assembled prompt directly. If host supports optional API evaluation and credentials exist, that can be a follow-on step, not required for base validity.

## Output Requirements

- Preserve file contents faithfully
- Use fenced code blocks with correct language where practical
- Keep section order stable
- Make output copy-paste ready for another LLM review

## Error Handling

- If `VALIDATION_PROMPT.md` is missing, report that caller is not in a kata directory
- If no source code is found, report that source code is missing
- If no tests are found, report that tests are missing
- If path argument does not resolve to a kata, stop and explain which path failed

## Final Checks

- Confirm prompt includes criteria, source, tests, and explicit scoring request
- Confirm Go handling excludes test files from source section
- Confirm command examples do not contradict actual host usage
