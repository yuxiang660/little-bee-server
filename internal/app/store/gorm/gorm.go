// Wrapper of gorm. Do not use other gorm package directly in the project.

package gorm

import (
	"os"
	"path/filepath"
	"time"

	igorm "github.com/jinzhu/gorm"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/store"

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

// SetDSN returns an action to set DSN for database connection.
func SetDSN(DSN string) Option {
	return func(o *options) {
		o.DSN = DSN
	}
}

// SetMaxLifetime returns an action to set max life time for a connection.
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

type storeGorm struct {
	db *igorm.DB
}

// New creates a SQL store object based on user configuration.
func New(opts ...Option) (store.SQL, error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	switch o.DBType {
	case "sqlite3":
		_ = os.MkdirAll(filepath.Dir(o.DSN), 0777)
	default:
		return nil, errors.ErrUnknownDatabase
	}

	db, err := igorm.Open(o.DBType, o.DSN)
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

	return &storeGorm{db: db}, nil
}

// Close close current db connection.  If database connection is not an io.Closer, returns an error.
func (s *storeGorm)Close() error{
	return s.db.Close()
}

// AutoMigrate run auto migration for given models.
func (s *storeGorm) AutoMigrate(values ...interface{}) error {
	return s.db.AutoMigrate(values...).Error
}

// Create insert the value into database
func (s *storeGorm) Create(value interface{}) error {
	return s.db.Create(value).Error
}

// Find find records that match given conditions
func (s *storeGorm) Find(out interface{}, where ...interface{}) error {
	return s.db.Find(out, where...).Error
}
