package database

import (
	"fmt"
	"log"
	"time"

	"github.com/meajudaaqui/user-service/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "sslmode=require host=ec2-34-233-115-14.compute-1.amazonaws.com port=5432 user=wwueqrowenxyjs dbname=djus4lp6a277j sslmode=disable password=d7fb29481b93549a5c5fc7d3384099f32e37c438b8dcee6e562f0a5b38fac9f6"

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "sslmode=require host=ec2-34-233-115-14.compute-1.amazonaws.com port=5432 user=wwueqrowenxyjs dbname=djus4lp6a277j sslmode=disable password=d7fb29481b93549a5c5fc7d3384099f32e37c438b8dcee6e562f0a5b38fac9f6", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	  }), &gorm.Config{})
	if err != nil {
		fmt.Println("nao foi possível conectar ao banco de dados")
		log.Fatal("Error: ", err)
	}

	db = database
	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
