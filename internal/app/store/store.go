package store

import (
	"github.com/jinzhu/gorm"
)

// Model is alias of gorm.Model
type Model gorm.Model

// Gorm defines interface to manage database.
type Gorm interface {
	AutoMigrate(values ...interface{}) error
	Create(value interface{}) error
	Find(out interface{}, where ...interface{}) error 
	Close() error
}