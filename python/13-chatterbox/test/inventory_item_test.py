from src.inventory_item import InventoryItem

def test_public_snapshot_contains_business_fields():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    assert item.public_snapshot() == {
        "id": 1,
        "name": "Widget",
        "batch_number": "B001",
        "quantity": 10,
        "stock_status": "available",
    }

def test_reserves_stock_and_reports_remaining_quantity():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    result = item.reserve(3)
    assert result == {"status": "reserved", "reserved": 3, "remaining": 7, "sku": "1-B001"}
    assert item.public_snapshot()["quantity"] == 7

def test_rejects_reservation_when_quantity_is_invalid():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    result = item.reserve(0)
    assert result == {"status": "rejected", "reason": "invalid_quantity", "remaining": 10}
    assert item.public_snapshot()["quantity"] == 10

def test_backorders_when_not_enough_stock():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    result = item.reserve(12)
    assert result == {"status": "backorder", "reserved": 0, "remaining": 10}
    assert item.public_snapshot()["quantity"] == 10

def test_receives_stock_after_low_quantity():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    item.reserve(8)
    assert item.public_snapshot()["stock_status"] == "low"
    assert item.receive_stock(5) == 7
    assert item.public_snapshot()["stock_status"] == "available"

def test_reports_out_of_stock_after_exact_reservation():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    item.reserve(10)
    assert item.public_snapshot()["stock_status"] == "out"
