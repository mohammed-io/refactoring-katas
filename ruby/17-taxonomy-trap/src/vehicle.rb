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

  def insurance_cost(days)
    days * 12
  end
end

class Truck < Vehicle
  def daily_rate
    80
  end

  def insurance_cost(days)
    days * 20
  end
end

class ElectricCar < Car
  def fuel_cost(_days)
    0
  end

  def rental_total(days, gps = false)
    (daily_rate * days) + fuel_cost(days) + insurance_cost(days) + (gps ? 8 * days : 0)
  end
end

class DieselCar < Car
  def fuel_cost(days)
    days * 5
  end

  def rental_total(days, gps = false)
    (daily_rate * days) + fuel_cost(days) + insurance_cost(days) + (gps ? 8 * days : 0)
  end
end

class ElectricTruck < Truck
  def fuel_cost(_days)
    0
  end

  def rental_total(days, gps = false)
    (daily_rate * days) + fuel_cost(days) + insurance_cost(days) + (gps ? 8 * days : 0)
  end
end

class DieselTruck < Truck
  def fuel_cost(days)
    days * 15
  end

  def rental_total(days, gps = false)
    (daily_rate * days) + fuel_cost(days) + insurance_cost(days) + (gps ? 8 * days : 0)
  end
end
