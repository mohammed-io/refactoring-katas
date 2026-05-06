# frozen_string_literal: true

class LegacySystem
  def process_everything(order)
    x = 0
    return { error: 'No items' } if !order[:items] || order[:items].empty?

    order[:items].each do |i|
      x += i[:price] * i[:quantity] if i[:price].positive?
    end
    d = 0
    d = x * 0.2 if order[:customer] && order[:customer][:type] == 'vip'
    d = x * 0.1 if order[:customer] && order[:customer][:type] == 'member'
    d += 10 if x > 100
    total = x - d + (x - d) * 0.07
    payment = { status: 'approved' }
    payment[:status] = 'manual_review' if total > 5000
    ship = { carrier: 'USPS' }
    ship[:carrier] = 'UPS' if total > 50
    email = { to: order[:customer] ? order[:customer][:email] : '', subject: "Order #{order[:id]}",
              body: format('Total: $%.2f', total) }
    log = "Order processed at #{Time.now}"
    { order_id: order[:id], total: total.round(2), payment_status: payment[:status], carrier: ship[:carrier],
      email: email, log: log }
  end
end
