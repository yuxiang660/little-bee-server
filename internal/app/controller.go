package app

import (
	"github.com/yuxiang660/little-bee-server/internal/app/controller/login"
	"github.com/yuxiang660/little-bee-server/internal/app/controller/user"
	"go.uber.org/dig"
)

// InjectController injects an controller constructor to dig container.
func InjectController(container *dig.Container) func() {

	err := container.Provide(login.New)
	handleError(err)
	err = container.Provide(user.New)
	handleError(err)

	return nil
}
