class Vehicle:
    def __init__(self, brand, model, year): self.brand=brand; self.model=model; self.year=year
class Car(Vehicle):
    def daily_rate(self): return 40
    def insurance_cost(self, days): return days * 12
class Truck(Vehicle):
    def daily_rate(self): return 80
    def insurance_cost(self, days): return days * 20
class ElectricCar(Car):
    def fuel_cost(self, days): return 0
    def rental_total(self, days, gps=False): return self.daily_rate() * days + self.fuel_cost(days) + self.insurance_cost(days) + (8 * days if gps else 0)
class DieselCar(Car):
    def fuel_cost(self, days): return days * 5
    def rental_total(self, days, gps=False): return self.daily_rate() * days + self.fuel_cost(days) + self.insurance_cost(days) + (8 * days if gps else 0)
class ElectricTruck(Truck):
    def fuel_cost(self, days): return 0
    def rental_total(self, days, gps=False): return self.daily_rate() * days + self.fuel_cost(days) + self.insurance_cost(days) + (8 * days if gps else 0)
class DieselTruck(Truck):
    def fuel_cost(self, days): return days * 15
    def rental_total(self, days, gps=False): return self.daily_rate() * days + self.fuel_cost(days) + self.insurance_cost(days) + (8 * days if gps else 0)
