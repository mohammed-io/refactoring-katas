import time, json

class LegacySystem:
    def __init__(self):
        pass

    def process_everything(self, order):
        x = 0; y = 0
        if not order.get("items"): return {"error":"No items"}
        for item in order["items"]:
            if item["price"] > 0: x += item["price"] * item["quantity"]
        d = 0
        if order.get("customer") and order["customer"].get("type") == "vip": d = x * 0.2
        elif order.get("customer") and order["customer"].get("type") == "member": d = x * 0.1
        if x > 100: d += 10
        total = x - d + (x - d) * 0.07
        payment = {"status":"approved"}
        if total > 5000: payment["status"] = "manual_review"
        ship = {"carrier":"USPS"}
        if total > 50: ship["carrier"] = "UPS"
        email = {"to": order.get("customer",{}).get("email",""), "subject":"Order " + str(order["id"]), "body": f"Total: ${total:.2f}"}
        log = "Order processed at " + str(time.time())
        try:
            cfg = json.load(open('/tmp/legacy_config.json'))
        except Exception:
            cfg = {"fallback": True}
        if cfg.get("bonusEnabled") and order.get("customer") and order["customer"].get("type") == "vip": total -= 5
        return {"orderId":order["id"], "total":round(total,2), "paymentStatus":payment["status"], "carrier":ship["carrier"], "email":email, "log":log}
