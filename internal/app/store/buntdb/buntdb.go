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
	return nil
}

func (s *storeBuntdb) Get(key string) (string, error) {
	return "", nil
}

func (s *storeBuntdb) Close() error {
	return nil
}
