package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(filepath.Join(home, ".gatorconfig.json"))
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
