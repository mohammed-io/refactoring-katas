class Order:
    def __init__(self, items, customer, table, special, rush=False):
        self.items = items
        self.customer = customer
        self.table = table
        self.special = special
        self.rush = rush


class KitchenTicket:
    def __init__(self):
        pass

    def print_ticket(self, order):
        lines = []
        lines.append("Table: " + str(order.table["number"]))
        lines.append("Zone: " + order.table.get("zone", "main"))
        lines.append("Server: " + order.table.get("server", "unassigned"))
        lines.append("Customer: " + order.customer["name"])
        if order.customer.get("vip"):
            lines.append("VIP")
        if order.rush:
            lines.append("RUSH")
        total_items = 0
        for item in order.items:
            total_items += item["qty"]
            line = item["name"] + " x" + str(item["qty"])
            if item.get("modifiers"):
                line += " [" + ", ".join(item["modifiers"]) + "]"
            if item.get("allergy"):
                line += " ALLERGY:" + item["allergy"]
            lines.append(line)
        lines.append("Items: " + str(total_items))
        if order.special and len(order.special) > 0:
            lines.append("Special: " + order.special)
        lines.append("---")
        return "\n".join(lines)
