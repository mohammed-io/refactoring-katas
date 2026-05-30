class NotificationBackend:
    def send(self, payload):
        if payload["channel"] == "fax":
            return {"status": "failed", "reason": "unsupported_channel", "payload": payload}
        return {"status": "sent", "delivery_id": f"{payload['channel']}-{payload['recipient']}", "payload": payload}


class NotificationGateway:
    def __init__(self):
        self.client = NotificationBackend()

    def dispatch(self, payload):
        return self.client.send(payload)


class NotificationAudit:
    def record(self, event, payload):
        return f"{event}:{payload['channel']}"


class NotificationClient:
    def __init__(self):
        self.gateway = NotificationGateway()
        self.audit = NotificationAudit()

    def send(self, payload):
        if not payload.get("message"):
            return {"status": "rejected", "reason": "missing_message", "payload": payload}

        normalized = {
            "recipient": payload.get("recipient", "unknown"),
            "message": payload["message"],
            "channel": payload.get("channel", "email"),
            "priority": payload.get("priority", "normal"),
        }
        result = self.gateway.dispatch(normalized)
        result["audit"] = [self.audit.record("queued", normalized), self.audit.record(result["status"], normalized)]
        if result["status"] == "failed" and normalized["priority"] == "high":
            result["status"] = "retrying"
            result["audit"].append(self.audit.record("retry_scheduled", normalized))
        return result
