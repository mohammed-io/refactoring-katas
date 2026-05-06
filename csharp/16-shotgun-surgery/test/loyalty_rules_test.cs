using Xunit;

public class LoyaltyRulesTest
{
    [Fact]
    public void BronzeDiscount()
    {
        Assert.Equal(0.05m, LoyaltyRules.get_discount_for_tier("bronze"));
    }

    [Fact]
    public void SilverDiscount()
    {
        Assert.Equal(0.1m, LoyaltyRules.get_discount_for_tier("silver"));
    }

    [Fact]
    public void GoldDiscount()
    {
        Assert.Equal(0.15m, LoyaltyRules.get_discount_for_tier("gold"));
    }

    [Fact]
    public void PlatinumDiscount()
    {
        Assert.Equal(0.2m, LoyaltyRules.get_discount_for_tier("platinum"));
    }

    [Fact]
    public void UnknownTierDiscount()
    {
        Assert.Equal(0m, LoyaltyRules.get_discount_for_tier("unknown"));
    }

    [Fact]
    public void BronzeLabel()
    {
        Assert.Equal("Bronze Member", LoyaltyRules.get_label_for_tier("bronze"));
    }

    [Fact]
    public void SilverLabel()
    {
        Assert.Equal("Silver Member", LoyaltyRules.get_label_for_tier("silver"));
    }

    [Fact]
    public void GoldLabel()
    {
        Assert.Equal("Gold Member", LoyaltyRules.get_label_for_tier("gold"));
    }

    [Fact]
    public void PlatinumLabel()
    {
        Assert.Equal("Platinum Member", LoyaltyRules.get_label_for_tier("platinum"));
    }

    [Fact]
    public void UnknownTierLabel()
    {
        Assert.Equal("Standard", LoyaltyRules.get_label_for_tier("unknown"));
    }

    [Fact]
    public void BronzeThreshold()
    {
        Assert.Equal(100, LoyaltyRules.get_threshold_for_tier("bronze"));
    }

    [Fact]
    public void SilverThreshold()
    {
        Assert.Equal(500, LoyaltyRules.get_threshold_for_tier("silver"));
    }

    [Fact]
    public void GoldThreshold()
    {
        Assert.Equal(2000, LoyaltyRules.get_threshold_for_tier("gold"));
    }

    [Fact]
    public void PlatinumThreshold()
    {
        Assert.Equal(10000, LoyaltyRules.get_threshold_for_tier("platinum"));
    }

    [Fact]
    public void UnknownTierThreshold()
    {
        Assert.Equal(0, LoyaltyRules.get_threshold_for_tier("unknown"));
    }

    [Fact]
    public void BronzeColor()
    {
        Assert.Equal("#CD7F32", LoyaltyRules.get_color_for_tier("bronze"));
    }

    [Fact]
    public void SilverColor()
    {
        Assert.Equal("#C0C0C0", LoyaltyRules.get_color_for_tier("silver"));
    }

    [Fact]
    public void GoldColor()
    {
        Assert.Equal("#FFD700", LoyaltyRules.get_color_for_tier("gold"));
    }

    [Fact]
    public void PlatinumColor()
    {
        Assert.Equal("#E5E4E2", LoyaltyRules.get_color_for_tier("platinum"));
    }

    [Fact]
    public void UnknownTierColor()
    {
        Assert.Equal("#000000", LoyaltyRules.get_color_for_tier("unknown"));
    }

    [Fact]
    public void CalculatesTotalForBronze()
    {
        Assert.Equal(95m, LoyaltyRules.calculate_total(100, "bronze"));
    }

    [Fact]
    public void CalculatesTotalForPlatinum()
    {
        Assert.Equal(80m, LoyaltyRules.calculate_total(100, "platinum"));
    }

    [Fact]
    public void CalculatesTotalForUnknownTier()
    {
        Assert.Equal(100m, LoyaltyRules.calculate_total(100, "unknown"));
    }
}
