package database

import (
	"fmt"
	"log"
	"task-5-pbi-btpns-febrian-syahroni/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "1234"
	dbPort   = "5432"
	dbname   = "task_btpn"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	fmt.Println("Connection success to database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
