using System.Text.Json;

public class ConfigLoader
{
    public Dictionary<string, object> load_config()
    {
        Dictionary<string, object> local = new();
        try
        {
            local = JsonSerializer.Deserialize<Dictionary<string, object>>(File.ReadAllText("/tmp/config.json")) ?? new();
        }
        catch
        {
        }

        var month = DateTime.Now.Month - 1;
        if (month == 11 || month == 0 || month == 1)
        {
            local["theme"] = "winter";
            local["discount"] = 0.1;
        }
        else if (month >= 5 && month <= 7)
        {
            local["theme"] = "summer";
            local["discount"] = 0.05;
        }

        return local;
    }
}
