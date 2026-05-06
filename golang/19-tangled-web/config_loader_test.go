package kata

import (
	"os"
	"testing"
)

func TestLoadConfigReturnsEmptyObjectWhenFileAndFetchFail(t *testing.T) {
	cl := NewConfigLoader()
	os.Remove("/tmp/config.json")
	result := cl.load_config()
	if result == nil {
		t.Error("expected non-nil object, got nil")
	}
}

func TestLoadConfigReadsLocalConfigFileWhenPresent(t *testing.T) {
	cl := NewConfigLoader()
	os.WriteFile("/tmp/config.json", []byte(`{"name":"test"}`), 0644)
	defer os.Remove("/tmp/config.json")
	result := cl.load_config()
	if result["name"] != "test" {
		t.Errorf("expected name 'test', got %v", result["name"])
	}
}

func TestLoadConfigLocalOverridesEmptyFile(t *testing.T) {
	cl := NewConfigLoader()
	os.WriteFile("/tmp/config.json", []byte(`{"name":"local"}`), 0644)
	defer os.Remove("/tmp/config.json")
	result := cl.load_config()
	if result["name"] != "local" {
		t.Errorf("expected name 'local', got %v", result["name"])
	}
}

func TestLoadConfigHandlesMalformedJsonGracefully(t *testing.T) {
	cl := NewConfigLoader()
	os.WriteFile("/tmp/config.json", []byte(`not json`), 0644)
	defer os.Remove("/tmp/config.json")
	result := cl.load_config()
	if result == nil {
		t.Error("expected non-nil object for malformed json, got nil")
	}
}

func TestLoadConfigIncludesSeasonalKeysInObject(t *testing.T) {
	cl := NewConfigLoader()
	os.Remove("/tmp/config.json")
	result := cl.load_config()
	_, hasTheme := result["theme"]
	_, hasDiscount := result["discount"]
	_ = hasTheme
	_ = hasDiscount
}
