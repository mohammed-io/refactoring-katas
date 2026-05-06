# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/loyalty_rules'

class LoyaltyRulesTest < Minitest::Test
  def setup
    @rules = LoyaltyRules.new
  end

  def test_bronze_discount
    assert_equal 0.05, @rules.get_discount_for_tier('bronze')
  end

  def test_silver_discount
    assert_equal 0.1, @rules.get_discount_for_tier('silver')
  end

  def test_gold_discount
    assert_equal 0.15, @rules.get_discount_for_tier('gold')
  end

  def test_platinum_discount
    assert_equal 0.2, @rules.get_discount_for_tier('platinum')
  end

  def test_unknown_tier_discount
    assert_equal 0, @rules.get_discount_for_tier('unknown')
  end

  def test_bronze_label
    assert_equal 'Bronze Member', @rules.get_label_for_tier('bronze')
  end

  def test_silver_label
    assert_equal 'Silver Member', @rules.get_label_for_tier('silver')
  end

  def test_gold_label
    assert_equal 'Gold Member', @rules.get_label_for_tier('gold')
  end

  def test_platinum_label
    assert_equal 'Platinum Member', @rules.get_label_for_tier('platinum')
  end

  def test_unknown_tier_label
    assert_equal 'Standard', @rules.get_label_for_tier('unknown')
  end

  def test_bronze_threshold
    assert_equal 100, @rules.get_threshold_for_tier('bronze')
  end

  def test_silver_threshold
    assert_equal 500, @rules.get_threshold_for_tier('silver')
  end

  def test_gold_threshold
    assert_equal 2000, @rules.get_threshold_for_tier('gold')
  end

  def test_platinum_threshold
    assert_equal 10_000, @rules.get_threshold_for_tier('platinum')
  end

  def test_unknown_tier_threshold
    assert_equal 0, @rules.get_threshold_for_tier('unknown')
  end

  def test_bronze_color
    assert_equal '#CD7F32', @rules.get_color_for_tier('bronze')
  end

  def test_silver_color
    assert_equal '#C0C0C0', @rules.get_color_for_tier('silver')
  end

  def test_gold_color
    assert_equal '#FFD700', @rules.get_color_for_tier('gold')
  end

  def test_platinum_color
    assert_equal '#E5E4E2', @rules.get_color_for_tier('platinum')
  end

  def test_unknown_tier_color
    assert_equal '#000000', @rules.get_color_for_tier('unknown')
  end

  def test_calculates_total_for_bronze
    assert_equal 95, @rules.calculate_total(100, 'bronze')
  end

  def test_calculates_total_for_platinum
    assert_equal 80, @rules.calculate_total(100, 'platinum')
  end

  def test_calculates_total_for_unknown_tier
    assert_equal 100, @rules.calculate_total(100, 'unknown')
  end
end
