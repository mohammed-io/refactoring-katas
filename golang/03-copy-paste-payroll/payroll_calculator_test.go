package kata

import "testing"

func TestGeneratePayslipsCalculatesFulltimePayslip(t *testing.T) {
	pc := NewPayrollCalculator()
	employees := []Employee{{ID: 1, Name: "Alice", Type: "fulltime", Salary: 60000}}
	result := pc.generate_payslips(employees)
	if len(result) != 1 {
		t.Fatalf("expected 1 payslip, got %d", len(result))
	}
	if result[0].Gross != 5000 {
		t.Errorf("expected gross 5000, got %v", result[0].Gross)
	}
	if result[0].Deductions != 1250 {
		t.Errorf("expected deductions 1250, got %v", result[0].Deductions)
	}
	if result[0].Net != 3750 {
		t.Errorf("expected net 3750, got %v", result[0].Net)
	}
}

func TestGeneratePayslipsCalculatesFulltimeWithBonus(t *testing.T) {
	pc := NewPayrollCalculator()
	employees := []Employee{{ID: 1, Name: "Alice", Type: "fulltime", Salary: 60000, Bonus: 12000}}
	result := pc.generate_payslips(employees)
	if len(result) != 1 {
		t.Fatalf("expected 1 payslip, got %d", len(result))
	}
	if result[0].Gross != 6000 {
		t.Errorf("expected gross 6000, got %v", result[0].Gross)
	}
	if result[0].Net != 4750 {
		t.Errorf("expected net 4750, got %v", result[0].Net)
	}
}

func TestGeneratePayslipsCalculatesParttimePayslip(t *testing.T) {
	pc := NewPayrollCalculator()
	employees := []Employee{{ID: 2, Name: "Bob", Type: "parttime", Hours: 80, Rate: 25}}
	result := pc.generate_payslips(employees)
	if len(result) != 1 {
		t.Fatalf("expected 1 payslip, got %d", len(result))
	}
	if result[0].Gross != 2000 {
		t.Errorf("expected gross 2000, got %v", result[0].Gross)
	}
	if result[0].Deductions != 300 {
		t.Errorf("expected deductions 300, got %v", result[0].Deductions)
	}
	if result[0].Net != 1700 {
		t.Errorf("expected net 1700, got %v", result[0].Net)
	}
}

func TestGeneratePayslipsCalculatesContractPayslip(t *testing.T) {
	pc := NewPayrollCalculator()
	employees := []Employee{{ID: 3, Name: "Carol", Type: "contract", FlatFee: 5000}}
	result := pc.generate_payslips(employees)
	if len(result) != 1 {
		t.Fatalf("expected 1 payslip, got %d", len(result))
	}
	if result[0].Gross != 5000 {
		t.Errorf("expected gross 5000, got %v", result[0].Gross)
	}
	if result[0].Deductions != 500 {
		t.Errorf("expected deductions 500, got %v", result[0].Deductions)
	}
	if result[0].Net != 4500 {
		t.Errorf("expected net 4500, got %v", result[0].Net)
	}
}

func TestGeneratePayslipsHandlesMultipleEmployees(t *testing.T) {
	pc := NewPayrollCalculator()
	employees := []Employee{
		{ID: 1, Name: "Alice", Type: "fulltime", Salary: 60000},
		{ID: 2, Name: "Bob", Type: "parttime", Hours: 80, Rate: 25},
	}
	result := pc.generate_payslips(employees)
	if len(result) != 2 {
		t.Fatalf("expected 2 payslips, got %d", len(result))
	}
	if result[0].Name != "Alice" {
		t.Errorf("expected first name Alice, got %q", result[0].Name)
	}
	if result[1].Name != "Bob" {
		t.Errorf("expected second name Bob, got %q", result[1].Name)
	}
}

func TestGeneratePayslipsReturnsEmptyForEmptyInput(t *testing.T) {
	pc := NewPayrollCalculator()
	result := pc.generate_payslips([]Employee{})
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d payslips", len(result))
	}
}
