using Xunit;

public class PackageValidatorTest
{
    [Fact]
    public void AllowsNormalPackage()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false });
        Assert.True(result.Allowed);
        Assert.Null(result.Warning);
    }

    [Fact]
    public void RejectsOverweight()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 60, Hazardous = false, Weekend = false });
        Assert.False(result.Allowed);
        Assert.Equal("Weight exceeded", result.Warning);
    }

    [Fact]
    public void RejectsHazardous()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 10, Hazardous = true, Weekend = false });
        Assert.False(result.Allowed);
        Assert.Equal("Hazardous material", result.Warning);
    }

    [Fact]
    public void RejectsWeekend()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = true });
        Assert.False(result.Allowed);
        Assert.Equal("No weekend delivery", result.Warning);
    }

    [Fact]
    public void RejectsExtremeTemperature()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false, TemperatureRequired = 50 });
        Assert.False(result.Allowed);
        Assert.Equal("Temperature out of range", result.Warning);
    }

    [Fact]
    public void AllowsValidTemperature()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false, TemperatureRequired = 20 });
        Assert.True(result.Allowed);
    }

    [Fact]
    public void AllowsRemoteSmallPackage()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 15, Hazardous = false, Weekend = false, RemoteArea = true });
        Assert.True(result.Allowed);
        Assert.Equal("Remote surcharge applies", result.Warning);
    }

    [Fact]
    public void RejectsRemoteHeavyPackage()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 25, Hazardous = false, Weekend = false, RemoteArea = true });
        Assert.False(result.Allowed);
        Assert.Equal("Too heavy for remote", result.Warning);
    }

    [Fact]
    public void RejectsNullPackage()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(null);
        Assert.False(result.Allowed);
        Assert.Equal("No package", result.Warning);
    }

    [Fact]
    public void RejectsMissingWeight()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Hazardous = false, Weekend = false });
        Assert.False(result.Allowed);
        Assert.Equal("No weight specified", result.Warning);
    }

    [Fact]
    public void Weight50AllowedAtBoundary()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 50, Hazardous = false, Weekend = false });
        Assert.True(result.Allowed);
        Assert.Null(result.Warning);
    }

    [Fact]
    public void Temperature40AllowedAtBoundary()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false, TemperatureRequired = 40 });
        Assert.True(result.Allowed);
    }

    [Fact]
    public void TemperatureMinus20AllowedAtBoundary()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false, TemperatureRequired = -20 });
        Assert.True(result.Allowed);
    }

    [Fact]
    public void RemoteWeight20AllowedAtBoundary()
    {
        var validator = new PackageValidator();
        var result = validator.can_deliver(new Package { Weight = 20, Hazardous = false, Weekend = false, RemoteArea = true });
        Assert.True(result.Allowed);
        Assert.Equal("Remote surcharge applies", result.Warning);
    }
}
