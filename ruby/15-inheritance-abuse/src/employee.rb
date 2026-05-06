# frozen_string_literal: true

class Employee
  def initialize(name, salary)
    @name = name
    @salary = salary
  end

  def calculate_bonus
    @salary * 0.02
  end
end

class Manager < Employee
  def calculate_bonus
    @salary * 0.05
  end
end

class SeniorManager < Manager
  def calculate_bonus
    base = @salary * 0.05
    base > 10_000 ? 10_000 : base
  end
end

class Director < SeniorManager
  def calculate_bonus
    base = @salary * 0.05
    base > 20_000 ? 20_000 : base
  end
end
