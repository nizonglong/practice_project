package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=study password=nizonglong sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println("fail:", err)
	} else {
		fmt.Println("successful")
	}
}
