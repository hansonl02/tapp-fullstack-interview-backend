package main

import (
	"github.com/gofiber/fiber/v2"
)

type GetDomains struct {
	Name string `json:"name"`
}

//	 GetDomains gets a list of available domains from a given business name
//	 {
//			"name": "name of business"
//	 }
func (a *App) GetDomains(c *fiber.Ctx) error {
	// TODO
	return c.JSON(SuccessResponse("Success!"))
}
