package user

import (
	"fmt"

	"github.com/yuxiang660/little-bee-server/internal/app/model"
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
	"github.com/yuxiang660/little-bee-server/internal/app/store"
)

// User defines the structure about user model.
type User struct {
	db store.Store
}

// New creates user model.
func New(db store.Store) model.IUser {
	return &User{
		db: db,
	}
}

// Create adds a user model to database.
func (u *User) Create(item schema.User) error {
	fmt.Println("user model creates one")
	return nil
}
