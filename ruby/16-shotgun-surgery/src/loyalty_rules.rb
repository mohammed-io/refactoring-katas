# frozen_string_literal: true

class LoyaltyRules
  def get_discount_for_tier(tier)
    case tier
    when 'bronze'
      0.05
    when 'silver'
      0.1
    when 'gold'
      0.15
    when 'platinum'
      0.2
    else
      0
    end
  end

  def get_label_for_tier(tier)
    case tier
    when 'bronze'
      'Bronze Member'
    when 'silver'
      'Silver Member'
    when 'gold'
      'Gold Member'
    when 'platinum'
      'Platinum Member'
    else
      'Standard'
    end
  end

  def get_threshold_for_tier(tier)
    case tier
    when 'bronze'
      100
    when 'silver'
      500
    when 'gold'
      2000
    when 'platinum'
      10_000
    else
      0
    end
  end

  def get_color_for_tier(tier)
    case tier
    when 'bronze'
      '#CD7F32'
    when 'silver'
      '#C0C0C0'
    when 'gold'
      '#FFD700'
    when 'platinum'
      '#E5E4E2'
    else
      '#000000'
    end
  end

  def calculate_total(spending, tier)
    spending * (1 - get_discount_for_tier(tier))
  end
end
