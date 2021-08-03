package database

import (
	"alisafdarirepo/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("UserName")
	password := os.Getenv("Password")
	dbName := os.Getenv("DBName")
	port := os.Getenv("Port")

	dsn := "user=" + username + "\tpassword=" + password + "\tdbname=" + dbName + "\tport=" + port + "\tsslmode=disable TimeZone=Asia/Shanghai"
	//"user=postgres password=52281374 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connecting To Database Successfully")

	err = db.AutoMigrate(&models.User{}, &models.Role{}, models.Permission{},
		&models.Product{}, &models.Order{})

	if err != nil {
		panic("Auto Migration Failed")
	}
	fmt.Println("AutoMigration Successfully")
	DB = db
}
