import time, random

class OrderProcessor:
    def __init__(self):
        pass

    def process_order(self, order):
        if not order.get("items"):
            return {"error": "No items"}
        if not order.get("customer") or not order["customer"].get("email"):
            return {"error": "Invalid customer"}
        if not order.get("address") or not order["address"].get("zip"):
            return {"error": "Invalid address"}
        inventory = True
        for item in order["items"]:
            if item.get("quantity", 0) > 100:
                inventory = False
        if not inventory:
            return {"error": "Out of stock"}
        subtotal = 0
        for item in order["items"]:
            subtotal += item["price"] * item["quantity"]
        shipping = 0
        if subtotal < 25:
            shipping = 5.99
        elif subtotal < 50:
            shipping = 3.99
        total = subtotal + subtotal * 0.07 + shipping
        label = {"to": order["address"], "weight": sum(i["quantity"] for i in order["items"]), "carrier": "USPS"}
        if total > 100:
            label["carrier"] = "UPS"
        return {"orderId": "ORD-" + str(int(time.time() * 1000)), "paymentId": "PAY-" + str(random.randint(0, 999999)), "paymentStatus": "pending_review" if total > 1000 else "approved", "total": round(total, 2), "shippingLabel": label, "email": {"to": order["customer"]["email"], "subject": "Order Confirmation", "body": f"Your order total is ${total:.2f}"}}
