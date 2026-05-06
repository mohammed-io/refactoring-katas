import assert from 'node:assert';
import test from 'node:test';
import { Account } from '../src/account.js';

test('logs in with correct password', () => {
  const acc = new Account('a@b.com', 'secret');
  assert.strictEqual(acc.login('secret'), true);
});

test('rejects wrong password', () => {
  const acc = new Account('a@b.com', 'secret');
  assert.strictEqual(acc.login('wrong'), false);
});

test('logs out', () => {
  const acc = new Account('a@b.com', 'secret');
  assert.strictEqual(acc.logout(), true);
});

test('updates profile', () => {
  const acc = new Account('a@b.com', 'secret');
  const result = acc.update_profile('Alice', 'Developer');
  assert.strictEqual(result.name, 'Alice');
  assert.strictEqual(result.bio, 'Developer');
});

test('changes password', () => {
  const acc = new Account('a@b.com', 'secret');
  assert.strictEqual(acc.change_password('secret', 'new'), true);
  assert.strictEqual(acc.login('new'), true);
});

test('rejects bad old password', () => {
  const acc = new Account('a@b.com', 'secret');
  assert.strictEqual(acc.change_password('wrong', 'new'), false);
});

test('adds payment method', () => {
  const acc = new Account('a@b.com', 'secret');
  const count = acc.add_payment_method('Visa-1234');
  assert.strictEqual(count, 1);
});

test('removes payment method', () => {
  const acc = new Account('a@b.com', 'secret');
  acc.add_payment_method('Visa-1234');
  const result = acc.remove_payment_method(0);
  assert.deepStrictEqual(result, []);
});

test('sets notification preference', () => {
  const acc = new Account('a@b.com', 'secret');
  const result = acc.set_notification_preference('sms', true);
  assert.strictEqual(result.sms, true);
});

test('exports all data', () => {
  const acc = new Account('a@b.com', 'secret');
  const result = acc.export_data();
  assert.strictEqual(result.email, 'a@b.com');
  assert.ok(result.auditLog);
});

test('logs access', () => {
  const acc = new Account('a@b.com', 'secret');
  const count = acc.log_access('view');
  assert.strictEqual(count, 1);
});

test('checks subscription', () => {
  const acc = new Account('a@b.com', 'secret');
  assert.strictEqual(acc.check_subscription(), 'basic');
});

test('upgrades subscription', () => {
  const acc = new Account('a@b.com', 'secret');
  const result = acc.upgrade_subscription('pro');
  assert.strictEqual(result, 'pro');
});
