package database_client

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"yofio/domain/model"
)

func GetConnectionPostgres() *gorm.DB {


	host := "ec2-174-129-229-106.compute-1.amazonaws.com"
	port := "5432"
	dbname := "d2m2hsofvodvak"
	user := "wbqenwzebxlhed"
	password := "78b52cbec481afe317aeaf4ecb9356a1969c1ec156bdad3dc2c451ab1d91b74f"

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