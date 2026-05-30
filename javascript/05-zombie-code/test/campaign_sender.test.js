import assert from 'node:assert';
import test from 'node:test';
import { CampaignSender } from '../src/campaign_sender.js';

test('counts EU customers with opt-in', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([{ active: true, region: 'EU', gdprOptIn: true }], 'Hi');
  assert.strictEqual(result.sent, 1);
});

test('skips EU customers without opt-in', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([{ active: true, region: 'EU', gdprOptIn: false }], 'Hi');
  assert.strictEqual(result.sent, 0);
});

test('counts non-EU active customers', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([{ active: true, region: 'US' }], 'Hi');
  assert.strictEqual(result.sent, 1);
});

test('skips inactive customers', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([{ active: false, region: 'US' }], 'Hi');
  assert.strictEqual(result.sent, 0);
  assert.strictEqual(result.skipped, 1);
});

test('handles mixed customers', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([
    { active: true, region: 'EU', gdprOptIn: true },
    { active: true, region: 'EU', gdprOptIn: false },
    { active: true, region: 'US' },
    { active: false, region: 'US' }
  ], 'Hi');
  assert.strictEqual(result.sent, 2);
  assert.strictEqual(result.skipped, 2);
});

test('skips unsubscribed customers', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([{ active: true, region: 'US', unsubscribed: true }], 'Hi');
  assert.strictEqual(result.sent, 0);
  assert.strictEqual(result.skipped, 1);
});

test('dry run never sends but keeps skip count', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([
    { active: true, region: 'US' },
    { active: false, region: 'US' }
  ], '__dry_run__');
  assert.strictEqual(result.sent, 0);
  assert.strictEqual(result.skipped, 1);
});

test('returns message in result', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([], 'Hello World');
  assert.strictEqual(result.message, 'Hello World');
});

test('returns timestamp in result', () => {
  const sender = new CampaignSender();
  const result = sender.send_campaign([], 'Hi');
  assert.ok(typeof result.timestamp === 'number');
});
