package app

import (
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/pkg/auth"
	"github.com/yuxiang660/little-bee-server/pkg/auth/jwtauth"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/dig"
)

// InjectAuther injects an auther constructor to dig container.
// InjectAuther returns a function to release auther resource. 
// The auther will be construct based on configuration from clients.
// For example, clients can set token expired time and so on.
func InjectAuther(container *dig.Container) (func(), error) {
	cfg := config.Global().JWTAuth

	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(cfg.Expired))
	opts = append(opts, jwtauth.SetSigningKey([]byte(cfg.SigningKey)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(cfg.SigningKey), nil
	}))

	var method jwt.SigningMethod
	switch cfg.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	opts = append(opts, jwtauth.SetSigningMethod(method))

	auther := jwtauth.New(opts...)

	_ = container.Provide(func() auth.Auther {
		return auther
	})
	
	return nil, nil
}