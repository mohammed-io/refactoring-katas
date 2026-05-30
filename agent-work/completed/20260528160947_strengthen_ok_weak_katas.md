# strengthen_ok_weak_katas

## Status: completed (20260528162959)

## Context
User asked to strengthen OK/Weak refactoring katas only, without cleaning starter code. Current worktree already has many uncommitted changes from prior kata improvements, so changes must layer on top and avoid reverting user work.

## Value Proposition
Turn warmup-quality katas into stronger behavior-preserving refactoring practice while keeping starter implementations intentionally smelly.

## Alternatives considered (with trade-offs)
1. Only add tests: lowest risk, but weak smells may remain too shallow.
2. Rewrite scenarios fully: strongest result, but high churn across 5 languages.
3. Add targeted behavior pressure plus smell-preserving source changes: best fit; improves practice value while preserving awkward legacy feel.
4. Add new katas: avoids existing churn, but does not fix identified weak exercises.

## Todos
- [x] Inspect current OK/Weak kata implementations and tests across languages.
- [x] Strengthen kata 04 without cleaning code.
- [x] Strengthen kata 08 without cleaning code.
- [x] Strengthen kata 14 without cleaning code.
- [x] Strengthen kata 15 without cleaning code.
- [x] Strengthen kata 17 without cleaning code.
- [x] Run focused tests and structure checks, ignoring warnings.
- [x] Review diff for accidental cleanups or unrelated edits.

## Acceptance Criteria
- Only katas 04, 08, 14, 15, and 17 receive new changes in this pass.
- Starter code remains intentionally smelly.
- Tests focus on observable behavior.
- Cross-language behavior stays aligned.
- Build/test warnings are ignored, but actual failures are investigated.

## Notes
- Existing dirty files in 04/05/13/14/19/20 are user/prior work; do not revert.
- Focused Python, JavaScript, Ruby, Go, and C# tests pass for katas 04, 08, 14, 15, and 17.
- C# tests required escalated execution because sandbox blocks dotnet test socket binding.
- Structure checks for all changed kata/language pairs report parse errors: 0.
- `git diff --check` passed.
