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
