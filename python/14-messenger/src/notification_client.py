class NotificationClient:
    def send(self, payload): return {"status":"sent", "payload":payload}
class NotificationGateway:
    def __init__(self): self.client = NotificationClient()
    def dispatch(self, payload): return self.client.send(payload)
class NotificationService:
    def __init__(self): self.gateway = NotificationGateway()
    def notify(self, payload): return self.gateway.dispatch(payload)
