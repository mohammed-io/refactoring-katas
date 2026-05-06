public record MonoItem(decimal Price, int Quantity);
public record MonoCustomer(string Type, string Email);
public record MonoOrder(int Id, List<MonoItem> Items, MonoCustomer Customer);

public class MonoResult
{
    public string? Error;
    public int OrderId;
    public decimal Total;
    public string PaymentStatus = "";
    public string Carrier = "";
}

public class LegacySystem
{
    public MonoResult process_everything(MonoOrder order)
    {
        decimal x = 0, y = 0;
        if (order.Items == null || order.Items.Count == 0)
        {
            return new MonoResult { Error = "No items" };
        }

        foreach (var i in order.Items)
        {
            if (i.Price > 0) x += i.Price * i.Quantity;
        }

        decimal d = 0;
        if (order.Customer.Type == "vip") d = x * 0.2m;
        else if (order.Customer.Type == "member") d = x * 0.1m;
        if (x > 100) d += 10;

        var total = x - d + (x - d) * 0.07m;
        var status = total > 5000 ? "manual_review" : "approved";
        var carrier = total > 50 ? "UPS" : "USPS";

        return new MonoResult
        {
            OrderId = order.Id,
            Total = Math.Round(total, 2),
            PaymentStatus = status,
            Carrier = carrier
        };
    }
}
