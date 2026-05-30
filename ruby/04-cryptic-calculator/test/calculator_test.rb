# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/calculator'

class CalculatorTest < Minitest::Test
  def setup
    @calc = Calculator.new
  end

  def test_normalizes_to_0_100_range
    assert_equal [0, 50, 100], @calc.normalize([10, 20, 30], 0, 100)
  end

  def test_normalizes_to_1_5_range
    assert_equal [1, 3, 5], @calc.normalize([10, 20, 30], 1, 5)
  end

  def test_handles_single_value
    result = @calc.normalize([50], 0, 100)
    assert result[0].nan?
  end

  def test_handles_negative_input_range
    assert_equal [0, 0.5, 1], @calc.normalize([-10, 0, 10], 0, 1)
  end

  def test_handles_same_min_and_max
    result = @calc.normalize([5, 5, 5], 0, 100)
    assert result[0].nan?
    assert result[1].nan?
    assert result[2].nan?
  end

  def test_handles_reversed_output_range
    assert_equal [100, 50, 0], @calc.normalize([10, 20, 30], 100, 0)
  end

  def test_rounds_fractional_results
    assert_equal [0, 33.33, 100], @calc.normalize([2, 3, 5], 0, 100)
  end

  def test_empty_input_returns_empty_result
    assert_equal [], @calc.normalize([], 0, 100)
  end

  def test_preserves_input_order_with_duplicates
    assert_equal [100, 0, 100, 50], @calc.normalize([30, 10, 30, 20], 0, 100)
  end

  def test_normalizes_decimal_output_range
    assert_equal [-1, 0, 1], @calc.normalize([1.5, 2.5, 3.5], -1, 1)
  end

  def test_rounds_negative_fractional_results
    assert_equal [-10, 0, 10], @calc.normalize([2, 5, 8], -10, 10)
  end
end
