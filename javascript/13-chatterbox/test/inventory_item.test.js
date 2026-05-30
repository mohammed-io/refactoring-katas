import assert from 'node:assert';
import test from 'node:test';
import { InventoryItem } from '../src/inventory_item.js';

test('public snapshot contains business fields', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  assert.deepStrictEqual(item.public_snapshot(), {
    id: 1,
    name: 'Widget',
    batch_number: 'B001',
    quantity: 10,
    stock_status: 'available',
  });
});

test('reserves stock and reports remaining quantity', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  const result = item.reserve(3);
  assert.deepStrictEqual(result, { status: 'reserved', reserved: 3, remaining: 7, sku: '1-B001' });
  assert.strictEqual(item.public_snapshot().quantity, 7);
});

test('rejects reservation when quantity is invalid', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  const result = item.reserve(0);
  assert.deepStrictEqual(result, { status: 'rejected', reason: 'invalid_quantity', remaining: 10 });
  assert.strictEqual(item.public_snapshot().quantity, 10);
});

test('backorders when not enough stock', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  const result = item.reserve(12);
  assert.deepStrictEqual(result, { status: 'backorder', reserved: 0, remaining: 10 });
  assert.strictEqual(item.public_snapshot().quantity, 10);
});

test('receives stock after low quantity', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  item.reserve(8);
  assert.strictEqual(item.public_snapshot().stock_status, 'low');
  assert.strictEqual(item.receive_stock(5), 7);
  assert.strictEqual(item.public_snapshot().stock_status, 'available');
});

test('reports out of stock after exact reservation', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  item.reserve(10);
  assert.strictEqual(item.public_snapshot().stock_status, 'out');
});
