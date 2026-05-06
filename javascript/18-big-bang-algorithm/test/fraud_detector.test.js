import assert from 'node:assert';
import test from 'node:test';
import { FraudDetector } from '../src/fraud_detector.js';

test('low risk small transaction', () => {
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 10, timestamp: Date.now(), history: [], merchant: 'grocery', country: 'US', cardCountry: 'US' });
  assert.strictEqual(result.rating, 'low');
  assert.strictEqual(result.level, 1);
});

test('medium risk for large amount', () => {
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 1100, timestamp: Date.now(), history: [], merchant: 'grocery', country: 'US', cardCountry: 'US' });
  assert.strictEqual(result.rating, 'medium');
  assert.strictEqual(result.level, 2);
});

test('elevated risk for gambling', () => {
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 100, timestamp: Date.now(), history: [], merchant: 'gambling', country: 'US', cardCountry: 'US' });
  assert.strictEqual(result.rating, 'medium');
  assert.strictEqual(result.level, 2);
});

test('high risk for cross-border', () => {
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 1000, timestamp: Date.now(), history: [], merchant: 'grocery', country: 'FR', cardCountry: 'US' });
  assert.strictEqual(result.rating, 'low');
  assert.strictEqual(result.level, 1);
});

test('critical risk for late night crypto', () => {
  const t = new Date('2024-01-01T02:00:00Z').getTime();
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 600, timestamp: t, history: [], merchant: 'crypto', country: 'CN', cardCountry: 'US' });
  assert.strictEqual(result.rating, 'elevated');
  assert.strictEqual(result.level, 3);
});

test('velocity increases risk', () => {
  const now = Date.now();
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 50, timestamp: now, history: [
    { amount: 10, timestamp: now - 1000 },
    { amount: 10, timestamp: now - 2000 },
    { amount: 10, timestamp: now - 3000 },
    { amount: 10, timestamp: now - 4000 }
  ], merchant: 'grocery', country: 'US', cardCountry: 'US' });
  assert.strictEqual(result.level, 2);
});

test('volume spikes increase risk', () => {
  const now = Date.now();
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 50, timestamp: now, history: [
    { amount: 200, timestamp: now - 10000 },
    { amount: 200, timestamp: now - 20000 },
    { amount: 200, timestamp: now - 30000 }
  ], merchant: 'grocery', country: 'US', cardCountry: 'US' });
  assert.strictEqual(result.level, 1);
});

test('includes score', () => {
  const detector = new FraudDetector();
  const result = detector.detect({ amount: 10, timestamp: Date.now(), history: [], merchant: 'grocery', country: 'US', cardCountry: 'US' });
  assert.strictEqual(typeof result.score, 'number');
});
