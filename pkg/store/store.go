package store

import (
	"time"
	"github.com/jinzhu/gorm"

	// gorm inject
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type options struct {
	debug        bool
	DBType       string
	DSN          string
	maxLifetime  int
	maxOpenConns int
	maxIdleConns int
}

var defaultOptions = options{
	debug: true,
	DBType: "sqlite3",
	DSN: "data/ginadmin.db",
	maxLifetime: 7200,
	maxOpenConns: 150,
	maxIdleConns: 50,
}

// Option defines function signature to set options.
type Option func(*options)

// SetDebug returns an action to set debug flag.
func SetDebug(debug bool) Option {
	return func(o *options) {
		o.debug = debug
	}
}

// SetDBType returns an action to set database type.
func SetDBType(DBType string) Option {
	return func(o *options) {
		o.DBType = DBType
	}
}

// SetDSN returns an action to set DSN for database connenction.
func SetDSN(DSN string) Option {
	return func(o *options) {
		o.DSN = DSN
	}
}

// SetMaxLifetime returns an action to set max life time for a connenction.
func SetMaxLifetime(maxLifetime int) Option {
	return func(o *options) {
		o.maxLifetime = maxLifetime
	}
}

// SetMaxOpenConns returns an action to set max number of connections.
func SetMaxOpenConns(maxOpenConns int) Option {
	return func(o *options) {
		o.maxOpenConns = maxOpenConns
	}
}

// SetMaxIdleConns returns an action to set max number of connections in the idle connection pool.
func SetMaxIdleConns(maxIdleConns int) Option {
	return func(o *options) {
		o.maxIdleConns = maxIdleConns
	}
}

// New creates an autherJWT object based on user configuration.
func New(opts ...Option) (*gorm.DB, error) {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	db, err := gorm.Open(o.DBType, o.DSN)
	if err != nil {
		return nil, err
	}

	if o.debug {
		db = db.Debug()
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(o.maxIdleConns)
	db.DB().SetMaxOpenConns(o.maxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(o.maxLifetime) * time.Second)
	return db, nil
}
