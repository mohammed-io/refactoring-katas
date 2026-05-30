## Code Smells to Look For

- [ ] Internal persistence fields are exposed directly
- [ ] Setters allow invalid external mutation
- [ ] The object behaves like a data bucket
- [ ] Reservation behavior sits beside raw field access instead of a clear inventory boundary

## Common Pitfalls

- Do not change tests while refactoring
- Do not rewrite from scratch before understanding behavior
- Run the test suite after small changes
