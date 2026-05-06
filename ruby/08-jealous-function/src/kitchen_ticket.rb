# frozen_string_literal: true

class Order
  attr_accessor :items, :customer, :table, :special

  def initialize(items, customer, table, special)
    @items = items
    @customer = customer
    @table = table
    @special = special
  end
end

class KitchenTicket
  def print_ticket(order)
    lines = []
    lines << "Table: #{order.table}"
    lines << "Customer: #{order.customer}"
    order.items.each { |item| lines << "#{item[:name]} x#{item[:qty]}" }
    lines << "Special: #{order.special}" if order.special&.length&.positive?
    lines << '---'
    lines.join("\n")
  end
end
