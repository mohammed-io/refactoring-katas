class Employee {
  constructor(name, salary) {
    this.name = name;
    this.salary = salary;
  }

  calculate_bonus() {
    return this.salary * 0.02;
  }

  calculate_total_reward(performance, years) {
    let x = this.calculate_bonus();
    if (performance === 'high') x += this.salary * 0.01;
    if (years >= 5) x += 500;
    return x;
  }
}

class Manager extends Employee {
  calculate_bonus() {
    return this.salary * 0.05;
  }

  calculate_total_reward(performance, years) {
    let x = this.calculate_bonus();
    if (performance === 'high') x += this.salary * 0.02;
    if (years >= 5) x += 1000;
    return x;
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

  calculate_total_reward(performance, years) {
    let x = this.calculate_bonus();
    if (performance === 'high') x += this.salary * 0.02;
    if (years >= 5) x += 1500;
    return x;
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

  calculate_total_reward(performance, years) {
    let x = this.calculate_bonus();
    if (performance === 'high') x += this.salary * 0.03;
    if (years >= 5) x += 2500;
    return x;
  }
}

export { Employee, Manager, SeniorManager, Director };
