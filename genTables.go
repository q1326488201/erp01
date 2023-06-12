package main

import (
	"app/entry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/erp?charset=utf8mb4&parseTime=true&loc=Local"))
	db.AutoMigrate(&entry.AuthUser{})
	db.AutoMigrate(&entry.AuthRole{})
	db.AutoMigrate(&entry.AuthUserRole{})
	db.AutoMigrate(&entry.AuthMenu{})
}
