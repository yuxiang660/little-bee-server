package model

import (
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
)

// IUser defines the interface to manager user model.
type IUser interface {
	Create(item schema.User) error
	Query(username string) (schema.UserQueryResults, error)
	GetRootUser() schema.User
}
