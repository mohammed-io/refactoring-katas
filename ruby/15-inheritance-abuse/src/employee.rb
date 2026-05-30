# frozen_string_literal: true

class Employee
  def initialize(name, salary)
    @name = name
    @salary = salary
  end

  def calculate_bonus
    @salary * 0.02
  end

  def calculate_total_reward(performance, years)
    x = calculate_bonus
    x += @salary * 0.01 if performance == 'high'
    x += 500 if years >= 5
    x
  end
end

class Manager < Employee
  def calculate_bonus
    @salary * 0.05
  end

  def calculate_total_reward(performance, years)
    x = calculate_bonus
    x += @salary * 0.02 if performance == 'high'
    x += 1000 if years >= 5
    x
  end
end

class SeniorManager < Manager
  def calculate_bonus
    base = @salary * 0.05
    base > 10_000 ? 10_000 : base
  end

  def calculate_total_reward(performance, years)
    x = calculate_bonus
    x += @salary * 0.02 if performance == 'high'
    x += 1500 if years >= 5
    x
  end
end

class Director < SeniorManager
  def calculate_bonus
    base = @salary * 0.05
    base > 20_000 ? 20_000 : base
  end

  def calculate_total_reward(performance, years)
    x = calculate_bonus
    x += @salary * 0.03 if performance == 'high'
    x += 2500 if years >= 5
    x
  end
end
