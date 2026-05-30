package kata

import "testing"

func TestSendCampaignCountsEUCustomersWithOptIn(t *testing.T) {
	cs := NewCampaignSender()
	customers := []CampaignCustomer{{Active: true, Region: "EU", GDPR: true}}
	result := cs.send_campaign(customers, "Hi")
	if result.Sent != 1 {
		t.Errorf("expected sent 1, got %d", result.Sent)
	}
}

func TestSendCampaignSkipsEUCustomersWithoutOptIn(t *testing.T) {
	cs := NewCampaignSender()
	customers := []CampaignCustomer{{Active: true, Region: "EU", GDPR: false}}
	result := cs.send_campaign(customers, "Hi")
	if result.Sent != 0 {
		t.Errorf("expected sent 0, got %d", result.Sent)
	}
	if result.Skipped != 1 {
		t.Errorf("expected skipped 1, got %d", result.Skipped)
	}
}

func TestSendCampaignCountsNonEUActiveCustomers(t *testing.T) {
	cs := NewCampaignSender()
	customers := []CampaignCustomer{{Active: true, Region: "US"}}
	result := cs.send_campaign(customers, "Hi")
	if result.Sent != 1 {
		t.Errorf("expected sent 1, got %d", result.Sent)
	}
}

func TestSendCampaignSkipsInactiveCustomers(t *testing.T) {
	cs := NewCampaignSender()
	customers := []CampaignCustomer{{Active: false, Region: "US"}}
	result := cs.send_campaign(customers, "Hi")
	if result.Sent != 0 {
		t.Errorf("expected sent 0, got %d", result.Sent)
	}
}

func TestSendCampaignHandlesMixedCustomers(t *testing.T) {
	cs := NewCampaignSender()
	customers := []CampaignCustomer{
		{Active: true, Region: "EU", GDPR: true},
		{Active: true, Region: "EU", GDPR: false},
		{Active: true, Region: "US"},
		{Active: false, Region: "US"},
	}
	result := cs.send_campaign(customers, "Hi")
	if result.Sent != 2 {
		t.Errorf("expected sent 2, got %d", result.Sent)
	}
	if result.Skipped != 2 {
		t.Errorf("expected skipped 2, got %d", result.Skipped)
	}
}

func TestSendCampaignSkipsUnsubscribedCustomers(t *testing.T) {
	cs := NewCampaignSender()
	customers := []CampaignCustomer{{Active: true, Region: "US", Unsubscribed: true}}
	result := cs.send_campaign(customers, "Hi")
	if result.Sent != 0 {
		t.Errorf("expected sent 0, got %d", result.Sent)
	}
	if result.Skipped != 1 {
		t.Errorf("expected skipped 1, got %d", result.Skipped)
	}
}

func TestSendCampaignDryRunNeverSendsButKeepsSkipCount(t *testing.T) {
	cs := NewCampaignSender()
	customers := []CampaignCustomer{
		{Active: true, Region: "US"},
		{Active: false, Region: "US"},
	}
	result := cs.send_campaign(customers, "__dry_run__")
	if result.Sent != 0 {
		t.Errorf("expected sent 0, got %d", result.Sent)
	}
	if result.Skipped != 1 {
		t.Errorf("expected skipped 1, got %d", result.Skipped)
	}
}

func TestSendCampaignReturnsMessageInResult(t *testing.T) {
	cs := NewCampaignSender()
	result := cs.send_campaign([]CampaignCustomer{}, "Hello World")
	if result.Message != "Hello World" {
		t.Errorf("expected message 'Hello World', got %q", result.Message)
	}
}

func TestSendCampaignReturnsTimestampInResult(t *testing.T) {
	cs := NewCampaignSender()
	result := cs.send_campaign([]CampaignCustomer{}, "Hi")
	if result.Timestamp == 0 {
		t.Error("expected non-zero timestamp")
	}
}
