# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/loan_approver'

class LoanApproverTest < Minitest::Test
  def setup
    @approver = LoanApprover.new
  end

  def test_allows_normal_package
    result = @approver.can_deliver(weight: 10, hazardous: false, weekend: false)
    assert_equal true, result[:allowed]
    assert_nil result[:warning]
  end

  def test_rejects_overweight
    result = @approver.can_deliver(weight: 60, hazardous: false, weekend: false)
    assert_equal false, result[:allowed]
    assert_equal 'Weight exceeded', result[:warning]
  end

  def test_rejects_hazardous
    result = @approver.can_deliver(weight: 10, hazardous: true, weekend: false)
    assert_equal false, result[:allowed]
    assert_equal 'Hazardous material', result[:warning]
  end

  def test_rejects_weekend
    result = @approver.can_deliver(weight: 10, hazardous: false, weekend: true)
    assert_equal false, result[:allowed]
    assert_equal 'No weekend delivery', result[:warning]
  end

  def test_rejects_extreme_temperature
    result = @approver.can_deliver(weight: 10, hazardous: false, weekend: false, temperature_required: 50)
    assert_equal false, result[:allowed]
    assert_equal 'Temperature out of range', result[:warning]
  end

  def test_allows_valid_temperature
    result = @approver.can_deliver(weight: 10, hazardous: false, weekend: false, temperature_required: 20)
    assert_equal true, result[:allowed]
  end

  def test_allows_remote_small_package
    result = @approver.can_deliver(weight: 15, hazardous: false, weekend: false, remote_area: true)
    assert_equal true, result[:allowed]
    assert_equal 'Remote surcharge applies', result[:warning]
  end

  def test_rejects_remote_heavy_package
    result = @approver.can_deliver(weight: 25, hazardous: false, weekend: false, remote_area: true)
    assert_equal false, result[:allowed]
    assert_equal 'Too heavy for remote', result[:warning]
  end

  def test_rejects_null_package
    result = @approver.can_deliver(nil)
    assert_equal false, result[:allowed]
    assert_equal 'No package', result[:warning]
  end

  def test_rejects_missing_weight
    result = @approver.can_deliver(hazardous: false)
    assert_equal false, result[:allowed]
    assert_equal 'No weight specified', result[:warning]
  end

  def test_weight_50_allowed_at_boundary
    result = @approver.can_deliver(weight: 50, hazardous: false, weekend: false)
    assert_equal true, result[:allowed]
    assert_nil result[:warning]
  end

  def test_temperature_40_allowed_at_boundary
    result = @approver.can_deliver(weight: 10, hazardous: false, weekend: false, temperature_required: 40)
    assert_equal true, result[:allowed]
  end

  def test_temperature_minus_20_allowed_at_boundary
    result = @approver.can_deliver(weight: 10, hazardous: false, weekend: false, temperature_required: -20)
    assert_equal true, result[:allowed]
  end

  def test_remote_weight_20_allowed_at_boundary
    result = @approver.can_deliver(weight: 20, hazardous: false, weekend: false, remote_area: true)
    assert_equal true, result[:allowed]
    assert_equal 'Remote surcharge applies', result[:warning]
  end
end
