package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// dsn = "address,password,database number"
	dsn = "127.0.0.1:6379,,1"
)

func TestRedis(t *testing.T) {
	store, err := New(SetDSN(dsn))
	assert.Nil(t, err)
	defer store.Close()

	key := "test_key"
	expectedVal := "test_value"
	actualVal := ""
	isExist := false

	err = store.Set(key, expectedVal, 0)
	assert.Nil(t, err)

	actualVal, err = store.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, actualVal)

	isExist, err = store.Exist(key)
	assert.Nil(t, err)
	assert.Equal(t, true, isExist)

	isExist, err = store.Exist("invalid_key")
	assert.Nil(t, err)
	assert.Equal(t, false, isExist)
}
