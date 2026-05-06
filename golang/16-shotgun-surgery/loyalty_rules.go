package kata

type LoyaltyRules struct{}

func NewLoyaltyRules() *LoyaltyRules {
	return &LoyaltyRules{}
}

func (lr *LoyaltyRules) get_discount_for_tier(t string) float64 {
	if t == "bronze" {
		return .05
	} else if t == "silver" {
		return .1
	} else if t == "gold" {
		return .15
	} else if t == "platinum" {
		return .2
	}
	return 0
}
func (lr *LoyaltyRules) get_label_for_tier(t string) string {
	if t == "bronze" {
		return "Bronze Member"
	} else if t == "silver" {
		return "Silver Member"
	} else if t == "gold" {
		return "Gold Member"
	} else if t == "platinum" {
		return "Platinum Member"
	}
	return "Standard"
}
func (lr *LoyaltyRules) get_threshold_for_tier(t string) int {
	if t == "bronze" {
		return 100
	} else if t == "silver" {
		return 500
	} else if t == "gold" {
		return 2000
	} else if t == "platinum" {
		return 10000
	}
	return 0
}
func (lr *LoyaltyRules) get_color_for_tier(t string) string {
	if t == "bronze" {
		return "#CD7F32"
	} else if t == "silver" {
		return "#C0C0C0"
	} else if t == "gold" {
		return "#FFD700"
	} else if t == "platinum" {
		return "#E5E4E2"
	}
	return "#000000"
}
func (lr *LoyaltyRules) calculate_total(s float64, t string) float64 { return s * (1 - lr.get_discount_for_tier(t)) }
