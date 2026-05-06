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
end
