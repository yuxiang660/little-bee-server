package main

import (
	"flag"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/yuxiang660/little-bee-server/internal/app"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
)

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

	releaseAPP := app.Open(app.SetConfigFile(configFile))

Loop:
	for {
		sig := <-sc
		logger.Info("Received a signal ", sig.String())

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

	logger.Info("Exit Service")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
