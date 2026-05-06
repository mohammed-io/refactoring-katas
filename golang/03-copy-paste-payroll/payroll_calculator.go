package kata

import "math"

type Employee struct {
	ID      int
	Name    string
	Type    string
	Salary  float64
	Bonus   float64
	Hours   float64
	Rate    float64
	FlatFee float64
}
type Payslip struct {
	ID         int
	Name       string
	Type       string
	Gross      float64
	Deductions float64
	Net        float64
}

type PayrollCalculator struct{}

func NewPayrollCalculator() *PayrollCalculator {
	return &PayrollCalculator{}
}

func (pc *PayrollCalculator) generate_payslips(employees []Employee) []Payslip {
	slips := []Payslip{}
	for _, emp := range employees {
		gross := 0.0
		deductions := 0.0
		net := 0.0
		if emp.Type == "fulltime" {
			gross = emp.Salary / 12
			deductions = gross * 0.25
			if emp.Bonus > 0 {
				gross += emp.Bonus / 12
			}
			net = gross - deductions
		} else if emp.Type == "parttime" {
			gross = emp.Hours * emp.Rate
			deductions = gross * 0.15
			net = gross - deductions
		} else if emp.Type == "contract" {
			gross = emp.FlatFee
			deductions = gross * 0.1
			net = gross - deductions
		}
		slips = append(slips, Payslip{emp.ID, emp.Name, emp.Type, math.Round(gross*100) / 100, math.Round(deductions*100) / 100, math.Round(net*100) / 100})
	}
	return slips
}
