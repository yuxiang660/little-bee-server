package app

import (
	"os"

	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
	"go.uber.org/dig"
)

type options struct {
	ConfigFile string
}

// Option defines function signature to set data in app options.
type Option func(*options)

// SetConfigFile returns an action to set configuration filename in app options.
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Open starts the web application after initialization.
// Open returns a function to release resources of the web application.
func Open(opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobal(o.ConfigFile)
	handleError(err)

	releaseLogger, err := ConfigLogger()
	handleError(err)

	logger.InfoWithFields("Start Server:", logger.Fields {
		"RunMode": config.Global().RunMode,
		"PID": os.Getpid(),
	})

	container, releaseContainer := BuildContainer()

	releaseHTTP := OpenHTTPServer(container)

	return func() {
		if releaseHTTP != nil {
			releaseHTTP()
		}

		if releaseContainer != nil {
			releaseContainer()
		}

		if releaseLogger != nil {
			releaseLogger()
		}
	}
}

// BuildContainer builds a dig container for dependency injection.
func BuildContainer() (*dig.Container, func()) {
	container := dig.New()

	releaseAuther, err := InjectAuther(container)
	handleError(err)

	releaseStore, err := InjectStore(container)
	handleError(err)

	return container, func() {
		if releaseStore != nil {
			releaseStore()
		}

		if releaseAuther != nil {
			releaseAuther()
		}
	}
}
