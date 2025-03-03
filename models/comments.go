package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Body   string    `json:"body"`
	BlogID uuid.UUID `json:"blog_id"`
	UserID uint      `json:"user_id"`
}
