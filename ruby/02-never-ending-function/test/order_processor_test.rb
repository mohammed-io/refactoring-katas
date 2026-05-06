# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/order_processor'

class OrderProcessorTest < Minitest::Test
  def setup
    @processor = OrderProcessor.new
  end

  def base_order(price: 10, quantity: 1)
    {
      items: [{ price: price, quantity: quantity }],
      customer: { email: 'a@b.com' },
      address: { zip: '12345' }
    }
  end

  def test_rejects_empty_items
    result = @processor.process_order(base_order.merge(items: []))
    assert_equal 'No items', result[:error]
  end

  def test_rejects_invalid_customer
    result = @processor.process_order(base_order.merge(customer: {}))
    assert_equal 'Invalid customer', result[:error]
  end

  def test_rejects_invalid_address
    result = @processor.process_order(base_order.merge(address: {}))
    assert_equal 'Invalid address', result[:error]
  end

  def test_rejects_out_of_stock
    result = @processor.process_order(base_order(quantity: 101))
    assert_equal 'Out of stock', result[:error]
  end

  def test_calculates_totals_for_small_order
    result = @processor.process_order(base_order)
    assert_equal 16.69, result[:total]
    assert_equal 'USPS', result[:shipping_label][:carrier]
  end

  def test_calculates_totals_for_medium_order
    result = @processor.process_order(base_order(price: 20))
    assert_equal 27.39, result[:total]
  end

  def test_uses_ups_for_large_orders
    result = @processor.process_order(base_order(price: 100))
    assert_equal 'UPS', result[:shipping_label][:carrier]
    assert_equal 'approved', result[:payment_status]
  end

  def test_flags_high_total_as_pending_review
    result = @processor.process_order(base_order(price: 1000))
    assert_equal 'pending_review', result[:payment_status]
  end

  def test_includes_email_confirmation
    result = @processor.process_order(base_order.merge(customer: { email: 'user@test.com' }))
    assert_equal 'user@test.com', result[:email][:to]
    assert_equal 'Order Confirmation', result[:email][:subject]
  end
end
