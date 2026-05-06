using Xunit;

public class TripBookingTest
{
    private TripBooking _booking = new TripBooking();

    private Booking Book(string? origin = "LAX", string? destination = "NYC", string? departure = "2024-01-01",
                        string? returnDate = null, string travelClass = "economy", string meal = "vegan",
                        string seat = "aisle", string? loyalty = null, bool insurance = false,
                        string? promo = null, bool flexible = false)
    {
        return _booking.book_trip(origin, destination, departure, returnDate, travelClass, meal, seat, loyalty, insurance, promo, flexible);
    }

    [Fact]
    public void RejectsMissingOrigin()
    {
        var result = Book(origin: null);
        Assert.Equal("Missing route", result.Error);
    }

    [Fact]
    public void RejectsMissingDestination()
    {
        var result = Book(destination: null);
        Assert.Equal("Missing route", result.Error);
    }

    [Fact]
    public void RejectsMissingDepartureDate()
    {
        var result = Book(departure: null);
        Assert.Equal("Missing departure", result.Error);
    }

    [Fact]
    public void CalculatesEconomyPrice()
    {
        var result = Book(travelClass: "economy");
        Assert.Equal(200m, result.Total);
        Assert.Equal("economy", result.Class);
    }

    [Fact]
    public void CalculatesBusinessPrice()
    {
        var result = Book(travelClass: "business");
        Assert.Equal(800m, result.Total);
    }

    [Fact]
    public void CalculatesFirstClassPrice()
    {
        var result = Book(travelClass: "first");
        Assert.Equal(2000m, result.Total);
    }

    [Fact]
    public void AppliesSave20Promo()
    {
        var result = Book(promo: "SAVE20");
        Assert.Equal(160m, result.Total);
    }

    [Fact]
    public void AppliesSave10Promo()
    {
        var result = Book(promo: "SAVE10");
        Assert.Equal(180m, result.Total);
    }

    [Fact]
    public void AddsInsurance()
    {
        var result = Book(insurance: true);
        Assert.Equal(250m, result.Total);
    }

    [Fact]
    public void AddsFlexibleDates()
    {
        var result = Book(flexible: true);
        Assert.Equal(230m, result.Total);
    }

    [Fact]
    public void AppliesGoldLoyaltyDiscount()
    {
        var result = Book(loyalty: "GOLD123");
        Assert.Equal(175m, result.Total);
    }

    [Fact]
    public void IncludesRouteInResult()
    {
        var result = Book();
        Assert.Equal("LAX", result.Origin);
        Assert.Equal("NYC", result.Destination);
    }

    [Fact]
    public void IncludesConfirmationCode()
    {
        var result = Book();
        Assert.StartsWith("BK-", result.Confirmation);
    }
}
