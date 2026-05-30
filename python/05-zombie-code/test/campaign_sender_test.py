from src.campaign_sender import CampaignSender

def test_eu_opt_in():
    sender = CampaignSender()
    result = sender.send_campaign([{"active": True, "region": "EU", "gdprOptIn": True}], "Hi")
    assert result["sent"] == 1

def test_eu_no_opt_in():
    sender = CampaignSender()
    result = sender.send_campaign([{"active": True, "region": "EU", "gdprOptIn": False}], "Hi")
    assert result["sent"] == 0

def test_non_eu_active():
    sender = CampaignSender()
    result = sender.send_campaign([{"active": True, "region": "US"}], "Hi")
    assert result["sent"] == 1

def test_inactive_customer():
    sender = CampaignSender()
    result = sender.send_campaign([{"active": False, "region": "US"}], "Hi")
    assert result["sent"] == 0
    assert result["skipped"] == 1

def test_mixed_customers():
    sender = CampaignSender()
    customers = [
        {"active": True, "region": "EU", "gdprOptIn": True},
        {"active": True, "region": "EU", "gdprOptIn": False},
        {"active": True, "region": "US"},
        {"active": False, "region": "US"}
    ]
    result = sender.send_campaign(customers, "Hi")
    assert result["sent"] == 2
    assert result["skipped"] == 2

def test_unsubscribed_customer():
    sender = CampaignSender()
    result = sender.send_campaign([{"active": True, "region": "US", "unsubscribed": True}], "Hi")
    assert result["sent"] == 0
    assert result["skipped"] == 1

def test_dry_run_never_sends_but_keeps_skip_count():
    sender = CampaignSender()
    result = sender.send_campaign([
        {"active": True, "region": "US"},
        {"active": False, "region": "US"}
    ], "__dry_run__")
    assert result["sent"] == 0
    assert result["skipped"] == 1

def test_message_in_result():
    sender = CampaignSender()
    result = sender.send_campaign([], "Hello World")
    assert result["message"] == "Hello World"

def test_timestamp_is_number():
    sender = CampaignSender()
    result = sender.send_campaign([], "Hi")
    assert isinstance(result["timestamp"], int)
