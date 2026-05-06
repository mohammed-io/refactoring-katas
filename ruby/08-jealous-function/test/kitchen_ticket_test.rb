# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/kitchen_ticket'

class KitchenTicketTest < Minitest::Test
  def setup
    @ticket = KitchenTicket.new
  end

  def test_prints_simple_ticket
    order = Order.new([{ name: 'Burger', qty: 1 }], 'Alice', 5, '')
    result = @ticket.print_ticket(order)
    assert_includes result, 'Table: 5'
    assert_includes result, 'Customer: Alice'
    assert_includes result, 'Burger x1'
  end

  def test_prints_ticket_with_multiple_items
    order = Order.new([{ name: 'Burger', qty: 2 }, { name: 'Fries', qty: 1 }], 'Bob', 12, '')
    result = @ticket.print_ticket(order)
    assert_includes result, 'Burger x2'
    assert_includes result, 'Fries x1'
  end

  def test_prints_ticket_with_special_instructions
    order = Order.new([{ name: 'Salad', qty: 1 }], 'Carol', 3, 'No onions')
    result = @ticket.print_ticket(order)
    assert_includes result, 'Special: No onions'
  end

  def test_omits_special_when_empty
    order = Order.new([{ name: 'Pizza', qty: 1 }], 'Dave', 7, '')
    result = @ticket.print_ticket(order)
    refute_includes result, 'Special:'
  end

  def test_includes_separator_line
    order = Order.new([{ name: 'Soup', qty: 1 }], 'Eve', 1, '')
    result = @ticket.print_ticket(order)
    assert_includes result, '---'
  end
end
