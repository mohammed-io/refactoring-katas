using Xunit;

public class KitchenTicketTest
{
    private KitchenOrder Order(
        List<TicketItem>? items = null,
        Dictionary<string, object>? customer = null,
        Dictionary<string, object>? table = null,
        string special = "",
        bool rush = false)
    {
        return new KitchenOrder(
            items ?? new List<TicketItem> { new TicketItem("Burger", 1) },
            customer ?? new Dictionary<string, object> { { "name", "Alice" }, { "vip", false } },
            table ?? new Dictionary<string, object> { { "number", 5 }, { "zone", "patio" }, { "server", "Sam" } },
            special,
            rush);
    }

    [Fact]
    public void PrintsTableCustomerAndServerDetails()
    {
        var result = new KitchenTicket().print_ticket(Order());
        Assert.Contains("Table: 5", result);
        Assert.Contains("Zone: patio", result);
        Assert.Contains("Server: Sam", result);
        Assert.Contains("Customer: Alice", result);
    }

    [Fact]
    public void PrintsTicketWithMultipleItemsAndCount()
    {
        var result = new KitchenTicket().print_ticket(Order(items: new List<TicketItem> { new TicketItem("Burger", 2), new TicketItem("Fries", 1) }));
        Assert.Contains("Burger x2", result);
        Assert.Contains("Fries x1", result);
        Assert.Contains("Items: 3", result);
    }

    [Fact]
    public void PrintsModifiersAndAllergyFlags()
    {
        var result = new KitchenTicket().print_ticket(Order(items: new List<TicketItem> { new TicketItem("Salad", 1, new List<string> { "no onion", "dressing side" }, "nuts") }));
        Assert.Contains("Salad x1 [no onion, dressing side] ALLERGY:nuts", result);
    }

    [Fact]
    public void PrintsVipAndRushMarkers()
    {
        var result = new KitchenTicket().print_ticket(Order(customer: new Dictionary<string, object> { { "name", "Carol" }, { "vip", true } }, rush: true));
        Assert.Contains("VIP", result);
        Assert.Contains("RUSH", result);
    }

    [Fact]
    public void OmitsSpecialWhenEmptyButKeepsSeparator()
    {
        var result = new KitchenTicket().print_ticket(Order(special: ""));
        Assert.DoesNotContain("Special:", result);
        Assert.Contains("---", result);
    }

    [Fact]
    public void PrintsSpecialInstructions()
    {
        var result = new KitchenTicket().print_ticket(Order(special: "Fire mains after starters"));
        Assert.Contains("Special: Fire mains after starters", result);
    }
}
