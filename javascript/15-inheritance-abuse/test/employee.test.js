import assert from 'node:assert';
import test from 'node:test';
import { Employee, Manager, SeniorManager, Director } from '../src/employee.js';

test('employee gets 2% bonus', () => {
  const emp = new Employee('Alice', 50000);
  assert.strictEqual(emp.calculate_bonus(), 1000);
});

test('manager gets 5% bonus', () => {
  const mgr = new Manager('Bob', 80000);
  assert.strictEqual(mgr.calculate_bonus(), 4000);
});

test('senior manager capped at 10000', () => {
  const sm = new SeniorManager('Carol', 300000);
  assert.strictEqual(sm.calculate_bonus(), 10000);
});

test('senior manager under cap', () => {
  const sm = new SeniorManager('Carol', 100000);
  assert.strictEqual(sm.calculate_bonus(), 5000);
});

test('director capped at 20000', () => {
  const dir = new Director('Dave', 600000);
  assert.strictEqual(dir.calculate_bonus(), 20000);
});

test('director under cap', () => {
  const dir = new Director('Dave', 200000);
  assert.strictEqual(dir.calculate_bonus(), 10000);
});

test('manager respects cap from senior manager', () => {
  const mgr = new Manager('Eve', 300000);
  assert.strictEqual(mgr.calculate_bonus(), 15000);
});

test('employee total reward adds high performance and tenure', () => {
  const emp = new Employee('Alice', 50000);
  assert.strictEqual(emp.calculate_total_reward('high', 5), 2000);
});

test('senior manager total reward uses capped bonus and tenure rule', () => {
  const sm = new SeniorManager('Carol', 300000);
  assert.strictEqual(sm.calculate_total_reward('normal', 7), 11500);
});

test('director total reward uses director performance rule', () => {
  const dir = new Director('Dana', 200000);
  assert.strictEqual(dir.calculate_total_reward('high', 3), 16000);
});
