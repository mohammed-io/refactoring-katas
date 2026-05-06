import assert from 'node:assert';
import test from 'node:test';
import { RoomLookup } from '../src/room_lookup.js';

test('returns room name for valid chain', () => {
  const db = {
    getStudent: () => ({ id: 1 }),
    getEnrollment: () => ({ courseId: 2 }),
    getCourse: () => ({ defaultSectionId: 3 }),
    getSection: () => ({ roomId: 4 }),
    getRoom: () => ({ name: 'Room A' })
  };
  const lookup = new RoomLookup();
  assert.strictEqual(lookup.get_room_for_student(1, db), 'Room A');
});

test('returns null for null db', () => {
  const lookup = new RoomLookup();
  assert.strictEqual(lookup.get_room_for_student(1, null), null);
});

test('returns null for null student', () => {
  const db = {
    getStudent: () => null,
    getEnrollment: () => ({}),
    getCourse: () => ({}),
    getSection: () => ({}),
    getRoom: () => ({})
  };
  const lookup = new RoomLookup();
  assert.strictEqual(lookup.get_room_for_student(1, db), null);
});

test('returns null for null enrollment', () => {
  const db = {
    getStudent: () => ({ id: 1 }),
    getEnrollment: () => null,
    getCourse: () => ({}),
    getSection: () => ({}),
    getRoom: () => ({})
  };
  const lookup = new RoomLookup();
  assert.strictEqual(lookup.get_room_for_student(1, db), null);
});

test('returns null for null course', () => {
  const db = {
    getStudent: () => ({ id: 1 }),
    getEnrollment: () => ({ courseId: 2 }),
    getCourse: () => null,
    getSection: () => ({}),
    getRoom: () => ({})
  };
  const lookup = new RoomLookup();
  assert.strictEqual(lookup.get_room_for_student(1, db), null);
});

test('returns null for null section', () => {
  const db = {
    getStudent: () => ({ id: 1 }),
    getEnrollment: () => ({ courseId: 2 }),
    getCourse: () => ({ defaultSectionId: 3 }),
    getSection: () => null,
    getRoom: () => ({})
  };
  const lookup = new RoomLookup();
  assert.strictEqual(lookup.get_room_for_student(1, db), null);
});

test('returns null for null room', () => {
  const db = {
    getStudent: () => ({ id: 1 }),
    getEnrollment: () => ({ courseId: 2 }),
    getCourse: () => ({ defaultSectionId: 3 }),
    getSection: () => ({ roomId: 4 }),
    getRoom: () => null
  };
  const lookup = new RoomLookup();
  assert.strictEqual(lookup.get_room_for_student(1, db), null);
});
