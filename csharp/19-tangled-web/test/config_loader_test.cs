using Xunit;

public class ConfigLoaderTest
{
    [Fact]
    public void ReturnsDefaultsWhenFileIsMissing()
    {
        using var env = new TestEnvironment();
        env.Set("APP_CONFIG_PATH", Path.Combine(Path.GetTempPath(), Guid.NewGuid().ToString(), "missing.json"));
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal(3, result["retries"]);
        Assert.Equal("standard", result["theme"]);
        Assert.Equal(0.0, result["discount"]);
    }

    [Fact]
    public void ReadsLocalConfigFileWhenPresent()
    {
        using var env = new TestEnvironment();
        var configPath = env.WriteConfig("{\"name\":\"test\",\"retries\":5}");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal("test", result["name"]);
        Assert.Equal(5.0, result["retries"]);
    }

    [Fact]
    public void EnvironmentOverridesLocalConfig()
    {
        using var env = new TestEnvironment();
        env.WriteConfig("{\"theme\":\"local\",\"retries\":5}");
        env.Set("APP_THEME", "env-theme");
        env.Set("APP_RETRIES", "9");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal("env-theme", result["theme"]);
        Assert.Equal(9, result["retries"]);
    }

    [Fact]
    public void HandlesMalformedJsonGracefully()
    {
        using var env = new TestEnvironment();
        env.WriteConfig("not json");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal("standard", result["theme"]);
    }

    [Fact]
    public void WinterSeasonalConfigHasHighestPrecedence()
    {
        using var env = new TestEnvironment();
        env.WriteConfig("{\"theme\":\"local\",\"discount\":0.25}");
        env.Set("APP_THEME", "env-theme");
        env.Set("APP_CURRENT_MONTH", "12");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal("winter", result["theme"]);
        Assert.Equal(0.1, result["discount"]);
    }

    [Fact]
    public void SummerSeasonalConfigIsDeterministic()
    {
        using var env = new TestEnvironment();
        env.Set("APP_CONFIG_PATH", Path.Combine(env.DirectoryPath, "missing.json"));
        env.Set("APP_CURRENT_MONTH", "7");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal("summer", result["theme"]);
        Assert.Equal(0.05, result["discount"]);
    }
}

public sealed class TestEnvironment : IDisposable
{
    private readonly Dictionary<string, string?> _oldValues = new();
    public string DirectoryPath { get; } = Path.Combine(Path.GetTempPath(), Guid.NewGuid().ToString());

    public TestEnvironment()
    {
        Directory.CreateDirectory(DirectoryPath);
    }

    public void Set(string key, string value)
    {
        if (!_oldValues.ContainsKey(key))
        {
            _oldValues[key] = Environment.GetEnvironmentVariable(key);
        }

        Environment.SetEnvironmentVariable(key, value);
    }

    public string WriteConfig(string body)
    {
        var path = Path.Combine(DirectoryPath, "config.json");
        File.WriteAllText(path, body);
        Set("APP_CONFIG_PATH", path);
        return path;
    }

    public void Dispose()
    {
        foreach (var entry in _oldValues)
        {
            Environment.SetEnvironmentVariable(entry.Key, entry.Value);
        }

        if (Directory.Exists(DirectoryPath))
        {
            Directory.Delete(DirectoryPath, true);
        }
    }
}
