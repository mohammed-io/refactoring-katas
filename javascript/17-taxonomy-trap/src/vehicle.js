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

  insurance_cost(days) {
    return days * 12;
  }
}

class Truck extends Vehicle {
  daily_rate() {
    return 80;
  }

  insurance_cost(days) {
    return days * 20;
  }
}

class ElectricCar extends Car {
  fuel_cost(days) {
    return 0;
  }

  rental_total(days, gps = false) {
    return this.daily_rate() * days + this.fuel_cost(days) + this.insurance_cost(days) + (gps ? 8 * days : 0);
  }
}

class DieselCar extends Car {
  fuel_cost(days) {
    return days * 5;
  }

  rental_total(days, gps = false) {
    return this.daily_rate() * days + this.fuel_cost(days) + this.insurance_cost(days) + (gps ? 8 * days : 0);
  }
}

class ElectricTruck extends Truck {
  fuel_cost(days) {
    return 0;
  }

  rental_total(days, gps = false) {
    return this.daily_rate() * days + this.fuel_cost(days) + this.insurance_cost(days) + (gps ? 8 * days : 0);
  }
}

class DieselTruck extends Truck {
  fuel_cost(days) {
    return days * 15;
  }

  rental_total(days, gps = false) {
    return this.daily_rate() * days + this.fuel_cost(days) + this.insurance_cost(days) + (gps ? 8 * days : 0);
  }
}

export { Vehicle, Car, Truck, ElectricCar, DieselCar, ElectricTruck, DieselTruck };
