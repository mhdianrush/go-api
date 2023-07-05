package models

import "time"

type Book struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	AuthorId    uint      `json:"author_id"`
	Writer      Author    `gorm:"foreignKey:AuthorId" json:"writer"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookResponse struct {
	Id          uint               `json:"id"`
	Title       string             `json:"title"`
	AuthorId    uint               `json:"-"`
	Writer      AuthorBookResponse `gorm:"foreignKey:AuthorId" json:"writer"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
