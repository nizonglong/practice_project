package src

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var DBHelper *gorm.DB
var err error

func InitDB() {
	DBHelper, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=nizonglong sslmode=disable")
	if err != nil {
		//fmt.Println(err)
		//log.Fatal("DB初始化错误：", err)
		ShutDownServer(err)
		return
	}
	DBHelper.LogMode(true)
	DBHelper.DB().SetMaxIdleConns(10)
	DBHelper.DB().SetMaxOpenConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)
}
