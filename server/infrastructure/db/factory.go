package db

import (
	"github.com/mddg/go-sm/server/infrastructure/db/mysql"
	"gorm.io/gorm"
)

type Type int64

const (
	MysqlDB Type = iota
)

func Factory(db Type) *gorm.DB {
	switch db {
	case MysqlDB:
		return mysql.Connection()
	default:
		return nil
	}
}
