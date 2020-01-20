package app

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/pkg/store/gormstore"
	"go.uber.org/dig"
)

// InjectStore injects store constructor to dig container.
func InjectStore(container *dig.Container) (func(), error) {
	cfg := config.Global()

	var dsn string
	switch cfg.Gorm.DBType {
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		_ = os.MkdirAll(filepath.Dir(dsn), 0777)
	default:
		return nil, errors.New("Unknown database")
	}

	gormCfg := cfg.Gorm

	var opts []gormstore.Option
	opts = append(opts, gormstore.SetDebug(gormCfg.Debug))
	opts = append(opts, gormstore.SetDBType(gormCfg.DBType))
	opts = append(opts, gormstore.SetDSN(dsn))
	opts = append(opts, gormstore.SetMaxLifetime(gormCfg.MaxLifetime))
	opts = append(opts, gormstore.SetMaxOpenConns(gormCfg.MaxOpenConns))
	opts = append(opts, gormstore.SetMaxIdleConns(gormCfg.MaxIdleConns))

	store, err := gormstore.New(opts...)
	if err != nil {
		return nil, err
	}

	releaseStore := func() {
		store.Close()
	}

	_ = container.Provide(func() *gorm.DB {
		return store
	})

	return releaseStore, nil
}