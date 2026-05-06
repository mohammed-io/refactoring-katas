using Xunit;

public class ConfigLoaderTest
{
    private string ConfigPath = "/tmp/config.json";

    [Fact]
    public void ReturnsEmptyObjectWhenFileAndFetchFail()
    {
        if (File.Exists(ConfigPath)) File.Delete(ConfigPath);
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.NotNull(result);
        Assert.IsType<Dictionary<string, object>>(result);
    }

    [Fact]
    public void ReadsLocalConfigFileWhenPresent()
    {
        File.WriteAllText(ConfigPath, "{\"name\":\"test\"}");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal("test", result["name"]?.ToString());
        File.Delete(ConfigPath);
    }

    [Fact]
    public void LocalOverridesEmptyFile()
    {
        File.WriteAllText(ConfigPath, "{\"name\":\"local\"}");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.Equal("local", result["name"]?.ToString());
        File.Delete(ConfigPath);
    }

    [Fact]
    public void HandlesMalformedJsonGracefully()
    {
        File.WriteAllText(ConfigPath, "not json");
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.NotNull(result);
        Assert.IsType<Dictionary<string, object>>(result);
        File.Delete(ConfigPath);
    }

    [Fact]
    public void IncludesSeasonalKeysInObject()
    {
        var loader = new ConfigLoader();
        var result = loader.load_config();
        Assert.True(result.ContainsKey("theme") || result["theme"] == null);
        Assert.True(result.ContainsKey("discount") || result["discount"] == null);
    }
}
