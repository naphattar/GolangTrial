package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naphattar/KaihorBackend/configs"
	"github.com/naphattar/KaihorBackend/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is runnning")
	})

	// connecting to Database
	configs.ConnectDB()

	// for admin to update the data
	app.Get("/admin/update", controllers.UpdateCampDataFromSpreadSheet)

	// for CRUD campData from Database
	app.Get("/camp", controllers.GetAllCampData)
	app.Get("/camp/:id", controllers.GetCampDatabyID)

	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}
}
