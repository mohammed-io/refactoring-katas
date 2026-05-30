# Kata Quality Fixes

## Context
Comprehensive evaluation of all 20 refactoring katas revealed 6 categories of quality issues: C# behavioral bugs, test file naming mismatches, misleading test names, missing cap/sort tests, missing boundary tests, and unaudited C# katas.

## Value Proposition
Fixing these ensures katas are professional, consistent, and provide reliable safety nets for practitioners.

## Todos

### Fix 1: C# Kata 07 Behavioral Bugs
- [ ] Fix workout sort (add OrderByDescending energy)
- [ ] Fix focus sort (add OrderBy tempo) and cap (30 not 20)
- [ ] Fix party filter thresholds (>110 not >=120, >6 not >=6), sort, and cap (20 not 30)

### Fix 2: Rename Mismatched Test Files (all 5 languages)
- [ ] Kata 04: shipping_test → calculator/normalizer test
- [ ] Kata 05: auth_test → campaign_sender test
- [ ] Kata 06: premium_test → package_delivery test
- [ ] Kata 10: customer_test → room_lookup test
- [ ] Kata 14: invoice_test → notification_client test
- [ ] Kata 16: product_test → loyalty_rules test
- [ ] Kata 17: shelter_test → vehicle test

### Fix 3: Misleading Test Names in Kata 18 (all 5 languages)
- [ ] Rename test_high_risk_for_cross_border → reflects actual low risk
- [ ] Rename test_elevated_risk_for_gambling → reflects actual medium risk
- [ ] Rename test_volume_spikes_increase_risk → reflects no increase

### Fix 4: Missing Cap/Sort Tests for Kata 07 (all 5 languages)
- [ ] Add workout cap test (25)
- [ ] Add workout sort test (desc energy)
- [ ] Add focus cap test (30)
- [ ] Add focus sort test (asc tempo)
- [ ] Add party cap test (20)
- [ ] Add party sort test (desc danceability)

### Fix 5: Boundary Tests for Katas 01, 02, 06 (all 5 languages)
- [ ] Kata 01: Add test for exactly $50 (no bonus)
- [ ] Kata 02: Add test for quantity=100 (not out of stock)
- [ ] Kata 06: Add tests for weight=50, temperature boundaries

### Fix 7: Audit All 20 C# Katas
- [ ] Compare all 20 C# katas against Python for behavioral parity
- [ ] Fix any divergences found

## Acceptance Criteria
- All tests pass in all 5 languages
- C# kata 07 matches other languages' behavior
- No test file has a name unrelated to its domain
- No test name contradicts its assertion
- Kata 07 has cap/sort tests for all 5 moods
- Katas 01/02/06 have boundary tests
- All C# katas produce same outputs as Python equivalents
