package auther

// TokenInfo defines the interface to access the contents of a token.
type TokenInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresAt() int64
	EncodeToJSON() ([]byte, error)
}

// Auther defines the infterface to manager a token.
type Auther interface {
	GenerateToken(userID string) (TokenInfo, error)
	DestroyToken(accessToken string) error
	ParseUserID(accessToken string) (string, error)
	Close() error
}
