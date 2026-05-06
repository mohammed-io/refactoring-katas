from src.loyalty_rules import LoyaltyRules

def test_bronze_discount():
    rules = LoyaltyRules()
    assert rules.get_discount_for_tier("bronze") == 0.05

def test_silver_discount():
    rules = LoyaltyRules()
    assert rules.get_discount_for_tier("silver") == 0.1

def test_gold_discount():
    rules = LoyaltyRules()
    assert rules.get_discount_for_tier("gold") == 0.15

def test_platinum_discount():
    rules = LoyaltyRules()
    assert rules.get_discount_for_tier("platinum") == 0.2

def test_unknown_tier_discount():
    rules = LoyaltyRules()
    assert rules.get_discount_for_tier("unknown") == 0

def test_bronze_label():
    rules = LoyaltyRules()
    assert rules.get_label_for_tier("bronze") == "Bronze Member"

def test_silver_label():
    rules = LoyaltyRules()
    assert rules.get_label_for_tier("silver") == "Silver Member"

def test_gold_label():
    rules = LoyaltyRules()
    assert rules.get_label_for_tier("gold") == "Gold Member"

def test_platinum_label():
    rules = LoyaltyRules()
    assert rules.get_label_for_tier("platinum") == "Platinum Member"

def test_unknown_tier_label():
    rules = LoyaltyRules()
    assert rules.get_label_for_tier("unknown") == "Standard"

def test_bronze_threshold():
    rules = LoyaltyRules()
    assert rules.get_threshold_for_tier("bronze") == 100

def test_silver_threshold():
    rules = LoyaltyRules()
    assert rules.get_threshold_for_tier("silver") == 500

def test_gold_threshold():
    rules = LoyaltyRules()
    assert rules.get_threshold_for_tier("gold") == 2000

def test_platinum_threshold():
    rules = LoyaltyRules()
    assert rules.get_threshold_for_tier("platinum") == 10000

def test_unknown_tier_threshold():
    rules = LoyaltyRules()
    assert rules.get_threshold_for_tier("unknown") == 0

def test_bronze_color():
    rules = LoyaltyRules()
    assert rules.get_color_for_tier("bronze") == "#CD7F32"

def test_silver_color():
    rules = LoyaltyRules()
    assert rules.get_color_for_tier("silver") == "#C0C0C0"

def test_gold_color():
    rules = LoyaltyRules()
    assert rules.get_color_for_tier("gold") == "#FFD700"

def test_platinum_color():
    rules = LoyaltyRules()
    assert rules.get_color_for_tier("platinum") == "#E5E4E2"

def test_unknown_tier_color():
    rules = LoyaltyRules()
    assert rules.get_color_for_tier("unknown") == "#000000"

def test_calculates_total_for_bronze():
    rules = LoyaltyRules()
    assert rules.calculate_total(100, "bronze") == 95

def test_calculates_total_for_platinum():
    rules = LoyaltyRules()
    assert rules.calculate_total(100, "platinum") == 80

def test_calculates_total_for_unknown_tier():
    rules = LoyaltyRules()
    assert rules.calculate_total(100, "unknown") == 100
