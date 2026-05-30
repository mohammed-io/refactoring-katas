package kata

import "strings"

type TicketItem struct {
	Name      string
	Qty       int
	Modifiers []string
	Allergy   string
}
type Order struct {
	Items    []TicketItem
	Customer map[string]any
	Table    map[string]any
	Special  string
	Rush     bool
}

type KitchenTicket struct{}

func NewKitchenTicket() *KitchenTicket {
	return &KitchenTicket{}
}

func (kt *KitchenTicket) print_ticket(order Order) string {
	lines := []string{}
	lines = append(lines, "Table: "+kt.itoa(order.Table["number"].(int)))
	if order.Table["zone"] != nil {
		lines = append(lines, "Zone: "+order.Table["zone"].(string))
	} else {
		lines = append(lines, "Zone: main")
	}
	if order.Table["server"] != nil {
		lines = append(lines, "Server: "+order.Table["server"].(string))
	} else {
		lines = append(lines, "Server: unassigned")
	}
	lines = append(lines, "Customer: "+order.Customer["name"].(string))
	if order.Customer["vip"] == true {
		lines = append(lines, "VIP")
	}
	if order.Rush {
		lines = append(lines, "RUSH")
	}
	totalItems := 0
	for _, item := range order.Items {
		totalItems += item.Qty
		line := item.Name + " x" + kt.itoa(item.Qty)
		if len(item.Modifiers) > 0 {
			line += " [" + strings.Join(item.Modifiers, ", ") + "]"
		}
		if item.Allergy != "" {
			line += " ALLERGY:" + item.Allergy
		}
		lines = append(lines, line)
	}
	lines = append(lines, "Items: "+kt.itoa(totalItems))
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
