## Code Smells to Look For

- [ ] Several layers only forward the same payload
- [ ] Middle objects add no behavior
- [ ] The call chain is longer than the work being done
- [ ] Validation, defaulting, auditing, and delivery boundaries are unclear

## Common Pitfalls

- Do not change tests while refactoring
- Do not rewrite from scratch before understanding behavior
- Run the test suite after small changes
