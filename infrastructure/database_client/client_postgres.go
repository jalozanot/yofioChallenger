package database_client

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"yofio/domain/model"
)

func GetConnectionPostgres() *gorm.DB {

	connection := os.Getenv("DATABASE_URL")


	db, err := gorm.Open("postgres",connection)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)

	db.AutoMigrate(&model.InvestmentEntity{})

	return db
}