package kata

import (
	"reflect"
	"testing"
)

func TestNotificationClientSendsNotificationThroughLayers(t *testing.T) {
	client := NewNotificationClient()
	payload := map[string]any{"recipient": "user-1", "message": "Hello", "channel": "sms"}
	result := client.send(payload)
	if result.Status != "sent" {
		t.Errorf("expected status 'sent', got %q", result.Status)
	}
	if result.DeliveryID != "sms-user-1" {
		t.Errorf("expected delivery id sms-user-1, got %q", result.DeliveryID)
	}
	if result.Payload["message"] != "Hello" {
		t.Errorf("expected message to be preserved")
	}
}

func TestNotificationClientDefaultsChannelAndPriority(t *testing.T) {
	client := NewNotificationClient()
	result := client.send(map[string]any{"message": "Test"})
	if result.Status != "sent" {
		t.Errorf("expected status 'sent', got %q", result.Status)
	}
	if result.Payload["channel"] != "email" {
		t.Errorf("expected default channel")
	}
	if result.Payload["priority"] != "normal" {
		t.Errorf("expected default priority")
	}
}

func TestNotificationClientPreservesExplicitPriority(t *testing.T) {
	client := NewNotificationClient()
	result := client.send(map[string]any{"recipient": "ops", "message": "Urgent", "priority": "high"})
	if result.Payload["priority"] != "high" {
		t.Errorf("expected explicit priority")
	}
	if result.Payload["recipient"] != "ops" {
		t.Errorf("expected explicit recipient")
	}
}

func TestNotificationClientRejectsMissingMessage(t *testing.T) {
	client := NewNotificationClient()
	result := client.send(map[string]any{"recipient": "ops"})
	if result.Status != "rejected" {
		t.Errorf("expected rejected status")
	}
	if result.Reason != "missing_message" {
		t.Errorf("expected missing message reason")
	}
}

func TestNotificationClientRecordsObservableAuditEvents(t *testing.T) {
	client := NewNotificationClient()
	result := client.send(map[string]any{"message": "Deploy complete", "channel": "push"})
	if !reflect.DeepEqual(result.Audit, []string{"queued:push", "sent:push"}) {
		t.Errorf("expected audit events, got %#v", result.Audit)
	}
}

func TestNotificationClientReportsFailedDeliveryForUnsupportedChannel(t *testing.T) {
	client := NewNotificationClient()
	result := client.send(map[string]any{"recipient": "ops", "message": "Legacy alert", "channel": "fax"})
	if result.Status != "failed" {
		t.Errorf("expected failed status, got %q", result.Status)
	}
	if result.Reason != "unsupported_channel" {
		t.Errorf("expected unsupported channel reason")
	}
	if !reflect.DeepEqual(result.Audit, []string{"queued:fax", "failed:fax"}) {
		t.Errorf("expected failed audit events, got %#v", result.Audit)
	}
}

func TestNotificationClientHighPriorityFailedDeliveryIsScheduledForRetry(t *testing.T) {
	client := NewNotificationClient()
	result := client.send(map[string]any{"recipient": "ops", "message": "Legacy alert", "channel": "fax", "priority": "high"})
	if result.Status != "retrying" {
		t.Errorf("expected retrying status, got %q", result.Status)
	}
	if result.Reason != "unsupported_channel" {
		t.Errorf("expected unsupported channel reason")
	}
	if !reflect.DeepEqual(result.Audit, []string{"queued:fax", "failed:fax", "retry_scheduled:fax"}) {
		t.Errorf("expected retry audit events, got %#v", result.Audit)
	}
}
