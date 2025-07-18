package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"` // 确保有autoIncrement
	Name     string `gorm:"size:255;not null;unique"`
	Email    string `gorm:"size:255;not null"`
	Password string `gorm:"type:text;not null"`
	Avatar   string `gorm:"type:text;not null"`
	Posts    []Post `gorm:"foreignKey:AuthorID"`
}

type UserReq struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

//

type Category struct {
	ID    uint   `gorm:"primaryKey"`
	Tag   string `gorm:"type:text;not null"`
	Posts []Post `gorm:"foreignKey:CategoryID"`
}

type Post struct {
	ID         uint      `gorm:"primaryKey"`
	Title      string    `gorm:"type:text;not null"`
	Content    string    `gorm:"type:text;not null"`
	Thumbnail  string    `gorm:"type:text;not null"`
	Date       time.Time `gorm:"type:date;not null"`
	AuthorID   uint      `gorm:"not null"`
	CategoryID uint      `gorm:"not null"`

	Author   User     `gorm:"foreignKey:AuthorID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

// 仅公开安全字段的作者信息
type AuthorDTO struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// 仅公开安全字段的分类信息
type CategoryDTO struct {
	ID  uint   `json:"id"`
	Tag string `json:"tag"`
}

// 文章 DTO，替代原始 Post
type PostDTO struct {
	ID        uint        `json:"id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	Thumbnail string      `json:"thumbnail"`
	Date      time.Time   `json:"date"`
	Author    AuthorDTO   `json:"author"`
	Category  CategoryDTO `json:"category"`
}
