using Xunit;

public class OrderProcessorTest
{
    private static Order CreateOrder(decimal price = 10, int quantity = 1, string email = "a@b.com", string zip = "12345")
    {
        return new Order(
            new List<LineItem> { new LineItem(price, quantity) },
            new Customer(email),
            new Address(zip)
        );
    }

    [Fact]
    public void RejectsEmptyItems()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(new Order(new List<LineItem>(), new Customer("a@b.com"), new Address("12345")));
        Assert.Equal("No items", result.Error);
    }

    [Fact]
    public void RejectsInvalidCustomer()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(email: ""));
        Assert.Equal("Invalid customer", result.Error);
    }

    [Fact]
    public void RejectsInvalidAddress()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(zip: ""));
        Assert.Equal("Invalid address", result.Error);
    }

    [Fact]
    public void RejectsOutOfStock()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(quantity: 101));
        Assert.Equal("Out of stock", result.Error);
    }

    [Fact]
    public void CalculatesTotalsForSmallOrder()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder());
        Assert.Equal(16.69m, result.Total);
        Assert.Equal("USPS", result.Carrier);
    }

    [Fact]
    public void CalculatesTotalsForMediumOrder()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(price: 20));
        Assert.Equal(27.39m, result.Total);
    }

    [Fact]
    public void UsesUpsForLargeOrders()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(price: 100));
        Assert.Equal("UPS", result.Carrier);
        Assert.Equal("approved", result.PaymentStatus);
    }

    [Fact]
    public void FlagsHighTotalAsPendingReview()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(price: 1000));
        Assert.Equal("pending_review", result.PaymentStatus);
    }

    [Fact]
    public void IncludesEmailConfirmation()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(email: "user@test.com"));
        Assert.Equal("user@test.com", result.EmailTo);
        Assert.Equal("Order Confirmation", result.EmailSubject);
    }

    [Fact]
    public void Quantity100NotOutOfStock()
    {
        var processor = new OrderProcessor();
        var result = processor.process_order(CreateOrder(price: 1, quantity: 100));
        Assert.Null(result.Error);
    }
}
