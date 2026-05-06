package kata

import "testing"

func TestInventoryItemStoresAndReturnsId(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	if item.get_id() != 1 {
		t.Errorf("expected id 1, got %d", item.get_id())
	}
}

func TestInventoryItemStoresAndReturnsName(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	if item.get_name() != "Widget" {
		t.Errorf("expected name 'Widget', got %q", item.get_name())
	}
}

func TestInventoryItemStoresAndReturnsBatchNumber(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	if item.get_batch_number() != "B001" {
		t.Errorf("expected batch number 'B001', got %q", item.get_batch_number())
	}
}

func TestInventoryItemStoresAndReturnsCacheTimestamp(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	if item.get_cache_timestamp() != 123 {
		t.Errorf("expected cache timestamp 123, got %d", item.get_cache_timestamp())
	}
}

func TestInventoryItemStoresAndReturnsRowId(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	if item.get_row_id() != 99 {
		t.Errorf("expected row id 99, got %d", item.get_row_id())
	}
}

func TestInventoryItemStoresAndReturnsQuantity(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	if item.get_quantity() != 10 {
		t.Errorf("expected quantity 10, got %d", item.get_quantity())
	}
}
