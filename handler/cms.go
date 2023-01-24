package handler

import (
	"fmt"
	"mscmsfinder/client"
	"mscmsfinder/database"
	"mscmsfinder/model"

	"github.com/gofiber/fiber/v2"
)

// Try to Parse CMS Information
func FindCMS(c *fiber.Ctx) error {

	message := "requesting"

	url := c.Query("url")
	message = TestEndpoints(url)

	return c.JSON(fiber.Map{"status": "success", "message": message, "data": nil})
}

func TestEndpoints(url string) string {
	message := "testing Endpoints"

	db := database.DB
	var systems []model.CMS
	db.Find(&systems)

	for _, cms := range systems {
		message = client.TestEndpoint(url, cms.UrlPath)
		fmt.Println(cms.Title)
	}

	return message
}
