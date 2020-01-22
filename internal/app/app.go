package app

import (
	"os"

	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
	"go.uber.org/dig"
)

type options struct {
	configFile string
}

// Option defines function signature to set data in app options.
type Option func(*options)

// SetConfigFile returns an action to set configuration filename in app options.
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.configFile = s
	}
}

// Open starts the web application, and returns a function to release resources of 
// the web application.
func Open(opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobal(o.configFile)
	handleError(err)

	releaseLogger := ConfigLogger()

	logger.InfoWithFields("Start Server:", logger.Fields {
		"RunMode": config.Global().RunMode,
		"PID": os.Getpid(),
	})

	container, releaseContainer := buildContainer()

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

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func buildContainer() (*dig.Container, func()) {
	container := dig.New()

	releaseAuther := InjectAuther(container)
	releaseStore := InjectStore(container)
	releaseController := InjectController(container)

	return container, func() {
		if releaseController != nil {
			releaseController()
		}

		if releaseStore != nil {
			releaseStore()
		}

		if releaseAuther != nil {
			releaseAuther()
		}
	}
}
