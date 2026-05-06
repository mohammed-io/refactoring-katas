from src.notification_client import NotificationService

def test_sends_notification_through_layers():
    service = NotificationService()
    result = service.notify({"message": "Hello"})
    assert result["status"] == "sent"
    assert result["payload"] == {"message": "Hello"}

def test_returns_sent_status():
    service = NotificationService()
    result = service.notify({"message": "Test"})
    assert result["status"] == "sent"

def test_preserves_payload():
    service = NotificationService()
    payload = {"alert": "Urgent", "level": 3}
    result = service.notify(payload)
    assert result["payload"] == payload
