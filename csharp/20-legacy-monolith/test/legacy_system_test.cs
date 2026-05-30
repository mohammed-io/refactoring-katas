using Xunit;

public class LegacySystemTest
{
    [Fact]
    public void RejectsEmptyOrder()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(1, new List<MonoItem>(), new MonoCustomer("", "")));
        Assert.Equal("No items", result.Error);
    }

    [Fact]
    public void CalculatesBasicTotal()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(1,
            new List<MonoItem> { new MonoItem(10, 2) },
            new MonoCustomer("", "a@b.com")));
        Assert.Equal(21.4m, result.Total);
        Assert.Equal("USPS", result.Carrier);
    }

    [Fact]
    public void AppliesMemberDiscount()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(2,
            new List<MonoItem> { new MonoItem(100, 1) },
            new MonoCustomer("member", "a@b.com")));
        Assert.Equal(96.3m, result.Total);
    }

    [Fact]
    public void AppliesVipDiscount()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(3,
            new List<MonoItem> { new MonoItem(100, 1) },
            new MonoCustomer("vip", "a@b.com")));
        Assert.Equal(85.6m, result.Total);
    }

    [Fact]
    public void AppliesBonusDiscountOver100()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(4,
            new List<MonoItem> { new MonoItem(200, 1) },
            new MonoCustomer("", "a@b.com")));
        Assert.Equal(203.3m, result.Total);
    }

    [Fact]
    public void UsesUpsForLargeTotal()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(5,
            new List<MonoItem> { new MonoItem(60, 1) },
            new MonoCustomer("", "a@b.com")));
        Assert.Equal("UPS", result.Carrier);
    }

    [Fact]
    public void FlagsHighTotalForReview()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(6,
            new List<MonoItem> { new MonoItem(5000, 1) },
            new MonoCustomer("", "a@b.com")));
        Assert.Equal("manual_review", result.PaymentStatus);
    }

    [Fact]
    public void IncludesEmailDetails()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(7,
            new List<MonoItem> { new MonoItem(10, 1) },
            new MonoCustomer("", "user@test.com")));
        Assert.Equal("user@test.com", result.Email.To);
        Assert.Contains("7", result.Email.Subject);
    }

    [Fact]
    public void IncludesLogEntry()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(8,
            new List<MonoItem> { new MonoItem(10, 1) },
            new MonoCustomer("", "a@b.com")));
        Assert.Contains("Order processed", result.Log);
    }

    [Fact]
    public void IncludesOrderId()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(99,
            new List<MonoItem> { new MonoItem(10, 1) },
            new MonoCustomer("", "a@b.com")));
        Assert.Equal(99, result.OrderId);
    }

    [Fact]
    public void IgnoresNonPositiveItemPricesButCountsPositiveQuantities()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(10,
            new List<MonoItem> { new MonoItem(20, 2), new MonoItem(-100, 1) },
            new MonoCustomer("", "a@b.com")));
        Assert.Equal(42.8m, result.Total);
        Assert.Equal(3, result.ShippingWeight);
    }

    [Fact]
    public void AppliesSave10Coupon()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(11,
            new List<MonoItem> { new MonoItem(100, 1) },
            new MonoCustomer("", "a@b.com"),
            "SAVE10"));
        Assert.Equal(96.3m, result.Total);
    }

    [Fact]
    public void TaxExemptCustomerPaysNoTax()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(12,
            new List<MonoItem> { new MonoItem(100, 1) },
            new MonoCustomer("", "a@b.com", TaxExempt: true)));
        Assert.Equal(100m, result.Total);
        Assert.Equal(0m, result.TaxRate);
    }

    [Fact]
    public void ExpressShippingOverridesCarrier()
    {
        var system = new LegacySystem();
        var result = system.process_everything(new MonoOrder(13,
            new List<MonoItem> { new MonoItem(10, 1) },
            new MonoCustomer("", "a@b.com"),
            Shipping: new MonoShipping("express")));
        Assert.Equal("FedEx", result.Carrier);
    }
}
