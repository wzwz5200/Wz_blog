package router

import (
	post "server/internal/handler/Post"
	hander "server/internal/handler/User"
	"server/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func InitRouter(app *fiber.App) {

	app.Get("/metrics", monitor.New(monitor.Config{Title: ""}))
	api := app.Group("/api")

	User := api.Group("/user")

	User.Post("reg", hander.User_reg)
	User.Post("login", hander.User_Login)

	User.Get("/hello", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
		return c.SendString("Hello JWT")
	})

	posts := api.Group("/post")

	posts.Get("posts", post.GetAll_Post)

}
