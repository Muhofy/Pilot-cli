package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds all user-configurable settings.
type Config struct {
	Lang  string `json:"lang"`  // e.g. "en", "tr"
	Model string `json:"model"` // e.g. "deepseek/deepseek-chat-v3.1:free"
}

// defaults returns the default configuration.
func defaults() Config {
	return Config{
		Lang:  "",
		Model: "",
	}
}

// configPath returns ~/.pilot/config.json
func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".pilot")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

// Load reads the config file, returning defaults if it doesn't exist.
func Load() Config {
	cfg := defaults()
	path, err := configPath()
	if err != nil {
		return cfg
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}
	json.Unmarshal(data, &cfg)
	return cfg
}

// Save writes the config to disk.
func Save(cfg Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

// Set updates a single key in the config.
func Set(key, value string) error {
	cfg := Load()
	switch key {
	case "lang":
		cfg.Lang = value
	case "model":
		cfg.Model = value
	default:
		return nil
	}
	return Save(cfg)
}