package store

// Store defines interface to manage storage.
type Store interface {
	Close() error
}