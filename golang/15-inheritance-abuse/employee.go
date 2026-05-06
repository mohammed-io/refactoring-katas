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
