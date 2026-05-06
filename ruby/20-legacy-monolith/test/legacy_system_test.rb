# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/legacy_system'

class LegacySystemTest < Minitest::Test
  def setup
    @system = LegacySystem.new
  end

  def test_rejects_empty_order
    result = @system.process_everything({ id: 1, items: [] })
    assert_equal 'No items', result[:error]
  end

  def test_calculates_basic_total
    result = @system.process_everything(
      id: 1,
      items: [{ price: 10, quantity: 2 }],
      customer: { email: 'a@b.com' }
    )
    assert_equal 21.4, result[:total]
    assert_equal 'USPS', result[:carrier]
  end

  def test_applies_member_discount
    result = @system.process_everything(
      id: 2,
      items: [{ price: 100, quantity: 1 }],
      customer: { type: 'member', email: 'a@b.com' }
    )
    assert_equal 96.3, result[:total]
  end

  def test_applies_vip_discount
    result = @system.process_everything(
      id: 3,
      items: [{ price: 100, quantity: 1 }],
      customer: { type: 'vip', email: 'a@b.com' }
    )
    assert_equal 85.6, result[:total]
  end

  def test_applies_bonus_discount_over_100
    result = @system.process_everything(
      id: 4,
      items: [{ price: 200, quantity: 1 }],
      customer: { email: 'a@b.com' }
    )
    assert_equal 203.3, result[:total]
  end

  def test_uses_ups_for_large_total
    result = @system.process_everything(
      id: 5,
      items: [{ price: 60, quantity: 1 }],
      customer: { email: 'a@b.com' }
    )
    assert_equal 'UPS', result[:carrier]
  end

  def test_flags_high_total_for_review
    result = @system.process_everything(
      id: 6,
      items: [{ price: 5000, quantity: 1 }],
      customer: { email: 'a@b.com' }
    )
    assert_equal 'manual_review', result[:payment_status]
  end

  def test_includes_email_details
    result = @system.process_everything(
      id: 7,
      items: [{ price: 10, quantity: 1 }],
      customer: { email: 'user@test.com' }
    )
    assert_equal 'user@test.com', result[:email][:to]
    assert_includes result[:email][:subject], '7'
  end

  def test_includes_log_entry
    result = @system.process_everything(
      id: 8,
      items: [{ price: 10, quantity: 1 }],
      customer: { email: 'a@b.com' }
    )
    assert_includes result[:log], 'Order processed'
  end

  def test_includes_order_id
    result = @system.process_everything(
      id: 99,
      items: [{ price: 10, quantity: 1 }],
      customer: { email: 'a@b.com' }
    )
    assert_equal 99, result[:order_id]
  end
end
