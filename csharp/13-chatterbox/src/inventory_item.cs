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
}
