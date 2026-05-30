# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/inventory_item'

class InventoryItemTest < Minitest::Test
  def setup
    @item = InventoryItem.new(1, 'Widget', 'B001', 123, 99, 10)
  end

  def test_public_snapshot_contains_business_fields
    assert_equal(
      { id: 1, name: 'Widget', batch_number: 'B001', quantity: 10, stock_status: 'available' },
      @item.public_snapshot
    )
  end

  def test_reserves_stock_and_reports_remaining_quantity
    result = @item.reserve(3)

    assert_equal({ status: 'reserved', reserved: 3, remaining: 7, sku: '1-B001' }, result)
    assert_equal 7, @item.public_snapshot[:quantity]
  end

  def test_rejects_reservation_when_quantity_is_invalid
    result = @item.reserve(0)

    assert_equal({ status: 'rejected', reason: 'invalid_quantity', remaining: 10 }, result)
    assert_equal 10, @item.public_snapshot[:quantity]
  end

  def test_backorders_when_not_enough_stock
    result = @item.reserve(12)

    assert_equal({ status: 'backorder', reserved: 0, remaining: 10 }, result)
    assert_equal 10, @item.public_snapshot[:quantity]
  end

  def test_receives_stock_after_low_quantity
    @item.reserve(8)

    assert_equal 'low', @item.public_snapshot[:stock_status]
    assert_equal 7, @item.receive_stock(5)
    assert_equal 'available', @item.public_snapshot[:stock_status]
  end

  def test_reports_out_of_stock_after_exact_reservation
    @item.reserve(10)

    assert_equal 'out', @item.public_snapshot[:stock_status]
  end
end
