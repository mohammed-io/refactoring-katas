package kata

import (
	"fmt"
	"math"
	"time"
)

type Item struct {
	Price    float64
	Quantity int
}
type Customer struct{ Email string }
type Address struct{ Zip string }
type Order struct {
	Items    []Item
	Customer Customer
	Address  Address
}
type OrderResult struct {
	Error         string
	Total         float64
	Carrier       string
	PaymentStatus string
	EmailTo       string
	OrderID       string
}

type OrderProcessor struct{}

func NewOrderProcessor() *OrderProcessor {
	return &OrderProcessor{}
}

func (op *OrderProcessor) process_order(order Order) OrderResult {
	if len(order.Items) == 0 {
		return OrderResult{Error: "No items"}
	}
	if order.Customer.Email == "" {
		return OrderResult{Error: "Invalid customer"}
	}
	if order.Address.Zip == "" {
		return OrderResult{Error: "Invalid address"}
	}
	inventory := true
	for _, item := range order.Items {
		if item.Quantity > 100 {
			inventory = false
		}
	}
	if !inventory {
		return OrderResult{Error: "Out of stock"}
	}
	subtotal := 0.0
	weight := 0
	for _, item := range order.Items {
		subtotal += item.Price * float64(item.Quantity)
		weight += item.Quantity
	}
	_ = weight
	shipping := 0.0
	if subtotal < 25 {
		shipping = 5.99
	} else if subtotal < 50 {
		shipping = 3.99
	}
	total := subtotal + subtotal*0.07 + shipping
	carrier := "USPS"
	if total > 100 {
		carrier = "UPS"
	}
	status := "approved"
	if total > 1000 {
		status = "pending_review"
	}
	return OrderResult{Total: math.Round(total*100) / 100, Carrier: carrier, PaymentStatus: status, EmailTo: order.Customer.Email, OrderID: fmt.Sprintf("ORD-%d", time.Now().UnixMilli())}
}
