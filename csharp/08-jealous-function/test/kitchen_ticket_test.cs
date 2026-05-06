using Xunit;

public class KitchenTicketTest
{
    [Fact]
    public void PrintsSimpleTicket()
    {
        var order = new KitchenOrder(
            new List<TicketItem> { new TicketItem("Burger", 1) },
            "Alice", 5, "");
        var result = new KitchenTicket().print_ticket(order);
        Assert.Contains("Table: 5", result);
        Assert.Contains("Customer: Alice", result);
        Assert.Contains("Burger x1", result);
    }

    [Fact]
    public void PrintsTicketWithMultipleItems()
    {
        var order = new KitchenOrder(
            new List<TicketItem> { new TicketItem("Burger", 2), new TicketItem("Fries", 1) },
            "Bob", 12, "");
        var result = new KitchenTicket().print_ticket(order);
        Assert.Contains("Burger x2", result);
        Assert.Contains("Fries x1", result);
    }

    [Fact]
    public void PrintsTicketWithSpecialInstructions()
    {
        var order = new KitchenOrder(
            new List<TicketItem> { new TicketItem("Salad", 1) },
            "Carol", 3, "No onions");
        var result = new KitchenTicket().print_ticket(order);
        Assert.Contains("Special: No onions", result);
    }

    [Fact]
    public void OmitsSpecialWhenEmpty()
    {
        var order = new KitchenOrder(
            new List<TicketItem> { new TicketItem("Pizza", 1) },
            "Dave", 7, "");
        var result = new KitchenTicket().print_ticket(order);
        Assert.DoesNotContain("Special:", result);
    }

    [Fact]
    public void IncludesSeparatorLine()
    {
        var order = new KitchenOrder(
            new List<TicketItem> { new TicketItem("Soup", 1) },
            "Eve", 1, "");
        var result = new KitchenTicket().print_ticket(order);
        Assert.Contains("---", result);
    }
}
