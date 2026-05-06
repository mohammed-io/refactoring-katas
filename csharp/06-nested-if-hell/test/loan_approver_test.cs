using Xunit;

public class LoanApproverTest
{
    [Fact]
    public void AllowsNormalPackage()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false });
        Assert.True(result.Allowed);
        Assert.Null(result.Warning);
    }

    [Fact]
    public void RejectsOverweight()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 60, Hazardous = false, Weekend = false });
        Assert.False(result.Allowed);
        Assert.Equal("Weight exceeded", result.Warning);
    }

    [Fact]
    public void RejectsHazardous()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 10, Hazardous = true, Weekend = false });
        Assert.False(result.Allowed);
        Assert.Equal("Hazardous material", result.Warning);
    }

    [Fact]
    public void RejectsWeekend()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = true });
        Assert.False(result.Allowed);
        Assert.Equal("No weekend delivery", result.Warning);
    }

    [Fact]
    public void RejectsExtremeTemperature()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false, TemperatureRequired = 50 });
        Assert.False(result.Allowed);
        Assert.Equal("Temperature out of range", result.Warning);
    }

    [Fact]
    public void AllowsValidTemperature()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 10, Hazardous = false, Weekend = false, TemperatureRequired = 20 });
        Assert.True(result.Allowed);
    }

    [Fact]
    public void AllowsRemoteSmallPackage()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 15, Hazardous = false, Weekend = false, RemoteArea = true });
        Assert.True(result.Allowed);
        Assert.Equal("Remote surcharge applies", result.Warning);
    }

    [Fact]
    public void RejectsRemoteHeavyPackage()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Weight = 25, Hazardous = false, Weekend = false, RemoteArea = true });
        Assert.False(result.Allowed);
        Assert.Equal("Too heavy for remote", result.Warning);
    }

    [Fact]
    public void RejectsNullPackage()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(null);
        Assert.False(result.Allowed);
        Assert.Equal("No package", result.Warning);
    }

    [Fact]
    public void RejectsMissingWeight()
    {
        var checker = new LoanApprover();
        var result = checker.can_deliver(new Package { Hazardous = false, Weekend = false });
        Assert.False(result.Allowed);
        Assert.Equal("No weight specified", result.Warning);
    }
}
