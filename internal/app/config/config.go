package config

import (
	"github.com/BurntSushi/toml"
)

// Config defines the structure of all configurations in config file(./configs/config.toml).
type Config struct {
	RunMode     string      `toml:"run_mode"`
	Log         Log         `toml:"log"`
	JWTAuth     JWTAuth     `toml:"jwt_auth"`
	Gorm        Gorm        `toml:"gorm"`
	Sqlite3     Sqlite3     `toml:"sqlite3"`
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

// Log defines the structure of log configuration in config file.
type Log struct {
	Level         int    `toml:"level"`
	Format        string `toml:"format"`
	Output        string `toml:"output"`
	OutputFile    string `toml:"output_file"`
}

// JWTAuth defines the structure of JWT Authentication configuration in config file.
type JWTAuth struct {
	SigningMethod string `toml:"signing_method"`
	SigningKey    string `toml:"signing_key"`
	Expired       int    `toml:"expired"`
}

// Gorm defines the structure of gorm configuration in config file.
type Gorm struct {
	Debug             bool   `toml:"debug"`
	DBType            string `toml:"db_type"`
	MaxLifetime       int    `toml:"max_lifetime"`
	MaxOpenConns      int    `toml:"max_open_conns"`
	MaxIdleConns      int    `toml:"max_idle_conns"`
}

// Sqlite3 defines the structure of sqlite3 configuration in config file.
type Sqlite3 struct {
	Path string `toml:"path"`
}

// DSN returns sqlite3 DSN string for database connenction. 
func (a Sqlite3) DSN() string {
	return a.Path
}