from src.kitchen_ticket import KitchenTicket

def test_simple_ticket():
    order = KitchenTicket([{"name": "Burger", "qty": 1}], "Alice", 5, "")
    result = order.print_ticket()
    assert "Table: 5" in result
    assert "Customer: Alice" in result
    assert "Burger x1" in result

def test_multiple_items():
    order = KitchenTicket([{"name": "Burger", "qty": 2}, {"name": "Fries", "qty": 1}], "Bob", 12, "")
    result = order.print_ticket()
    assert "Burger x2" in result
    assert "Fries x1" in result

def test_special_instructions():
    order = KitchenTicket([{"name": "Salad", "qty": 1}], "Carol", 3, "No onions")
    result = order.print_ticket()
    assert "Special: No onions" in result

def test_omit_special_when_empty():
    order = KitchenTicket([{"name": "Pizza", "qty": 1}], "Dave", 7, "")
    result = order.print_ticket()
    assert "Special:" not in result

def test_includes_separator():
    order = KitchenTicket([{"name": "Soup", "qty": 1}], "Eve", 1, "")
    result = order.print_ticket()
    assert "---" in result
