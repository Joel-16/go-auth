package models

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    uint      `json:"user_id"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
