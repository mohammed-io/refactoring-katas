package kata

import (
	"strings"
	"testing"
)

func TestPrintTicketPrintsSimpleTicket(t *testing.T) {
	kp := NewKitchenTicket()
	order := Order{Items: []TicketItem{{Name: "Burger", Qty: 1}}, Customer: "Alice", Table: 5, Special: ""}
	result := kp.print_ticket(order)
	if !strings.Contains(result, "Table: 5") {
		t.Error("expected ticket to contain 'Table: 5'")
	}
	if !strings.Contains(result, "Customer: Alice") {
		t.Error("expected ticket to contain 'Customer: Alice'")
	}
	if !strings.Contains(result, "Burger x1") {
		t.Error("expected ticket to contain 'Burger x1'")
	}
}

func TestPrintTicketPrintsTicketWithMultipleItems(t *testing.T) {
	kp := NewKitchenTicket()
	order := Order{Items: []TicketItem{{Name: "Burger", Qty: 2}, {Name: "Fries", Qty: 1}}, Customer: "Bob", Table: 12, Special: ""}
	result := kp.print_ticket(order)
	if !strings.Contains(result, "Burger x2") {
		t.Error("expected ticket to contain 'Burger x2'")
	}
	if !strings.Contains(result, "Fries x1") {
		t.Error("expected ticket to contain 'Fries x1'")
	}
}

func TestPrintTicketPrintsTicketWithSpecialInstructions(t *testing.T) {
	kp := NewKitchenTicket()
	order := Order{Items: []TicketItem{{Name: "Salad", Qty: 1}}, Customer: "Carol", Table: 3, Special: "No onions"}
	result := kp.print_ticket(order)
	if !strings.Contains(result, "Special: No onions") {
		t.Error("expected ticket to contain 'Special: No onions'")
	}
}

func TestPrintTicketOmitsSpecialWhenEmpty(t *testing.T) {
	kp := NewKitchenTicket()
	order := Order{Items: []TicketItem{{Name: "Pizza", Qty: 1}}, Customer: "Dave", Table: 7, Special: ""}
	result := kp.print_ticket(order)
	if strings.Contains(result, "Special:") {
		t.Error("expected ticket to NOT contain 'Special:'")
	}
}

func TestPrintTicketIncludesSeparatorLine(t *testing.T) {
	kp := NewKitchenTicket()
	order := Order{Items: []TicketItem{{Name: "Soup", Qty: 1}}, Customer: "Eve", Table: 1, Special: ""}
	result := kp.print_ticket(order)
	if !strings.Contains(result, "---") {
		t.Error("expected ticket to contain '---'")
	}
}
