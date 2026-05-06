from src.legacy_system import LegacySystem

def test_rejects_empty_order():
    system = LegacySystem()
    result = system.process_everything({"id": 1, "items": []})
    assert result["error"] == "No items"

def test_calculates_basic_total():
    system = LegacySystem()
    result = system.process_everything({
        "id": 1,
        "items": [{"price": 10, "quantity": 2}],
        "customer": {"email": "a@b.com"}
    })
    assert result["total"] == 21.4
    assert result["carrier"] == "USPS"

def test_applies_member_discount():
    system = LegacySystem()
    result = system.process_everything({
        "id": 2,
        "items": [{"price": 100, "quantity": 1}],
        "customer": {"type": "member", "email": "a@b.com"}
    })
    assert result["total"] == 96.3

def test_applies_vip_discount():
    system = LegacySystem()
    result = system.process_everything({
        "id": 3,
        "items": [{"price": 100, "quantity": 1}],
        "customer": {"type": "vip", "email": "a@b.com"}
    })
    assert result["total"] == 85.6

def test_applies_bonus_discount_over_100():
    system = LegacySystem()
    result = system.process_everything({
        "id": 4,
        "items": [{"price": 200, "quantity": 1}],
        "customer": {"email": "a@b.com"}
    })
    assert result["total"] == 203.3

def test_uses_ups_for_large_total():
    system = LegacySystem()
    result = system.process_everything({
        "id": 5,
        "items": [{"price": 60, "quantity": 1}],
        "customer": {"email": "a@b.com"}
    })
    assert result["carrier"] == "UPS"

def test_flags_high_total_for_review():
    system = LegacySystem()
    result = system.process_everything({
        "id": 6,
        "items": [{"price": 5000, "quantity": 1}],
        "customer": {"email": "a@b.com"}
    })
    assert result["paymentStatus"] == "manual_review"

def test_includes_email_details():
    system = LegacySystem()
    result = system.process_everything({
        "id": 7,
        "items": [{"price": 10, "quantity": 1}],
        "customer": {"email": "user@test.com"}
    })
    assert result["email"]["to"] == "user@test.com"
    assert "7" in result["email"]["subject"]

def test_includes_log_entry():
    system = LegacySystem()
    result = system.process_everything({
        "id": 8,
        "items": [{"price": 10, "quantity": 1}],
        "customer": {"email": "a@b.com"}
    })
    assert "Order processed" in result["log"]

def test_includes_order_id():
    system = LegacySystem()
    result = system.process_everything({
        "id": 99,
        "items": [{"price": 10, "quantity": 1}],
        "customer": {"email": "a@b.com"}
    })
    assert result["orderId"] == 99
