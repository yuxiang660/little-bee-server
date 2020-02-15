package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testJWTData struct {
	store string
	dsn string
	expectedID string
}

func TestJWT(t *testing.T) {
	cases := []testJWTData{
		{ "buntdb", ":memory:", "test_buntdb_id" },
		{ "redis", "127.0.0.1:6379,,1", "test_redis_id" },
	}

	for _, data := range cases {
		t.Run(fmt.Sprintf("JWT Test with %s", data.store), func(t *testing.T) {
			jwtAuth, err := New(
				SetExpired(1),
				SetSigningKey("GINADMIN"),
				SetSigningMethod("HS512"),
				SetStore(data.store),
				SetDSN(data.dsn),
			)
			assert.Nil(t, err)
			defer jwtAuth.Close()

			token, err := jwtAuth.GenerateToken(data.expectedID)
			assert.Nil(t, err)
			assert.NotNil(t, token)

			actualID, err := jwtAuth.ParseUserID(token.AccessToken)
			assert.Nil(t, err)
			assert.Equal(t, data.expectedID, actualID)

			err = jwtAuth.DestroyToken(token.AccessToken)
			assert.Nil(t, err)
		
			actualID, err = jwtAuth.ParseUserID(token.AccessToken)
			assert.NotNil(t, err)
			assert.EqualError(t, err, "Invalid Token")
			assert.Empty(t, actualID)
		})
	}
}

type testJWTErrorData struct {
	store string
	dsn string
	errorMessage string
}

func TestJWTError(t *testing.T) {
	cases := []testJWTErrorData{
		{ "buntdb", ":memory:", "token is expired by 1s" },
		{ "redis", "127.0.0.1:6379,,1", "token is expired by 1s" },
	}

	for _, data := range cases {
		t.Run(fmt.Sprintf("JWT Test Error with %s", data.store), func(t *testing.T) {
			jwtAuth, err := New(
				SetExpired(1),
				SetSigningKey("GINADMIN"),
				SetSigningMethod("HS512"),
				SetStore(data.store),
				SetDSN(data.dsn),
			)
			assert.Nil(t, err)
			defer jwtAuth.Close()
		
			token, err := jwtAuth.GenerateToken("test_id")
			assert.Nil(t, err)
			assert.NotNil(t, token)
		
			time.Sleep(2 * time.Second)
		
			actualID, err := jwtAuth.ParseUserID(token.AccessToken)
			assert.NotNil(t, err)
			assert.EqualError(t, err, data.errorMessage)
			assert.Empty(t, actualID)
		})
	}
}
