package buntdb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	expiration = time.Millisecond
)

func TestBuntDB(t *testing.T) {
	store, err := New(SetDSN(":memory:"))
	assert.Nil(t, err)
	defer store.Close()

	key := "test_key"
	expectedVal := "test_value"
	actualVal := ""
	isExist := false

	err = store.Set(key, expectedVal, expiration)
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

	time.Sleep(expiration)
	isExist, err = store.Exist(key)
	assert.Nil(t, err)
	assert.Equal(t, false, isExist)
}
