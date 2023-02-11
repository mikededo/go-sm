package mysql

import (
	"sync"

	"github.com/mddg/go-sm/server/infrastructure/db/mysql/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var lock = sync.Mutex{}

func Connection() *gorm.DB {
	lock.Lock()
	defer lock.Unlock()

	if db != nil {
		return db
	}

	dsn := "root:test-db@tcp(127.0.0.1:3306)/go-sm-db?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to go-sm-db database")
	}

	// Add schemas
	schema.AttachUserToDatabase(_db)
	schema.AttachPostToDatabase(_db)
	db = _db

	return _db
}
