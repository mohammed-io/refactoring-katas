from src.notification_client import NotificationClient

def test_sends_notification_through_layers():
    client = NotificationClient()
    result = client.send({"recipient": "user-1", "message": "Hello", "channel": "sms"})
    assert result["status"] == "sent"
    assert result["delivery_id"] == "sms-user-1"
    assert result["payload"]["message"] == "Hello"

def test_defaults_channel_and_priority():
    client = NotificationClient()
    result = client.send({"message": "Test"})
    assert result["status"] == "sent"
    assert result["payload"]["channel"] == "email"
    assert result["payload"]["priority"] == "normal"

def test_preserves_explicit_priority():
    client = NotificationClient()
    result = client.send({"recipient": "ops", "message": "Urgent", "priority": "high"})
    assert result["payload"]["priority"] == "high"
    assert result["payload"]["recipient"] == "ops"

def test_rejects_missing_message():
    client = NotificationClient()
    result = client.send({"recipient": "ops"})
    assert result["status"] == "rejected"
    assert result["reason"] == "missing_message"

def test_records_observable_audit_events():
    client = NotificationClient()
    result = client.send({"message": "Deploy complete", "channel": "push"})
    assert result["audit"] == ["queued:push", "sent:push"]

def test_reports_failed_delivery_for_unsupported_channel():
    client = NotificationClient()
    result = client.send({"recipient": "ops", "message": "Legacy alert", "channel": "fax"})
    assert result["status"] == "failed"
    assert result["reason"] == "unsupported_channel"
    assert result["audit"] == ["queued:fax", "failed:fax"]

def test_high_priority_failed_delivery_is_scheduled_for_retry():
    client = NotificationClient()
    result = client.send({"recipient": "ops", "message": "Legacy alert", "channel": "fax", "priority": "high"})
    assert result["status"] == "retrying"
    assert result["reason"] == "unsupported_channel"
    assert result["audit"] == ["queued:fax", "failed:fax", "retry_scheduled:fax"]
