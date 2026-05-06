using Xunit;

public class CampaignSenderTest
{
    [Fact]
    public void CountsEuCustomersWithOptIn()
    {
        var estimator = new CampaignSender();
        var result = estimator.send_campaign(new List<CampaignCustomer>
        {
            new CampaignCustomer(true, "EU", true)
        }, "Hi");
        Assert.Equal(1, result.Sent);
    }

    [Fact]
    public void SkipsEuCustomersWithoutOptIn()
    {
        var estimator = new CampaignSender();
        var result = estimator.send_campaign(new List<CampaignCustomer>
        {
            new CampaignCustomer(true, "EU", false)
        }, "Hi");
        Assert.Equal(0, result.Sent);
    }

    [Fact]
    public void CountsNonEuActiveCustomers()
    {
        var estimator = new CampaignSender();
        var result = estimator.send_campaign(new List<CampaignCustomer>
        {
            new CampaignCustomer(true, "US", false)
        }, "Hi");
        Assert.Equal(1, result.Sent);
    }

    [Fact]
    public void SkipsInactiveCustomers()
    {
        var estimator = new CampaignSender();
        var result = estimator.send_campaign(new List<CampaignCustomer>
        {
            new CampaignCustomer(false, "US", false)
        }, "Hi");
        Assert.Equal(0, result.Sent);
    }

    [Fact]
    public void HandlesMixedCustomers()
    {
        var estimator = new CampaignSender();
        var result = estimator.send_campaign(new List<CampaignCustomer>
        {
            new CampaignCustomer(true, "EU", true),
            new CampaignCustomer(true, "EU", false),
            new CampaignCustomer(true, "US", false),
            new CampaignCustomer(false, "US", false)
        }, "Hi");
        Assert.Equal(2, result.Sent);
    }

    [Fact]
    public void ReturnsMessageInResult()
    {
        var estimator = new CampaignSender();
        var result = estimator.send_campaign(new List<CampaignCustomer>(), "Hello World");
        Assert.Equal("Hello World", result.Message);
    }

    [Fact]
    public void ReturnsTimestampInResult()
    {
        var estimator = new CampaignSender();
        var result = estimator.send_campaign(new List<CampaignCustomer>(), "Hi");
        Assert.True(result.Timestamp > 0);
    }
}
