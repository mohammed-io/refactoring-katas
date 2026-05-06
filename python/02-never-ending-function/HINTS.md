## Code Smells to Look For

- [ ] One function owns validation, math, side effects, and response building
- [ ] Temporary variables carry state across unrelated steps
- [ ] Hard-to-test behavior mixed with deterministic calculations

## Common Pitfalls

- Do not change tests while refactoring
- Do not rewrite from scratch before understanding behavior
- Run the test suite after small changes
