package kata

import "math"

type MonoItem struct {
	Price    float64
	Quantity int
}
type MonoCustomer struct {
	Type  string
	Email string
}
type MonoOrder struct {
	ID       int
	Items    []MonoItem
	Customer MonoCustomer
}
type MonoResult struct {
	Error         string
	OrderID       int
	Total         float64
	PaymentStatus string
	Carrier       string
	EmailTo       string
	Log           string
}

type LegacySystem struct{}

func NewLegacySystem() *LegacySystem {
	return &LegacySystem{}
}

func (ls *LegacySystem) process_everything(order MonoOrder) MonoResult {
	x := 0.0
	y := 0
	_ = y
	if len(order.Items) == 0 {
		return MonoResult{Error: "No items"}
	}
	for _, i := range order.Items {
		if i.Price > 0 {
			x += i.Price * float64(i.Quantity)
		}
	}
	d := 0.0
	if order.Customer.Type == "vip" {
		d = x * .2
	} else if order.Customer.Type == "member" {
		d = x * .1
	}
	if x > 100 {
		d += 10
	}
	total := x - d + (x-d)*.07
	status := "approved"
	if total > 5000 {
		status = "manual_review"
	}
	carrier := "USPS"
	if total > 50 {
		carrier = "UPS"
	}
	return MonoResult{OrderID: order.ID, Total: math.Round(total*100) / 100, PaymentStatus: status, Carrier: carrier, EmailTo: order.Customer.Email, Log: "Order processed"}
}
