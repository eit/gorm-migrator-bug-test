package main

import (
	"fmt"
	"log"

	"database/sql"

	"gorm-test/config"
	"gorm-test/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	sqlDB, err := sql.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	config.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	if err := config.DB.AutoMigrate(&model.User{}, &model.Profile{}); err != nil {
		log.Fatal(err.Error())
	}

}
