public record LineItem(decimal Price, int Quantity);
public record Customer(string Email);
public record Address(string Zip);
public record Order(List<LineItem> Items, Customer Customer, Address Address);

public class OrderResult
{
    public string? Error;
    public decimal Total;
    public string Carrier = "";
    public string PaymentStatus = "";
    public string EmailTo = "";
    public string EmailSubject = "";
}

public class OrderProcessor
{
    public OrderResult process_order(Order order)
    {
        if (order.Items == null || order.Items.Count == 0)
        {
            return new OrderResult
            {
                Error = "No items"
            };
        }

        if (order.Customer == null || string.IsNullOrEmpty(order.Customer.Email))
        {
            return new OrderResult
            {
                Error = "Invalid customer"
            };
        }

        if (order.Address == null || string.IsNullOrEmpty(order.Address.Zip))
        {
            return new OrderResult
            {
                Error = "Invalid address"
            };
        }

        var inventory = true;
        foreach (var item in order.Items)
        {
            if (item.Quantity > 100)
            {
                inventory = false;
            }
        }

        if (!inventory)
        {
            return new OrderResult
            {
                Error = "Out of stock"
            };
        }

        decimal subtotal = 0;
        var weight = 0;
        foreach (var item in order.Items)
        {
            subtotal += item.Price * item.Quantity;
            weight += item.Quantity;
        }

        decimal shipping = 0;
        if (subtotal < 25)
        {
            shipping = 5.99m;
        }
        else if (subtotal < 50)
        {
            shipping = 3.99m;
        }

        var total = subtotal + subtotal * 0.07m + shipping;
        var carrier = total > 100 ? "UPS" : "USPS";
        var status = total > 1000 ? "pending_review" : "approved";

        return new OrderResult
        {
            Total = Math.Round(total, 2),
            Carrier = carrier,
            PaymentStatus = status,
            EmailTo = order.Customer.Email,
            EmailSubject = "Order Confirmation"
        };
    }
}
