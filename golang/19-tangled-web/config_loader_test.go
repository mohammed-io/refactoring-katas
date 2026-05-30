package kata

import (
	"os"
	"testing"
)

func TestLoadConfigReturnsDefaultsWhenFileIsMissing(t *testing.T) {
	cl := NewConfigLoader()
	t.Setenv("APP_CONFIG_PATH", t.TempDir()+"/missing.json")
	result := cl.load_config()
	if result["retries"] != float64(3) || result["theme"] != "standard" || result["discount"] != float64(0) {
		t.Errorf("expected defaults, got %#v", result)
	}
}

func TestLoadConfigReadsLocalConfigFileWhenPresent(t *testing.T) {
	cl := NewConfigLoader()
	path := t.TempDir() + "/config.json"
	t.Setenv("APP_CONFIG_PATH", path)
	writeFile(t, path, `{"name":"test","retries":5}`)
	result := cl.load_config()
	if result["name"] != "test" || result["retries"] != float64(5) {
		t.Errorf("expected local config, got %#v", result)
	}
}

func TestLoadConfigEnvironmentOverridesLocalConfig(t *testing.T) {
	cl := NewConfigLoader()
	path := t.TempDir() + "/config.json"
	t.Setenv("APP_CONFIG_PATH", path)
	t.Setenv("APP_THEME", "env-theme")
	t.Setenv("APP_RETRIES", "9")
	writeFile(t, path, `{"theme":"local","retries":5}`)
	result := cl.load_config()
	if result["theme"] != "env-theme" || result["retries"] != float64(9) {
		t.Errorf("expected environment override, got %#v", result)
	}
}

func TestLoadConfigHandlesMalformedJsonGracefully(t *testing.T) {
	cl := NewConfigLoader()
	path := t.TempDir() + "/config.json"
	t.Setenv("APP_CONFIG_PATH", path)
	writeFile(t, path, `not json`)
	result := cl.load_config()
	if result["theme"] != "standard" {
		t.Errorf("expected defaults after malformed json, got %#v", result)
	}
}

func TestLoadConfigWinterSeasonalConfigHasHighestPrecedence(t *testing.T) {
	cl := NewConfigLoader()
	path := t.TempDir() + "/config.json"
	t.Setenv("APP_CONFIG_PATH", path)
	t.Setenv("APP_THEME", "env-theme")
	t.Setenv("APP_CURRENT_MONTH", "12")
	writeFile(t, path, `{"theme":"local","discount":0.25}`)
	result := cl.load_config()
	if result["theme"] != "winter" || result["discount"] != 0.1 {
		t.Errorf("expected winter seasonal override, got %#v", result)
	}
}

func TestLoadConfigSummerSeasonalConfigIsDeterministic(t *testing.T) {
	cl := NewConfigLoader()
	t.Setenv("APP_CONFIG_PATH", t.TempDir()+"/missing.json")
	t.Setenv("APP_CURRENT_MONTH", "7")
	result := cl.load_config()
	if result["theme"] != "summer" || result["discount"] != 0.05 {
		t.Errorf("expected summer seasonal config, got %#v", result)
	}
}

func writeFile(t *testing.T, path string, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}
}
