class Vehicle:
    def __init__(self, brand, model, year): self.brand=brand; self.model=model; self.year=year
class Car(Vehicle):
    def daily_rate(self): return 40
class Truck(Vehicle):
    def daily_rate(self): return 80
class ElectricCar(Car):
    def fuel_cost(self, days): return 0
class DieselCar(Car):
    def fuel_cost(self, days): return days * 5
class ElectricTruck(Truck):
    def fuel_cost(self, days): return 0
class DieselTruck(Truck):
    def fuel_cost(self, days): return days * 15
