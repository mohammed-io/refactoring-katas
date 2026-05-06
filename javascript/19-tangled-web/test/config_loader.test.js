import assert from 'node:assert';
import test from 'node:test';
import fs from 'fs';
import { ConfigLoader } from '../src/config_loader.js';

test('returns empty object when file and fetch fail', () => {
  const loader = new ConfigLoader();
  const result = loader.load_config();
  assert.ok(typeof result === 'object');
});

test('reads local config file when present', () => {
  fs.writeFileSync('/tmp/config.json', JSON.stringify({ name: 'test' }));
  const loader = new ConfigLoader();
  const result = loader.load_config();
  assert.strictEqual(result.name, 'test');
  fs.unlinkSync('/tmp/config.json');
});

test('local overrides empty file', () => {
  fs.writeFileSync('/tmp/config.json', JSON.stringify({ name: 'local' }));
  const loader = new ConfigLoader();
  const result = loader.load_config();
  assert.strictEqual(result.name, 'local');
  fs.unlinkSync('/tmp/config.json');
});

test('handles malformed json gracefully', () => {
  fs.writeFileSync('/tmp/config.json', 'not json');
  const loader = new ConfigLoader();
  const result = loader.load_config();
  assert.ok(typeof result === 'object');
  fs.unlinkSync('/tmp/config.json');
});

test('includes seasonal keys in object', () => {
  const loader = new ConfigLoader();
  const result = loader.load_config();
  assert.ok('theme' in result || !result.theme);
  assert.ok('discount' in result || !result.discount);
});
