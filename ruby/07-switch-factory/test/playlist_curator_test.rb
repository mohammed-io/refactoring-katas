# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/playlist_curator'

class PlaylistCuratorTest < Minitest::Test
  def setup
    @curator = PlaylistCurator.new
  end

  def test_filters_happy_tracks
    result = @curator.create_playlist('happy', [{ title: 'A', tempo: 130 }, { title: 'B', tempo: 100 }])
    assert_equal 1, result.length
    assert_equal 'A', result[0][:title]
  end

  def test_filters_sad_tracks
    result = @curator.create_playlist('sad', [{ title: 'A', tempo: 80 }, { title: 'B', tempo: 100 }])
    assert_equal 1, result.length
    assert_equal 'A', result[0][:title]
  end

  def test_filters_workout_tracks
    result = @curator.create_playlist('workout', [{ title: 'A', tempo: 140, energy: 8 }, { title: 'B', tempo: 140, energy: 5 }])
    assert_equal 1, result.length
    assert_equal 'A', result[0][:title]
  end

  def test_filters_focus_tracks
    result = @curator.create_playlist('focus', [{ title: 'A', instrumental: true }, { title: 'B', instrumental: false }])
    assert_equal 1, result.length
    assert_equal 'A', result[0][:title]
  end

  def test_filters_party_tracks
    result = @curator.create_playlist('party', [{ title: 'A', tempo: 120, danceability: 7 }, { title: 'B', tempo: 100, danceability: 5 }])
    assert_equal 1, result.length
    assert_equal 'A', result[0][:title]
  end

  def test_caps_happy_playlist_at_20
    tracks = Array.new(25) { |i| { title: i.to_s, tempo: 130 } }
    result = @curator.create_playlist('happy', tracks)
    assert_equal 20, result.length
  end

  def test_caps_sad_playlist_at_15
    tracks = Array.new(20) { |i| { title: i.to_s, tempo: 80 } }
    result = @curator.create_playlist('sad', tracks)
    assert_equal 15, result.length
  end

  def test_defaults_to_first_10_for_unknown_mood
    tracks = Array.new(20) { |i| { title: i.to_s, tempo: 100 } }
    result = @curator.create_playlist('mysterious', tracks)
    assert_equal 10, result.length
  end

  def test_sorts_happy_by_descending_tempo
    result = @curator.create_playlist('happy', [{ title: 'A', tempo: 130 }, { title: 'B', tempo: 150 }])
    assert_equal 'B', result[0][:title]
    assert_equal 'A', result[1][:title]
  end

  def test_sorts_sad_by_ascending_tempo
    result = @curator.create_playlist('sad', [{ title: 'A', tempo: 80 }, { title: 'B', tempo: 60 }])
    assert_equal 'B', result[0][:title]
    assert_equal 'A', result[1][:title]
  end
end
