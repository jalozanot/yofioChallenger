package database_client

import (

	"yofio/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func GetConnectionPostgres() *gorm.DB {


	host := "127.0.0.1"
	port := "5432"
	dbname := "test"
	user := "postgres"
	password := "postgres"

	db, err := gorm.Open("postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbname+
			" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)

	db.AutoMigrate(&model.InvestmentEntity{})

	return db
}