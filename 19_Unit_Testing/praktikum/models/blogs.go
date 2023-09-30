package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `json:"user_id" form:"user_id"`
	Judul     string `json:"judul" form:"judul"`
	Konten    string `json:"konten" form:"konten"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BlogResponse struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	Judul     string `json:"judul" form:"judul"`
	Konten    string `json:"konten" form:"konten"`
}