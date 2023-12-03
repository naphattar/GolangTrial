package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/naphattar/KaihorBackend/configs"
	"github.com/naphattar/KaihorBackend/controllers"
)

func main() {
	app := fiber.New()

	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is runnning")
	})

	// connecting to Database
	configs.ConnectDB()

	// for admin to update the data
	app.Get("/admin/update", controllers.UpdateCampDataFromSpreadSheet)

	// for CRUD campData from Database
	app.Get("/camp", controllers.GetAllCampData)
	app.Get("/camp/id/:id", controllers.GetCampDatabyID)
	app.Get("/camp/location/:location", controllers.GetCampDatabyLocation)
	app.Get("/camp/keyword/:keyword", controllers.GetCampDatabyKeyword)
	app.Get("/camp/year/:year", controllers.GetCampDatabyYear)

	err := app.Listen(configs.EnvPORT())
	if err != nil {
		panic(err)
	}
}
