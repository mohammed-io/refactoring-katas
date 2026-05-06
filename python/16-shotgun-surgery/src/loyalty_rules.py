class LoyaltyRules:
    def __init__(self):
        pass

    def get_discount_for_tier(self, tier):
        if tier == "bronze": return 0.05
        elif tier == "silver": return 0.1
        elif tier == "gold": return 0.15
        elif tier == "platinum": return 0.2
        return 0

    def get_label_for_tier(self, tier):
        if tier == "bronze": return "Bronze Member"
        elif tier == "silver": return "Silver Member"
        elif tier == "gold": return "Gold Member"
        elif tier == "platinum": return "Platinum Member"
        return "Standard"

    def get_threshold_for_tier(self, tier):
        if tier == "bronze": return 100
        elif tier == "silver": return 500
        elif tier == "gold": return 2000
        elif tier == "platinum": return 10000
        return 0

    def get_color_for_tier(self, tier):
        if tier == "bronze": return "#CD7F32"
        elif tier == "silver": return "#C0C0C0"
        elif tier == "gold": return "#FFD700"
        elif tier == "platinum": return "#E5E4E2"
        return "#000000"

    def calculate_total(self, spending, tier):
        return spending * (1 - self.get_discount_for_tier(tier))
