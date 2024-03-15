package database

import (
	"fmt"
	"go_net_http/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func GetDB() *gorm.DB {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		"localhost", "5432", "postgres", "ian",
		"testdb")
	db, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	dbClient, _ := db.DB()
	err = dbClient.Ping()
	if err != nil {
		log.Fatal("error occured while acquiring database connection: ", err)
	}
	fmt.Println("✅ Successfully configured DB.")
	Database = db
	db.AutoMigrate(&model.Language{})
	return db
}
