using System.Collections.Generic;

public class InventoryItem
{
    public int Id;
    public string Name;
    public string BatchNumber;
    public int CacheTimestamp;
    public int RowId;
    public int Quantity;

    public InventoryItem(int id, string name, string batch, int cache, int row, int qty)
    {
        Id = id;
        Name = name;
        BatchNumber = batch;
        CacheTimestamp = cache;
        RowId = row;
        Quantity = qty;
    }

    public int get_id() => Id;

    public void set_id(int v)
    {
        Id = v;
    }

    public string get_name() => Name;

    public void set_name(string v)
    {
        Name = v;
    }

    public int get_quantity() => Quantity;

    public void set_quantity(int v)
    {
        Quantity = v;
    }

    public Dictionary<string, object> reserve(int units)
    {
        if (units <= 0)
        {
            return new Dictionary<string, object>
            {
                ["status"] = "rejected",
                ["reason"] = "invalid_quantity",
                ["remaining"] = Quantity
            };
        }

        if (units > Quantity)
        {
            return new Dictionary<string, object>
            {
                ["status"] = "backorder",
                ["reserved"] = 0,
                ["remaining"] = Quantity
            };
        }

        Quantity -= units;
        return new Dictionary<string, object>
        {
            ["status"] = "reserved",
            ["reserved"] = units,
            ["remaining"] = Quantity,
            ["sku"] = $"{Id}-{BatchNumber}"
        };
    }

    public int receive_stock(int units)
    {
        if (units <= 0)
        {
            return Quantity;
        }

        Quantity += units;
        return Quantity;
    }

    public Dictionary<string, object> public_snapshot()
    {
        return new Dictionary<string, object>
        {
            ["id"] = Id,
            ["name"] = Name,
            ["batch_number"] = BatchNumber,
            ["quantity"] = Quantity,
            ["stock_status"] = stock_status()
        };
    }

    public string stock_status()
    {
        if (Quantity == 0)
        {
            return "out";
        }

        if (Quantity < 5)
        {
            return "low";
        }

        return "available";
    }
}
