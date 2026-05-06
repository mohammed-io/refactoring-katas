package kata

import "strings"

type TicketItem struct {
	Name string
	Qty  int
}
type Order struct {
	Items    []TicketItem
	Customer string
	Table    int
	Special  string
}

type KitchenTicket struct{}

func NewKitchenTicket() *KitchenTicket {
	return &KitchenTicket{}
}

func (kt *KitchenTicket) print_ticket(order Order) string {
	lines := []string{}
	lines = append(lines, "Table: "+kt.itoa(order.Table))
	lines = append(lines, "Customer: "+order.Customer)
	for _, item := range order.Items {
		lines = append(lines, item.Name+" x"+kt.itoa(item.Qty))
	}
	if len(order.Special) > 0 {
		lines = append(lines, "Special: "+order.Special)
	}
	lines = append(lines, "---")
	return strings.Join(lines, "\n")
}
func (kt *KitchenTicket) itoa(n int) string { return kt.fmtInt(n) }
func (kt *KitchenTicket) fmtInt(n int) string {
	if n == 0 {
		return "0"
	}
	digits := ""
	for n > 0 {
		digits = string(rune('0'+n%10)) + digits
		n /= 10
	}
	return digits
}
