using Xunit;

public class PayrollCalculatorTest
{
    [Fact]
    public void CalculatesFulltimePayslip()
    {
        var calculator = new PayrollCalculator();
        var result = calculator.generate_payslips(new List<Employee>
        {
            new Employee { Id = 1, Name = "Alice", Type = "fulltime", Salary = 60000 }
        });
        Assert.Equal(5000m, result[0].Gross);
        Assert.Equal(1250m, result[0].Deductions);
        Assert.Equal(3750m, result[0].Net);
    }

    [Fact]
    public void CalculatesFulltimeWithBonus()
    {
        var calculator = new PayrollCalculator();
        var result = calculator.generate_payslips(new List<Employee>
        {
            new Employee { Id = 1, Name = "Alice", Type = "fulltime", Salary = 60000, Bonus = 12000 }
        });
        Assert.Equal(6000m, result[0].Gross);
        Assert.Equal(4750m, result[0].Net);
    }

    [Fact]
    public void CalculatesParttimePayslip()
    {
        var calculator = new PayrollCalculator();
        var result = calculator.generate_payslips(new List<Employee>
        {
            new Employee { Id = 2, Name = "Bob", Type = "parttime", Hours = 80, Rate = 25 }
        });
        Assert.Equal(2000m, result[0].Gross);
        Assert.Equal(300m, result[0].Deductions);
        Assert.Equal(1700m, result[0].Net);
    }

    [Fact]
    public void CalculatesContractPayslip()
    {
        var calculator = new PayrollCalculator();
        var result = calculator.generate_payslips(new List<Employee>
        {
            new Employee { Id = 3, Name = "Carol", Type = "contract", FlatFee = 5000 }
        });
        Assert.Equal(5000m, result[0].Gross);
        Assert.Equal(500m, result[0].Deductions);
        Assert.Equal(4500m, result[0].Net);
    }

    [Fact]
    public void HandlesMultipleEmployees()
    {
        var calculator = new PayrollCalculator();
        var result = calculator.generate_payslips(new List<Employee>
        {
            new Employee { Id = 1, Name = "Alice", Type = "fulltime", Salary = 60000 },
            new Employee { Id = 2, Name = "Bob", Type = "parttime", Hours = 80, Rate = 25 }
        });
        Assert.Equal(2, result.Count);
        Assert.Equal("Alice", result[0].Name);
        Assert.Equal("Bob", result[1].Name);
    }

    [Fact]
    public void ReturnsEmptyArrayForEmptyInput()
    {
        var calculator = new PayrollCalculator();
        var result = calculator.generate_payslips(new List<Employee>());
        Assert.Empty(result);
    }
}
