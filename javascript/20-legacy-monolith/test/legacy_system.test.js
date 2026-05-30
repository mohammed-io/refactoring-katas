import assert from 'node:assert';
import test from 'node:test';
import { LegacySystem } from '../src/legacy_system.js';

test('rejects empty order', () => {
  const system = new LegacySystem();
  const result = system.process_everything({ id: 1, items: [] });
  assert.strictEqual(result.error, 'No items');
});

test('calculates basic total', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 1,
    items: [{ price: 10, quantity: 2 }],
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.total, 21.4);
  assert.strictEqual(result.carrier, 'USPS');
});

test('applies member discount', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 2,
    items: [{ price: 100, quantity: 1 }],
    customer: { type: 'member', email: 'a@b.com' }
  });
  assert.strictEqual(result.total, 96.3);
});

test('applies vip discount', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 3,
    items: [{ price: 100, quantity: 1 }],
    customer: { type: 'vip', email: 'a@b.com' }
  });
  assert.strictEqual(result.total, 85.6);
});

test('applies bonus discount over 100', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 4,
    items: [{ price: 200, quantity: 1 }],
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.total, 203.3);
});

test('uses UPS for large total', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 5,
    items: [{ price: 60, quantity: 1 }],
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.carrier, 'UPS');
});

test('flags high total for review', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 6,
    items: [{ price: 5000, quantity: 1 }],
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.paymentStatus, 'manual_review');
});

test('includes email details', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 7,
    items: [{ price: 10, quantity: 1 }],
    customer: { email: 'user@test.com' }
  });
  assert.strictEqual(result.email.to, 'user@test.com');
  assert.ok(result.email.subject.includes('7'));
});

test('includes log entry', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 8,
    items: [{ price: 10, quantity: 1 }],
    customer: { email: 'a@b.com' }
  });
  assert.ok(result.log.includes('Order processed'));
});

test('includes order id', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 99,
    items: [{ price: 10, quantity: 1 }],
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.orderId, 99);
});

test('ignores non-positive item prices but counts positive quantities', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 10,
    items: [{ price: 20, quantity: 2 }, { price: -100, quantity: 1 }],
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.total, 42.8);
  assert.strictEqual(result.shippingWeight, 3);
});

test('applies SAVE10 coupon', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 11,
    items: [{ price: 100, quantity: 1 }],
    coupon: 'SAVE10',
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.total, 96.3);
});

test('tax exempt customer pays no tax', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 12,
    items: [{ price: 100, quantity: 1 }],
    customer: { email: 'a@b.com', taxExempt: true }
  });
  assert.strictEqual(result.total, 100);
  assert.strictEqual(result.taxRate, 0);
});

test('express shipping overrides carrier', () => {
  const system = new LegacySystem();
  const result = system.process_everything({
    id: 13,
    items: [{ price: 10, quantity: 1 }],
    shipping: { speed: 'express' },
    customer: { email: 'a@b.com' }
  });
  assert.strictEqual(result.carrier, 'FedEx');
});
