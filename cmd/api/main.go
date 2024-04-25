package main

import (
	"go-online-shop/external/database"
	"go-online-shop/internal/config"
	"log"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("Successfully connected to database")
	}
}
