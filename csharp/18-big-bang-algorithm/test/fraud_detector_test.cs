using Xunit;

public class FraudDetectorTest
{
    private long Now = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds();

    private Tx CreateTx(decimal amount, long timestamp, List<Tx>? history = null, string merchant = "grocery", string country = "US", string cardCountry = "US")
    {
        return new Tx(amount, timestamp, history ?? new List<Tx>(), merchant, country, cardCountry);
    }

    [Fact]
    public void LowRiskSmallTransaction()
    {
        var detector = new FraudDetector();
        var result = detector.detect(CreateTx(10, Now));
        Assert.Equal("low", result.Rating);
        Assert.Equal(1, result.Level);
    }

    [Fact]
    public void MediumRiskForLargeAmount()
    {
        var detector = new FraudDetector();
        var result = detector.detect(CreateTx(1100, Now));
        Assert.Equal("medium", result.Rating);
        Assert.Equal(2, result.Level);
    }

    [Fact]
    public void GamblingMerchantIsMediumRisk()
    {
        var detector = new FraudDetector();
        var result = detector.detect(CreateTx(100, Now, merchant: "gambling"));
        Assert.Equal("medium", result.Rating);
        Assert.Equal(2, result.Level);
    }

    [Fact]
    public void CrossBorderAloneIsLowRisk()
    {
        var detector = new FraudDetector();
        var result = detector.detect(CreateTx(1000, Now, country: "FR", cardCountry: "US"));
        Assert.Equal("low", result.Rating);
        Assert.Equal(1, result.Level);
    }

    [Fact]
    public void CriticalRiskForLateNightCrypto()
    {
        var detector = new FraudDetector();
        var t = new DateTimeOffset(2024, 1, 1, 2, 0, 0, TimeSpan.Zero).ToUnixTimeMilliseconds();
        var result = detector.detect(CreateTx(600, t, merchant: "crypto", country: "CN", cardCountry: "US"));
        Assert.Equal("elevated", result.Rating);
        Assert.Equal(3, result.Level);
    }

    [Fact]
    public void VelocityIncreasesRisk()
    {
        var detector = new FraudDetector();
        var history = new List<Tx>
        {
            CreateTx(10, Now - 1000, merchant: "", country: "", cardCountry: ""),
            CreateTx(10, Now - 2000, merchant: "", country: "", cardCountry: ""),
            CreateTx(10, Now - 3000, merchant: "", country: "", cardCountry: ""),
            CreateTx(10, Now - 4000, merchant: "", country: "", cardCountry: "")
        };
        var result = detector.detect(CreateTx(50, Now, history));
        Assert.Equal(2, result.Level);
    }

    [Fact]
    public void VolumeSpikesAloneStayLowRisk()
    {
        var detector = new FraudDetector();
        var history = new List<Tx>
        {
            CreateTx(200, Now - 10000, merchant: "", country: "", cardCountry: ""),
            CreateTx(200, Now - 20000, merchant: "", country: "", cardCountry: ""),
            CreateTx(200, Now - 30000, merchant: "", country: "", cardCountry: "")
        };
        var result = detector.detect(CreateTx(50, Now, history));
        Assert.Equal(1, result.Level);
    }

    [Fact]
    public void IncludesScore()
    {
        var detector = new FraudDetector();
        var result = detector.detect(CreateTx(10, Now));
        Assert.IsType<int>(result.Score);
    }
}
