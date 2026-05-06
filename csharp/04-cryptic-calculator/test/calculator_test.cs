using Xunit;

public class CalculatorTest
{
    [Fact]
    public void NormalizesScoresTo0To100Range()
    {
        var calculator = new Calculator();
        var result = calculator.normalize(new[] { 10d, 20d, 30d }, 0, 100);
        Assert.Equal(new[] { 0d, 50d, 100d }, result);
    }

    [Fact]
    public void NormalizesScoresTo1To5Range()
    {
        var calculator = new Calculator();
        var result = calculator.normalize(new[] { 10d, 20d, 30d }, 1, 5);
        Assert.Equal(new[] { 1d, 3d, 5d }, result);
    }

    [Fact]
    public void HandlesSingleValue()
    {
        var calculator = new Calculator();
        var result = calculator.normalize(new[] { 50d }, 0, 100);
        Assert.True(double.IsNaN(result[0]));
    }

    [Fact]
    public void HandlesNegativeInputRange()
    {
        var calculator = new Calculator();
        var result = calculator.normalize(new[] { -10d, 0d, 10d }, 0, 1);
        Assert.Equal(new[] { 0d, 0.5d, 1d }, result);
    }

    [Fact]
    public void HandlesSameMinAndMax()
    {
        var calculator = new Calculator();
        var result = calculator.normalize(new[] { 5d, 5d, 5d }, 0, 100);
        Assert.True(double.IsNaN(result[0]));
        Assert.True(double.IsNaN(result[1]));
        Assert.True(double.IsNaN(result[2]));
    }
}
