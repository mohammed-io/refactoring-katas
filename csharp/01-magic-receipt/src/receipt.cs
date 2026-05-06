public class Receipt
{
    public decimal calculate_total(decimal[] items, string customerType = "")
    {
        decimal total = 0;
        for (int i = 0; i < items.Length; i++)
        {
            total += items[i];
        }

        decimal discount = 0;
        if (customerType == "member")
        {
            discount = total * 0.05m;
        }
        else if (customerType == "vip")
        {
            discount = total * 0.15m;
        }

        if (total > 50)
        {
            discount += 5;
        }

        var final = (total - discount) * 1.08m;
        if (customerType == "vip")
        {
            final -= 2;
        }

        return Math.Round(final, 2);
    }
}
