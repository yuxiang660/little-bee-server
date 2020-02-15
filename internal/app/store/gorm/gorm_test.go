package gorm

import (
	"testing"
	"os"

	igorm "github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

const (
	sqliteDSN = "gorm_sqlite_test.db"
)

type TestData struct {
	igorm.Model
	Name string
	Age uint
}

func TestSqlite(t *testing.T) {
	store, err := New(
		SetDebug(false),
		SetDBType("sqlite3"),
		SetDSN(sqliteDSN),
		SetMaxLifetime(7200),
		SetMaxOpenConns(150),
		SetMaxIdleConns(50),
	)
	assert.Nil(t, err)

	defer func() {
		store.Close()
		os.Remove(sqliteDSN)
	}()

	err = store.AutoMigrate(&TestData{})
	assert.Nil(t, err)

	expectedData := []TestData{{Name: "Tom", Age: 10}, {Name: "Jack", Age: 16}}
	err = store.Create(&expectedData[0])
	assert.Nil(t, err)
	err = store.Create(&expectedData[1])
	assert.Nil(t, err)

	var actualData []TestData
	err = store.Find(&actualData)
	assert.Nil(t, err)
	for i, data := range expectedData {
		assert.Equal(t, data.Name, actualData[i].Name)
		assert.Equal(t, data.Age, actualData[i].Age)
	}

	err = store.Find(&actualData, 2)
	assert.Nil(t, err)
	assert.Equal(t, expectedData[1].Name, actualData[0].Name)
	assert.Equal(t, expectedData[1].Age, actualData[0].Age)
}
