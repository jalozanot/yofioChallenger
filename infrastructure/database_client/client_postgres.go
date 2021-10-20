package database_client

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"yofio/domain/model"
)

func GetConnectionPostgres() *gorm.DB {


	host := "localhost"
	port := "5432"
	dbname := "test"
	user := "yofio"
	password := "yofio"

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