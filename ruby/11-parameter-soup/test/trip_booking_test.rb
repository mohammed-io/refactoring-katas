# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/trip_booking'

class TripBookingTest < Minitest::Test
  def setup
    @booking = TripBooking.new
  end

  def base_args
    ['LAX', 'NYC', '2024-01-01', nil, 'economy', 'vegan', 'aisle', nil, false, nil, false]
  end

  def test_rejects_missing_origin
    result = @booking.book_trip(nil, 'NYC', '2024-01-01', nil, 'economy', 'vegan', 'aisle', nil, false, nil, false)
    assert_equal 'Missing route', result[:error]
  end

  def test_rejects_missing_destination
    result = @booking.book_trip('LAX', nil, '2024-01-01', nil, 'economy', 'vegan', 'aisle', nil, false, nil, false)
    assert_equal 'Missing route', result[:error]
  end

  def test_rejects_missing_departure_date
    result = @booking.book_trip('LAX', 'NYC', nil, nil, 'economy', 'vegan', 'aisle', nil, false, nil, false)
    assert_equal 'Missing departure', result[:error]
  end

  def test_calculates_economy_price
    result = @booking.book_trip(*base_args)
    assert_equal 200, result[:total]
    assert_equal 'economy', result[:class]
  end

  def test_calculates_business_price
    args = base_args; args[4] = 'business'
    result = @booking.book_trip(*args)
    assert_equal 800, result[:total]
  end

  def test_calculates_first_class_price
    args = base_args; args[4] = 'first'
    result = @booking.book_trip(*args)
    assert_equal 2000, result[:total]
  end

  def test_applies_save20_promo
    args = base_args; args[9] = 'SAVE20'
    result = @booking.book_trip(*args)
    assert_equal 160, result[:total]
  end

  def test_applies_save10_promo
    args = base_args; args[9] = 'SAVE10'
    result = @booking.book_trip(*args)
    assert_equal 180, result[:total]
  end

  def test_adds_insurance
    args = base_args; args[8] = true
    result = @booking.book_trip(*args)
    assert_equal 250, result[:total]
  end

  def test_adds_flexible_dates
    args = base_args; args[10] = true
    result = @booking.book_trip(*args)
    assert_equal 230, result[:total]
  end

  def test_applies_gold_loyalty_discount
    args = base_args; args[7] = 'GOLD123'
    result = @booking.book_trip(*args)
    assert_equal 175, result[:total]
  end

  def test_includes_route_in_result
    result = @booking.book_trip(*base_args)
    assert_equal 'LAX', result[:origin]
    assert_equal 'NYC', result[:destination]
  end

  def test_includes_confirmation_code
    result = @booking.book_trip(*base_args)
    assert result[:confirmation].start_with?('BK-')
  end
end
