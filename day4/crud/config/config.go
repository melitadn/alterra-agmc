package config

import (
	"crud/lib/static"
	"crud/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "password",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "alterra",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)
	log.Println(connectionString)
	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		log.Fatal(e)
	}
	InitMigrate()
	static.BookInit = static.Init()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
