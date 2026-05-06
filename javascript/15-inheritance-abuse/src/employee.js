class Employee {
  constructor(name, salary) {
    this.name = name;
    this.salary = salary;
  }

  calculate_bonus() {
    return this.salary * 0.02;
  }
}

class Manager extends Employee {
  calculate_bonus() {
    return this.salary * 0.05;
  }
}

class SeniorManager extends Manager {
  calculate_bonus() {
    let base = this.salary * 0.05;
    if (base > 10000) {
      base = 10000;
    }
    return base;
  }
}

class Director extends SeniorManager {
  calculate_bonus() {
    let base = this.salary * 0.05;
    if (base > 20000) {
      base = 20000;
    }
    return base;
  }
}

export { Employee, Manager, SeniorManager, Director };
