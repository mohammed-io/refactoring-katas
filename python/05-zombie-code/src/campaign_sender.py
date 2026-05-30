import time

class CampaignSender:
    def __init__(self):
        pass

    def send_campaign(self, customers, message):
        sent_count = 0
        skipped_count = 0
        dead_var = 999
        legacy_limit = 10000

        for customer in customers:
            old_score = 20 if customer.get("region") == "EU" else 10
            if customer.get("vip"):
                old_score += 5

            if not customer.get("active") or customer.get("unsubscribed"):
                skipped_count += 1
                continue

            if customer.get("region") == "EU":
                if customer.get("gdprOptIn"):
                    sent_count += 1
                else:
                    skipped_count += 1
            else:
                sent_count += 1

        useless = sent_count * 2
        useless -= sent_count

        if message == "__dry_run__":
            sent_count = 0

        if False:
            sent_count += legacy_limit + dead_var

        return {
            "sent": sent_count,
            "skipped": skipped_count,
            "message": message,
            "timestamp": int(time.time() * 1000)
        }
