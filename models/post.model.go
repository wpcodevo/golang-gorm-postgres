package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title     string    ``
	Content   string    ``
	Image     string    ``
	User      string    ``
	CreateAt  time.Time ``
	UpdatedAt time.Time ``
}

type CreatePostRequest struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	Content   string    `json:"content" bson:"content" binding:"required"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty"`
	User      string    `json:"user" bson:"user" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdatePost struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Content   string    `json:"content,omitempty" bson:"content,omitempty"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty"`
	User      string    `json:"user,omitempty" bson:"user,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
