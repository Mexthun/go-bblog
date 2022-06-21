package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "methun:Password123@tcp(sh-cynosdbmysql-grp-9g61hf16.sql.tencentcdb.com:20027)/mmblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
