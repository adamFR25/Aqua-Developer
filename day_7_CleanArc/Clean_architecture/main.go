package main

import (
	"fmt"
	"log"
	"os"
	"tesla/storage"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic(err)
	}

	database := &storage.Database{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Dbname:   os.Getenv("DBNAME"),
		Port:     os.Getenv("PORT"),
		Sslmode:  os.Getenv("SSLMODE"),
		TimeZone: os.Getenv("TIMEZONE"),
	}

	db, errdb := storage.Connection(database)
	if errdb != nil {
		log.Panic()
	}
	fmt.Println("Success", db)
	e := echo.New()

	e.Logger.Fatal(e.Start("127.0.0.1:5002"))
}
