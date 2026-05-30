using Xunit;

public class EmployeeTest
{
    [Fact]
    public void EmployeeGets2PercentBonus()
    {
        var emp = new Employee("Alice", 50000);
        Assert.Equal(1000m, emp.calculate_bonus());
    }

    [Fact]
    public void ManagerGets5PercentBonus()
    {
        var mgr = new Manager("Bob", 80000);
        Assert.Equal(4000m, mgr.calculate_bonus());
    }

    [Fact]
    public void SeniorManagerCappedAt10000()
    {
        var sm = new SeniorManager("Carol", 300000);
        Assert.Equal(10000m, sm.calculate_bonus());
    }

    [Fact]
    public void SeniorManagerUnderCap()
    {
        var sm = new SeniorManager("Carol", 100000);
        Assert.Equal(5000m, sm.calculate_bonus());
    }

    [Fact]
    public void DirectorCappedAt20000()
    {
        var dir = new Director("Dave", 600000);
        Assert.Equal(20000m, dir.calculate_bonus());
    }

    [Fact]
    public void DirectorUnderCap()
    {
        var dir = new Director("Dave", 200000);
        Assert.Equal(10000m, dir.calculate_bonus());
    }

    [Fact]
    public void ManagerRespectsCapFromSeniorManager()
    {
        var mgr = new Manager("Eve", 300000);
        Assert.Equal(15000m, mgr.calculate_bonus());
    }

    [Fact]
    public void EmployeeTotalRewardAddsHighPerformanceAndTenure()
    {
        var emp = new Employee("Alice", 50000);
        Assert.Equal(2000m, emp.calculate_total_reward("high", 5));
    }

    [Fact]
    public void SeniorManagerTotalRewardUsesCappedBonusAndTenureRule()
    {
        var sm = new SeniorManager("Carol", 300000);
        Assert.Equal(11500m, sm.calculate_total_reward("normal", 7));
    }

    [Fact]
    public void DirectorTotalRewardUsesDirectorPerformanceRule()
    {
        var dir = new Director("Dana", 200000);
        Assert.Equal(16000m, dir.calculate_total_reward("high", 3));
    }
}
