from src.receipt import Receipt

def test_regular_customer():
    receipt = Receipt()
    assert receipt.calculate_total([10, 20, 30]) == 59.4

def test_member_customer():
    receipt = Receipt()
    assert receipt.calculate_total([10, 20, 30], "member") == 56.16

def test_vip_customer():
    receipt = Receipt()
    assert receipt.calculate_total([10, 20, 30], "vip") == 47.68

def test_bonus_discount_over_50():
    receipt = Receipt()
    assert receipt.calculate_total([60]) == 59.4

def test_vip_extra_discount():
    receipt = Receipt()
    assert receipt.calculate_total([100], "vip") == 84.4

def test_empty_items():
    receipt = Receipt()
    assert receipt.calculate_total([]) == 0


def test_exactly_50_no_bonus():
    receipt = Receipt()
    assert receipt.calculate_total([50]) == 54.0
