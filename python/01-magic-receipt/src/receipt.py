class Receipt:
    def __init__(self):
        pass

    def calculate_total(self, items, customer_type=None):
        total = 0
        for price in items:
            total += price
        discount = 0
        if customer_type == "member":
            discount = total * 0.05
        elif customer_type == "vip":
            discount = total * 0.15
        if total > 50:
            discount += 5
        final = (total - discount) * 1.08
        if customer_type == "vip":
            final -= 2
        return round(final, 2)
