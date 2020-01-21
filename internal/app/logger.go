package app

import (
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
)

// ConfigLogger configures the logger based on config file for the web application.
// Clients can set the logger level and output place.
func ConfigLogger() func() {
	cfg := config.Global().Log
	logger.SetLevel(cfg.Level)
	logger.SetFormatter(cfg.Format)
	releaseCall, err := logger.SetOutput(cfg.Output, cfg.OutputFile)
	handleError(err)
	
	return releaseCall
}