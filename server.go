package main

import (
	"github.com/gofiber/fiber/v2"
)

var HOST string = "0.0.0.0"
var PORT string = "3000"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	// fmt.Printf("Server running at %s:%s", HOST, PORT)
	app.Listen(HOST + ":" + PORT)
}
