package User

import (
	"server/initdb"
	User "server/internal/service/user"
	"server/model"

	"github.com/gofiber/fiber/v2"
)

func User_reg(c *fiber.Ctx) error {

	db := initdb.DB

	Req := model.UserReq{}

	c.BodyParser(&Req)

	if User.Reg_service(db, Req, c) {

		return c.SendStatus(200)
	}

	return c.SendStatus(400)

}



func User_Login(c *fiber.Ctx) error {

	db := initdb.DB
	Req := model.UserReq{}

	c.BodyParser(&Req)

	if User.Login_Service(db, Req, c) {
		return c.SendStatus(200)

	}
	c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "用户密码或账户错误",
	})

	return c.SendStatus(400)
}
