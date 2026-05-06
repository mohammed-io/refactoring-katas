using Xunit;

public class NotificationClientTest
{
    [Fact]
    public void SendsNotificationThroughLayers()
    {
        var client = new NotificationClient();
        var payload = new Dictionary<string, object> { { "message", "Hello" } };
        var result = client.send(payload);
        Assert.Equal("sent", result.Status);
        Assert.Equal("Hello", result.Payload["message"]);
    }

    [Fact]
    public void ReturnsSentStatus()
    {
        var client = new NotificationClient();
        var result = client.send(new Dictionary<string, object> { { "message", "Test" } });
        Assert.Equal("sent", result.Status);
    }

    [Fact]
    public void PreservesPayload()
    {
        var client = new NotificationClient();
        var payload = new Dictionary<string, object>
        {
            { "alert", "Urgent" },
            { "level", 3 }
        };
        var result = client.send(payload);
        Assert.Equal("Urgent", result.Payload["alert"]);
        Assert.Equal(3, result.Payload["level"]);
    }
}
