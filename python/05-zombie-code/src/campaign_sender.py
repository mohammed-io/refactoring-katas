import time

class CampaignSender:
    def __init__(self):
        pass

    def send_campaign(self, customers, message):
        sent_count = 0

        for customer in customers:
            if not customer.get("active"):
                continue

            if customer.get("region") == "EU":
                if customer.get("gdprOptIn"):
                    sent_count += 1
            else:
                sent_count += 1

        return {
            "sent": sent_count,
            "message": message,
            "timestamp": int(time.time() * 1000)
        }
