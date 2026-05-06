import assert from 'node:assert';
import test from 'node:test';
import { KitchenTicket } from '../src/kitchen_ticket.js';

test('prints simple ticket', () => {
  const order = new KitchenTicket([{ name: 'Burger', qty: 1 }], 'Alice', 5, '');
  const result = order.print_ticket();
  assert.ok(result.includes('Table: 5'));
  assert.ok(result.includes('Customer: Alice'));
  assert.ok(result.includes('Burger x1'));
});

test('prints ticket with multiple items', () => {
  const order = new KitchenTicket([{ name: 'Burger', qty: 2 }, { name: 'Fries', qty: 1 }], 'Bob', 12, '');
  const result = order.print_ticket();
  assert.ok(result.includes('Burger x2'));
  assert.ok(result.includes('Fries x1'));
});

test('prints ticket with special instructions', () => {
  const order = new KitchenTicket([{ name: 'Salad', qty: 1 }], 'Carol', 3, 'No onions');
  const result = order.print_ticket();
  assert.ok(result.includes('Special: No onions'));
});

test('omits special when empty', () => {
  const order = new KitchenTicket([{ name: 'Pizza', qty: 1 }], 'Dave', 7, '');
  const result = order.print_ticket();
  assert.ok(!result.includes('Special:'));
});

test('includes separator line', () => {
  const order = new KitchenTicket([{ name: 'Soup', qty: 1 }], 'Eve', 1, '');
  const result = order.print_ticket();
  assert.ok(result.includes('---'));
});
