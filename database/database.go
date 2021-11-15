package database

import (
	"log"
	"time"

	"github.com/douglasandradeee/web_api_gin/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	connString := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error ao abrir conexão no banco de dados: ERROR - ", err.Error())
	}

	db = database
	config, _ := db.DB()
	//SetMaxIdleConns - define o número máximo de conexões no modo inativo pool de conexão.
	config.SetMaxIdleConns(10)
	//SetMaxOpenConns - define o número máximo de conexões abertas com o banco de dados.
	config.SetMaxOpenConns(100)
	//SetConnMaxLifetime - define a quantidade máxima de tempo que uma conexão pode ser reutilizada.
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDB() *gorm.DB {
	return db
}
