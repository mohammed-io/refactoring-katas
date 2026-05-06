import assert from 'node:assert';
import test from 'node:test';
import { PlaylistCurator } from '../src/playlist_curator.js';

test('filters happy tracks', () => {
  const curator = new PlaylistCurator();
  const tracks = [{ title: 'A', tempo: 130 }, { title: 'B', tempo: 100 }];
  const result = curator.create_playlist('happy', tracks);
  assert.strictEqual(result.length, 1);
  assert.strictEqual(result[0].title, 'A');
});

test('filters sad tracks', () => {
  const curator = new PlaylistCurator();
  const tracks = [{ title: 'A', tempo: 80 }, { title: 'B', tempo: 100 }];
  const result = curator.create_playlist('sad', tracks);
  assert.strictEqual(result.length, 1);
  assert.strictEqual(result[0].title, 'A');
});

test('filters workout tracks', () => {
  const curator = new PlaylistCurator();
  const tracks = [{ title: 'A', tempo: 140, energy: 8 }, { title: 'B', tempo: 140, energy: 5 }];
  const result = curator.create_playlist('workout', tracks);
  assert.strictEqual(result.length, 1);
  assert.strictEqual(result[0].title, 'A');
});

test('filters focus tracks', () => {
  const curator = new PlaylistCurator();
  const tracks = [{ title: 'A', instrumental: true }, { title: 'B', instrumental: false }];
  const result = curator.create_playlist('focus', tracks);
  assert.strictEqual(result.length, 1);
  assert.strictEqual(result[0].title, 'A');
});

test('filters party tracks', () => {
  const curator = new PlaylistCurator();
  const tracks = [{ title: 'A', tempo: 120, danceability: 7 }, { title: 'B', tempo: 100, danceability: 5 }];
  const result = curator.create_playlist('party', tracks);
  assert.strictEqual(result.length, 1);
  assert.strictEqual(result[0].title, 'A');
});

test('caps happy playlist at 20', () => {
  const curator = new PlaylistCurator();
  const tracks = Array(25).fill(null).map((_, i) => ({ title: String(i), tempo: 130 }));
  const result = curator.create_playlist('happy', tracks);
  assert.strictEqual(result.length, 20);
});

test('caps sad playlist at 15', () => {
  const curator = new PlaylistCurator();
  const tracks = Array(20).fill(null).map((_, i) => ({ title: String(i), tempo: 80 }));
  const result = curator.create_playlist('sad', tracks);
  assert.strictEqual(result.length, 15);
});

test('defaults to first 10 for unknown mood', () => {
  const curator = new PlaylistCurator();
  const tracks = Array(20).fill(null).map((_, i) => ({ title: String(i), tempo: 100 }));
  const result = curator.create_playlist('mysterious', tracks);
  assert.strictEqual(result.length, 10);
});

test('sorts happy by descending tempo', () => {
  const curator = new PlaylistCurator();
  const tracks = [{ title: 'A', tempo: 130 }, { title: 'B', tempo: 150 }];
  const result = curator.create_playlist('happy', tracks);
  assert.strictEqual(result[0].title, 'B');
  assert.strictEqual(result[1].title, 'A');
});

test('sorts sad by ascending tempo', () => {
  const curator = new PlaylistCurator();
  const tracks = [{ title: 'A', tempo: 80 }, { title: 'B', tempo: 60 }];
  const result = curator.create_playlist('sad', tracks);
  assert.strictEqual(result[0].title, 'B');
  assert.strictEqual(result[1].title, 'A');
});
