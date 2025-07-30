package main

import (
	"RalfDalfs/go-fiber/config"
	"RalfDalfs/go-fiber/internal/pages"
	"RalfDalfs/go-fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	app := fiber.New()
	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(":3000")
}
