package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(home, ".gatorconfig.json"), data, 0644)
}
