import assert from 'node:assert';
import test from 'node:test';
import { OrderProcessor } from '../src/order_processor.js';

test('rejects empty items', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [], customer: { email: 'a@b.com' }, address: { zip: '12345' } });
  assert.strictEqual(result.error, 'No items');
});

test('rejects invalid customer', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 10, quantity: 1 }], customer: {}, address: { zip: '12345' } });
  assert.strictEqual(result.error, 'Invalid customer');
});

test('rejects invalid address', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 10, quantity: 1 }], customer: { email: 'a@b.com' }, address: {} });
  assert.strictEqual(result.error, 'Invalid address');
});

test('rejects out of stock', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 10, quantity: 101 }], customer: { email: 'a@b.com' }, address: { zip: '12345' } });
  assert.strictEqual(result.error, 'Out of stock');
});

test('calculates totals for small order', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 10, quantity: 1 }], customer: { email: 'a@b.com' }, address: { zip: '12345' } });
  assert.strictEqual(result.total, 16.69);
  assert.strictEqual(result.shippingLabel.carrier, 'USPS');
});

test('calculates totals for medium order', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 20, quantity: 1 }], customer: { email: 'a@b.com' }, address: { zip: '12345' } });
  assert.strictEqual(result.total, 27.39);
});

test('uses UPS for large orders', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 100, quantity: 1 }], customer: { email: 'a@b.com' }, address: { zip: '12345' } });
  assert.strictEqual(result.shippingLabel.carrier, 'UPS');
  assert.strictEqual(result.paymentStatus, 'approved');
});

test('flags high total as pending review', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 1000, quantity: 1 }], customer: { email: 'a@b.com' }, address: { zip: '12345' } });
  assert.strictEqual(result.paymentStatus, 'pending_review');
});

test('includes email confirmation', () => {
  const processor = new OrderProcessor();
  const result = processor.process_order({ items: [{ price: 10, quantity: 1 }], customer: { email: 'user@test.com' }, address: { zip: '12345' } });
  assert.strictEqual(result.email.to, 'user@test.com');
  assert.strictEqual(result.email.subject, 'Order Confirmation');
});
