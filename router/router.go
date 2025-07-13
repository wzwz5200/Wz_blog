package router

import (
	hander "server/internal/handler/User"
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {

	api := app.Group("/api")
	User := api.Group("/user")

	User.Post("reg", hander.User_reg)
	User.Post("login", hander.User_Login)

	User.Get("/hello", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
		return c.SendString("Hello JWT")
	})

	

}
