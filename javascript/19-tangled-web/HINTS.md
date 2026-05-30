## Code Smells to Look For

- [ ] Filesystem path, environment variables, current date, and defaults are hardcoded
- [ ] Error handling and config merging are tangled
- [ ] Config precedence is implicit inside one method

## Common Pitfalls

- Do not change tests while refactoring
- Do not rewrite from scratch before understanding behavior
- Run the test suite after small changes
