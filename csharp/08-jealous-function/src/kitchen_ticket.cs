public record TicketItem(string Name, int Qty);

public class KitchenOrder
{
    public List<TicketItem> Items;
    public string Customer;
    public int Table;
    public string Special;

    public KitchenOrder(List<TicketItem> items, string customer, int table, string special)
    {
        Items = items;
        Customer = customer;
        Table = table;
        Special = special;
    }
}

public class KitchenTicket
{
    public string print_ticket(KitchenOrder order)
    {
        var lines = new List<string>();
        lines.Add("Table: " + order.Table);
        lines.Add("Customer: " + order.Customer);
        foreach (var item in order.Items)
        {
            lines.Add(item.Name + " x" + item.Qty);
        }
        if (!string.IsNullOrEmpty(order.Special))
        {
            lines.Add("Special: " + order.Special);
        }
        lines.Add("---");
        return string.Join("\n", lines);
    }
}
