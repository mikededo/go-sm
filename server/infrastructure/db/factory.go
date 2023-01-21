package db

import (
	"github.com/mddg/go-sm/server/infrastructure/db/mysql"
	"gorm.io/gorm"
)

type DbType int64

const (
	MysqlDb DbType = iota
)

func DbFactory(db DbType) *gorm.DB {
	switch db {
	case MysqlDb:
		return mysql.Connection()
	}

	return nil
}
