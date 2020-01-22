package store

// Store defines interface to manage storage.
type Store interface {
	AutoMigrate(values ...interface{}) error
	Close() error
}