package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	// host := "localhost"
	// username := "postgres"
	// password := "database123"
	database := "e_commerce"
	// port := "5432"
	// sslmode := "disable"

	// dsn := "host=" + host + "user=" + username + "password=" + password + "dbname=" + database + "port=" + port + "sslmode=disable TimeZone=Asia/Shanghai"
	var dsn string = "host=localhost user=postgres password=database123 dbname=e_commerce port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		panic("Can't connect to database" + database + "")
	}
}
