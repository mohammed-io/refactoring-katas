import assert from 'node:assert';
import test from 'node:test';
import { TripBooking } from '../src/trip_booking.js';

test('rejects missing origin', () => {
  const booking = new TripBooking();
  const result = booking.book_trip(null, 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, null, false);
  assert.strictEqual(result.error, 'Missing route');
});

test('rejects missing destination', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', null, '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, null, false);
  assert.strictEqual(result.error, 'Missing route');
});

test('rejects missing departure date', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', null, null, 'economy', 'vegan', 'aisle', null, false, null, false);
  assert.strictEqual(result.error, 'Missing departure');
});

test('calculates economy price', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, null, false);
  assert.strictEqual(result.total, 200);
  assert.strictEqual(result.class, 'economy');
});

test('calculates business price', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'business', 'kosher', 'window', null, false, null, false);
  assert.strictEqual(result.total, 800);
});

test('calculates first class price', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'first', 'halal', 'window', null, false, null, false);
  assert.strictEqual(result.total, 2000);
});

test('applies SAVE20 promo', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, 'SAVE20', false);
  assert.strictEqual(result.total, 160);
});

test('applies SAVE10 promo', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, 'SAVE10', false);
  assert.strictEqual(result.total, 180);
});

test('adds insurance', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, true, null, false);
  assert.strictEqual(result.total, 250);
});

test('adds flexible dates', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, null, true);
  assert.strictEqual(result.total, 230);
});

test('applies gold loyalty discount', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', 'GOLD123', false, null, false);
  assert.strictEqual(result.total, 175);
});

test('includes route in result', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, null, false);
  assert.strictEqual(result.origin, 'LAX');
  assert.strictEqual(result.destination, 'NYC');
});

test('includes confirmation code', () => {
  const booking = new TripBooking();
  const result = booking.book_trip('LAX', 'NYC', '2024-01-01', null, 'economy', 'vegan', 'aisle', null, false, null, false);
  assert.ok(result.confirmation.startsWith('BK-'));
});
