package app

import (
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/auther"
	"github.com/yuxiang660/little-bee-server/internal/app/auther/jwt"
	"go.uber.org/dig"
)

// InjectAuther injects an auther constructor to dig container.
// InjectAuther returns a function to release auther resource. 
// The auther will be construct based on configuration from clients.
// For example, clients can set token expired time and so on.
func InjectAuther(container *dig.Container) func() {
	cfg := config.Global().JWTAuth

	a, err := jwt.New(
		jwt.SetExpired(cfg.Expired),
		jwt.SetSigningKey(cfg.SigningKey),
		jwt.SetSigningMethod(cfg.SigningMethod),
	)
	handleError(err)

	err = container.Provide(func() auther.Auther {
		return a
	})
	handleError(err)

	return nil
}