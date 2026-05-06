# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/inventory_item'

class InventoryItemTest < Minitest::Test
  def setup
    @item = InventoryItem.new(1, 'Widget', 'B001', 123, 99, 10)
  end

  def test_stores_and_returns_id
    assert_equal 1, @item.get_id
    @item.set_id(2)
    assert_equal 2, @item.get_id
  end

  def test_stores_and_returns_name
    assert_equal 'Widget', @item.get_name
    @item.set_name('Gadget')
    assert_equal 'Gadget', @item.get_name
  end

  def test_stores_and_returns_batch_number
    assert_equal 'B001', @item.get_batch_number
    @item.set_batch_number('B002')
    assert_equal 'B002', @item.get_batch_number
  end

  def test_stores_and_returns_cache_timestamp
    assert_equal 123, @item.get_cache_timestamp
    @item.set_cache_timestamp(456)
    assert_equal 456, @item.get_cache_timestamp
  end

  def test_stores_and_returns_row_id
    assert_equal 99, @item.get_row_id
    @item.set_row_id(100)
    assert_equal 100, @item.get_row_id
  end

  def test_stores_and_returns_quantity
    assert_equal 10, @item.get_quantity
    @item.set_quantity(5)
    assert_equal 5, @item.get_quantity
  end
end
