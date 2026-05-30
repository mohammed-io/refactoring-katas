public record MonoItem(decimal Price, int Quantity);
public record MonoCustomer(string Type, string Email, string Country = "", bool TaxExempt = false);
public record MonoShipping(string Speed = "");
public record MonoOrder(int Id, List<MonoItem> Items, MonoCustomer Customer, string Coupon = "", MonoShipping? Shipping = null);

public class MonoEmail
{
    public string To = "";
    public string Subject = "";
    public string Body = "";
}

public class MonoResult
{
    public string? Error;
    public int OrderId;
    public decimal Total;
    public string PaymentStatus = "";
    public string Carrier = "";
    public MonoEmail Email = new();
    public string Log = "";
    public int LoyaltyPoints;
    public decimal TaxRate;
    public decimal ShippingWeight;
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
            if (i.Quantity > 0) y += i.Quantity;
        }

        if (x <= 0)
        {
            return new MonoResult { Error = "Invalid total" };
        }

        decimal d = 0;
        if (order.Customer.Type == "vip") d = x * 0.2m;
        else if (order.Customer.Type == "member") d = x * 0.1m;
        if (x > 100) d += 10;
        if (order.Coupon == "SAVE10") d += x * 0.1m;

        var taxRate = 0.07m;
        if (order.Customer.Country == "EU") taxRate = 0.2m;
        if (order.Customer.TaxExempt) taxRate = 0m;
        var total = x - d + (x - d) * taxRate;
        var status = total > 5000 ? "manual_review" : "approved";
        var carrier = total > 50 ? "UPS" : "USPS";
        if (order.Shipping?.Speed == "express") carrier = "FedEx";
        var email = new MonoEmail
        {
            To = order.Customer.Email,
            Subject = $"Order {order.Id}",
            Body = $"Total: ${total:0.00}"
        };

        var log = $"Order processed at {DateTimeOffset.UtcNow.ToUnixTimeSeconds()}";
        Dictionary<string, object>? cfg = null;
        try
        {
            var json = File.ReadAllText("/tmp/legacy_config.json");
            cfg = System.Text.Json.JsonSerializer.Deserialize<Dictionary<string, object>>(json);
        }
        catch
        {
            cfg = new Dictionary<string, object> { { "fallback", true } };
        }
        if (cfg.TryGetValue("bonusEnabled", out var bonusVal) && bonusVal is System.Text.Json.JsonElement { ValueKind: System.Text.Json.JsonValueKind.True } && order.Customer?.Type == "vip")
        {
            total -= 5;
        }

        return new MonoResult
        {
            OrderId = order.Id,
            Total = Math.Round(total, 2),
            PaymentStatus = status,
            Carrier = carrier,
            Email = email,
            Log = log,
            LoyaltyPoints = (int)(total / 10),
            TaxRate = taxRate,
            ShippingWeight = y
        };
    }
}
