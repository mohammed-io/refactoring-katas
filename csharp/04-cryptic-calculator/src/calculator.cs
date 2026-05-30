public class Calculator
{
    public double[] normalize(double[] a, double b, double c)
    {
        var d = new List<double>();
        var e = double.PositiveInfinity;
        var f = double.NegativeInfinity;

        for (int g = 0; g < a.Length; g++)
        {
            if (a[g] < e) e = a[g];
            if (a[g] > f) f = a[g];
        }

        var h = f - e;
        var i = c - b;

        for (int g = 0; g < a.Length; g++)
        {
            var k = b + ((a[g] - e) / h) * i;
            d.Add(Math.Round(k, 2));
        }

        return d.ToArray();
    }
}
