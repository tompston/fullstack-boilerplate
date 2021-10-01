package main

import (
	"fiber_server/settings"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/helmet/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()    // create the server
	app.Use(logger.New()) // logs the requests
	app.Use(csrf.New())   // csrf protection
	app.Use(helmet.New()) // security headers

	//// Enable cors + use the url from which the frontend makes API requests if needed
	// app.Use(cors.New(cors.Config{
	// 	// AllowOrigins: settings.Config("BASE_URL"),
	// }))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Howdy")
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("the endpoint for the api")
	})

	port, _ := strconv.Atoi(settings.Config("GOLANG_PORT")) // import the .env int

	app.Listen(fmt.Sprintf("%s%d", ":", port)) // run the server
}
