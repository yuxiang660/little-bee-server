package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
	"github.com/yuxiang660/little-bee-server/internal/app/routers"
	"go.uber.org/dig"
)

// OpenHTTPServer opens HTTP service, and returns a function to release resources of HTTP service.
func OpenHTTPServer(container *dig.Container) func() {
	r, err := routers.InitRouters(container)
	handleError(err)

	cfg := config.Global().HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr: addr,
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	go func() {
		logger.Info("HTTP service start, listen at: ", addr)
		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()
	
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error(err.Error())
		}
	}
}