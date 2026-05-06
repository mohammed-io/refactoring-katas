import assert from 'node:assert';
import test from 'node:test';
import { LoyaltyRules } from '../src/loyalty_rules.js';

test('bronze discount', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_discount_for_tier('bronze'), 0.05);
});

test('silver discount', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_discount_for_tier('silver'), 0.1);
});

test('gold discount', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_discount_for_tier('gold'), 0.15);
});

test('platinum discount', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_discount_for_tier('platinum'), 0.2);
});

test('unknown tier discount', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_discount_for_tier('unknown'), 0);
});

test('bronze label', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_label_for_tier('bronze'), 'Bronze Member');
});

test('silver label', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_label_for_tier('silver'), 'Silver Member');
});

test('gold label', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_label_for_tier('gold'), 'Gold Member');
});

test('platinum label', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_label_for_tier('platinum'), 'Platinum Member');
});

test('unknown tier label', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_label_for_tier('unknown'), 'Standard');
});

test('bronze threshold', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_threshold_for_tier('bronze'), 100);
});

test('silver threshold', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_threshold_for_tier('silver'), 500);
});

test('gold threshold', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_threshold_for_tier('gold'), 2000);
});

test('platinum threshold', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_threshold_for_tier('platinum'), 10000);
});

test('unknown tier threshold', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_threshold_for_tier('unknown'), 0);
});

test('bronze color', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_color_for_tier('bronze'), '#CD7F32');
});

test('silver color', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_color_for_tier('silver'), '#C0C0C0');
});

test('gold color', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_color_for_tier('gold'), '#FFD700');
});

test('platinum color', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_color_for_tier('platinum'), '#E5E4E2');
});

test('unknown tier color', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.get_color_for_tier('unknown'), '#000000');
});

test('calculates total for bronze', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.calculate_total(100, 'bronze'), 95);
});

test('calculates total for platinum', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.calculate_total(100, 'platinum'), 80);
});

test('calculates total for unknown tier', () => {
  const rules = new LoyaltyRules();
  assert.strictEqual(rules.calculate_total(100, 'unknown'), 100);
});
