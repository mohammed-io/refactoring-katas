# fix_kata_format_and_parity

## Status: complete

## Context
The user provided an audit table questioning whether current katas still have domain drift and uneven formatting/coverage.
Current verification shows README scenarios are aligned, but old file/class names still imply different domains.
Generated C# files are minified into single lines, Ruby has semicolon-compressed branches, Go has very small source/tests, and test coverage volume varies heavily against JavaScript.

## Value Proposition
Make kata parity credible beyond README text: names, formatting, and tests should look like deliberate, maintainable kata starting points in every language.

## Alternatives considered (with trade-offs)
1. Only answer the audit.
Trade-off: fast, but leaves valid quality issues.
2. Rename/reformat only the examples mentioned.
Trade-off: fixes visible symptoms but not repo-wide consistency.
3. Regenerate style/parity across languages with canonical names and expanded tests.
Trade-off: larger change, but best matches the user's stated bar.

Chosen approach: fix repo-wide names/formatting and raise non-JS tests to a reasonable baseline while keeping the existing canonical kata behavior.

## Todos
- [x] Verify which audit items are still true
- [x] Normalize old-domain source/test filenames where practical
- [x] Reformat C# and Ruby away from minified/compressed style
- [x] Expand Go/Ruby/Python/C# tests for parity-critical cases
- [x] Run all kata test suites
- [x] Archive work

## Acceptance Criteria
- Kata README scenarios remain aligned across all languages
- Old-domain filenames/classes no longer make kata numbers look like different domains
- C# files are multi-line readable source files
- Ruby source avoids semicolon-compressed multi-statement lines
- Non-JS tests cover core canonical behaviors rather than one tiny smoke test
- All 100 kata test suites pass

## Notes
Audit table is partly correct:
- README business scenarios are already aligned.
- Old source/test filenames remain misleading for several languages.
- Formatting and coverage quality complaints are valid.

Final verification:
- Old-domain source/test name search is clean.
- No test file remains below the minimum coverage-size baseline.
- JavaScript, Python, Ruby, Go, and C# all pass 20/20 kata test suites.
- Python Makefiles now run Python through `mise exec` so `make install` does not use system Python.
