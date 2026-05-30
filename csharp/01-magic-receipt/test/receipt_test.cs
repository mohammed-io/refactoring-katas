using Xunit;

public class ReceiptTest
{
    [Fact]
    public void CalculatesTotalForRegularCustomer()
    {
        var receipt = new Receipt();
        Assert.Equal(59.4m, receipt.calculate_total(new[] { 10m, 20m, 30m }));
    }

    [Fact]
    public void CalculatesTotalForMemberCustomer()
    {
        var receipt = new Receipt();
        Assert.Equal(56.16m, receipt.calculate_total(new[] { 10m, 20m, 30m }, "member"));
    }

    [Fact]
    public void CalculatesTotalForVipCustomer()
    {
        var receipt = new Receipt();
        Assert.Equal(47.68m, receipt.calculate_total(new[] { 10m, 20m, 30m }, "vip"));
    }

    [Fact]
    public void AppliesBonusDiscountOver50()
    {
        var receipt = new Receipt();
        Assert.Equal(59.4m, receipt.calculate_total(new[] { 60m }));
    }

    [Fact]
    public void AppliesVipExtraDiscount()
    {
        var receipt = new Receipt();
        Assert.Equal(84.4m, receipt.calculate_total(new[] { 100m }, "vip"));
    }

    [Fact]
    public void Returns0ForEmptyItems()
    {
        var receipt = new Receipt();
        Assert.Equal(0m, receipt.calculate_total(Array.Empty<decimal>()));
    }

    [Fact]
    public void Exactly50NoBonus()
    {
        var receipt = new Receipt();
        Assert.Equal(54.0m, receipt.calculate_total(new[] { 50m }));
    }
}
