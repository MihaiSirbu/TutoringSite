package initializers

import (
    "log"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {
    

	var err error

	var DB_CONN_URL string = "host=suleiman.db.elephantsql.com user=cxyuowhs password=q1PQaJuep4qslGI0FbdWs-qzENKVgKE2 dbname=cxyuowhs port=5432 sslmode=disable"
	dsn := DB_CONN_URL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal("Cannot connect to DB")
	}


	  
    

    
    

    return nil
}


