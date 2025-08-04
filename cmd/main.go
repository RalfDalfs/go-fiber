package main

import (
	"RalfDalfs/go-fiber/config"
	"RalfDalfs/go-fiber/internal/pages"
	"RalfDalfs/go-fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(":3000")
}
