using Xunit;
using System.Collections.Generic;

public class InventoryItemTest
{
    [Fact]
    public void PublicSnapshotContainsBusinessFields()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(new Dictionary<string, object>
        {
            ["id"] = 1,
            ["name"] = "Widget",
            ["batch_number"] = "B001",
            ["quantity"] = 10,
            ["stock_status"] = "available"
        }, item.public_snapshot());
    }

    [Fact]
    public void ReservesStockAndReportsRemainingQuantity()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(new Dictionary<string, object>
        {
            ["status"] = "reserved",
            ["reserved"] = 3,
            ["remaining"] = 7,
            ["sku"] = "1-B001"
        }, item.reserve(3));
        Assert.Equal(7, item.public_snapshot()["quantity"]);
    }

    [Fact]
    public void RejectsReservationWhenQuantityIsInvalid()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(new Dictionary<string, object>
        {
            ["status"] = "rejected",
            ["reason"] = "invalid_quantity",
            ["remaining"] = 10
        }, item.reserve(0));
        Assert.Equal(10, item.public_snapshot()["quantity"]);
    }

    [Fact]
    public void BackordersWhenNotEnoughStock()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        Assert.Equal(new Dictionary<string, object>
        {
            ["status"] = "backorder",
            ["reserved"] = 0,
            ["remaining"] = 10
        }, item.reserve(12));
        Assert.Equal(10, item.public_snapshot()["quantity"]);
    }

    [Fact]
    public void ReceivesStockAfterLowQuantity()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        item.reserve(8);
        Assert.Equal("low", item.public_snapshot()["stock_status"]);
        Assert.Equal(7, item.receive_stock(5));
        Assert.Equal("available", item.public_snapshot()["stock_status"]);
    }

    [Fact]
    public void ReportsOutOfStockAfterExactReservation()
    {
        var item = new InventoryItem(1, "Widget", "B001", 123, 99, 10);
        item.reserve(10);
        Assert.Equal("out", item.public_snapshot()["stock_status"]);
    }
}
