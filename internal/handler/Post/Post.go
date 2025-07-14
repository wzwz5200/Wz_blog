package post

import (
	"server/initdb"
	post "server/internal/service/Post"

	"github.com/gofiber/fiber/v2"
)

func GetAll_Post(c *fiber.Ctx) error {

	db := initdb.DB

	if post.GetAll_Post_Service(db, c) {

		
		return c.SendStatus(200)
	}

	return c.SendStatus(400)
}
