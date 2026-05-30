# frozen_string_literal: true

class LegacySystem
  def process_everything(order)
    x = 0
    y = 0
    return { error: 'No items' } if !order[:items] || order[:items].empty?

    order[:items].each do |i|
      x += i[:price] * i[:quantity] if i[:price].positive?
      y += i[:quantity] if i[:quantity].positive?
    end
    return { error: 'Invalid total' } if x <= 0

    d = 0
    d = x * 0.2 if order[:customer] && order[:customer][:type] == 'vip'
    d = x * 0.1 if order[:customer] && order[:customer][:type] == 'member'
    d += 10 if x > 100
    d += x * 0.1 if order[:coupon] == 'SAVE10'
    taxable = x - d
    tax_rate = 0.07
    tax_rate = 0.2 if order[:customer] && order[:customer][:country] == 'EU'
    tax_rate = 0 if order[:customer] && order[:customer][:tax_exempt]
    total = taxable + taxable * tax_rate
    payment = { status: 'approved' }
    payment[:status] = 'manual_review' if total > 5000
    ship = { carrier: 'USPS' }
    ship[:carrier] = 'UPS' if total > 50
    ship[:carrier] = 'FedEx' if order[:shipping] && order[:shipping][:speed] == 'express'
    ship[:weight] = y
    email = { to: order[:customer] ? order[:customer][:email] : '', subject: "Order #{order[:id]}",
              body: format('Total: $%.2f', total) }
    log = "Order processed at #{Time.now}"
    loyalty = { points: (total / 10).floor }
    { order_id: order[:id], total: total.round(2), payment_status: payment[:status], carrier: ship[:carrier],
      email: email, log: log, loyalty_points: loyalty[:points], tax_rate: tax_rate, shipping_weight: ship[:weight] }
  end
end
