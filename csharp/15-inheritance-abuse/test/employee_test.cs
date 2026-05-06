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
}
