package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IUser struct {
	gorm.Model
	ID       int
	Age      int
	Name     string
	Birthday string
	Email    string
}

//DSN
const DSN = "host=localhost port=54321 user=postgres dbname=study sslmode=disable password=123456"

//指定驱动
const DRIVER = "postgres"

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(DRIVER, DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func add() {
	user := &IUser{ID: 6, Name: "li 6", Age: 26}
	db.NewRecord(user)
	db.Create(&user)
}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	add()
}
