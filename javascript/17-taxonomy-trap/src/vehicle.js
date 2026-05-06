class Vehicle {
  constructor(brand, model, year) {
    this.brand = brand;
    this.model = model;
    this.year = year;
  }
}

class Car extends Vehicle {
  daily_rate() {
    return 40;
  }
}

class Truck extends Vehicle {
  daily_rate() {
    return 80;
  }
}

class ElectricCar extends Car {
  fuel_cost(days) {
    return 0;
  }
}

class DieselCar extends Car {
  fuel_cost(days) {
    return days * 5;
  }
}

class ElectricTruck extends Truck {
  fuel_cost(days) {
    return 0;
  }
}

class DieselTruck extends Truck {
  fuel_cost(days) {
    return days * 15;
  }
}

export { Vehicle, Car, Truck, ElectricCar, DieselCar, ElectricTruck, DieselTruck };
