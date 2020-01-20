package config

import (
	"github.com/BurntSushi/toml"
)

// Config defines the structure to store all configurations in config file(./configs/config.toml).
type Config struct {
	RunMode     string      `toml:"run_mode"`
}

var (
	global *Config
)

// Parse returns configurations after parsing .toml config file.
func Parse(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// LoadGlobal loads all configurations globally from the config file.
func LoadGlobal(fpath string) error {
	c, err := Parse(fpath)
	if err != nil {
		return err
	}

	global = c
	return nil
}

// Global exports configuration handler.
// If no configuration was loaded, returns empty config handler.
func Global() *Config {
	if global == nil {
		return &Config{}
	}

	return global
}
