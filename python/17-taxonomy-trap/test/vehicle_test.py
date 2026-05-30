from src.vehicle import *

def test_car_daily_rate():
    v = Car("Toyota", "Camry", 2020)
    assert v.daily_rate() == 40

def test_truck_daily_rate():
    v = Truck("Ford", "F-150", 2020)
    assert v.daily_rate() == 80

def test_electric_car_fuel_cost():
    v = ElectricCar("Tesla", "Model 3", 2020)
    assert v.fuel_cost(5) == 0

def test_diesel_car_fuel_cost():
    v = DieselCar("VW", "Jetta", 2020)
    assert v.fuel_cost(5) == 25

def test_electric_truck_fuel_cost():
    v = ElectricTruck("Rivian", "R1T", 2020)
    assert v.fuel_cost(5) == 0

def test_diesel_truck_fuel_cost():
    v = DieselTruck("Ford", "F-250", 2020)
    assert v.fuel_cost(5) == 75

def test_car_stores_brand():
    v = Car("Toyota", "Camry", 2020)
    assert v.brand == "Toyota"

def test_truck_stores_model():
    v = Truck("Ford", "F-150", 2020)
    assert v.model == "F-150"

def test_diesel_car_rental_total_combines_rate_fuel_insurance_and_gps():
    v = DieselCar("VW", "Jetta", 2020)
    assert v.rental_total(3, gps=True) == 195

def test_electric_truck_rental_total_combines_truck_rate_and_insurance():
    v = ElectricTruck("Rivian", "R1T", 2020)
    assert v.rental_total(2, gps=False) == 200

def test_diesel_truck_rental_total_includes_higher_fuel_cost():
    v = DieselTruck("Ford", "F-250", 2020)
    assert v.rental_total(2, gps=True) == 246
