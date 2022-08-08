package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name  string    `gorm:"type:varchar(255)"`
	Email string    `gorm:"uniqueIndex"`
}

func (User) UsersTable() string {
	return "users"
}
