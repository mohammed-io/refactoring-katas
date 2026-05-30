using System.Text.Json;

public class ConfigLoader
{
    public Dictionary<string, object> load_config()
    {
        Dictionary<string, object> defaults = new()
        {
            ["retries"] = 3,
            ["theme"] = "standard",
            ["discount"] = 0.0
        };
        Dictionary<string, object> local = new();
        try
        {
            var path = Environment.GetEnvironmentVariable("APP_CONFIG_PATH") ?? "/tmp/config.json";
            var parsed = JsonDocument.Parse(File.ReadAllText(path)).RootElement;
            foreach (var property in parsed.EnumerateObject())
            {
                local[property.Name] = property.Value.ValueKind == JsonValueKind.Number
                    ? property.Value.GetDouble()
                    : property.Value.ToString();
            }
        }
        catch
        {
        }

        Dictionary<string, object> envConfig = new();
        if (Environment.GetEnvironmentVariable("APP_THEME") is string theme)
        {
            envConfig["theme"] = theme;
        }

        if (Environment.GetEnvironmentVariable("APP_RETRIES") is string retries)
        {
            envConfig["retries"] = int.Parse(retries);
        }

        var month = int.Parse(Environment.GetEnvironmentVariable("APP_CURRENT_MONTH") ?? DateTime.Now.Month.ToString());
        Dictionary<string, object> seasonal = new();
        if (month == 12 || month == 1 || month == 2)
        {
            seasonal["theme"] = "winter";
            seasonal["discount"] = 0.1;
        }
        else if (month >= 6 && month <= 8)
        {
            seasonal["theme"] = "summer";
            seasonal["discount"] = 0.05;
        }

        var result = new Dictionary<string, object>();
        foreach (var source in new[] { defaults, local, envConfig, seasonal })
        {
            foreach (var pair in source)
            {
                result[pair.Key] = pair.Value;
            }
        }

        return result;
    }
}
