package schema

// LoginParam defines the structure of credential for client login.
type LoginParam struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
