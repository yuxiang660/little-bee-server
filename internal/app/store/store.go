package store

import (
	"github.com/jinzhu/gorm"
)

// Model is alias of gorm.Model
type Model gorm.Model

// Store defines interface to manage storage.
type Store interface {
	AutoMigrate(values ...interface{}) error
	Close() error
}