package auther

import (
	"errors"
)

// ErrInvalidToken defines an error for invalid token.
var (
	ErrInvalidToken = errors.New("invalid token")
)

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
	ParseUserID(accessToken string) (string, error)
}
