package kata

import (
	"reflect"
	"testing"
)

func TestInventoryItemPublicSnapshotContainsBusinessFields(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	expected := map[string]interface{}{"id": 1, "name": "Widget", "batch_number": "B001", "quantity": 10, "stock_status": "available"}
	if !reflect.DeepEqual(expected, item.public_snapshot()) {
		t.Errorf("expected %#v, got %#v", expected, item.public_snapshot())
	}
}

func TestInventoryItemReservesStockAndReportsRemainingQuantity(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	expected := map[string]interface{}{"status": "reserved", "reserved": 3, "remaining": 7, "sku": "1-B001"}
	result := item.reserve(3)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected reservation %#v, got %#v", expected, result)
	}
	if item.public_snapshot()["quantity"] != 7 {
		t.Errorf("expected quantity to reflect reservation")
	}
}

func TestInventoryItemRejectsReservationWhenQuantityIsInvalid(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	expected := map[string]interface{}{"status": "rejected", "reason": "invalid_quantity", "remaining": 10}
	if !reflect.DeepEqual(expected, item.reserve(0)) {
		t.Errorf("expected rejected reservation")
	}
	if item.public_snapshot()["quantity"] != 10 {
		t.Errorf("expected unchanged quantity")
	}
}

func TestInventoryItemBackordersWhenNotEnoughStock(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	expected := map[string]interface{}{"status": "backorder", "reserved": 0, "remaining": 10}
	if !reflect.DeepEqual(expected, item.reserve(12)) {
		t.Errorf("expected backorder")
	}
	if item.public_snapshot()["quantity"] != 10 {
		t.Errorf("expected unchanged quantity")
	}
}

func TestInventoryItemReceivesStockAfterLowQuantity(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	item.reserve(8)
	if item.public_snapshot()["stock_status"] != "low" {
		t.Errorf("expected low stock")
	}
	if item.receive_stock(5) != 7 {
		t.Errorf("expected quantity 7")
	}
	if item.public_snapshot()["stock_status"] != "available" {
		t.Errorf("expected available stock")
	}
}

func TestInventoryItemReportsOutOfStockAfterExactReservation(t *testing.T) {
	item := NewInventoryItem(1, "Widget", "B001", 123, 99, 10)
	item.reserve(10)
	if item.public_snapshot()["stock_status"] != "out" {
		t.Errorf("expected out of stock")
	}
}
