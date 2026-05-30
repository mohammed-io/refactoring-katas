package kata

import "testing"

func TestVehicleDailyRateCar(t *testing.T) {
	v := NewVehicle("Toyota", "Camry", "car", "diesel", 2020)
	if got := v.daily_rate(); got != 40 {
		t.Errorf("Vehicle.daily_rate() = %v, want 40", got)
	}
}

func TestVehicleDailyRateTruck(t *testing.T) {
	v := NewVehicle("Ford", "F-150", "truck", "diesel", 2020)
	if got := v.daily_rate(); got != 80 {
		t.Errorf("Vehicle.daily_rate() = %v, want 80", got)
	}
}

func TestVehicleFuelCostElectricCar(t *testing.T) {
	v := NewVehicle("Tesla", "Model 3", "car", "electric", 2020)
	if got := v.fuel_cost(5); got != 0 {
		t.Errorf("Vehicle.fuel_cost(5) = %v, want 0", got)
	}
}

func TestVehicleFuelCostDieselCar(t *testing.T) {
	v := NewVehicle("VW", "Jetta", "car", "diesel", 2020)
	if got := v.fuel_cost(5); got != 25 {
		t.Errorf("Vehicle.fuel_cost(5) = %v, want 25", got)
	}
}

func TestVehicleFuelCostElectricTruck(t *testing.T) {
	v := NewVehicle("Rivian", "R1T", "truck", "electric", 2020)
	if got := v.fuel_cost(5); got != 0 {
		t.Errorf("Vehicle.fuel_cost(5) = %v, want 0", got)
	}
}

func TestVehicleFuelCostDieselTruck(t *testing.T) {
	v := NewVehicle("Ford", "F-250", "truck", "diesel", 2020)
	if got := v.fuel_cost(5); got != 75 {
		t.Errorf("Vehicle.fuel_cost(5) = %v, want 75", got)
	}
}

func TestVehicleStoresBrand(t *testing.T) {
	v := NewVehicle("Toyota", "Camry", "car", "diesel", 2020)
	if v.Brand != "Toyota" {
		t.Errorf("Vehicle.Brand = %q, want 'Toyota'", v.Brand)
	}
}

func TestVehicleStoresModel(t *testing.T) {
	v := NewVehicle("Ford", "F-150", "truck", "diesel", 2020)
	if v.Model != "F-150" {
		t.Errorf("Vehicle.Model = %q, want 'F-150'", v.Model)
	}
}

func TestVehicleDieselCarRentalTotalCombinesRateFuelInsuranceAndGPS(t *testing.T) {
	v := NewVehicle("VW", "Jetta", "car", "diesel", 2020)
	if got := v.rental_total(3, true); got != 195 {
		t.Errorf("Vehicle.rental_total() = %v, want 195", got)
	}
}

func TestVehicleElectricTruckRentalTotalCombinesTruckRateAndInsurance(t *testing.T) {
	v := NewVehicle("Rivian", "R1T", "truck", "electric", 2020)
	if got := v.rental_total(2, false); got != 200 {
		t.Errorf("Vehicle.rental_total() = %v, want 200", got)
	}
}

func TestVehicleDieselTruckRentalTotalIncludesHigherFuelCost(t *testing.T) {
	v := NewVehicle("Ford", "F-250", "truck", "diesel", 2020)
	if got := v.rental_total(2, true); got != 246 {
		t.Errorf("Vehicle.rental_total() = %v, want 246", got)
	}
}
