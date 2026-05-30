# improve_weak_ok_katas

## Status: completed (20260528115557)

## Context
Previous evaluation found weak/OK refactoring practice value in katas 04, 05, 13, 14, 19, and 20. User asked to improve only those katas, across Ruby, Python, JavaScript, Go, and C#, using repository structure tooling for parity. No unrelated kata changes.

## Value Proposition
Make targeted katas better practice for behavior-preserving refactoring: enough design pressure, tests focused on observable behavior, consistent cross-language scenarios, and clearer safety nets.

## Alternatives considered (with trade-offs)
1. Patch only docs and hints: fastest, but leaves weak exercise design intact.
2. Rewrite all weak/OK katas from scratch: high quality potential, but larger risk and broad churn across 30 kata-language pairs.
3. Targeted scenario/test/source upgrades for weak/OK katas: best current fit; improves practice value while preserving repo structure and limiting blast radius.
4. Add new katas instead: avoids disturbing existing exercises, but does not fix identified weak spots.

## Todos
- [x] Inspect target kata structure across languages with tree-sitter command.
- [x] Improve base weak/OK katas 04, 05, 13, 14, 19, 20 across all languages.
- [x] Select 3-4 target katas for excellent-level pass and document what improves them.
- [x] Run relevant tests/structure checks for changed katas.
- [x] Review diff for accidental unrelated changes and edge cases.

## Acceptance Criteria
- Only weak/OK katas are changed.
- All five language implementations stay behaviorally aligned per kata.
- Tests exercise observable behavior and avoid locking unnecessary internals where possible.
- Relevant tests pass or any inability is reported.

## Notes
- Git status was clean before starting.
- Focused Ruby/Python/JavaScript/Go tests passed for katas 04, 05, 13, 14, 19, 20.
- Focused C# tests passed for katas 04, 05, 13, 14, 19, 20 outside sandbox due dotnet test socket binding.
- Structure checks passed with zero parse errors and aligned test counts across languages.
- Intentional smell rule: starter implementations must stay smelly; improvements should strengthen observable tests, parity, scenario pressure, and docs/hints.
- Excellent candidates now: 13 Chatterbox, 14 Messenger, 19 Tangled Web, 20 Legacy Monolith.
- What improves them further:
  - 13: add invalid mutation examples to make encapsulation payoff clearer without requiring a clean starter.
  - 14: add retry/failure result cases so call-chain and responsibility boundaries matter more.
  - 19: add one precedence-conflict case per source (defaults/local/env/seasonal) and malformed env value behavior.
  - 20: add one fraud/payment edge and one shipment pricing edge to increase monolith pressure while keeping fast tests.
