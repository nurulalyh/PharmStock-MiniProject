package response

import (
	"time"

	"gorm.io/gorm"
)

type LoginResponse struct {
	Id           string `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Username     string `json:"username" form:"username"`
	Role         string `json:"role" form:"role"`
	Token        any    `json:"token" form:"token"`
	RefreshToken any    `json:"refresh_token" form:"refresh_token"`
}

type UsersResponse struct {
	Id        string         `json:"id" form:"id"`
	Name      string         `json:"name" form:"name"`
	Username  string         `json:"username" form:"username"`
	Password  string         `json:"password" form:"password"`
	Email     string         `json:"email" form:"email"`
	Phone     string         `json:"phone" form:"phone"`
	Address   string         `json:"address" form:"address"`
	Role      string         `json:"role" form:"role"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
