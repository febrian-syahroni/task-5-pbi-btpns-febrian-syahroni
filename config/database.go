package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// variabel global
var DB *gorm.DB

func ConnectDB() *gorm.DB {
	host := "localhost"
	username := "postgres"
	password := "1234"
	database := "foto-profil"
	port := "5432"

	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		panic("Can't connect to database")
	}

	DB.AutoMigrate(&Product{})

}
