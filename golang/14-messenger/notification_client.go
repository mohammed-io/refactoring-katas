package kata

type NotifyResult struct {
	Status  string
	Payload map[string]any
}
type NotificationClient struct{}

func NewNotificationClient() NotificationClient {
	return NotificationClient{}
}
func (c NotificationClient) send(p map[string]any) NotifyResult { return NotifyResult{"sent", p} }

// Note: original Go had gateway/service layers. Flattened to NotificationClient per canonical.
