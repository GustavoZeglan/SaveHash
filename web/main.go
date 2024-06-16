package main

import (
	"github.com/GustavoZeglan/SaveHash/core/db"
	"github.com/GustavoZeglan/SaveHash/web/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router.Initialize(database)
}
