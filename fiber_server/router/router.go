package router

import (
	"fiber_server/settings/database"

	"github.com/gofiber/fiber/v2"
)

func Url(app *fiber.App) {
	app.Get("/", Home)
	app.Get("/debug", Debug)

}

func Home(c *fiber.Ctx) error {
	return c.SendString("Howdy")
}

// a random route that can be used to test things
func Debug(c *fiber.Ctx) error {

	db, err := database.GetDbConnSqlx()
	if err != nil {
		panic(err)
	}
	db.Exec("")
	db.Close()

	return c.SendString("hello there")
}
