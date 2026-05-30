import assert from 'node:assert';
import test from 'node:test';
import { Receipt } from '../src/receipt.js';

test('calculates total for regular customer', () => {
  const receipt = new Receipt();
  assert.strictEqual(receipt.calculate_total([10, 20, 30]), 59.4);
});

test('calculates total for member customer', () => {
  const receipt = new Receipt();
  assert.strictEqual(receipt.calculate_total([10, 20, 30], 'member'), 56.16);
});

test('calculates total for vip customer', () => {
  const receipt = new Receipt();
  assert.strictEqual(receipt.calculate_total([10, 20, 30], 'vip'), 47.68);
});

test('applies bonus discount over 50', () => {
  const receipt = new Receipt();
  assert.strictEqual(receipt.calculate_total([60]), 59.4);
});

test('applies vip extra discount', () => {
  const receipt = new Receipt();
  assert.strictEqual(receipt.calculate_total([100], 'vip'), 84.4);
});

test('returns 0 for empty items', () => {
  const receipt = new Receipt();
  assert.strictEqual(receipt.calculate_total([]), 0);
});

test('exactly 50 gets no bonus discount', () => {
  const receipt = new Receipt();
  assert.strictEqual(receipt.calculate_total([50]), 54);
});
