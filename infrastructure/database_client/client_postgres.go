package database_client

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"yofio/domain/model"
)

func GetConnectionPostgres() *gorm.DB {


	host := "ec2-54-92-230-7.compute-1.amazonaws.com"
	port := "5432"
	dbname := "d12sdavvqmev9l"
	user := "mwhdjvrlkijbnh"
	password := "fba5a673cb8a35c2114a9ce803060b858c0e65ef6011228b4a6e78bcab93954f"

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