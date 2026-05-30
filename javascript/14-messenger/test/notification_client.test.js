import assert from 'node:assert';
import test from 'node:test';
import { NotificationClient } from '../src/notification_client.js';

test('sends notification through layers', () => {
  const client = new NotificationClient();
  const result = client.send({ recipient: 'user-1', message: 'Hello', channel: 'sms' });
  assert.strictEqual(result.status, 'sent');
  assert.strictEqual(result.delivery_id, 'sms-user-1');
  assert.strictEqual(result.payload.message, 'Hello');
});

test('defaults channel and priority', () => {
  const client = new NotificationClient();
  const result = client.send({ message: 'Test' });
  assert.strictEqual(result.status, 'sent');
  assert.strictEqual(result.payload.channel, 'email');
  assert.strictEqual(result.payload.priority, 'normal');
});

test('preserves explicit priority', () => {
  const client = new NotificationClient();
  const result = client.send({ recipient: 'ops', message: 'Urgent', priority: 'high' });
  assert.strictEqual(result.payload.priority, 'high');
  assert.strictEqual(result.payload.recipient, 'ops');
});

test('rejects missing message', () => {
  const client = new NotificationClient();
  const result = client.send({ recipient: 'ops' });
  assert.strictEqual(result.status, 'rejected');
  assert.strictEqual(result.reason, 'missing_message');
});

test('records observable audit events', () => {
  const client = new NotificationClient();
  const result = client.send({ message: 'Deploy complete', channel: 'push' });
  assert.deepStrictEqual(result.audit, ['queued:push', 'sent:push']);
});

test('reports failed delivery for unsupported channel', () => {
  const client = new NotificationClient();
  const result = client.send({ recipient: 'ops', message: 'Legacy alert', channel: 'fax' });
  assert.strictEqual(result.status, 'failed');
  assert.strictEqual(result.reason, 'unsupported_channel');
  assert.deepStrictEqual(result.audit, ['queued:fax', 'failed:fax']);
});

test('high priority failed delivery is scheduled for retry', () => {
  const client = new NotificationClient();
  const result = client.send({ recipient: 'ops', message: 'Legacy alert', channel: 'fax', priority: 'high' });
  assert.strictEqual(result.status, 'retrying');
  assert.strictEqual(result.reason, 'unsupported_channel');
  assert.deepStrictEqual(result.audit, ['queued:fax', 'failed:fax', 'retry_scheduled:fax']);
});
