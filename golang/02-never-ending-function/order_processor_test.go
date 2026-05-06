package kata

import (
	"strings"
	"testing"
)

func TestProcessOrderRejectsEmptyItems(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{}, Customer: Customer{Email: "a@b.com"}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.Error != "No items" {
		t.Errorf("expected error 'No items', got %q", result.Error)
	}
}

func TestProcessOrderRejectsInvalidCustomer(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 10, Quantity: 1}}, Customer: Customer{}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.Error != "Invalid customer" {
		t.Errorf("expected error 'Invalid customer', got %q", result.Error)
	}
}

func TestProcessOrderRejectsInvalidAddress(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 10, Quantity: 1}}, Customer: Customer{Email: "a@b.com"}, Address: Address{}}
	result := op.process_order(order)
	if result.Error != "Invalid address" {
		t.Errorf("expected error 'Invalid address', got %q", result.Error)
	}
}

func TestProcessOrderRejectsOutOfStock(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 10, Quantity: 101}}, Customer: Customer{Email: "a@b.com"}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.Error != "Out of stock" {
		t.Errorf("expected error 'Out of stock', got %q", result.Error)
	}
}

func TestProcessOrderCalculatesTotalsForSmallOrder(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 10, Quantity: 1}}, Customer: Customer{Email: "a@b.com"}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.Total != 16.69 {
		t.Errorf("expected total 16.69, got %v", result.Total)
	}
	if result.Carrier != "USPS" {
		t.Errorf("expected carrier USPS, got %q", result.Carrier)
	}
}

func TestProcessOrderCalculatesTotalsForMediumOrder(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 20, Quantity: 1}}, Customer: Customer{Email: "a@b.com"}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.Total != 27.39 {
		t.Errorf("expected total 27.39, got %v", result.Total)
	}
}

func TestProcessOrderUsesUPSForLargeOrders(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 100, Quantity: 1}}, Customer: Customer{Email: "a@b.com"}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.Carrier != "UPS" {
		t.Errorf("expected carrier UPS, got %q", result.Carrier)
	}
	if result.PaymentStatus != "approved" {
		t.Errorf("expected payment status 'approved', got %q", result.PaymentStatus)
	}
}

func TestProcessOrderFlagsHighTotalAsPendingReview(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 1000, Quantity: 1}}, Customer: Customer{Email: "a@b.com"}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.PaymentStatus != "pending_review" {
		t.Errorf("expected payment status 'pending_review', got %q", result.PaymentStatus)
	}
}

func TestProcessOrderIncludesEmailConfirmation(t *testing.T) {
	op := NewOrderProcessor()
	order := Order{Items: []Item{{Price: 10, Quantity: 1}}, Customer: Customer{Email: "user@test.com"}, Address: Address{Zip: "12345"}}
	result := op.process_order(order)
	if result.EmailTo != "user@test.com" {
		t.Errorf("expected email to 'user@test.com', got %q", result.EmailTo)
	}
	if !strings.HasPrefix(result.OrderID, "ORD-") {
		t.Errorf("expected OrderID to start with 'ORD-', got %q", result.OrderID)
	}
}
