from src.order_processor import OrderProcessor

def order(price=10, quantity=1, email="a@b.com", zip_code="12345"):
    return {
        "items": [{"price": price, "quantity": quantity}],
        "customer": {"email": email},
        "address": {"zip": zip_code}
    }

def test_rejects_empty_items():
    processor = OrderProcessor()
    result = processor.process_order({"items": [], "customer": {"email": "a@b.com"}, "address": {"zip": "12345"}})
    assert result["error"] == "No items"

def test_rejects_invalid_customer():
    processor = OrderProcessor()
    result = processor.process_order({"items": [{"price": 10, "quantity": 1}], "customer": {}, "address": {"zip": "12345"}})
    assert result["error"] == "Invalid customer"

def test_rejects_invalid_address():
    processor = OrderProcessor()
    result = processor.process_order({"items": [{"price": 10, "quantity": 1}], "customer": {"email": "a@b.com"}, "address": {}})
    assert result["error"] == "Invalid address"

def test_rejects_out_of_stock():
    processor = OrderProcessor()
    result = processor.process_order(order(quantity=101))
    assert result["error"] == "Out of stock"

def test_calculates_totals_small_order():
    processor = OrderProcessor()
    result = processor.process_order(order())
    assert result["total"] == 16.69
    assert result["shippingLabel"]["carrier"] == "USPS"

def test_calculates_totals_medium_order():
    processor = OrderProcessor()
    result = processor.process_order(order(20))
    assert result["total"] == 27.39

def test_uses_ups_large_orders():
    processor = OrderProcessor()
    result = processor.process_order(order(100))
    assert result["shippingLabel"]["carrier"] == "UPS"
    assert result["paymentStatus"] == "approved"

def test_flags_high_total_pending_review():
    processor = OrderProcessor()
    result = processor.process_order(order(1000))
    assert result["paymentStatus"] == "pending_review"

def test_includes_email_confirmation():
    processor = OrderProcessor()
    result = processor.process_order(order(email="user@test.com"))
    assert result["email"]["to"] == "user@test.com"
    assert result["email"]["subject"] == "Order Confirmation"
