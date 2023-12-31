package user

import (
	"time"
)

type User struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Occupation     string    `json:"occupation"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"password_hash"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateAt       time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP()"`
}
