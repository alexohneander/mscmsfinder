package middleware

import (
	"mscmsfinder/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupMonitoring(app *fiber.App) {
	isMonitoringActive := config.Config("MONITORING")

	if isMonitoringActive == "true" {
		app.Get("/monitor", monitor.New(monitor.Config{Title: "msContact Monitoring"}))
	}
}
