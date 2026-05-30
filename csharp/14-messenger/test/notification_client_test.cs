using Xunit;

public class NotificationClientTest
{
    [Fact]
    public void SendsNotificationThroughLayers()
    {
        var client = new NotificationClient();
        var payload = new Dictionary<string, object> { { "recipient", "user-1" }, { "message", "Hello" }, { "channel", "sms" } };
        var result = client.send(payload);
        Assert.Equal("sent", result.Status);
        Assert.Equal("sms-user-1", result.DeliveryId);
        Assert.Equal("Hello", result.Payload["message"]);
    }

    [Fact]
    public void DefaultsChannelAndPriority()
    {
        var client = new NotificationClient();
        var result = client.send(new Dictionary<string, object> { { "message", "Test" } });
        Assert.Equal("sent", result.Status);
        Assert.Equal("email", result.Payload["channel"]);
        Assert.Equal("normal", result.Payload["priority"]);
    }

    [Fact]
    public void PreservesExplicitPriority()
    {
        var client = new NotificationClient();
        var payload = new Dictionary<string, object> { { "recipient", "ops" }, { "message", "Urgent" }, { "priority", "high" } };
        var result = client.send(payload);
        Assert.Equal("high", result.Payload["priority"]);
        Assert.Equal("ops", result.Payload["recipient"]);
    }

    [Fact]
    public void RejectsMissingMessage()
    {
        var client = new NotificationClient();
        var result = client.send(new Dictionary<string, object> { { "recipient", "ops" } });
        Assert.Equal("rejected", result.Status);
        Assert.Equal("missing_message", result.Reason);
    }

    [Fact]
    public void RecordsObservableAuditEvents()
    {
        var client = new NotificationClient();
        var result = client.send(new Dictionary<string, object> { { "message", "Deploy complete" }, { "channel", "push" } });
        Assert.Equal(new List<string> { "queued:push", "sent:push" }, result.Audit);
    }

    [Fact]
    public void ReportsFailedDeliveryForUnsupportedChannel()
    {
        var client = new NotificationClient();
        var result = client.send(new Dictionary<string, object> { { "recipient", "ops" }, { "message", "Legacy alert" }, { "channel", "fax" } });
        Assert.Equal("failed", result.Status);
        Assert.Equal("unsupported_channel", result.Reason);
        Assert.Equal(new List<string> { "queued:fax", "failed:fax" }, result.Audit);
    }

    [Fact]
    public void HighPriorityFailedDeliveryIsScheduledForRetry()
    {
        var client = new NotificationClient();
        var result = client.send(new Dictionary<string, object> { { "recipient", "ops" }, { "message", "Legacy alert" }, { "channel", "fax" }, { "priority", "high" } });
        Assert.Equal("retrying", result.Status);
        Assert.Equal("unsupported_channel", result.Reason);
        Assert.Equal(new List<string> { "queued:fax", "failed:fax", "retry_scheduled:fax" }, result.Audit);
    }
}
