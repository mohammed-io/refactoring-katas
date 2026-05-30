# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/fraud_detector'

class FraudDetectorTest < Minitest::Test
  def setup
    @detector = FraudDetector.new
    @now = (Time.now.to_f * 1000).to_i
  end

  def test_low_risk_small_transaction
    result = @detector.detect(
      amount: 10, timestamp: @now, history: [], merchant: 'grocery', country: 'US', card_country: 'US'
    )
    assert_equal 'low', result[:rating]
    assert_equal 1, result[:level]
  end

  def test_medium_risk_for_large_amount
    result = @detector.detect(
      amount: 1100, timestamp: @now, history: [], merchant: 'grocery', country: 'US', card_country: 'US'
    )
    assert_equal 'medium', result[:rating]
    assert_equal 2, result[:level]
  end

  def test_gambling_merchant_is_medium_risk
    result = @detector.detect(
      amount: 100, timestamp: @now, history: [], merchant: 'gambling', country: 'US', card_country: 'US'
    )
    assert_equal 'medium', result[:rating]
    assert_equal 2, result[:level]
  end

  def test_cross_border_alone_is_low_risk
    result = @detector.detect(
      amount: 1000, timestamp: @now, history: [], merchant: 'grocery', country: 'FR', card_country: 'US'
    )
    assert_equal 'low', result[:rating]
    assert_equal 1, result[:level]
  end

  def test_critical_risk_for_late_night_crypto
    t = Time.utc(2024, 1, 1, 2, 0, 0).to_i * 1000
    result = @detector.detect(
      amount: 600, timestamp: t, history: [], merchant: 'crypto', country: 'CN', card_country: 'US'
    )
    assert_equal 'high', result[:rating]
    assert_equal 4, result[:level]
  end

  def test_velocity_increases_risk
    hist = (1..4).map { |i| { amount: 10, timestamp: @now - i * 1000 } }
    result = @detector.detect(
      amount: 50, timestamp: @now, history: hist, merchant: 'grocery', country: 'US', card_country: 'US'
    )
    assert_equal 2, result[:level]
  end

  def test_volume_spikes_alone_stay_low_risk
    hist = [
      { amount: 200, timestamp: @now - 10_000 },
      { amount: 200, timestamp: @now - 20_000 },
      { amount: 200, timestamp: @now - 30_000 }
    ]
    result = @detector.detect(
      amount: 50, timestamp: @now, history: hist, merchant: 'grocery', country: 'US', card_country: 'US'
    )
    assert_equal 1, result[:level]
  end

  def test_includes_score
    result = @detector.detect(
      amount: 10, timestamp: @now, history: [], merchant: 'grocery', country: 'US', card_country: 'US'
    )
    assert_kind_of Integer, result[:score]
  end
end
