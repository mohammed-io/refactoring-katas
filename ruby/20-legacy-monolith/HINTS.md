## Code Smells to Look For

- [ ] Many previous smells appear in one workflow
- [ ] Pricing, payment, shipping, email, logging, tax, and loyalty are coupled
- [ ] Small rule changes require reading the whole function
- [ ] Temporary names make it hard to see which rule changed the total

## Common Pitfalls

- Do not change tests while refactoring
- Do not rewrite from scratch before understanding behavior
- Run the test suite after small changes
