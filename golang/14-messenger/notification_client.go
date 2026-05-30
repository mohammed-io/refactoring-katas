package kata

type NotifyResult struct {
	Status     string
	Reason     string
	DeliveryID string
	Payload    map[string]any
	Audit      []string
}
type NotificationBackend struct{}

func (b NotificationBackend) send(p map[string]any) NotifyResult {
	if p["channel"] == "fax" {
		return NotifyResult{Status: "failed", Reason: "unsupported_channel", Payload: p}
	}
	return NotifyResult{Status: "sent", DeliveryID: p["channel"].(string) + "-" + p["recipient"].(string), Payload: p}
}

type NotificationGateway struct {
	client NotificationBackend
}

func (g NotificationGateway) dispatch(p map[string]any) NotifyResult {
	return g.client.send(p)
}

type NotificationAudit struct{}

func (a NotificationAudit) record(event string, p map[string]any) string {
	return event + ":" + p["channel"].(string)
}

type NotificationClient struct {
	gateway NotificationGateway
	audit   NotificationAudit
}

func NewNotificationClient() NotificationClient {
	return NotificationClient{NotificationGateway{NotificationBackend{}}, NotificationAudit{}}
}

func (c NotificationClient) send(p map[string]any) NotifyResult {
	message, ok := p["message"].(string)
	if !ok || message == "" {
		return NotifyResult{Status: "rejected", Reason: "missing_message", Payload: p}
	}

	normalized := map[string]any{
		"recipient": "unknown",
		"message":   message,
		"channel":   "email",
		"priority":  "normal",
	}
	for key, value := range p {
		normalized[key] = value
	}
	result := c.gateway.dispatch(normalized)
	result.Audit = []string{c.audit.record("queued", normalized), c.audit.record(result.Status, normalized)}
	if result.Status == "failed" && normalized["priority"] == "high" {
		result.Status = "retrying"
		result.Audit = append(result.Audit, c.audit.record("retry_scheduled", normalized))
	}
	return result
}
