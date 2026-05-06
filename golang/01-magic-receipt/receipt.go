package kata

import "math"

type Receipt struct{}

func NewReceipt() *Receipt {
	return &Receipt{}
}

func (r *Receipt) calculate_total(items []float64, customerType string) float64 {
	total := 0.0
	for i := 0; i < len(items); i++ {
		total += items[i]
	}
	discount := 0.0
	if customerType == "member" {
		discount = total * 0.05
	} else if customerType == "vip" {
		discount = total * 0.15
	}
	if total > 50 {
		discount += 5
	}
	final := (total - discount) * 1.08
	if customerType == "vip" {
		final -= 2
	}
	return math.Round(final*100) / 100
}
