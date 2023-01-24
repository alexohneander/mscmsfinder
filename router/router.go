package router

import (
	"mscmsfinder/handler"
	"mscmsfinder/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/check", handler.Hello)

	// CMS Finder
	cms := api.Group("/cms")
	cms.Get("/find", handler.FindCMS)

	// Monitoring
	middleware.SetupMonitoring(app)
}
