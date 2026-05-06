# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/payroll_calculator'

class PayrollCalculatorTest < Minitest::Test
  def setup
    @calculator = PayrollCalculator.new
  end

  def test_calculates_fulltime_payslip
    result = @calculator.generate_payslips([{ id: 1, name: 'Alice', type: 'fulltime', salary: 60_000 }])
    assert_equal 5000, result[0][:gross]
    assert_equal 1250, result[0][:deductions]
    assert_equal 3750, result[0][:net]
  end

  def test_calculates_fulltime_with_bonus
    result = @calculator.generate_payslips([{ id: 1, name: 'Alice', type: 'fulltime', salary: 60_000, bonus: 12_000 }])
    assert_equal 6000, result[0][:gross]
    assert_equal 4750, result[0][:net]
  end

  def test_calculates_parttime_payslip
    result = @calculator.generate_payslips([{ id: 2, name: 'Bob', type: 'parttime', hours: 80, rate: 25 }])
    assert_equal 2000, result[0][:gross]
    assert_equal 300, result[0][:deductions]
    assert_equal 1700, result[0][:net]
  end

  def test_calculates_contract_payslip
    result = @calculator.generate_payslips([{ id: 3, name: 'Carol', type: 'contract', flat_fee: 5000 }])
    assert_equal 5000, result[0][:gross]
    assert_equal 500, result[0][:deductions]
    assert_equal 4500, result[0][:net]
  end

  def test_handles_multiple_employees
    result = @calculator.generate_payslips([
      { id: 1, name: 'Alice', type: 'fulltime', salary: 60_000 },
      { id: 2, name: 'Bob', type: 'parttime', hours: 80, rate: 25 }
    ])
    assert_equal 2, result.length
    assert_equal 'Alice', result[0][:name]
    assert_equal 'Bob', result[1][:name]
  end

  def test_returns_empty_array_for_empty_input
    assert_equal [], @calculator.generate_payslips([])
  end
end
