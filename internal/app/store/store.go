package store

import (
	"github.com/jinzhu/gorm"
)

// Model is alias of gorm.Model
type Model gorm.Model

// SQL defines interface to manage SQL database.
type SQL interface {
	AutoMigrate(values ...interface{}) error
	Create(value interface{}) error
	Find(out interface{}, where ...interface{}) error 
	Close() error
}