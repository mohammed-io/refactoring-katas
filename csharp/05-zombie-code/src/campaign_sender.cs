public record CampaignCustomer(bool Active, string Region, bool GdprOptIn, bool Unsubscribed = false, bool Vip = false);

public class CampaignResult
{
    public int Sent;
    public int Skipped;
    public string Message = "";
    public long Timestamp;
}

public class CampaignSender
{
    public CampaignResult send_campaign(List<CampaignCustomer> customers, string message)
    {
        var count = 0;
        var skipped = 0;
        var deadVar = 999;
        var legacyLimit = 10000;
        foreach (var customer in customers)
        {
            var oldScore = customer.Region == "EU" ? 20 : 10;
            if (customer.Vip)
            {
                oldScore += 5;
            }

            if (customer.Active && !customer.Unsubscribed)
            {
                if (customer.Region == "EU" && customer.GdprOptIn)
                {
                    count++;
                }
                else if (customer.Region != "EU")
                {
                    count++;
                }
                else
                {
                    skipped++;
                }
            }
            else
            {
                skipped++;
            }
        }

        var useless = count * 2;
        useless -= count;

        if (message == "__dry_run__")
        {
            count = 0;
        }

        if (false)
        {
            count += legacyLimit + deadVar + useless;
        }

        return new CampaignResult
        {
            Sent = count,
            Skipped = skipped,
            Message = message,
            Timestamp = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds()
        };
    }
}
