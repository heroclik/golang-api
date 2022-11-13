package main

import (
	"example/go-orm-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Username: "admin", Fname: "adminfname", Lname: "adminlname", Role: "1"})
	db.Create(&model.User{Username: "user", Fname: "userfname", Lname: "userlname", Role: "2"})

}
