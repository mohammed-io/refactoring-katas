import assert from 'node:assert';
import test from 'node:test';
import { Calculator } from '../src/calculator.js';

test('normalizes scores to 0-100 range', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([10, 20, 30], 0, 100);
  assert.deepStrictEqual(result, [0, 50, 100]);
});

test('normalizes scores to 1-5 range', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([10, 20, 30], 1, 5);
  assert.deepStrictEqual(result, [1, 3, 5]);
});

test('handles single value', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([50], 0, 100);
  assert.ok(Number.isNaN(result[0]));
});

test('handles negative input range', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([-10, 0, 10], 0, 1);
  assert.deepStrictEqual(result, [0, 0.5, 1]);
});

test('handles same min and max', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([5, 5, 5], 0, 100);
  assert.ok(Number.isNaN(result[0]));
  assert.ok(Number.isNaN(result[1]));
  assert.ok(Number.isNaN(result[2]));
});

test('handles reversed output range', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([10, 20, 30], 100, 0);
  assert.deepStrictEqual(result, [100, 50, 0]);
});

test('rounds fractional results', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([2, 3, 5], 0, 100);
  assert.deepStrictEqual(result, [0, 33.33, 100]);
});

test('empty input returns empty result', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([], 0, 100);
  assert.deepStrictEqual(result, []);
});

test('preserves input order with duplicates', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([30, 10, 30, 20], 0, 100);
  assert.deepStrictEqual(result, [100, 0, 100, 50]);
});

test('normalizes decimal output range', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([1.5, 2.5, 3.5], -1, 1);
  assert.deepStrictEqual(result, [-1, 0, 1]);
});

test('rounds negative fractional results', () => {
  const calculator = new Calculator();
  const result = calculator.normalize([2, 5, 8], -10, 10);
  assert.deepStrictEqual(result, [-10, 0, 10]);
});
