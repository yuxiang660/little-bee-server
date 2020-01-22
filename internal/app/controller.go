package app

import (
	"github.com/yuxiang660/little-bee-server/internal/app/controller/login"
	"go.uber.org/dig"
)

// InjectController injects an controller constructor to dig container.
func InjectController(container *dig.Container) func() {

	_ = container.Provide(login.New)

	return nil
}
