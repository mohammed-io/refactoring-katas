class Employee:
    def __init__(self, name, salary): self.name=name; self.salary=salary
    def calculate_bonus(self): return self.salary * 0.02
    def calculate_total_reward(self, performance, years):
        x = self.calculate_bonus()
        if performance == "high": x += self.salary * 0.01
        if years >= 5: x += 500
        return x
class Manager(Employee):
    def calculate_bonus(self): return self.salary * 0.05
    def calculate_total_reward(self, performance, years):
        x = self.calculate_bonus()
        if performance == "high": x += self.salary * 0.02
        if years >= 5: x += 1000
        return x
class SeniorManager(Manager):
    def calculate_bonus(self):
        base = self.salary * 0.05
        if base > 10000: base = 10000
        return base
    def calculate_total_reward(self, performance, years):
        x = self.calculate_bonus()
        if performance == "high": x += self.salary * 0.02
        if years >= 5: x += 1500
        return x
class Director(SeniorManager):
    def calculate_bonus(self):
        base = self.salary * 0.05
        if base > 20000: base = 20000
        return base
    def calculate_total_reward(self, performance, years):
        x = self.calculate_bonus()
        if performance == "high": x += self.salary * 0.03
        if years >= 5: x += 2500
        return x
