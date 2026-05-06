# audit_kata_validity

## Status: completed (20260506080801)

## Context
The repository contains 20 refactoring katas across Ruby, Python, JavaScript, Go, and C#.
The user wants all katas checked and repaired so they are useful refactoring exercises rather than polished solutions.
README files should briefly explain component usage/domain context, while HINTS files should give only small hints and never full solutions.
Every kata Makefile should support dependency installation and test execution.
The user clarified that `mise` is the preferred installation tool.
The user also clarified that the katas should be inspired by Emily Bache's Tennis Refactoring Kata: the same kata behavior should exist across languages, with intentionally awkward but working code and fast tests.

## Value Proposition
Ensure every kata is worth solving, has clear learner-facing context, preserves solution-free hints, and can be installed/tested consistently.

## Alternatives considered (with trade-offs)
1. Manually inspect and edit every kata file one by one.
Trade-off: highest precision, too slow and high risk of inconsistency across 100 directories.
2. Use scripts only to audit structure, then hand-fix problem patterns.
Trade-off: fast enough, preserves judgment for actual content quality.
3. Replace all README/HINTS/Makefiles from templates.
Trade-off: consistent, but may erase useful kata-specific detail and user edits.

Chosen approach: audit mechanically for completeness and common invalid patterns, sample content across languages, then patch the smallest set of repo-wide problems.

## Todos
- [x] Inventory all kata directories and required files
- [x] Audit README and HINTS content for role mismatch or solution leakage
- [x] Audit Makefiles for install and test targets
- [x] Inspect source/test quality to confirm exercises need refactoring and tests are runnable
- [x] Patch invalid or weak kata files
- [x] Run verification checks and archive work

## Acceptance Criteria
- All 100 expected language/kata directories are present
- Each kata has README, HINTS, VALIDATION_PROMPT, Makefile, source files, and tests
- README content explains business/component usage instead of giving refactoring steps
- HINTS content gives small smell-oriented hints without proper solutions
- Makefiles expose valid `install` and `test` targets
- Starter implementations remain intentionally refactorable while tests pass
- Each kata number has the same domain, behavior, public API, and test intent across all five languages
- Kata shape matches Tennis Kata thinking: small behavior surface, comprehensive tests, awkward design, no prescribed refactoring solution

## Notes
Need protect existing user changes: current worktree is fully untracked, so edits must stay targeted.

Audit results so far:
- All 100 expected language/kata directories exist.
- No required README, HINTS, VALIDATION_PROMPT, or Makefile files are missing or empty.
- No source/test directories are missing files.
- Every Makefile exposes at least `install` and `test`.
- README files contain refactoring instructions in many directories; this violates the requested README role.
- HINTS files do not contain copy-paste full solutions, but many include explicit Fowler techniques and step-by-step suggestions that are stronger than "small hints."

Tennis Kata comparison:
- Upstream Tennis Kata has one domain/behavior across languages and multiple intentionally smelly implementations.
- Tests are the safety net and learners should not need to edit them during refactoring.
- README describes scenario and exercise context, not a specific design solution.
- Current repo fails parity: several kata numbers use different domains across languages, for example kata 15 is shapes in Ruby/Python, employee bonuses in JavaScript, rental plans in Go, and zoo animals in C#.

Implementation direction:
- Use the JavaScript sequence as the canonical 20-kata behavior set.
- Port each kata number to Ruby, Python, Go, and C# with the same domain and test intent.
- Keep implementations intentionally awkward and behavior-preserving, matching Tennis Kata style.
- Make all Makefile `install` targets call `mise install`.

Verification:
- `make test` passed for all 20 JavaScript katas.
- `make test` passed for all 20 Python katas.
- `make test` passed for all 20 Ruby katas.
- `make test` passed for all 20 Go katas with local `GOCACHE=$(CURDIR)/.gocache`.
- `make test` passed for all 20 C# katas using `mise exec -- dotnet test`; this required escalated execution because the .NET test runner opens a local socket.
- Final structure check found 100 kata directories and no missing README/HINTS/VALIDATION_PROMPT/Makefile files.
- README/HINTS scans found no remaining solution-oriented README sections or explicit solution sections in HINTS.
