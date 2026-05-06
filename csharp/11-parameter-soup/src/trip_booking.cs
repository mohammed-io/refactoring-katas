public class Booking
{
    public string? Error;
    public string Origin = "";
    public string Destination = "";
    public decimal Total;
    public string Class = "";
    public string Confirmation = "";
}

public class TripBooking
{
    public Booking book_trip(string? origin, string? destination, string? departureDate, string? returnDate, string travelClass, string meal, string seat, string? loyalty, bool insurance, string? promo, bool flexible)
    {
        if (string.IsNullOrEmpty(origin) || string.IsNullOrEmpty(destination))
        {
            return new Booking { Error = "Missing route" };
        }

        if (string.IsNullOrEmpty(departureDate))
        {
            return new Booking { Error = "Missing departure" };
        }

        decimal b = travelClass == "business" ? 800 : travelClass == "first" ? 2000 : 200;
        decimal d = promo == "SAVE20" ? 0.2m : promo == "SAVE10" ? 0.1m : 0;
        var total = b * (1 - d);

        if (insurance)
        {
            total += 50;
        }

        if (flexible)
        {
            total += 30;
        }

        if (loyalty != null && loyalty.StartsWith("GOLD"))
        {
            total -= 25;
        }

        return new Booking
        {
            Origin = origin,
            Destination = destination,
            Total = total,
            Class = travelClass,
            Confirmation = "BK-" + Random.Shared.Next(1000000)
        };
    }
}
