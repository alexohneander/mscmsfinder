package handler

import (
	"mscmsfinder/client"

	"github.com/gofiber/fiber/v2"
)

// Try to Parse CMS Information
func FindCMS(c *fiber.Ctx) error {

	message := "requesting"

	url := c.Query("url")
	message = client.RequestURL(url)

	return c.JSON(fiber.Map{"status": "success", "message": message, "data": nil})
}
