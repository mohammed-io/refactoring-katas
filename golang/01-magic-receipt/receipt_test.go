package kata

import "testing"

func TestCalculateTotalForRegularCustomer(t *testing.T) {
	r := NewReceipt()
	got := r.calculate_total([]float64{10, 20, 30}, "")
	if got != 59.4 {
		t.Errorf("calculate_total([10,20,30], '') = %v, want 59.4", got)
	}
}

func TestCalculateTotalForMemberCustomer(t *testing.T) {
	r := NewReceipt()
	got := r.calculate_total([]float64{10, 20, 30}, "member")
	if got != 56.16 {
		t.Errorf("calculate_total([10,20,30], 'member') = %v, want 56.16", got)
	}
}

func TestCalculateTotalForVipCustomer(t *testing.T) {
	r := NewReceipt()
	got := r.calculate_total([]float64{10, 20, 30}, "vip")
	if got != 47.68 {
		t.Errorf("calculate_total([10,20,30], 'vip') = %v, want 47.68", got)
	}
}

func TestCalculateTotalAppliesBonusDiscountOver50(t *testing.T) {
	r := NewReceipt()
	got := r.calculate_total([]float64{60}, "")
	if got != 59.4 {
		t.Errorf("calculate_total([60], '') = %v, want 59.4", got)
	}
}

func TestCalculateTotalAppliesVipExtraDiscount(t *testing.T) {
	r := NewReceipt()
	got := r.calculate_total([]float64{100}, "vip")
	if got != 84.4 {
		t.Errorf("calculate_total([100], 'vip') = %v, want 84.4", got)
	}
}

func TestCalculateTotalReturns0ForEmptyItems(t *testing.T) {
	r := NewReceipt()
	got := r.calculate_total([]float64{}, "")
	if got != 0 {
		t.Errorf("calculate_total([], '') = %v, want 0", got)
	}
}
