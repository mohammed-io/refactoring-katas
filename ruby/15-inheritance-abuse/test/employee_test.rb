# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/employee'

class EmployeeTest < Minitest::Test
  def test_employee_gets_2_percent_bonus
    emp = Employee.new('Alice', 50_000)
    assert_equal 1000, emp.calculate_bonus
  end

  def test_manager_gets_5_percent_bonus
    mgr = Manager.new('Bob', 80_000)
    assert_equal 4000, mgr.calculate_bonus
  end

  def test_senior_manager_capped_at_10000
    sm = SeniorManager.new('Carol', 300_000)
    assert_equal 10_000, sm.calculate_bonus
  end

  def test_senior_manager_under_cap
    sm = SeniorManager.new('Carol', 100_000)
    assert_equal 5000, sm.calculate_bonus
  end

  def test_director_capped_at_20000
    dir = Director.new('Dave', 600_000)
    assert_equal 20_000, dir.calculate_bonus
  end

  def test_director_under_cap
    dir = Director.new('Dave', 200_000)
    assert_equal 10_000, dir.calculate_bonus
  end

  def test_manager_respects_cap_from_senior_manager
    mgr = Manager.new('Eve', 300_000)
    assert_equal 15_000, mgr.calculate_bonus
  end

  def test_employee_total_reward_adds_high_performance_and_tenure
    emp = Employee.new('Alice', 50_000)
    assert_equal 2000, emp.calculate_total_reward('high', 5)
  end

  def test_senior_manager_total_reward_uses_capped_bonus_and_tenure_rule
    sm = SeniorManager.new('Carol', 300_000)
    assert_equal 11_500, sm.calculate_total_reward('normal', 7)
  end

  def test_director_total_reward_uses_director_performance_rule
    dir = Director.new('Dana', 200_000)
    assert_equal 16_000, dir.calculate_total_reward('high', 3)
  end
end
