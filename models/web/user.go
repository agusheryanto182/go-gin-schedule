package web

import "time"

type UserCreateRequest struct {
	Email string `binding:"required,email,max=30" validate:"required,max=30,email" json:"email"`
}

type UserResponse struct {
	UserId    int       `gorm:"primaryKey" json:"user_id"`
	Email     string    `binding:"required,email,max=30" validate:"required,max=30,email" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
