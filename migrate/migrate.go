package main

import (
	"gin-api/models/usermodel"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_crud_5?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&usermodel.User{})
	log.Println("Migration Completed")
}
