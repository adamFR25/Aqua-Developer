package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
	TimeZone string
}

// var DB *gorm.DB

func ConnectDB(database *Database) (gorm.DB, error) {
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		database.Host, database.User, database.Password, database.Dbname, database.Port, database.Sslmode, database.TimeZone)
	

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})
	if err != nil {
		log.Panic("error ", dsn)
	}

	return DB, nil
}
