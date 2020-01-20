package config

import (
	"github.com/BurntSushi/toml"
)

// Config defines the structure to store all configurations in config file(./configs/config.toml).
type Config struct {
	RunMode     string      `toml:"run_mode"`
	Log         Log         `toml:"log"`
	JWTAuth     JWTAuth     `toml:"jwt_auth"`
}

// Log defines the structure to store log configuration in config file.
type Log struct {
	Level         int    `toml:"level"`
	Format        string `toml:"format"`
	Output        string `toml:"output"`
	OutputFile    string `toml:"output_file"`
}

// JWTAuth defines the structure to store JWT Authentication configuration in config file.
type JWTAuth struct {
	SigningMethod string `toml:"signing_method"`
	SigningKey    string `toml:"signing_key"`
	Expired       int    `toml:"expired"`
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
