package kata

type Employee struct {
	Name   string
	Salary float64
	Type   string
}

func NewEmployee(name string, salary float64, employeeType string) *Employee {
	return &Employee{Name: name, Salary: salary, Type: employeeType}
}

func (e *Employee) calculate_bonus() float64 {
	if e.Type == "director" {
		b := e.Salary * 0.05
		if b > 20000 {
			return 20000
		}
		return b
	}
	if e.Type == "senior_manager" {
		b := e.Salary * 0.05
		if b > 10000 {
			return 10000
		}
		return b
	}
	if e.Type == "manager" {
		return e.Salary * 0.05
	}
	return e.Salary * 0.02
}

func (e *Employee) calculate_total_reward(performance string, years int) float64 {
	x := e.calculate_bonus()
	if e.Type == "director" {
		if performance == "high" {
			x += e.Salary * 0.03
		}
		if years >= 5 {
			x += 2500
		}
		return x
	}
	if e.Type == "senior_manager" {
		if performance == "high" {
			x += e.Salary * 0.02
		}
		if years >= 5 {
			x += 1500
		}
		return x
	}
	if e.Type == "manager" {
		if performance == "high" {
			x += e.Salary * 0.02
		}
		if years >= 5 {
			x += 1000
		}
		return x
	}
	if performance == "high" {
		x += e.Salary * 0.01
	}
	if years >= 5 {
		x += 500
	}
	return x
}
