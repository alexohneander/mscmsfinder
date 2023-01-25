package handler

import (
	"mscmsfinder/client"
	"mscmsfinder/database"
	"mscmsfinder/model"
	"mscmsfinder/types"

	"github.com/gofiber/fiber/v2"
)

// Try to Parse CMS Information
func FindCMS(c *fiber.Ctx) error {

	var res types.ParseResponse

	url := c.Query("url")
	res = testEndpoints(url)

	return c.JSON(fiber.Map{"status": res.Status, "message": res.Message, "data": res.Payload})
}

func testEndpoints(url string) types.ParseResponse {
	var res types.ParseResponse
	res.Message = "testing Endpoints"

	db := database.DB
	var systems []model.CMS
	db.Find(&systems)

	for _, cms := range systems {
		if !res.Found {
			res = client.TestEndpoint(url, cms)
		}
	}

	return res
}
