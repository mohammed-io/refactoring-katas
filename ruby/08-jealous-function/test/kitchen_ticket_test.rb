# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/kitchen_ticket'

class KitchenTicketTest < Minitest::Test
  def setup
    @ticket = KitchenTicket.new
  end

  def order(items: [{ name: 'Burger', qty: 1 }], customer: { name: 'Alice', vip: false },
            table: { number: 5, zone: 'patio', server: 'Sam' }, special: '', rush: false)
    Order.new(items, customer, table, special, rush)
  end

  def test_prints_table_customer_and_server_details
    result = @ticket.print_ticket(order)
    assert_includes result, 'Table: 5'
    assert_includes result, 'Zone: patio'
    assert_includes result, 'Server: Sam'
    assert_includes result, 'Customer: Alice'
  end

  def test_prints_ticket_with_multiple_items_and_count
    result = @ticket.print_ticket(order(items: [{ name: 'Burger', qty: 2 }, { name: 'Fries', qty: 1 }]))
    assert_includes result, 'Burger x2'
    assert_includes result, 'Fries x1'
    assert_includes result, 'Items: 3'
  end

  def test_prints_modifiers_and_allergy_flags
    result = @ticket.print_ticket(order(items: [{ name: 'Salad', qty: 1, modifiers: ['no onion', 'dressing side'], allergy: 'nuts' }]))
    assert_includes result, 'Salad x1 [no onion, dressing side] ALLERGY:nuts'
  end

  def test_prints_vip_and_rush_markers
    result = @ticket.print_ticket(order(customer: { name: 'Carol', vip: true }, rush: true))
    assert_includes result, 'VIP'
    assert_includes result, 'RUSH'
  end

  def test_omits_special_when_empty_but_keeps_separator
    result = @ticket.print_ticket(order(special: ''))
    refute_includes result, 'Special:'
    assert_includes result, '---'
  end

  def test_prints_special_instructions
    result = @ticket.print_ticket(order(special: 'Fire mains after starters'))
    assert_includes result, 'Special: Fire mains after starters'
  end
end
