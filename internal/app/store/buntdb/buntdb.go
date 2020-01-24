// Wrapper of buntdb. Do not use other buntdb directly in the project.

package buntdb

import (
	"os"
	"path/filepath"
	"time"

	"github.com/tidwall/buntdb"
	"github.com/yuxiang660/little-bee-server/internal/app/store"
)

type options struct {
	dsn string
}

// Option defines function signature to set options.
type Option func(*options)

// SetDSN returns an action to set file DSN.
func SetDSN(dsn string) Option {
	return func(o *options) {
		o.dsn = dsn
	}
}

type storeBuntdb struct {
	db *buntdb.DB
}

// New creates a noSQL buntdb store based on user configuration.
func New(opts ...Option) (store.NoSQL, error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	if o.dsn != ":memory:" {
		os.MkdirAll(filepath.Dir(o.dsn), 0777)
	}

	db, err := buntdb.Open(o.dsn)
	if err != nil {
		return nil, err
	}

	return &storeBuntdb{
		db: db,
	}, nil
}

func (s *storeBuntdb) Set(key, value string, expiration time.Duration) error {
	return s.db.Update(func(tx *buntdb.Tx) error {
		var opts *buntdb.SetOptions
		if expiration > 0 {
			opts = &buntdb.SetOptions{Expires: true, TTL: expiration}
		}
		_, _, err := tx.Set(key, value, opts)
		return err
	})
}

func (s *storeBuntdb) Get(key string) (string, error) {
	var value string
	err := s.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		value = val
		return nil 
	})

	return value, err
}

func (s *storeBuntdb) Exist(key string) (bool, error) {
	_, err := s.Get(key)

	if err == buntdb.ErrNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *storeBuntdb) Close() error {
	return s.db.Close()
}
