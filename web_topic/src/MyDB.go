package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=nizonglong sslmode=disable")
	if err != nil {
		fmt.Println(err)
		DB.Close()
	}
}
