package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name  string    `json:"name" gorm:"type:varchar(255)" binding:"required"`
	Email string    `json:"email" gorm:"uniqueIndex" binding:"required"`
}

func (User) UsersTable() string {
	return "users"
}
