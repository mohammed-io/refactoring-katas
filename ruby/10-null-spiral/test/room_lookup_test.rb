# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/room_lookup'

class Db
  def initialize(missing = nil)
    @missing = missing
  end

  def get_student(_)
    @missing == 'student' ? nil : { id: 1 }
  end

  def get_enrollment(_)
    @missing == 'enrollment' ? nil : { course_id: 2 }
  end

  def get_course(_)
    @missing == 'course' ? nil : { default_section_id: 3 }
  end

  def get_section(_)
    @missing == 'section' ? nil : { room_id: 4 }
  end

  def get_room(_)
    @missing == 'room' ? nil : { name: 'Room A' }
  end
end

class RoomLookupTest < Minitest::Test
  def setup
    @lookup = RoomLookup.new
  end

  def test_returns_room_name_for_valid_chain
    assert_equal 'Room A', @lookup.get_room_for_student(1, Db.new)
  end

  def test_returns_nil_for_nil_db
    assert_nil @lookup.get_room_for_student(1, nil)
  end

  def test_returns_nil_for_nil_student
    assert_nil @lookup.get_room_for_student(1, Db.new('student'))
  end

  def test_returns_nil_for_nil_enrollment
    assert_nil @lookup.get_room_for_student(1, Db.new('enrollment'))
  end

  def test_returns_nil_for_nil_course
    assert_nil @lookup.get_room_for_student(1, Db.new('course'))
  end

  def test_returns_nil_for_nil_section
    assert_nil @lookup.get_room_for_student(1, Db.new('section'))
  end

  def test_returns_nil_for_nil_room
    assert_nil @lookup.get_room_for_student(1, Db.new('room'))
  end
end
