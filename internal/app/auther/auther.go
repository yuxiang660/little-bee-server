package auther

import (
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
)

// Auther defines the infterface to manager a token.
type Auther interface {
	GenerateToken(userID string) (schema.LoginTokenInfo, error)
	DestroyToken(accessToken string) error
	ParseUserID(accessToken string) (string, error)
	Close() error
}
