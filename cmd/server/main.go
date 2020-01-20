package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/yuxiang660/little-bee-server/internal/app"
	"github.com/yuxiang660/little-bee-server/pkg/logger"
	"github.com/yuxiang660/little-bee-server/pkg/util"
)

// VERSION indicates the version of the project.
var VERSION = "0.1.0"

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "", "Configuration File(.json, .yaml, .toml)")
}

func main() {
	flag.Parse()

	if configFile == "" {
		panic("Please input configuration file using -c")
	}

	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	logger.SetVersion(VERSION)
	ctx := logger.AddTraceIDToContext(context.Background(), util.NewTraceID())

	releaseAPP := app.Open(ctx,
		app.SetConfigFile(configFile),
		app.SetVersion(VERSION))

Loop:
	for {
		sig := <-sc
		logger.Printf(ctx, "Received a signal [%s]", sig.String())

		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break Loop
		case syscall.SIGHUP:
		default:
			break Loop
		}
	}

	if releaseAPP != nil {
		releaseAPP()
	}

	logger.Printf(ctx, "Exit Service")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
