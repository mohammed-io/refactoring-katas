# frozen_string_literal: true

class OrderProcessor
  def process_order(order)
    return { error: 'No items' } if !order[:items] || order[:items].empty?
    return { error: 'Invalid customer' } if !order[:customer] || !order[:customer][:email]
    return { error: 'Invalid address' } if !order[:address] || !order[:address][:zip]

    inventory = true
    order[:items].each { |item| inventory = false if item[:quantity] > 100 }
    return { error: 'Out of stock' } unless inventory

    subtotal = 0
    order[:items].each { |item| subtotal += item[:price] * item[:quantity] }
    shipping = 0
    shipping = 5.99 if subtotal < 25
    shipping = 3.99 if subtotal >= 25 && subtotal < 50
    total = subtotal + subtotal * 0.07 + shipping
    label = { to: order[:address], weight: order[:items].map { |i| i[:quantity] }.sum, carrier: 'USPS' }
    label[:carrier] = 'UPS' if total > 100
    { order_id: "ORD-#{Time.now.to_i}", payment_id: "PAY-#{rand(1_000_000)}",
      payment_status: total > 1000 ? 'pending_review' : 'approved', total: total.round(2), shipping_label: label, email: { to: order[:customer][:email], subject: 'Order Confirmation', body: format('Your order total is $%.2f', total) } }
  end
end
