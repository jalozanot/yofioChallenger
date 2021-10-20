package database_client

/*import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"yofio/domain/model"
)

func GetConnectionPostgres() *gorm.DB {


	host := "ec2-52-200-68-5.compute-1.amazonaws.com"
	port := "5432"
	dbname := "d3qf19uk0iggoe"
	user := "zbtituqdyzpmyg"
	password := "1b85c87b510d5e43352a9c2a9a35ca5e713f84b3e9daa5c28fa383109b711ca6"

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
}*/