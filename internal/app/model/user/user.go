package user

import (
	"github.com/yuxiang660/little-bee-server/internal/app/logger"
	"github.com/yuxiang660/little-bee-server/internal/app/model"
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
	"github.com/yuxiang660/little-bee-server/internal/app/store"
)

// User defines the structure about user model.
type User struct {
	db store.Store
}

// New creates user model.
func New(db store.Store) (model.IUser, error) {
	err := db.AutoMigrate(&schema.User{})
	if err != nil {
		logger.Error("User model fails to auto migrate!")
		return nil, err
	}

	return &User{
		db: db,
	}, err
}

// Create adds a user model to database.
func (u *User) Create(item schema.User) error {
	if err := u.db.Create(&item); err != nil {
		return err
	}

	return nil
}

// Query returns all users in database with the username.
func (u *User) Query(username string) ([]schema.User, error) {
	var users []schema.User
	err := u.db.Find(&users, "user_name = ?", username)

	return users, err
}
