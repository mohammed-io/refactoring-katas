package kata

import (
	"strings"
	"testing"
)

func makeOrder() Order {
	return Order{
		Items:    []TicketItem{{Name: "Burger", Qty: 1}},
		Customer: map[string]any{"name": "Alice", "vip": false},
		Table:    map[string]any{"number": 5, "zone": "patio", "server": "Sam"},
		Special:  "",
		Rush:     false,
	}
}

func TestPrintTicketPrintsTableCustomerAndServerDetails(t *testing.T) {
	result := NewKitchenTicket().print_ticket(makeOrder())
	for _, want := range []string{"Table: 5", "Zone: patio", "Server: Sam", "Customer: Alice"} {
		if !strings.Contains(result, want) {
			t.Errorf("expected ticket to contain %q", want)
		}
	}
}

func TestPrintTicketPrintsTicketWithMultipleItemsAndCount(t *testing.T) {
	order := makeOrder()
	order.Items = []TicketItem{{Name: "Burger", Qty: 2}, {Name: "Fries", Qty: 1}}
	result := NewKitchenTicket().print_ticket(order)
	for _, want := range []string{"Burger x2", "Fries x1", "Items: 3"} {
		if !strings.Contains(result, want) {
			t.Errorf("expected ticket to contain %q", want)
		}
	}
}

func TestPrintTicketPrintsModifiersAndAllergyFlags(t *testing.T) {
	order := makeOrder()
	order.Items = []TicketItem{{Name: "Salad", Qty: 1, Modifiers: []string{"no onion", "dressing side"}, Allergy: "nuts"}}
	result := NewKitchenTicket().print_ticket(order)
	if !strings.Contains(result, "Salad x1 [no onion, dressing side] ALLERGY:nuts") {
		t.Errorf("expected modifiers and allergy, got %q", result)
	}
}

func TestPrintTicketPrintsVipAndRushMarkers(t *testing.T) {
	order := makeOrder()
	order.Customer = map[string]any{"name": "Carol", "vip": true}
	order.Rush = true
	result := NewKitchenTicket().print_ticket(order)
	if !strings.Contains(result, "VIP") || !strings.Contains(result, "RUSH") {
		t.Errorf("expected vip and rush markers, got %q", result)
	}
}

func TestPrintTicketOmitsSpecialWhenEmptyButKeepsSeparator(t *testing.T) {
	result := NewKitchenTicket().print_ticket(makeOrder())
	if strings.Contains(result, "Special:") {
		t.Error("expected ticket to omit special")
	}
	if !strings.Contains(result, "---") {
		t.Error("expected separator")
	}
}

func TestPrintTicketPrintsSpecialInstructions(t *testing.T) {
	order := makeOrder()
	order.Special = "Fire mains after starters"
	result := NewKitchenTicket().print_ticket(order)
	if !strings.Contains(result, "Special: Fire mains after starters") {
		t.Errorf("expected special instruction, got %q", result)
	}
}
