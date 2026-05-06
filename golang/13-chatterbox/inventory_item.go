package kata

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
func (i *InventoryItem) get_name() string          { return i.Name }
func (i *InventoryItem) get_batch_number() string { return i.BatchNumber }
func (i *InventoryItem) get_cache_timestamp() int  { return i.CacheTimestamp }
func (i *InventoryItem) get_row_id() int           { return i.RowID }
func (i *InventoryItem) get_quantity() int         { return i.Quantity }
