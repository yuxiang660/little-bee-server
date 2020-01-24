package store

import (
	"time"

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

// NoSQL defines interface to manage NoSQL database.
type NoSQL interface {
	Set(key, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Close() error
}
