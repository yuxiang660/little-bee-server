package schema

// User defines the structure of user information in memroy.
type User struct {
	RecordID  string    `json:"record_id"`
	UserName  string    `json:"user_name" binding:"required"`
	Password  string    `json:"password"`
}
