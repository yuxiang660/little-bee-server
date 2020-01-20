package app

import (
	"context"
	"os"

	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/pkg/logger"
)

type options struct {
	ConfigFile string
	Version    string
}

// Option defines function signature to set data in app options.
type Option func(*options)

// SetConfigFile returns an action to set configuration filename in app options.
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetVersion returns an action to set version of the project.
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Open starts the web application after initialization.
// Open returns a function to close the web application.
func Open(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobal(o.ConfigFile)
	handleError(err)

	cfg := config.Global()
	logger.Printf(ctx, "Start Server, Run Mode: %s, Version: %s, PID: %d", cfg.RunMode, o.Version, os.Getpid())

	releaseLogger, err := ConfigLogger()
	handleError(err)

	return func() {
		if releaseLogger != nil {
			releaseLogger()
		}
	}
}
