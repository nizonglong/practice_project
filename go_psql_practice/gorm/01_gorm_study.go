package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "practice_project/go_psql_practice/gorm/model"
	"time"
)

func initTable(db *gorm.DB) {
	// 自动迁移模式
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Language{})
	db.AutoMigrate(&CreditCard{})
}

func update(db *gorm.DB, user User) {
	// update
	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=test_02 password=nizonglong sslmode=disable")
	db.LogMode(true)
	db.SingularTable(true)
	defer db.Close()

	if err != nil {
		fmt.Println("fail:", err)
	} else {
		fmt.Println("successful")
	}

	// 建表
	initTable(db)

	// insert user
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.Create(&user)

	//update(db, user)

	// 匹配一条记录
	db.Where("name = ?", "jinzhu").First(&user)

	users := make([]User, 3)
	users[0] = User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	users[1] = User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	users[2] = User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// 匹配全部记录
	db.Where("name = ?", "jinzhu").Find(&users)

	fmt.Println("where: ", user)
	fmt.Println("where all: ", users)

	count := 0
	db.Table("user").Count(&count)
	fmt.Println(count)
}
