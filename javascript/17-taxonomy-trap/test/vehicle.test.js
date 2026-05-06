import assert from 'node:assert';
import test from 'node:test';
import { Car, Truck, ElectricCar, DieselCar, ElectricTruck, DieselTruck } from '../src/vehicle.js';

test('car daily rate', () => {
  const v = new Car('Toyota', 'Camry', 2020);
  assert.strictEqual(v.daily_rate(), 40);
});

test('truck daily rate', () => {
  const v = new Truck('Ford', 'F-150', 2020);
  assert.strictEqual(v.daily_rate(), 80);
});

test('electric car fuel cost', () => {
  const v = new ElectricCar('Tesla', 'Model 3', 2020);
  assert.strictEqual(v.fuel_cost(5), 0);
});

test('diesel car fuel cost', () => {
  const v = new DieselCar('VW', 'Jetta', 2020);
  assert.strictEqual(v.fuel_cost(5), 25);
});

test('electric truck fuel cost', () => {
  const v = new ElectricTruck('Rivian', 'R1T', 2020);
  assert.strictEqual(v.fuel_cost(5), 0);
});

test('diesel truck fuel cost', () => {
  const v = new DieselTruck('Ford', 'F-250', 2020);
  assert.strictEqual(v.fuel_cost(5), 75);
});

test('car stores brand', () => {
  const v = new Car('Toyota', 'Camry', 2020);
  assert.strictEqual(v.brand, 'Toyota');
});

test('truck stores model', () => {
  const v = new Truck('Ford', 'F-150', 2020);
  assert.strictEqual(v.model, 'F-150');
});
