package kata

import "fmt"

type InventoryItem struct {
	ID             int
	Name           string
	BatchNumber    string
	CacheTimestamp int
	RowID          int
	Quantity       int
}

func NewInventoryItem(id int, name string, batchNumber string, cacheTimestamp int, rowID int, quantity int) *InventoryItem {
	return &InventoryItem{ID: id, Name: name, BatchNumber: batchNumber, CacheTimestamp: cacheTimestamp, RowID: rowID, Quantity: quantity}
}

func (i *InventoryItem) get_id() int              { return i.ID }
func (i *InventoryItem) get_name() string         { return i.Name }
func (i *InventoryItem) get_batch_number() string { return i.BatchNumber }
func (i *InventoryItem) get_cache_timestamp() int { return i.CacheTimestamp }
func (i *InventoryItem) get_row_id() int          { return i.RowID }
func (i *InventoryItem) get_quantity() int        { return i.Quantity }

func (i *InventoryItem) reserve(units int) map[string]interface{} {
	if units <= 0 {
		return map[string]interface{}{"status": "rejected", "reason": "invalid_quantity", "remaining": i.Quantity}
	}
	if units > i.Quantity {
		return map[string]interface{}{"status": "backorder", "reserved": 0, "remaining": i.Quantity}
	}

	i.Quantity -= units
	return map[string]interface{}{
		"status":    "reserved",
		"reserved":  units,
		"remaining": i.Quantity,
		"sku":       fmt.Sprintf("%d-%s", i.ID, i.BatchNumber),
	}
}

func (i *InventoryItem) receive_stock(units int) int {
	if units <= 0 {
		return i.Quantity
	}

	i.Quantity += units
	return i.Quantity
}

func (i *InventoryItem) public_snapshot() map[string]interface{} {
	return map[string]interface{}{
		"id":           i.ID,
		"name":         i.Name,
		"batch_number": i.BatchNumber,
		"quantity":     i.Quantity,
		"stock_status": i.stock_status(),
	}
}

func (i *InventoryItem) stock_status() string {
	if i.Quantity == 0 {
		return "out"
	}
	if i.Quantity < 5 {
		return "low"
	}
	return "available"
}
