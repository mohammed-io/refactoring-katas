package kata

import (
	"maps"
	"testing"
)

func TestNotificationClientSendsNotificationThroughLayers(t *testing.T) {
	client := NewNotificationClient()
	payload := map[string]any{"message": "Hello"}
	result := client.send(payload)
	if result.Status != "sent" {
		t.Errorf("expected status 'sent', got %q", result.Status)
	}
	if !maps.Equal(result.Payload, map[string]any{"message": "Hello"}) {
		t.Errorf("expected payload {message: Hello}, got %v", result.Payload)
	}
}

func TestNotificationClientReturnsSentStatus(t *testing.T) {
	client := NewNotificationClient()
	result := client.send(map[string]any{"message": "Test"})
	if result.Status != "sent" {
		t.Errorf("expected status 'sent', got %q", result.Status)
	}
}

func TestNotificationClientPreservesPayload(t *testing.T) {
	client := NewNotificationClient()
	payload := map[string]any{"alert": "Urgent", "level": 3}
	result := client.send(payload)
	if !maps.Equal(result.Payload, payload) {
		t.Errorf("expected payload to be preserved, got %v, want %v", result.Payload, payload)
	}
}
