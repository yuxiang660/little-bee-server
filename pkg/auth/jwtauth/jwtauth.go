package jwtauth

import (
	"context"
	"time"

	"github.com/yuxiang660/little-bee-server/pkg/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	keyfunc       jwt.Keyfunc
	expired       int
	tokenType     string
}

const defaultKey = "GINADMIN"

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       7200,
	signingMethod: jwt.SigningMethodHS512,
	signingKey:    []byte(defaultKey),
	keyfunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(defaultKey), nil
	},
}

// Option defines function signature to set options.
type Option func(*options)

// SetSigningMethod returns an action to set signing method.
func SetSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// SetSigningKey returns an action to set signing key.
func SetSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

// SetExpired returns an action to set token expired time(s).
func SetExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// SetKeyfunc returns an action to verify token key.
func SetKeyfunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyfunc = keyFunc
	}
}

// autherJWT defines a structure to store JWT Authentication properties.
type autherJWT struct {
	opts  *options
}

// New creates an autherJWT object based on user configuration.
func New(opts ...Option) auth.Auther {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	return &autherJWT{opts: &o}
}

// GenerateToken generates a token for a user.
func (a *autherJWT) GenerateToken(ctx context.Context, userID string) (auth.TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opts.expired) * time.Second).Unix()

	token := jwt.NewWithClaims(a.opts.signingMethod, &jwt.StandardClaims{
		IssuedAt: now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject: userID,
	})

	tokenString, err := token.SignedString(a.opts.signingKey)
	if err != nil {
		return nil, err
	}

	tokenInfo := &tokenInfo{
		ExpiresAt:   expiresAt,
		TokenType:   a.opts.tokenType,
		AccessToken: tokenString,
	}
	return tokenInfo, nil
}

func (a *autherJWT) parseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, a.opts.keyfunc)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, auth.ErrInvalidToken
	}

	return token.Claims.(*jwt.StandardClaims), nil
}

// ParseUserID parses a token.
// If the token is invalid, returns auth.ErrInvalidToken error.
// If the token is valid, returns user id string of the token user. 
func (a *autherJWT) ParseUserID(ctx context.Context, tokenString string) (string, error) {
	claims, err := a.parseToken(tokenString)
	if err != nil {
		return "", err
	}

	return claims.Subject, nil
}

