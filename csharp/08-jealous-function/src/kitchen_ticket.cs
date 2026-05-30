public class TicketItem
{
    public string Name;
    public int Qty;
    public List<string> Modifiers;
    public string Allergy;

    public TicketItem(string name, int qty, List<string>? modifiers = null, string allergy = "")
    {
        Name = name;
        Qty = qty;
        Modifiers = modifiers ?? new List<string>();
        Allergy = allergy;
    }
}

public class KitchenOrder
{
    public List<TicketItem> Items;
    public Dictionary<string, object> Customer;
    public Dictionary<string, object> Table;
    public string Special;
    public bool Rush;

    public KitchenOrder(List<TicketItem> items, Dictionary<string, object> customer, Dictionary<string, object> table, string special, bool rush = false)
    {
        Items = items;
        Customer = customer;
        Table = table;
        Special = special;
        Rush = rush;
    }
}

public class KitchenTicket
{
    public string print_ticket(KitchenOrder order)
    {
        var lines = new List<string>();
        lines.Add("Table: " + order.Table["number"]);
        lines.Add("Zone: " + (order.Table.ContainsKey("zone") ? order.Table["zone"] : "main"));
        lines.Add("Server: " + (order.Table.ContainsKey("server") ? order.Table["server"] : "unassigned"));
        lines.Add("Customer: " + order.Customer["name"]);
        if (order.Customer.ContainsKey("vip") && (bool)order.Customer["vip"])
        {
            lines.Add("VIP");
        }
        if (order.Rush)
        {
            lines.Add("RUSH");
        }
        var totalItems = 0;
        foreach (var item in order.Items)
        {
            totalItems += item.Qty;
            var line = item.Name + " x" + item.Qty;
            if (item.Modifiers.Count > 0)
            {
                line += " [" + string.Join(", ", item.Modifiers) + "]";
            }
            if (!string.IsNullOrEmpty(item.Allergy))
            {
                line += " ALLERGY:" + item.Allergy;
            }
            lines.Add(line);
        }
        lines.Add("Items: " + totalItems);
        if (!string.IsNullOrEmpty(order.Special))
        {
            lines.Add("Special: " + order.Special);
        }
        lines.Add("---");
        return string.Join("\n", lines);
    }
}
