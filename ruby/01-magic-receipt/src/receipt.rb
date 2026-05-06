# frozen_string_literal: true

class Receipt
  def calculate_total(items, customer_type = nil)
    total = 0
    items.each { |price| total += price }
    discount = 0
    if customer_type == 'member'
      discount = total * 0.05
    elsif customer_type == 'vip'
      discount = total * 0.15
    end
    discount += 5 if total > 50
    final = (total - discount) * 1.08
    final -= 2 if customer_type == 'vip'
    final.round(2)
  end
end
