package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=172.17.0.3 port=5432 user=postgres dbname=study password=123456")
	defer db.Close()

	if err != nil {
		fmt.Println("fail:", err)
	} else {
		fmt.Println("successful")
	}
}
