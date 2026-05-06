using Xunit;

public class InventoryItemTest
{
    [Fact]
    public void StoresAndReturnsId()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(1, item.get_id());
        item.set_id(2);
        Assert.Equal(2, item.get_id());
    }

    [Fact]
    public void StoresAndReturnsName()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal("Widget", item.get_name());
        item.set_name("Gadget");
        Assert.Equal("Gadget", item.get_name());
    }

    [Fact]
    public void StoresAndReturnsBatchNumber()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal("B001", item.GetBatchNumber());
        item.SetBatchNumber("B002");
        Assert.Equal("B002", item.GetBatchNumber());
    }

    [Fact]
    public void StoresAndReturnsCacheTimestamp()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(123, item.GetCacheTimestamp());
        item.SetCacheTimestamp(456);
        Assert.Equal(456, item.GetCacheTimestamp());
    }

    [Fact]
    public void StoresAndReturnsRowId()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(99, item.GetRowId());
        item.SetRowId(100);
        Assert.Equal(100, item.GetRowId());
    }

    [Fact]
    public void StoresAndReturnsQuantity()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(10, item.get_quantity());
        item.set_quantity(5);
        Assert.Equal(5, item.get_quantity());
    }
}
