# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/receipt'

class ReceiptTest < Minitest::Test
  def setup
    @receipt = Receipt.new
  end

  def test_regular_customer
    assert_equal 59.4, @receipt.calculate_total([10, 20, 30])
  end

  def test_member_customer
    assert_equal 56.16, @receipt.calculate_total([10, 20, 30], 'member')
  end

  def test_vip_customer
    assert_equal 47.68, @receipt.calculate_total([10, 20, 30], 'vip')
  end

  def test_bonus_discount_over_50
    assert_equal 59.4, @receipt.calculate_total([60])
  end

  def test_vip_extra_discount
    assert_equal 84.4, @receipt.calculate_total([100], 'vip')
  end

  def test_empty_items
    assert_equal 0, @receipt.calculate_total([])
  end

  def test_exactly_50_no_bonus
    assert_equal 54.0, @receipt.calculate_total([50])
  end
end
