public class NotificationResult
{
    public string Status = "";
    public Dictionary<string, object> Payload = new();
}

public class _NotificationClient
{
    public NotificationResult send(Dictionary<string, object> p)
    {
        return new NotificationResult
        {
            Status = "sent",
            Payload = p
        };
    }
}

public class NotificationGateway
{
    private _NotificationClient _client = new();

    public NotificationResult dispatch(Dictionary<string, object> p)
    {
        return _client.send(p);
    }
}

public class NotificationClient
{
    private NotificationGateway _gateway = new();

    public NotificationResult send(Dictionary<string, object> p)
    {
        return _gateway.dispatch(p);
    }
}
