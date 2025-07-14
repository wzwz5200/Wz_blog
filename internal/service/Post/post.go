package post

import (
	"math"
	"server/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAll_Post_Service(db *gorm.DB, c *fiber.Ctx) bool {
	// 默认分页参数
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 获取总数量
	var total int64
	db.Model(&model.Post{}).Count(&total)

	// 查询分页内容
	var posts []model.Post
	result := db.Preload("Author").Preload("Category").
		Limit(pageSize).Offset(offset).Order("date DESC").Find(&posts)

	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
		return false
	}

	// 转换为 DTO
	var postDTOs []model.PostDTO
	for _, post := range posts {
		postDTOs = append(postDTOs, model.PostDTO{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			Thumbnail: post.Thumbnail,
			Date:      post.Date,
			Author: model.AuthorDTO{
				ID:     post.Author.ID,
				Name:   post.Author.Name,
				Avatar: post.Author.Avatar,
			},
			Category: model.CategoryDTO{
				ID:  post.Category.ID,
				Tag: post.Category.Tag,
			},
		})
	}

	// 返回分页结构
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"page":      page,
		"pageSize":  pageSize,
		"total":     total,
		"totalPage": int(math.Ceil(float64(total) / float64(pageSize))),
		"data":      postDTOs,
	})

	return true
}
