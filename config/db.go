package config

import (
	"fmt"
	"log"
	"os"

	"project-2/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// local
// var (
// 	host     = "localhost"
// 	user     = "postgres"
// 	password = "postgres"
// 	dbPort   = "5432"
// 	dbname   = "hacktiv8-project-2"
// 	db       *gorm.DB
// 	err      error
// )

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbPort   = os.Getenv("DB_PORT")
	dbname   = os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

func StartDB() {
	// using ssl
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)
	// without ssl
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: InitLog(),
	})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	_ = db.AutoMigrate(model.User{}, model.Photo{}, model.Comment{}, model.SocialMedia{})
	fmt.Println("Database connection success.")
}

func GetDB() *gorm.DB {
	return db
}
