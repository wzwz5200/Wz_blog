package main

import (
	"server/initdb"
	"server/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	initdb.Initdb()
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: false,
	})
	router.InitRouter(app)

	app.Listen(":3000")
}
