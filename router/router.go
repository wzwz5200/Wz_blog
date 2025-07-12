package router

import (
	hander "server/internal/handler/User"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {

	api := app.Group("/api")
	User := api.Group("/user")

	User.Get("reg", hander.User_reg)

}
