package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func Connection(database *Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		database.Host, database.User, database.Password, database.Dbname, database.Port, database.Sslmode, database.TimeZone)
	db, errdb := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errdb != nil {
		log.Panic("error")
		return db, errdb
	}

	return db, nil
}
