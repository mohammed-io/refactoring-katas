public class Calculator
{
    public double[] normalize(double[] scores, double newMin, double newMax)
    {
        var result = new List<double>();
        var min = double.PositiveInfinity;
        var max = double.NegativeInfinity;

        // Find min and max of input scores
        for (int i = 0; i < scores.Length; i++)
        {
            if (scores[i] < min) min = scores[i];
            if (scores[i] > max) max = scores[i];
        }

        var inputRange = max - min;
        var outputRange = newMax - newMin;

        for (int i = 0; i < scores.Length; i++)
        {
            var normalizedValue = newMin + ((scores[i] - min) / inputRange) * outputRange;
            result.Add(Math.Round(normalizedValue, 2));
        }

        return result.ToArray();
    }
}
