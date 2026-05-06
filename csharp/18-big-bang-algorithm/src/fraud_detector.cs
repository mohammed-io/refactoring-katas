public record Tx(decimal Amount, long Timestamp, List<Tx> History, string Merchant, string Country, string CardCountry);

public record Risk(int Score, int Level, string Rating);

public class FraudDetector
{
    public Risk detect(Tx tx)
    {
        int s = 0, v = 0, m = 0;
        var h = DateTimeOffset.FromUnixTimeMilliseconds(tx.Timestamp).UtcDateTime.Hour;

        foreach (var x in tx.History)
        {
            if (x.Amount > tx.Amount * 2) v++;
            if (Math.Abs(tx.Timestamp - x.Timestamp) < 3600000) s++;
        }

        if (tx.Amount > 500 && h >= 0 && h < 6) m += 30;
        if (tx.Amount > 1000) m += 20;
        if (tx.Merchant == "gambling" || tx.Merchant == "crypto") m += 25;
        if (tx.Country != tx.CardCountry) m += 15;
        if (s > 3) m += 20;
        if (v > 2) m += 15;

        var level = m < 20 ? 1 : m < 40 ? 2 : m < 60 ? 3 : m < 80 ? 4 : 5;
        var rating = new[] { "", "low", "medium", "elevated", "high", "critical" }[level];

        return new Risk(m, level, rating);
    }
}
