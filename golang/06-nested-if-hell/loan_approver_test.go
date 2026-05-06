package kata

import "testing"

func TestCanDeliverAllowsNormalPackage(t *testing.T) {
	la := NewLoanApprover()
	pkg := &Package{Weight: 10, HasWeight: true, Hazardous: false, Weekend: false}
	result := la.can_deliver(pkg)
	if !result.Allowed {
		t.Errorf("expected allowed true, got false")
	}
	if result.Warning != "" {
		t.Errorf("expected no warning, got %q", result.Warning)
	}
}

func TestCanDeliverRejectsOverweight(t *testing.T) {
	la := NewLoanApprover()
	pkg := &Package{Weight: 60, HasWeight: true, Hazardous: false, Weekend: false}
	result := la.can_deliver(pkg)
	if result.Allowed {
		t.Errorf("expected allowed false, got true")
	}
	if result.Warning != "Weight exceeded" {
		t.Errorf("expected warning 'Weight exceeded', got %q", result.Warning)
	}
}

func TestCanDeliverRejectsHazardous(t *testing.T) {
	la := NewLoanApprover()
	pkg := &Package{Weight: 10, HasWeight: true, Hazardous: true, Weekend: false}
	result := la.can_deliver(pkg)
	if result.Allowed {
		t.Errorf("expected allowed false, got true")
	}
	if result.Warning != "Hazardous material" {
		t.Errorf("expected warning 'Hazardous material', got %q", result.Warning)
	}
}

func TestCanDeliverRejectsWeekend(t *testing.T) {
	la := NewLoanApprover()
	pkg := &Package{Weight: 10, HasWeight: true, Hazardous: false, Weekend: true}
	result := la.can_deliver(pkg)
	if result.Allowed {
		t.Errorf("expected allowed false, got true")
	}
	if result.Warning != "No weekend delivery" {
		t.Errorf("expected warning 'No weekend delivery', got %q", result.Warning)
	}
}

func TestCanDeliverRejectsExtremeTemperature(t *testing.T) {
	la := NewLoanApprover()
	temp50 := 50
	pkg := &Package{Weight: 10, HasWeight: true, Hazardous: false, Weekend: false, TemperatureRequired: &temp50}
	result := la.can_deliver(pkg)
	if result.Allowed {
		t.Errorf("expected allowed false, got true")
	}
	if result.Warning != "Temperature out of range" {
		t.Errorf("expected warning 'Temperature out of range', got %q", result.Warning)
	}
}

func TestCanDeliverAllowsValidTemperature(t *testing.T) {
	la := NewLoanApprover()
	temp20 := 20
	pkg := &Package{Weight: 10, HasWeight: true, Hazardous: false, Weekend: false, TemperatureRequired: &temp20}
	result := la.can_deliver(pkg)
	if !result.Allowed {
		t.Errorf("expected allowed true, got false")
	}
}

func TestCanDeliverAllowsRemoteSmallPackage(t *testing.T) {
	la := NewLoanApprover()
	pkg := &Package{Weight: 15, HasWeight: true, Hazardous: false, Weekend: false, RemoteArea: true}
	result := la.can_deliver(pkg)
	if !result.Allowed {
		t.Errorf("expected allowed true, got false")
	}
	if result.Warning != "Remote surcharge applies" {
		t.Errorf("expected warning 'Remote surcharge applies', got %q", result.Warning)
	}
}

func TestCanDeliverRejectsRemoteHeavyPackage(t *testing.T) {
	la := NewLoanApprover()
	pkg := &Package{Weight: 25, HasWeight: true, Hazardous: false, Weekend: false, RemoteArea: true}
	result := la.can_deliver(pkg)
	if result.Allowed {
		t.Errorf("expected allowed false, got true")
	}
	if result.Warning != "Too heavy for remote" {
		t.Errorf("expected warning 'Too heavy for remote', got %q", result.Warning)
	}
}

func TestCanDeliverRejectsNullPackage(t *testing.T) {
	la := NewLoanApprover()
	result := la.can_deliver(nil)
	if result.Allowed {
		t.Errorf("expected allowed false, got true")
	}
	if result.Warning != "No package" {
		t.Errorf("expected warning 'No package', got %q", result.Warning)
	}
}

func TestCanDeliverRejectsMissingWeight(t *testing.T) {
	la := NewLoanApprover()
	pkg := &Package{Hazardous: false}
	result := la.can_deliver(pkg)
	if result.Allowed {
		t.Errorf("expected allowed false, got true")
	}
	if result.Warning != "No weight specified" {
		t.Errorf("expected warning 'No weight specified', got %q", result.Warning)
	}
}
