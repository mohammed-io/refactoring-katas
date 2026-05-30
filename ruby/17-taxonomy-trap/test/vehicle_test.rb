# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/vehicle'

class VehicleTest < Minitest::Test
  def test_car_daily_rate
    v = Car.new('Toyota', 'Camry', 2020)
    assert_equal 40, v.daily_rate
  end

  def test_truck_daily_rate
    v = Truck.new('Ford', 'F-150', 2020)
    assert_equal 80, v.daily_rate
  end

  def test_electric_car_fuel_cost
    v = ElectricCar.new('Tesla', 'Model 3', 2020)
    assert_equal 0, v.fuel_cost(5)
  end

  def test_diesel_car_fuel_cost
    v = DieselCar.new('VW', 'Jetta', 2020)
    assert_equal 25, v.fuel_cost(5)
  end

  def test_electric_truck_fuel_cost
    v = ElectricTruck.new('Rivian', 'R1T', 2020)
    assert_equal 0, v.fuel_cost(5)
  end

  def test_diesel_truck_fuel_cost
    v = DieselTruck.new('Ford', 'F-250', 2020)
    assert_equal 75, v.fuel_cost(5)
  end

  def test_car_stores_brand
    v = Car.new('Toyota', 'Camry', 2020)
    assert_equal 'Toyota', v.brand
  end

  def test_truck_stores_model
    v = Truck.new('Ford', 'F-150', 2020)
    assert_equal 'F-150', v.model
  end

  def test_diesel_car_rental_total_combines_rate_fuel_insurance_and_gps
    v = DieselCar.new('VW', 'Jetta', 2020)
    assert_equal 195, v.rental_total(3, true)
  end

  def test_electric_truck_rental_total_combines_truck_rate_and_insurance
    v = ElectricTruck.new('Rivian', 'R1T', 2020)
    assert_equal 200, v.rental_total(2, false)
  end

  def test_diesel_truck_rental_total_includes_higher_fuel_cost
    v = DieselTruck.new('Ford', 'F-250', 2020)
    assert_equal 246, v.rental_total(2, true)
  end
end
