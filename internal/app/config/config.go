package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
)

// Config defines the structure of all configurations in config file(./configs/config.toml).
type Config struct {
	RunMode     string      `toml:"run_mode"`
	DocDir      string      `toml:"doc_dir"`
	Root        Root        `toml:"root"`
	HTTP        HTTP        `toml:"http"`
	CORS        CORS        `toml:"cors"`
	Log         Log         `toml:"log"`
	JWTAuth     JWTAuth     `toml:"jwt_auth"`
	Redis       Redis       `toml:"redis"`
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

// Root defines username and password of root user.
type Root struct {
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
}

// HTTP defines the structure of http configuration in config file.
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	CertFile        string `toml:"cert_file"`
	KeyFile         string `toml:"key_file"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// CORS defines the structure of CORS configuration in config file.
type CORS struct {
	Enable           bool     `toml:"enable"`
	AllowOrigins     []string `toml:"allow_origins"`
	AllowMethods     []string `toml:"allow_methods"`
	AllowHeaders     []string `toml:"allow_headers"`
	AllowCredentials bool     `toml:"allow_credentials"`
	MaxAge           int      `toml:"max_age"`
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
	Store         string `toml:"store"`
	BuntdbPath    string `toml:"buntdb_path"`
	RedisDB       int    `toml:"redis_db"`
}

// Redis defines the structure of Buntdb about Redis storage configuration in config file.
type Redis struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
}

// DSN returns DSN string for JWTAuth database connenction. 
func (j JWTAuth) DSN() string {
	var dsn string

	switch j.Store {
	case "buntdb":
		dsn = global.JWTAuth.BuntdbPath
	case "redis":
		dsn = fmt.Sprintf("%s,%s,%d", global.Redis.Addr, global.Redis.Password, global.JWTAuth.RedisDB)
	default:
		logger.Error("Unknow JWTAuth Database")
		return ""
	}

	return dsn
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

// IsDebugMode checks whether the project is in debug mode or not.
func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

// DSN returns DSN string for Gorm database connenction. 
func (g Gorm) DSN() string {
	var dsn string

	switch g.DBType {
	case "sqlite3":
		dsn = global.Sqlite3.Path
	default:
		logger.Error("Unknow Gorm Database")
		return ""
	}

	return dsn
}