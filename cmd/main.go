package main

import (
	"RalfDalfs/go-fiber/config"
	"RalfDalfs/go-fiber/internal/pages"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"log"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(":3000")
}
