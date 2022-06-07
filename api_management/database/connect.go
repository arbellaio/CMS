package database

import (
	"api_management/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=1234 dbname=go_admin port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
	DB = db
	db.AutoMigrate(&models.User{}, &models.Role{})
	fmt.Println(db)
}
