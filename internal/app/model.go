package app

import (
	"go.uber.org/dig"
	"github.com/yuxiang660/little-bee-server/internal/app/model/user"
)

// InjectModel injects model constructor to dig container.
func InjectModel(container *dig.Container) func() {

	err := container.Provide(user.New)
	handleError(err)
	
	return nil
}
