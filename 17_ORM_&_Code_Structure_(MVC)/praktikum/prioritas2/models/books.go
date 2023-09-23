package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        int    `json:"ID" form:"ID" gorm:"primaryKey"`
	Judul     string `json:"judul" form:"judul"`
	Penulis   string `json:"penulis" form:"penulis"`
	Penerbit  string `json:"penerbit" form:"penerbit"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
