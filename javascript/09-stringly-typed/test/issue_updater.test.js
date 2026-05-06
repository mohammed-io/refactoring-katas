import assert from 'node:assert';
import test from 'node:test';
import { IssueUpdater } from '../src/issue_updater.js';

test('closes issue', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(42, 'close');
  assert.ok(result.includes('status=closed'));
});

test('opens issue', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(42, 'open');
  assert.ok(result.includes('status=open'));
});

test('sets in_progress', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(42, 'progress');
  assert.ok(result.includes('status=in_progress'));
});

test('sets priority 1', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(42, 'close:1');
  assert.ok(result.includes('priority=1'));
});

test('sets priority 2', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(42, 'open:2');
  assert.ok(result.includes('priority=2'));
});

test('defaults to priority 3', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(42, 'progress');
  assert.ok(result.includes('priority=3'));
});

test('includes issue id', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(99, 'close');
  assert.ok(result.includes('Issue 99'));
});

test('ignores invalid priority', () => {
  const updater = new IssueUpdater();
  const result = updater.update_issue(42, 'close:99');
  assert.ok(result.includes('priority=3'));
});
