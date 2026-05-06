# improve_root_readme

## Status: completed (20260506081628)

## Context
The root README should explain how to use the refactoring-katas repo after the Tennis-Kata-style cleanup.
It currently mentions `mise install`, `make test`, and kata structure, but it is ambiguous about where to run setup and still frames hints as techniques.

## Value Proposition
Make first-run usage clear: install `mise`, enter a kata directory, run `make install`, run `make test`, refactor behavior-preservingly, and use hints only as small nudges.

## Alternatives considered (with trade-offs)
1. Leave README as-is.
Trade-off: fewer edits, but users may run `mise install` from the wrong directory and miss per-kata `make install`.
2. Patch only Quick Start.
Trade-off: fixes immediate setup confusion, but leaves Philosophy and command table slightly inconsistent.
3. Rewrite the README around the current repo shape.
Trade-off: more edit surface, but gives accurate requirements and expectations.

Chosen approach: targeted README rewrite for requirements, quick start, philosophy, and command table.

## Todos
- [x] Review current root README for setup gaps
- [x] Patch README requirements and workflow wording
- [x] Verify README no longer contradicts kata/hint requirements
- [x] Archive work

## Acceptance Criteria
- README states `mise` is required
- README tells users to run `make install` inside the kata directory
- README explains `make test`, `make doctor`, and `make verify`
- README describes HINTS as small hints, not solutions or technique recipes
- README reflects the Tennis-Kata-inspired workflow accurately

## Notes
User asked whether README is good and whether it tells requirements such as `mise` and `make`.

Findings:
- Previous README was close, but it implied root-level `mise install` instead of the per-kata `make install` flow.
- Previous README called HINTS "techniques," which conflicted with the small-hints-only rule.
- Previous kata table exposed exact refactoring names, so it was too solution-oriented for a Tennis-Kata-style exercise.

Changes:
- Added explicit requirements: `mise`, `make`, and shell/toolchain access.
- Reworked Quick Start around `cd <kata>`, `make install`, `make test`, refactor, `make verify`, and `make doctor`.
- Changed "Key Refactorings" to neutral "Practice Focus."
