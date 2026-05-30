import assert from 'node:assert';
import test from 'node:test';
import fs from 'fs';
import os from 'os';
import path from 'path';
import { ConfigLoader } from '../src/config_loader.js';

function withEnv(values, fn) {
  const oldEnv = { ...process.env };
  Object.assign(process.env, values);
  try {
    fn();
  } finally {
    process.env = oldEnv;
  }
}

test('returns defaults when file is missing', () => {
  withEnv({ APP_CONFIG_PATH: path.join(os.tmpdir(), 'missing-config.json') }, () => {
    const loader = new ConfigLoader();
    const result = loader.load_config();
    assert.strictEqual(result.retries, 3);
    assert.strictEqual(result.theme, 'standard');
    assert.strictEqual(result.discount, 0);
  });
});

test('reads local config file when present', () => {
  const dir = fs.mkdtempSync(path.join(os.tmpdir(), 'config-loader-'));
  const configPath = path.join(dir, 'config.json');
  fs.writeFileSync(configPath, JSON.stringify({ name: 'test', retries: 5 }));
  withEnv({ APP_CONFIG_PATH: configPath }, () => {
    const loader = new ConfigLoader();
    const result = loader.load_config();
    assert.strictEqual(result.name, 'test');
    assert.strictEqual(result.retries, 5);
  });
  fs.rmSync(dir, { recursive: true, force: true });
});

test('environment overrides local config', () => {
  const dir = fs.mkdtempSync(path.join(os.tmpdir(), 'config-loader-'));
  const configPath = path.join(dir, 'config.json');
  fs.writeFileSync(configPath, JSON.stringify({ theme: 'local', retries: 5 }));
  withEnv({ APP_CONFIG_PATH: configPath, APP_THEME: 'env-theme', APP_RETRIES: '9' }, () => {
    const loader = new ConfigLoader();
    const result = loader.load_config();
    assert.strictEqual(result.theme, 'env-theme');
    assert.strictEqual(result.retries, 9);
  });
  fs.rmSync(dir, { recursive: true, force: true });
});

test('handles malformed json gracefully', () => {
  const dir = fs.mkdtempSync(path.join(os.tmpdir(), 'config-loader-'));
  const configPath = path.join(dir, 'config.json');
  fs.writeFileSync(configPath, 'not json');
  withEnv({ APP_CONFIG_PATH: configPath }, () => {
    const loader = new ConfigLoader();
    const result = loader.load_config();
    assert.strictEqual(result.theme, 'standard');
  });
  fs.rmSync(dir, { recursive: true, force: true });
});

test('winter seasonal config has highest precedence', () => {
  const dir = fs.mkdtempSync(path.join(os.tmpdir(), 'config-loader-'));
  const configPath = path.join(dir, 'config.json');
  fs.writeFileSync(configPath, JSON.stringify({ theme: 'local', discount: 0.25 }));
  withEnv({ APP_CONFIG_PATH: configPath, APP_THEME: 'env-theme', APP_CURRENT_MONTH: '12' }, () => {
    const loader = new ConfigLoader();
    const result = loader.load_config();
    assert.strictEqual(result.theme, 'winter');
    assert.strictEqual(result.discount, 0.1);
  });
  fs.rmSync(dir, { recursive: true, force: true });
});

test('summer seasonal config is deterministic', () => {
  withEnv({ APP_CONFIG_PATH: path.join(os.tmpdir(), 'missing-config.json'), APP_CURRENT_MONTH: '7' }, () => {
    const loader = new ConfigLoader();
    const result = loader.load_config();
    assert.strictEqual(result.theme, 'summer');
    assert.strictEqual(result.discount, 0.05);
  });
});
