package kata

import "testing"

func TestCalculateBonusEmployeeGets2Percent(t *testing.T) {
	emp := NewEmployee("Alice", 50000, "employee")
	if emp.calculate_bonus() != 1000 {
		t.Errorf("expected bonus 1000, got %v", emp.calculate_bonus())
	}
}

func TestCalculateBonusManagerGets5Percent(t *testing.T) {
	mgr := NewEmployee("Bob", 80000, "manager")
	if mgr.calculate_bonus() != 4000 {
		t.Errorf("expected bonus 4000, got %v", mgr.calculate_bonus())
	}
}

func TestCalculateBonusSeniorManagerCappedAt10000(t *testing.T) {
	sm := NewEmployee("Carol", 300000, "senior_manager")
	if sm.calculate_bonus() != 10000 {
		t.Errorf("expected bonus 10000 (capped), got %v", sm.calculate_bonus())
	}
}

func TestCalculateBonusSeniorManagerUnderCap(t *testing.T) {
	sm := NewEmployee("Carol", 100000, "senior_manager")
	if sm.calculate_bonus() != 5000 {
		t.Errorf("expected bonus 5000 (under cap), got %v", sm.calculate_bonus())
	}
}

func TestCalculateBonusDirectorCappedAt20000(t *testing.T) {
	dir := NewEmployee("Dave", 600000, "director")
	if dir.calculate_bonus() != 20000 {
		t.Errorf("expected bonus 20000 (capped), got %v", dir.calculate_bonus())
	}
}

func TestCalculateBonusDirectorUnderCap(t *testing.T) {
	dir := NewEmployee("Dave", 200000, "director")
	if dir.calculate_bonus() != 10000 {
		t.Errorf("expected bonus 10000 (under cap), got %v", dir.calculate_bonus())
	}
}

func TestCalculateBonusManagerRespectsCapFromSeniorManager(t *testing.T) {
	mgr := NewEmployee("Eve", 300000, "manager")
	if mgr.calculate_bonus() != 15000 {
		t.Errorf("expected bonus 15000 (no cap for Manager), got %v", mgr.calculate_bonus())
	}
}

func TestCalculateTotalRewardEmployeeAddsHighPerformanceAndTenure(t *testing.T) {
	emp := NewEmployee("Alice", 50000, "employee")
	if got := emp.calculate_total_reward("high", 5); got != 2000 {
		t.Errorf("expected reward 2000, got %v", got)
	}
}

func TestCalculateTotalRewardSeniorManagerUsesCappedBonusAndTenureRule(t *testing.T) {
	sm := NewEmployee("Carol", 300000, "senior_manager")
	if got := sm.calculate_total_reward("normal", 7); got != 11500 {
		t.Errorf("expected reward 11500, got %v", got)
	}
}

func TestCalculateTotalRewardDirectorUsesDirectorPerformanceRule(t *testing.T) {
	dir := NewEmployee("Dana", 200000, "director")
	if got := dir.calculate_total_reward("high", 3); got != 16000 {
		t.Errorf("expected reward 16000, got %v", got)
	}
}
