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
	local := map[string]any{}
	raw, err := os.ReadFile("/tmp/config.json")
	if err == nil {
		json.Unmarshal(raw, &local)
	}
	seasonal := map[string]any{}
	month := int(time.Now().Month()) - 1
	if month == 11 || month == 0 || month == 1 {
		seasonal = map[string]any{"theme": "winter", "discount": 0.1}
	} else if month >= 5 && month <= 7 {
		seasonal = map[string]any{"theme": "summer", "discount": 0.05}
	}
	for k, v := range seasonal {
		local[k] = v
	}
	return local
}
