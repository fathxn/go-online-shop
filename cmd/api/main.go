package main

import (
	"github.com/gofiber/fiber/v2"
	"go-online-shop/apps/auth"
	"go-online-shop/apps/product"
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

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)
	product.Init(router, db)
	router.Listen(config.Cfg.App.Port)
}
