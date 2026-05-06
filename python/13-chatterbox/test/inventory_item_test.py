from src.inventory_item import InventoryItem

def test_id():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    assert item.get_id() == 1
    item.set_id(2)
    assert item.get_id() == 2

def test_name():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    assert item.get_name() == "Widget"
    item.set_name("Gadget")
    assert item.get_name() == "Gadget"

def test_batch_number():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    assert item.get_batch_number() == "B001"
    item.set_batch_number("B002")
    assert item.get_batch_number() == "B002"

def test_cache_timestamp():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    assert item.get_cache_timestamp() == 123
    item.set_cache_timestamp(456)
    assert item.get_cache_timestamp() == 456

def test_row_id():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    assert item.get_row_id() == 99
    item.set_row_id(100)
    assert item.get_row_id() == 100

def test_quantity():
    item = InventoryItem(1, "Widget", "B001", 123, 99, 10)
    assert item.get_quantity() == 10
    item.set_quantity(5)
    assert item.get_quantity() == 5
