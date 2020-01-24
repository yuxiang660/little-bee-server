package user

import (
	"fmt"
	"crypto/sha1"

	"github.com/google/uuid"
	"github.com/yuxiang660/little-bee-server/internal/app/config"
	"github.com/yuxiang660/little-bee-server/internal/app/model"
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
	"github.com/yuxiang660/little-bee-server/internal/app/store"
)

// User defines the structure about user model.
type User struct {
	db store.SQL
}

// New creates user model.
func New(db store.SQL) (model.IUser, error) {
	err := db.AutoMigrate(&schema.User{})

	return &User{
		db: db,
	}, err
}

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Create adds a user model to database.
func (u *User) Create(item schema.LoginParam) error {
	var user schema.User
	
	user.UserName = item.UserName
	user.Password = hash(item.Password)
	uuid, _ := uuid.NewRandom()
	user.RecordID = uuid.String()

	if err := u.db.Create(&user); err != nil {
		return err
	}

	return nil
}

// Query returns all users in database with the username.
func (u *User) Query(username string) (schema.UserQueryResults, error) {
	var users []schema.User
	err := u.db.Find(&users, "user_name = ?", username)

	return schema.UserQueryResults{Users : users}, err
}

// GetRootUser returns root user info which comes from configuration file.
func (u *User) GetRootUser() schema.User {
	root := config.Global().Root
	return schema.User{
		RecordID: root.UserName,
		UserName: root.UserName,
		Password: hash(root.Password),
	}
}

// VerifyCredential verifies the login credential with the user info.
func (u *User) VerifyCredential(cred schema.LoginParam, user schema.User) bool {
	return (cred.UserName == user.UserName) && (hash(cred.Password) == user.Password)
}
