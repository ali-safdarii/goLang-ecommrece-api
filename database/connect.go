package database

import (
	"alisafdarirepo/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// Connect to mysql /*
/*
func Connect() {
	dsn := "root:52281374ali*058058@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	fmt.Println("Database Connection Successfully")

	DB=db

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Auto Migration Failed")
	}

}
*/

/*Connect to postgres sql*/
func Connect() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=52281374 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic("Auto Migration Failed")
	}

	DB = db

}
