package model

import (
	"github.com/yuxiang660/little-bee-server/internal/app/schema"
)

// IUser defines the interface to manager user model.
type IUser interface {
	Create(item schema.User) error
}
