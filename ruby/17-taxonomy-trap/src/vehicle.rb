# frozen_string_literal: true

class Vehicle
  attr_reader :brand, :model, :year

  def initialize(brand, model, year)
    @brand = brand
    @model = model
    @year = year
  end
end

class Car < Vehicle
  def daily_rate
    40
  end
end

class Truck < Vehicle
  def daily_rate
    80
  end
end

class ElectricCar < Car
  def fuel_cost(_days)
    0
  end
end

class DieselCar < Car
  def fuel_cost(days)
    days * 5
  end
end

class ElectricTruck < Truck
  def fuel_cost(_days)
    0
  end
end

class DieselTruck < Truck
  def fuel_cost(days)
    days * 15
  end
end
