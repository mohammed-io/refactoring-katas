import assert from 'node:assert';
import test from 'node:test';
import { KitchenTicket, Order } from '../src/kitchen_ticket.js';

function makeOrder({ items, customer, table, special = '', rush = false } = {}) {
  return new Order(
    items || [{ name: 'Burger', qty: 1 }],
    customer || { name: 'Alice', vip: false },
    table || { number: 5, zone: 'patio', server: 'Sam' },
    special,
    rush,
  );
}

test('prints table customer and server details', () => {
  const result = new KitchenTicket().print_ticket(makeOrder());
  assert.ok(result.includes('Table: 5'));
  assert.ok(result.includes('Zone: patio'));
  assert.ok(result.includes('Server: Sam'));
  assert.ok(result.includes('Customer: Alice'));
});

test('prints ticket with multiple items and count', () => {
  const result = new KitchenTicket().print_ticket(makeOrder({ items: [{ name: 'Burger', qty: 2 }, { name: 'Fries', qty: 1 }] }));
  assert.ok(result.includes('Burger x2'));
  assert.ok(result.includes('Fries x1'));
  assert.ok(result.includes('Items: 3'));
});

test('prints modifiers and allergy flags', () => {
  const result = new KitchenTicket().print_ticket(makeOrder({ items: [{ name: 'Salad', qty: 1, modifiers: ['no onion', 'dressing side'], allergy: 'nuts' }] }));
  assert.ok(result.includes('Salad x1 [no onion, dressing side] ALLERGY:nuts'));
});

test('prints vip and rush markers', () => {
  const result = new KitchenTicket().print_ticket(makeOrder({ customer: { name: 'Carol', vip: true }, rush: true }));
  assert.ok(result.includes('VIP'));
  assert.ok(result.includes('RUSH'));
});

test('omits special when empty but keeps separator', () => {
  const result = new KitchenTicket().print_ticket(makeOrder({ special: '' }));
  assert.ok(!result.includes('Special:'));
  assert.ok(result.includes('---'));
});

test('prints special instructions', () => {
  const result = new KitchenTicket().print_ticket(makeOrder({ special: 'Fire mains after starters' }));
  assert.ok(result.includes('Special: Fire mains after starters'));
});
