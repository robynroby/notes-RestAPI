package database

import (
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)