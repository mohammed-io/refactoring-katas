package kata

type Vehicle struct {
	Brand  string
	Model  string
	Year   int
	Type   string
	Fuel   string
}

func NewVehicle(brand, model, vehicleType, fuel string, year int) *Vehicle {
	return &Vehicle{Brand: brand, Model: model, Year: year, Type: vehicleType, Fuel: fuel}
}

func (v *Vehicle) daily_rate() int {
	if v.Type == "truck" {
		return 80
	}
	return 40
}

func (v *Vehicle) insurance_cost(days int) int {
	if v.Type == "truck" {
		return days * 20
	}
	return days * 12
}

func (v *Vehicle) fuel_cost(days int) int {
	if v.Fuel == "electric" {
		return 0
	}
	if v.Type == "truck" {
		return days * 15
	}
	return days * 5
}

func (v *Vehicle) rental_total(days int, gps bool) int {
	x := v.daily_rate()*days + v.fuel_cost(days) + v.insurance_cost(days)
	if gps {
		x += 8 * days
	}
	return x
}
