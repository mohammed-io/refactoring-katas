public record CampaignCustomer(bool Active, string Region, bool GdprOptIn);

public class CampaignResult
{
    public int Sent;
    public string Message = "";
    public long Timestamp;
}

public class CampaignSender
{
    public CampaignResult send_campaign(List<CampaignCustomer> customers, string message)
    {
        var count = 0;
        foreach (var customer in customers)
        {
            if (customer.Active)
            {
                if (customer.Region == "EU" && customer.GdprOptIn)
                {
                    count++;
                }
                else if (customer.Region != "EU")
                {
                    count++;
                }
            }
        }

        return new CampaignResult
        {
            Sent = count,
            Message = message,
            Timestamp = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds()
        };
    }
}
