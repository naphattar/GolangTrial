package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/naphattar/KaihorBackend/configs"
	"github.com/naphattar/KaihorBackend/controllers"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// AllowOrigin defines a list of origins that may access the resource.
	//
	// Optional. Default value "*"
	AllowOrigins string

	// AllowMethods defines a list of methods allowed when accessing the resource.
	// This is used in response to a preflight request.
	//
	// Optional. Default value "GET,POST,HEAD,PUT,DELETE,PATCH"
	AllowMethods string

	// AllowHeaders defines a list of request headers that can be used when
	// making the actual request. This is in response to a preflight request.
	//
	// Optional. Default value "".
	AllowHeaders string

	// AllowCredentials indicates whether or not the response to the request
	// can be exposed when the credentials flag is true. When used as part of
	// a response to a preflight request, this indicates whether or not the
	// actual request can be made using credentials.
	//
	// Optional. Default value false.
	AllowCredentials bool

	// ExposeHeaders defines a whitelist headers that clients are allowed to
	// access.
	//
	// Optional. Default value "".
	ExposeHeaders string

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached.
	//
	// Optional. Default value 0.
	MaxAge int
}

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
