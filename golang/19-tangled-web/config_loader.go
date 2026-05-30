package kata

import (
	"encoding/json"
	"os"
	"time"
)

type ConfigLoader struct{}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (cl *ConfigLoader) load_config() map[string]any {
	defaults := map[string]any{"retries": float64(3), "theme": "standard", "discount": float64(0)}
	local := map[string]any{}
	path := os.Getenv("APP_CONFIG_PATH")
	if path == "" {
		path = "/tmp/config.json"
	}
	raw, err := os.ReadFile(path)
	if err == nil {
		json.Unmarshal(raw, &local)
	}
	envConfig := map[string]any{}
	if theme := os.Getenv("APP_THEME"); theme != "" {
		envConfig["theme"] = theme
	}
	if retries := os.Getenv("APP_RETRIES"); retries != "" {
		var parsed float64
		json.Unmarshal([]byte(retries), &parsed)
		envConfig["retries"] = parsed
	}
	seasonal := map[string]any{}
	month := int(time.Now().Month())
	if forced := os.Getenv("APP_CURRENT_MONTH"); forced != "" {
		var parsed float64
		json.Unmarshal([]byte(forced), &parsed)
		month = int(parsed)
	}
	if month == 12 || month == 1 || month == 2 {
		seasonal = map[string]any{"theme": "winter", "discount": 0.1}
	} else if month >= 6 && month <= 8 {
		seasonal = map[string]any{"theme": "summer", "discount": 0.05}
	}
	result := map[string]any{}
	for k, v := range defaults {
		result[k] = v
	}
	for k, v := range local {
		result[k] = v
	}
	for k, v := range envConfig {
		result[k] = v
	}
	for k, v := range seasonal {
		result[k] = v
	}
	return result
}
