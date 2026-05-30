from src.kitchen_ticket import KitchenTicket, Order


def order(items=None, customer=None, table=None, special="", rush=False):
    return Order(
        items or [{"name": "Burger", "qty": 1}],
        customer or {"name": "Alice", "vip": False},
        table or {"number": 5, "zone": "patio", "server": "Sam"},
        special,
        rush,
    )


def test_prints_table_customer_and_server_details():
    result = KitchenTicket().print_ticket(order())
    assert "Table: 5" in result
    assert "Zone: patio" in result
    assert "Server: Sam" in result
    assert "Customer: Alice" in result


def test_prints_ticket_with_multiple_items_and_count():
    result = KitchenTicket().print_ticket(order(items=[{"name": "Burger", "qty": 2}, {"name": "Fries", "qty": 1}]))
    assert "Burger x2" in result
    assert "Fries x1" in result
    assert "Items: 3" in result


def test_prints_modifiers_and_allergy_flags():
    result = KitchenTicket().print_ticket(order(items=[{"name": "Salad", "qty": 1, "modifiers": ["no onion", "dressing side"], "allergy": "nuts"}]))
    assert "Salad x1 [no onion, dressing side] ALLERGY:nuts" in result


def test_prints_vip_and_rush_markers():
    result = KitchenTicket().print_ticket(order(customer={"name": "Carol", "vip": True}, rush=True))
    assert "VIP" in result
    assert "RUSH" in result


def test_omits_special_when_empty_but_keeps_separator():
    result = KitchenTicket().print_ticket(order(special=""))
    assert "Special:" not in result
    assert "---" in result


def test_prints_special_instructions():
    result = KitchenTicket().print_ticket(order(special="Fire mains after starters"))
    assert "Special: Fire mains after starters" in result
