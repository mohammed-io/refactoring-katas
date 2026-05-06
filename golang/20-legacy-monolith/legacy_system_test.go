package kata

import (
	"strings"
	"testing"
)

func TestProcessEverythingRejectsEmptyOrder(t *testing.T) {
	ls := NewLegacySystem()
	result := ls.process_everything(MonoOrder{ID: 1, Items: []MonoItem{}})
	if result.Error != "No items" {
		t.Errorf("expected error 'No items', got %q", result.Error)
	}
}

func TestProcessEverythingCalculatesBasicTotal(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 1, Items: []MonoItem{{Price: 10, Quantity: 2}}, Customer: MonoCustomer{Email: "a@b.com"}}
	result := ls.process_everything(order)
	if result.Total != 21.4 {
		t.Errorf("expected total 21.4, got %v", result.Total)
	}
	if result.Carrier != "USPS" {
		t.Errorf("expected carrier 'USPS', got %q", result.Carrier)
	}
}

func TestProcessEverythingAppliesMemberDiscount(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 2, Items: []MonoItem{{Price: 100, Quantity: 1}}, Customer: MonoCustomer{Type: "member", Email: "a@b.com"}}
	result := ls.process_everything(order)
	if result.Total != 96.3 {
		t.Errorf("expected total 96.3, got %v", result.Total)
	}
}

func TestProcessEverythingAppliesVipDiscount(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 3, Items: []MonoItem{{Price: 100, Quantity: 1}}, Customer: MonoCustomer{Type: "vip", Email: "a@b.com"}}
	result := ls.process_everything(order)
	if result.Total != 85.6 {
		t.Errorf("expected total 85.6, got %v", result.Total)
	}
}

func TestProcessEverythingAppliesBonusDiscountOver100(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 4, Items: []MonoItem{{Price: 200, Quantity: 1}}, Customer: MonoCustomer{Email: "a@b.com"}}
	result := ls.process_everything(order)
	if result.Total != 203.3 {
		t.Errorf("expected total 203.3, got %v", result.Total)
	}
}

func TestProcessEverythingUsesUPSForLargeTotal(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 5, Items: []MonoItem{{Price: 60, Quantity: 1}}, Customer: MonoCustomer{Email: "a@b.com"}}
	result := ls.process_everything(order)
	if result.Carrier != "UPS" {
		t.Errorf("expected carrier 'UPS', got %q", result.Carrier)
	}
}

func TestProcessEverythingFlagsHighTotalForReview(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 6, Items: []MonoItem{{Price: 5000, Quantity: 1}}, Customer: MonoCustomer{Email: "a@b.com"}}
	result := ls.process_everything(order)
	if result.PaymentStatus != "manual_review" {
		t.Errorf("expected payment status 'manual_review', got %q", result.PaymentStatus)
	}
}

func TestProcessEverythingIncludesEmailDetails(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 7, Items: []MonoItem{{Price: 10, Quantity: 1}}, Customer: MonoCustomer{Email: "user@test.com"}}
	result := ls.process_everything(order)
	if result.EmailTo != "user@test.com" {
		t.Errorf("expected email to 'user@test.com', got %q", result.EmailTo)
	}
}

func TestProcessEverythingIncludesLogEntry(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 8, Items: []MonoItem{{Price: 10, Quantity: 1}}, Customer: MonoCustomer{Email: "a@b.com"}}
	result := ls.process_everything(order)
	if !strings.Contains(result.Log, "Order processed") {
		t.Errorf("expected log to contain 'Order processed', got %q", result.Log)
	}
}

func TestProcessEverythingIncludesOrderId(t *testing.T) {
	ls := NewLegacySystem()
	order := MonoOrder{ID: 99, Items: []MonoItem{{Price: 10, Quantity: 1}}, Customer: MonoCustomer{Email: "a@b.com"}}
	result := ls.process_everything(order)
	if result.OrderID != 99 {
		t.Errorf("expected order id 99, got %d", result.OrderID)
	}
}
