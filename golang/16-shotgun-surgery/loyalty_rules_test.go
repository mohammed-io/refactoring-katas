package kata

import "testing"

func TestGetDiscountForTierBronze(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_discount_for_tier("bronze"); got != 0.05 {
		t.Errorf("get_discount_for_tier(bronze) = %v, want 0.05", got)
	}
}

func TestGetDiscountForTierSilver(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_discount_for_tier("silver"); got != 0.1 {
		t.Errorf("get_discount_for_tier(silver) = %v, want 0.1", got)
	}
}

func TestGetDiscountForTierGold(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_discount_for_tier("gold"); got != 0.15 {
		t.Errorf("get_discount_for_tier(gold) = %v, want 0.15", got)
	}
}

func TestGetDiscountForTierPlatinum(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_discount_for_tier("platinum"); got != 0.2 {
		t.Errorf("get_discount_for_tier(platinum) = %v, want 0.2", got)
	}
}

func TestGetDiscountForTierUnknown(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_discount_for_tier("unknown"); got != 0 {
		t.Errorf("get_discount_for_tier(unknown) = %v, want 0", got)
	}
}

func TestGetLabelForTierBronze(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_label_for_tier("bronze"); got != "Bronze Member" {
		t.Errorf("get_label_for_tier(bronze) = %q, want 'Bronze Member'", got)
	}
}

func TestGetLabelForTierSilver(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_label_for_tier("silver"); got != "Silver Member" {
		t.Errorf("get_label_for_tier(silver) = %q, want 'Silver Member'", got)
	}
}

func TestGetLabelForTierGold(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_label_for_tier("gold"); got != "Gold Member" {
		t.Errorf("get_label_for_tier(gold) = %q, want 'Gold Member'", got)
	}
}

func TestGetLabelForTierPlatinum(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_label_for_tier("platinum"); got != "Platinum Member" {
		t.Errorf("get_label_for_tier(platinum) = %q, want 'Platinum Member'", got)
	}
}

func TestGetLabelForTierUnknown(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_label_for_tier("unknown"); got != "Standard" {
		t.Errorf("get_label_for_tier(unknown) = %q, want 'Standard'", got)
	}
}

func TestGetThresholdForTierBronze(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_threshold_for_tier("bronze"); got != 100 {
		t.Errorf("get_threshold_for_tier(bronze) = %v, want 100", got)
	}
}

func TestGetThresholdForTierSilver(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_threshold_for_tier("silver"); got != 500 {
		t.Errorf("get_threshold_for_tier(silver) = %v, want 500", got)
	}
}

func TestGetThresholdForTierGold(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_threshold_for_tier("gold"); got != 2000 {
		t.Errorf("get_threshold_for_tier(gold) = %v, want 2000", got)
	}
}

func TestGetThresholdForTierPlatinum(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_threshold_for_tier("platinum"); got != 10000 {
		t.Errorf("get_threshold_for_tier(platinum) = %v, want 10000", got)
	}
}

func TestGetThresholdForTierUnknown(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_threshold_for_tier("unknown"); got != 0 {
		t.Errorf("get_threshold_for_tier(unknown) = %v, want 0", got)
	}
}

func TestGetColorForTierBronze(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_color_for_tier("bronze"); got != "#CD7F32" {
		t.Errorf("get_color_for_tier(bronze) = %q, want '#CD7F32'", got)
	}
}

func TestGetColorForTierSilver(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_color_for_tier("silver"); got != "#C0C0C0" {
		t.Errorf("get_color_for_tier(silver) = %q, want '#C0C0C0'", got)
	}
}

func TestGetColorForTierGold(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_color_for_tier("gold"); got != "#FFD700" {
		t.Errorf("get_color_for_tier(gold) = %q, want '#FFD700'", got)
	}
}

func TestGetColorForTierPlatinum(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_color_for_tier("platinum"); got != "#E5E4E2" {
		t.Errorf("get_color_for_tier(platinum) = %q, want '#E5E4E2'", got)
	}
}

func TestGetColorForTierUnknown(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.get_color_for_tier("unknown"); got != "#000000" {
		t.Errorf("get_color_for_tier(unknown) = %q, want '#000000'", got)
	}
}

func TestCalculateTotalForBronze(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.calculate_total(100, "bronze"); got != 95 {
		t.Errorf("calculate_total(100, bronze) = %v, want 95", got)
	}
}

func TestCalculateTotalForPlatinum(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.calculate_total(100, "platinum"); got != 80 {
		t.Errorf("calculate_total(100, platinum) = %v, want 80", got)
	}
}

func TestCalculateTotalForUnknownTier(t *testing.T) {
	lr := NewLoyaltyRules()
	if got := lr.calculate_total(100, "unknown"); got != 100 {
		t.Errorf("calculate_total(100, unknown) = %v, want 100", got)
	}
}
