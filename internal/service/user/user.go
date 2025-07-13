package User

import (
	"errors"
	"fmt"
	"server/middleware"
	"server/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

func ValidateUser(req model.UserReq) error {
	return validate.Struct(req)
}

func HashPassword(password string) (string, error) {

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("加密密码失败: %w", err)
	}
	return string(hashedBytes), nil
}

// CheckPassword 验证密码
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func Reg_service(db *gorm.DB, req model.UserReq, c *fiber.Ctx) bool {
	// 1. 表单验证
	if err := validate.Struct(req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "用户信息格式错误",
			"details": err.Error(),
		})
		return false
	}

	// 2. 检查用户名是否已存在
	var existingUser model.User
	if err := db.Where("name = ?", req.Name).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 不是"记录未找到"的其他数据库错误
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "数据库查询错误",
			})
			return false
		}
		// 如果是"记录未找到"错误，继续执行注册流程
	} else {
		// 找到记录，用户名已存在
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "用户名已存在",
		})
		return false
	}

	//	3. 检查邮箱是否已存在
	if err := db.Where("email = ?", req.Email).Take(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "数据库查询错误",
			})
			return false
		}
	} else {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "邮箱已注册",
		})
		return false
	}

	// 4. 密码加密
	hashed, err := HashPassword(req.Password)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "密码加密失败",
		})
		return false
	}

	// 5. 创建用户
	newUser := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Avatar:   "1.png",
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "用户创建失败",
			"details": err.Error(),
		})
		return false
	}

	// 6. 成功响应
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "用户注册成功",
		"user": fiber.Map{
			"id":    newUser.ID,
			"name":  newUser.Name,
			"email": newUser.Email,
		},
	})
	return true
}

var SecretKey = middleware.SecretKey

func Login_Service(db *gorm.DB, req model.UserReq, c *fiber.Ctx) bool {
	var NewUser model.User

	// 查找用户
	result := db.Where("name = ?", req.Name).First(&NewUser)
	if result.Error != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "用户不存在",
		})
		return false
	}

	// 校验密码
	match := CheckPassword(req.Password, NewUser.Password)
	if !match {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "密码错误",
		})
		return false
	}

	//密码正确，生成 JWT
	token, err := middleware.GenerateJWT(NewUser.Name)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "生成 Token 失败",
		})
		return false
	}

	// 返回 Token 和用户信息
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "用户登录成功",
		"token":   token,
		"user": fiber.Map{
			"id":    NewUser.ID,
			"name":  NewUser.Name,
			"email": NewUser.Email,
		},
	})

	return true
}
