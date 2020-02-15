// A wrapper for jwt-go. Don't use jwt-go directly in the project

package jwt

import (
	"time"

	"github.com/yuxiang660/little-bee-server/internal/app/auther"
	"github.com/yuxiang660/little-bee-server/internal/app/errors"
	"github.com/yuxiang660/little-bee-server/internal/app/store"
	"github.com/yuxiang660/little-bee-server/internal/app/store/buntdb"
	"github.com/yuxiang660/little-bee-server/internal/app/store/redis"
	"github.com/yuxiang660/little-bee-server/internal/app/model/schema"
	ijwt "github.com/dgrijalva/jwt-go"
)

type options struct {
	tokenType     string
	expired       int
	signingKey    string
	signingMethod string
	store         string
	dsn           string
}

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       7200,
	signingKey:    "GINADMIN",
	signingMethod: "HS512",
	store:         "buntdb",
	dsn:           "export/data/little-bee-auther.db",
}

// Option defines function signature to set options.
type Option func(*options)

// SetExpired returns an action to set token expired time(s).
func SetExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// SetSigningKey returns an action to set signing key.
func SetSigningKey(key string) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

// SetSigningMethod returns an action to set signing method.
func SetSigningMethod(method string) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// SetStore returns an action to set store type for auther.
func SetStore(strore string) Option {
	return func(o *options) {
		o.store = strore
	}
}

// SetDSN returns an action to set dsn string for store connection.
func SetDSN(dsn string) Option {
	return func(o *options) {
		o.dsn = dsn
	}
}

// autherJWT defines a structure to store JWT Authentication properties.
type autherJWT struct {
	tokenType     string
	expired       int
	signingKey    interface{}
	signingMethod ijwt.SigningMethod
	keyfunc       ijwt.Keyfunc
	db            store.NoSQL
}

// New creates an autherJWT object based on user configuration.
func New(opts ...Option) (auther.Auther, error) {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	var a autherJWT
	a.tokenType = o.tokenType
	a.expired = o.expired
	a.signingKey = []byte(o.signingKey)
	
	var method ijwt.SigningMethod
	switch o.signingMethod {
	case "HS256":
		method = ijwt.SigningMethodHS256
	case "HS384":
		method = ijwt.SigningMethodHS384
	default:
		method = ijwt.SigningMethodHS512
	}
	a.signingMethod = method

	a.keyfunc = func(t *ijwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*ijwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidToken
		}
		return []byte(o.signingKey), nil
	}

	var db store.NoSQL
	var err error
	switch o.store {
	case "buntdb":
		db, err = buntdb.New(buntdb.SetDSN(o.dsn))
	case "redis":
		db, err = redis.New(redis.SetDSN(o.dsn))
	default:
		err = errors.ErrUnknownDatabase
	}

	if err != nil {
		return nil, err
	}

	a.db = db

	return &a, nil
}

// GenerateToken generates a token for a user.
func (a *autherJWT) GenerateToken(userID string) (schema.LoginTokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.expired) * time.Second).Unix()

	token := ijwt.NewWithClaims(a.signingMethod, &ijwt.StandardClaims{
		IssuedAt: now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject: userID,
	})

	tokenString, err := token.SignedString(a.signingKey)
	if err != nil {
		return schema.LoginTokenInfo{}, err
	}

	return schema.LoginTokenInfo{
		AccessToken: tokenString,
		TokenType: a.tokenType,
		ExpiresAt: expiresAt,
	}, nil
}

func (a *autherJWT) parseToken(tokenString string) (*ijwt.StandardClaims, error) {
	token, err := ijwt.ParseWithClaims(tokenString, &ijwt.StandardClaims{}, a.keyfunc)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.ErrInvalidToken
	}

	return token.Claims.(*ijwt.StandardClaims), nil
}

// DestroyToken cancels a valid access token after clients logout.
// Store the canceled token into database for the check if clients use the token again.
func (a *autherJWT) DestroyToken(accessToken string) error {
	claims, err := a.parseToken(accessToken)
	if err != nil {
		return err
	}

	return a.db.Set(accessToken, "place_holder", time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))
}

// ParseUserID parses a token.
// If the token is invalid, returns ErrInvalidToken error.
// If the token is valid, returns user id string of the token user. 
func (a *autherJWT) ParseUserID(accessToken string) (string, error) {
	claims, err := a.parseToken(accessToken)
	if err != nil {
		return "", err
	}

	var exists bool
	exists, err = a.db.Exist(accessToken)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.ErrInvalidToken
	}

	return claims.Subject, nil
}

// Close releases the resources of Auther after close. 
func (a *autherJWT) Close() error {
	return a.db.Close()
}
