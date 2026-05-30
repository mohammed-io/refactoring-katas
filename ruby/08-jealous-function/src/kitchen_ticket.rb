# frozen_string_literal: true

class Order
  attr_accessor :items, :customer, :table, :special, :rush

  def initialize(items, customer, table, special, rush = false)
    @items = items
    @customer = customer
    @table = table
    @special = special
    @rush = rush
  end
end

class KitchenTicket
  def print_ticket(order)
    lines = []
    lines << "Table: #{order.table[:number]}"
    lines << "Zone: #{order.table[:zone] || 'main'}"
    lines << "Server: #{order.table[:server] || 'unassigned'}"
    lines << "Customer: #{order.customer[:name]}"
    lines << 'VIP' if order.customer[:vip]
    lines << 'RUSH' if order.rush
    total_items = 0
    order.items.each do |item|
      total_items += item[:qty]
      line = "#{item[:name]} x#{item[:qty]}"
      line += " [#{item[:modifiers].join(', ')}]" if item[:modifiers]&.any?
      line += " ALLERGY:#{item[:allergy]}" if item[:allergy]
      lines << line
    end
    lines << "Items: #{total_items}"
    lines << "Special: #{order.special}" if order.special&.length&.positive?
    lines << '---'
    lines.join("\n")
  end
end
