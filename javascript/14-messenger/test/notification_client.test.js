import assert from 'node:assert';
import test from 'node:test';
import { NotificationClient } from '../src/notification_client.js';

test('sends notification', () => {
  const client = new NotificationClient();
  const result = client.send({ message: 'Hello' });
  assert.strictEqual(result.status, 'sent');
  assert.deepStrictEqual(result.payload, { message: 'Hello' });
});

test('returns sent status', () => {
  const client = new NotificationClient();
  const result = client.send({ message: 'Test' });
  assert.strictEqual(result.status, 'sent');
});

test('preserves payload', () => {
  const client = new NotificationClient();
  const payload = { alert: 'Urgent', level: 3 };
  const result = client.send(payload);
  assert.deepStrictEqual(result.payload, payload);
});
