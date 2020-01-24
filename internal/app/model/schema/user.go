package schema

import (
	"github.com/yuxiang660/little-bee-server/internal/app/store"
)

// User defines the structure of user information in memroy.
type User struct {
	store.Model
	RecordID  string    `json:"record_id"`
	UserName  string    `json:"user_name" binding:"required"`
	Password  string    `json:"password"`
}

// UserQuery defines query form to query users with a user name.
type UserQuery struct {
	UserName  string    `form:"user_name"`
}

// UserQueryResults defines the return data from user query function.
type UserQueryResults struct {
	Users []User
}
