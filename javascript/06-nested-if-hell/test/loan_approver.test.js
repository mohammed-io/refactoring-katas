import assert from 'node:assert';
import test from 'node:test';
import { LoanApprover } from '../src/loan_approver.js';

test('allows normal package', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 10, hazardous: false, weekend: false });
  assert.strictEqual(result.allowed, true);
  assert.strictEqual(result.warning, null);
});

test('rejects overweight', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 60, hazardous: false, weekend: false });
  assert.strictEqual(result.allowed, false);
  assert.strictEqual(result.warning, 'Weight exceeded');
});

test('rejects hazardous', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 10, hazardous: true, weekend: false });
  assert.strictEqual(result.allowed, false);
  assert.strictEqual(result.warning, 'Hazardous material');
});

test('rejects weekend', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 10, hazardous: false, weekend: true });
  assert.strictEqual(result.allowed, false);
  assert.strictEqual(result.warning, 'No weekend delivery');
});

test('rejects extreme temperature', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 10, hazardous: false, weekend: false, temperatureRequired: 50 });
  assert.strictEqual(result.allowed, false);
  assert.strictEqual(result.warning, 'Temperature out of range');
});

test('allows valid temperature', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 10, hazardous: false, weekend: false, temperatureRequired: 20 });
  assert.strictEqual(result.allowed, true);
});

test('allows remote small package', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 15, hazardous: false, weekend: false, remoteArea: true });
  assert.strictEqual(result.allowed, true);
  assert.strictEqual(result.warning, 'Remote surcharge applies');
});

test('rejects remote heavy package', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 25, hazardous: false, weekend: false, remoteArea: true });
  assert.strictEqual(result.allowed, false);
  assert.strictEqual(result.warning, 'Too heavy for remote');
});

test('rejects null package', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver(null);
  assert.strictEqual(result.allowed, false);
  assert.strictEqual(result.warning, 'No package');
});

test('rejects missing weight', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ hazardous: false });
  assert.strictEqual(result.allowed, false);
  assert.strictEqual(result.warning, 'No weight specified');
});

test('weight 50 is allowed at boundary', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 50, hazardous: false, weekend: false });
  assert.strictEqual(result.allowed, true);
  assert.strictEqual(result.warning, null);
});

test('temperature 40 is allowed at boundary', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 10, hazardous: false, weekend: false, temperatureRequired: 40 });
  assert.strictEqual(result.allowed, true);
});

test('temperature -20 is allowed at boundary', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 10, hazardous: false, weekend: false, temperatureRequired: -20 });
  assert.strictEqual(result.allowed, true);
});

test('remote area weight 20 is allowed at boundary', () => {
  const approver = new LoanApprover();
  const result = approver.can_deliver({ weight: 20, hazardous: false, weekend: false, remoteArea: true });
  assert.strictEqual(result.allowed, true);
  assert.strictEqual(result.warning, 'Remote surcharge applies');
});
