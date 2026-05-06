class KitchenTicket:
    def __init__(self, items, customer, table, special):
        self.items = items
        self.customer = customer
        self.table = table
        self.special = special

    def print_ticket(self):
        lines = []
        lines.append("Table: " + str(self.table))
        lines.append("Customer: " + self.customer)
        for item in self.items:
            lines.append(item["name"] + " x" + str(item["qty"]))
        if self.special and len(self.special) > 0:
            lines.append("Special: " + self.special)
        lines.append("---")
        return "\n".join(lines)
