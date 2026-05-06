import assert from 'node:assert';
import test from 'node:test';
import { PayrollCalculator } from '../src/payroll_calculator.js';

test('calculates fulltime payslip', () => {
  const calculator = new PayrollCalculator();
  const result = calculator.generate_payslips([{ id: 1, name: 'Alice', type: 'fulltime', salary: 60000 }]);
  assert.strictEqual(result[0].gross, 5000);
  assert.strictEqual(result[0].deductions, 1250);
  assert.strictEqual(result[0].net, 3750);
});

test('calculates fulltime with bonus', () => {
  const calculator = new PayrollCalculator();
  const result = calculator.generate_payslips([{ id: 1, name: 'Alice', type: 'fulltime', salary: 60000, bonus: 12000 }]);
  assert.strictEqual(result[0].gross, 6000);
  assert.strictEqual(result[0].net, 4750);
});

test('calculates parttime payslip', () => {
  const calculator = new PayrollCalculator();
  const result = calculator.generate_payslips([{ id: 2, name: 'Bob', type: 'parttime', hours: 80, rate: 25 }]);
  assert.strictEqual(result[0].gross, 2000);
  assert.strictEqual(result[0].deductions, 300);
  assert.strictEqual(result[0].net, 1700);
});

test('calculates contract payslip', () => {
  const calculator = new PayrollCalculator();
  const result = calculator.generate_payslips([{ id: 3, name: 'Carol', type: 'contract', flatFee: 5000 }]);
  assert.strictEqual(result[0].gross, 5000);
  assert.strictEqual(result[0].deductions, 500);
  assert.strictEqual(result[0].net, 4500);
});

test('handles multiple employees', () => {
  const calculator = new PayrollCalculator();
  const result = calculator.generate_payslips([
    { id: 1, name: 'Alice', type: 'fulltime', salary: 60000 },
    { id: 2, name: 'Bob', type: 'parttime', hours: 80, rate: 25 }
  ]);
  assert.strictEqual(result.length, 2);
  assert.strictEqual(result[0].name, 'Alice');
  assert.strictEqual(result[1].name, 'Bob');
});

test('returns empty array for empty input', () => {
  const calculator = new PayrollCalculator();
  const result = calculator.generate_payslips([]);
  assert.deepStrictEqual(result, []);
});
