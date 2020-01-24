// Wrapper of redis. Do not use other redis package directly in the project.

package redis

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
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

type parsedOptions struct {
	addr     string
	password string
	db       int
}

var defaultOptions = parsedOptions{
	addr:     ":6379",
	password: "",
	db:       0,
}

func parseDSN(dsn string) parsedOptions {
	o := defaultOptions

	opts := strings.Split(dsn, ",")
	if len(opts) == 3 {
		o.addr = opts[0]
		o.password = opts[1]
		o.db, _ = strconv.Atoi(opts[2])
	}

	return o
}

type storeRedis struct {
	db *redis.Client
}

// New creates a noSQL redis store based on user configuration.
func New(opts ...Option) (store.NoSQL, error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	dsn := parseDSN(o.dsn)
	db := redis.NewClient(&redis.Options{
		Addr:     dsn.addr,
		Password: dsn.password,
		DB:       dsn.db,
	})

	_, err := db.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &storeRedis{
		db: db,
	}, nil
}

func (s *storeRedis) Set(key, value string, expiration time.Duration) error {
	return nil
}

func (s *storeRedis) Get(key string) (string, error) {
	return "", nil
}

func (s *storeRedis) Close() error {
	return nil
}
