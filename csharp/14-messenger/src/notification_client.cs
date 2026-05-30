public class NotificationResult
{
    public string Status = "";
    public string Reason = "";
    public string DeliveryId = "";
    public Dictionary<string, object> Payload = new();
    public List<string> Audit = new();
}

public class NotificationBackend
{
    public NotificationResult send(Dictionary<string, object> p)
    {
        if (p["channel"].ToString() == "fax")
        {
            return new NotificationResult
            {
                Status = "failed",
                Reason = "unsupported_channel",
                Payload = p
            };
        }

        return new NotificationResult
        {
            Status = "sent",
            DeliveryId = $"{p["channel"]}-{p["recipient"]}",
            Payload = p
        };
    }
}

public class NotificationGateway
{
    private NotificationBackend _client = new();

    public NotificationResult dispatch(Dictionary<string, object> p)
    {
        return _client.send(p);
    }
}

public class NotificationAudit
{
    public string record(string eventName, Dictionary<string, object> payload)
    {
        return $"{eventName}:{payload["channel"]}";
    }
}

public class NotificationClient
{
    private NotificationGateway _gateway = new();
    private NotificationAudit _audit = new();

    public NotificationResult send(Dictionary<string, object> p)
    {
        if (!p.ContainsKey("message") || p["message"].ToString() == "")
        {
            return new NotificationResult
            {
                Status = "rejected",
                Reason = "missing_message",
                Payload = p
            };
        }

        var normalized = new Dictionary<string, object>
        {
            ["recipient"] = p.ContainsKey("recipient") ? p["recipient"] : "unknown",
            ["message"] = p["message"],
            ["channel"] = p.ContainsKey("channel") ? p["channel"] : "email",
            ["priority"] = p.ContainsKey("priority") ? p["priority"] : "normal"
        };
        var result = _gateway.dispatch(normalized);
        result.Audit = new List<string> { _audit.record("queued", normalized), _audit.record(result.Status, normalized) };
        if (result.Status == "failed" && normalized["priority"].ToString() == "high")
        {
            result.Status = "retrying";
            result.Audit.Add(_audit.record("retry_scheduled", normalized));
        }
        return result;
    }
}
