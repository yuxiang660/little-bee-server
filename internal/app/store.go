package app

import (
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/store"
	"github.com/yuxiang660/little-bee-server/internal/app/store/gorm"
	"go.uber.org/dig"
)

// InjectStore injects store constructor to dig container.
func InjectStore(container *dig.Container) func() {
	cfg := config.Global()
	db, err := gorm.New(
		gorm.SetDebug(cfg.Gorm.Debug),
		gorm.SetDBType(cfg.Gorm.DBType),
		gorm.SetDSN(cfg.Gorm.DSN()),
		gorm.SetMaxLifetime(cfg.Gorm.MaxLifetime),
		gorm.SetMaxOpenConns(cfg.Gorm.MaxOpenConns),
		gorm.SetMaxIdleConns(cfg.Gorm.MaxIdleConns),
	)
	handleError(err)

	_ = container.Provide(func() store.Store {
		return db
	})

	return func() {
		db.Close()
	}
}