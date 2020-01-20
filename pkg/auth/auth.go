package auth

import (
	"context"
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
	GenerateToken(ctx context.Context, userID string) (TokenInfo, error)
	ParseUserID(ctx context.Context, accessToken string) (string, error)
}
