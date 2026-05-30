import time, json

class LegacySystem:
    def __init__(self):
        pass

    def process_everything(self, order):
        x = 0; y = 0
        if not order.get("items"): return {"error":"No items"}
        for item in order["items"]:
            if item["price"] > 0: x += item["price"] * item["quantity"]
            if item["quantity"] > 0: y += item["quantity"]
        if x <= 0: return {"error":"Invalid total"}
        d = 0
        if order.get("customer") and order["customer"].get("type") == "vip": d = x * 0.2
        elif order.get("customer") and order["customer"].get("type") == "member": d = x * 0.1
        if x > 100: d += 10
        if order.get("coupon") == "SAVE10": d += x * 0.1
        taxable = x - d
        tax_rate = 0.07
        if order.get("customer") and order["customer"].get("country") == "EU": tax_rate = 0.2
        if order.get("customer") and order["customer"].get("taxExempt"): tax_rate = 0
        total = taxable + taxable * tax_rate
        payment = {"status":"approved"}
        if total > 5000: payment["status"] = "manual_review"
        ship = {"carrier":"USPS"}
        if total > 50: ship["carrier"] = "UPS"
        if order.get("shipping", {}).get("speed") == "express": ship["carrier"] = "FedEx"
        ship["weight"] = y
        email = {"to": order.get("customer",{}).get("email",""), "subject":"Order " + str(order["id"]), "body": f"Total: ${total:.2f}"}
        log = "Order processed at " + str(time.time())
        try:
            cfg = json.load(open('/tmp/legacy_config.json'))
        except Exception:
            cfg = {"fallback": True}
        if cfg.get("bonusEnabled") and order.get("customer") and order["customer"].get("type") == "vip": total -= 5
        loyalty = {"points": int(total / 10)}
        return {"orderId":order["id"], "total":round(total,2), "paymentStatus":payment["status"], "carrier":ship["carrier"], "email":email, "log":log, "loyaltyPoints":loyalty["points"], "taxRate":tax_rate, "shippingWeight":ship["weight"]}
