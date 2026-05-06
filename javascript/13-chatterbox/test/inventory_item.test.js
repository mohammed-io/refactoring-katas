import assert from 'node:assert';
import test from 'node:test';
import { InventoryItem } from '../src/inventory_item.js';

test('stores and returns id', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  assert.strictEqual(item.get_id(), 1);
  item.set_id(2);
  assert.strictEqual(item.get_id(), 2);
});

test('stores and returns name', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  assert.strictEqual(item.get_name(), 'Widget');
  item.set_name('Gadget');
  assert.strictEqual(item.get_name(), 'Gadget');
});

test('stores and returns batch number', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  assert.strictEqual(item.get_batch_number(), 'B001');
  item.set_batch_number('B002');
  assert.strictEqual(item.get_batch_number(), 'B002');
});

test('stores and returns cache timestamp', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  assert.strictEqual(item.get_cache_timestamp(), 123);
  item.set_cache_timestamp(456);
  assert.strictEqual(item.get_cache_timestamp(), 456);
});

test('stores and returns row id', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  assert.strictEqual(item.get_row_id(), 99);
  item.set_row_id(100);
  assert.strictEqual(item.get_row_id(), 100);
});

test('stores and returns quantity', () => {
  const item = new InventoryItem(1, 'Widget', 'B001', 123, 99, 10);
  assert.strictEqual(item.get_quantity(), 10);
  item.set_quantity(5);
  assert.strictEqual(item.get_quantity(), 5);
});
