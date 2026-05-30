# Kata 19: Tangled Web

## Business Scenario

This kata is about a configuration loader. The implementation works and the tests describe the current behavior, but the code is intentionally awkward so it can be improved through refactoring.

## Component Usage

The component loads defaults, local config, environment overrides, and seasonal rules through hardcoded environment dependencies. Treat the tests as the contract while you improve the internal design.

## How to Run

```bash
make test
```
