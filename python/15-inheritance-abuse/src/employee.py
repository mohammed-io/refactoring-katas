class Employee:
    def __init__(self, name, salary): self.name=name; self.salary=salary
    def calculate_bonus(self): return self.salary * 0.02
class Manager(Employee):
    def calculate_bonus(self): return self.salary * 0.05
class SeniorManager(Manager):
    def calculate_bonus(self):
        base = self.salary * 0.05
        if base > 10000: base = 10000
        return base
class Director(SeniorManager):
    def calculate_bonus(self):
        base = self.salary * 0.05
        if base > 20000: base = 20000
        return base
